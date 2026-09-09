package dolt

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/steveyegge/beads/internal/configfile"
)

// TestApplyConfigDefaultsAppliesPoolTimeoutLadder pins the pool-deadline knobs
// at the seam every open path shares. applyResolvedConfig applied the
// BEADS_DOLT_POOL_READ_TIMEOUT / dolt.pool-read-timeout ladder (#5089), but
// only callers of NewFromConfig* pass through it: the CLI's own store open and
// bd serve's provider hand-build their Config and go straight to New →
// applyConfigDefaults, so every `bd` command in server mode kept the built-in
// 10s deadline whatever the knob said (gastownhall/beads#6144). New is the one
// constructor all of them call, so the ladder has to hold there.
func TestApplyConfigDefaultsAppliesPoolTimeoutLadder(t *testing.T) {
	t.Run("env vars populate the pool deadlines on the constructor path", func(t *testing.T) {
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "90s")
		t.Setenv("BEADS_DOLT_POOL_WRITE_TIMEOUT", "45")
		cfg := &Config{ServerMode: true, Database: "ladder", Path: t.TempDir()}

		applyConfigDefaults(cfg)

		if cfg.PoolReadTimeout != 90*time.Second {
			t.Fatalf("PoolReadTimeout = %v, want 90s from BEADS_DOLT_POOL_READ_TIMEOUT", cfg.PoolReadTimeout)
		}
		if cfg.PoolWriteTimeout != 45*time.Second {
			t.Fatalf("PoolWriteTimeout = %v, want 45s (bare number = seconds)", cfg.PoolWriteTimeout)
		}
		if dsn := buildServerDSN(cfg, cfg.Database); !strings.Contains(dsn, "readTimeout=1m30s") {
			t.Fatalf("buildServerDSN did not carry the env deadline: %s", dsn)
		}
	})

	t.Run("caller-set pool deadlines win over env vars", func(t *testing.T) {
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "90s")
		cfg := &Config{ServerMode: true, Database: "ladder", Path: t.TempDir(), PoolReadTimeout: 2 * time.Minute}

		applyConfigDefaults(cfg)

		if cfg.PoolReadTimeout != 2*time.Minute {
			t.Fatalf("PoolReadTimeout = %v, want the caller's 2m", cfg.PoolReadTimeout)
		}
	})

	t.Run("unset knobs leave the built-in default in place", func(t *testing.T) {
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "")
		t.Setenv("BEADS_DOLT_POOL_WRITE_TIMEOUT", "")
		cfg := &Config{ServerMode: true, Database: "ladder", Path: t.TempDir()}

		applyConfigDefaults(cfg)

		if cfg.PoolReadTimeout != 0 || cfg.PoolWriteTimeout != 0 {
			t.Fatalf("pool deadlines = %v/%v, want 0/0 so buildServerDSN applies its default", cfg.PoolReadTimeout, cfg.PoolWriteTimeout)
		}
		if dsn := buildServerDSN(cfg, cfg.Database); !strings.Contains(dsn, "readTimeout=10s") {
			t.Fatalf("buildServerDSN default deadline missing: %s", dsn)
		}
	})
}

// TestApplyResolvedConfigReadsPoolTimeoutsFromDir pins the config.yaml rung for
// library consumers. The GetString rungs inside ApplyPoolTimeouts read a
// package-global viper populated only by cmd/bd's config.Initialize(), so for a
// process that links beads as a library (gc, including its supervisor) they
// always return "" and the rig's configured pool deadlines were silently
// ignored — the process ran the 10s default whatever config.yaml said
// (ga-dwobeb). applyResolvedConfig now falls back to a direct read of
// <beadsDir>/config.yaml, the same pattern dolt.auto-start has carried since it
// hit this exact hole.
func TestApplyResolvedConfigReadsPoolTimeoutsFromDir(t *testing.T) {
	writeConfigYAML := func(t *testing.T, beadsDir, body string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(beadsDir, "config.yaml"), []byte(body), 0o644); err != nil {
			t.Fatalf("writing config.yaml: %v", err)
		}
	}
	baseFileCfg := func() *configfile.Config {
		return &configfile.Config{Backend: configfile.BackendDolt, DoltDatabase: "ladder"}
	}

	t.Run("config.yaml populates the deadlines with no env and no global viper", func(t *testing.T) {
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "")
		t.Setenv("BEADS_DOLT_POOL_WRITE_TIMEOUT", "")
		beadsDir := t.TempDir()
		writeConfigYAML(t, beadsDir, "dolt:\n  pool-read-timeout: 120s\n  pool-write-timeout: 45\n")
		cfg := &Config{}

		if err := applyResolvedConfig(context.Background(), beadsDir, baseFileCfg(), cfg); err != nil {
			t.Fatalf("applyResolvedConfig: %v", err)
		}

		if cfg.PoolReadTimeout != 120*time.Second {
			t.Fatalf("PoolReadTimeout = %v, want 120s from config.yaml", cfg.PoolReadTimeout)
		}
		if cfg.PoolWriteTimeout != 45*time.Second {
			t.Fatalf("PoolWriteTimeout = %v, want 45s from config.yaml (bare number = seconds)", cfg.PoolWriteTimeout)
		}
	})

	t.Run("env var wins over config.yaml", func(t *testing.T) {
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "90s")
		t.Setenv("BEADS_DOLT_POOL_WRITE_TIMEOUT", "")
		beadsDir := t.TempDir()
		writeConfigYAML(t, beadsDir, "dolt:\n  pool-read-timeout: 120s\n")
		cfg := &Config{}

		if err := applyResolvedConfig(context.Background(), beadsDir, baseFileCfg(), cfg); err != nil {
			t.Fatalf("applyResolvedConfig: %v", err)
		}

		if cfg.PoolReadTimeout != 90*time.Second {
			t.Fatalf("PoolReadTimeout = %v, want the env's 90s over the file's 120s", cfg.PoolReadTimeout)
		}
	})

	t.Run("hand-built Config with BeadsDir gets the fallback via ApplyPoolTimeouts", func(t *testing.T) {
		// Direct New() callers skip applyResolvedConfig; the fallback must
		// still reach them through ApplyPoolTimeouts when they identified
		// their project by setting BeadsDir (codex review finding on
		// ga-dwobeb: Config exposes BeadsDir, so "hand-built means dir-less"
		// was not a safe assumption).
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "")
		t.Setenv("BEADS_DOLT_POOL_WRITE_TIMEOUT", "")
		beadsDir := t.TempDir()
		writeConfigYAML(t, beadsDir, "dolt:\n  pool-read-timeout: 75s\n")
		cfg := &Config{BeadsDir: beadsDir}

		ApplyPoolTimeouts(cfg)

		if cfg.PoolReadTimeout != 75*time.Second {
			t.Fatalf("PoolReadTimeout = %v, want 75s from BeadsDir config.yaml", cfg.PoolReadTimeout)
		}
	})

	t.Run("no config.yaml leaves the deadlines unset", func(t *testing.T) {
		t.Setenv("BEADS_DOLT_POOL_READ_TIMEOUT", "")
		t.Setenv("BEADS_DOLT_POOL_WRITE_TIMEOUT", "")
		cfg := &Config{}

		if err := applyResolvedConfig(context.Background(), t.TempDir(), baseFileCfg(), cfg); err != nil {
			t.Fatalf("applyResolvedConfig: %v", err)
		}

		if cfg.PoolReadTimeout != 0 || cfg.PoolWriteTimeout != 0 {
			t.Fatalf("pool deadlines = %v/%v, want 0/0 so buildServerDSN applies its default", cfg.PoolReadTimeout, cfg.PoolWriteTimeout)
		}
	})
}
