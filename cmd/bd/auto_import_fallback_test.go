package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/steveyegge/beads/internal/storage"
	"github.com/steveyegge/beads/internal/types"
)

// fakeFallbackStore is a storage.DoltStorage that deliberately does NOT
// implement the jsonlImporter (ImportJSONLData) interface, so
// maybeAutoImportJSONL routes it through the non-embedded fallback path —
// the same path a server-mode (dolt sql-server) store takes. It records
// whether the fallback actually attempted an import so tests can assert the
// emptiness guard.
type fakeFallbackStore struct {
	storage.DoltStorage // embedded interface; unimplemented methods panic if called

	totalIssues  int
	importCalled bool
}

func (f *fakeFallbackStore) GetStatistics(_ context.Context) (*types.Statistics, error) {
	return &types.Statistics{TotalIssues: f.totalIssues}, nil
}

func (f *fakeFallbackStore) GetConfig(_ context.Context, _ string) (string, error) {
	f.importCalled = true
	return "", nil
}

func (f *fakeFallbackStore) SetConfig(_ context.Context, _, _ string) error {
	f.importCalled = true
	return nil
}

func (f *fakeFallbackStore) CreateIssuesWithFullOptions(_ context.Context, _ []*types.Issue, _ string, _ storage.BatchCreateOptions) error {
	f.importCalled = true
	return nil
}

func (f *fakeFallbackStore) Commit(_ context.Context, _ string) error {
	return nil
}

// writeJSONL drops a minimal valid issues.jsonl into beadsDir for the
// auto-import path to discover.
func writeJSONL(t *testing.T, beadsDir string) {
	t.Helper()
	line := `{"id":"ff-1","title":"Imported","status":"open","priority":2,"issue_type":"task","created_at":"2026-06-02T00:00:00Z","updated_at":"2026-06-02T00:00:00Z"}` + "\n"
	if err := os.WriteFile(filepath.Join(beadsDir, "issues.jsonl"), []byte(line), 0o644); err != nil {
		t.Fatalf("writing issues.jsonl: %v", err)
	}
}

// TestFallbackAutoImportSkipsNonEmpty is the regression test for gt-5mha:
// the non-embedded (server-mode) fallback in maybeAutoImportJSONL must NOT
// re-import the JSONL when the database already has issues. Without the
// emptiness guard, every non-read-only command re-imports the entire JSONL —
// for a large town JSONL (hq at 5+ MB) that turns routine commands like
// 'gt mail inbox' into multi-second hangs.
func TestFallbackAutoImportSkipsNonEmpty(t *testing.T) {
	beadsDir := t.TempDir()
	writeJSONL(t, beadsDir)

	store := &fakeFallbackStore{totalIssues: 5} // database is NOT empty
	maybeAutoImportJSONL(context.Background(), store, beadsDir)

	if store.importCalled {
		t.Fatal("fallback auto-import re-imported into a non-empty database; " +
			"the emptiness guard should have skipped it (gt-5mha hang root cause)")
	}
}

// TestFallbackAutoImportRunsWhenEmpty verifies the guard does not regress the
// intended GH#2994 upgrade-recovery behavior: a genuinely empty database still
// auto-imports from the JSONL.
func TestFallbackAutoImportRunsWhenEmpty(t *testing.T) {
	beadsDir := t.TempDir()
	writeJSONL(t, beadsDir)

	store := &fakeFallbackStore{totalIssues: 0} // database IS empty
	maybeAutoImportJSONL(context.Background(), store, beadsDir)

	if !store.importCalled {
		t.Fatal("fallback auto-import skipped an empty database; " +
			"upgrade-recovery import (GH#2994) must still run when empty")
	}
}
