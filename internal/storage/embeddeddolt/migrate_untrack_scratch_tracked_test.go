//go:build cgo

package embeddeddolt_test

import (
	"database/sql"
	"strconv"
	"strings"
	"testing"
)

// TestIgnoredCursorScratchTrackedAtHeadIsSweptOnOpen covers the one window in
// which the #4356 untrack scratch table can end up TRACKED: a concurrent
// writer's blanket commit sweeps it into HEAD while a repair is in flight
// (dropIgnoredCursorScratch's own doc comment). The next writable open must
// then drop it AND commit the deletion, or a permanent delete delta is left
// in the working set.
//
// The drop is committed through a plain DOLT_ADD, which is a silent no-op on a
// table dolt_ignore covers. The scratch name is therefore only safe to sweep
// while no dolt_ignore pattern matches it; this pins that the sweep works on
// the real engine with whatever patterns the open itself seeds.
func TestIgnoredCursorScratchTrackedAtHeadIsSweptOnOpen(t *testing.T) {
	ctx := t.Context()
	pristine := newPristineEmbeddedDoltFixture(t, "scratchtracked")
	closeEmbeddedDoltStore(t, pristine.store)
	f := &legacyTrackedFixture{
		beadsDir: pristine.beadsDir,
		dataDir:  pristine.dataDir,
		database: "scratchtracked",
	}

	f.withRawConn(t, func(conn *sql.Conn) {
		backupCursorRows(t, ctx, conn)
		mustExecOn(t, ctx, conn, "CALL DOLT_ADD('-f', '"+untrackScratchTable+"')")
		mustExecOn(t, ctx, conn, "CALL DOLT_COMMIT('-m', 'unrelated blanket commit sweeps the scratch into HEAD')")
		if !trackedAtHead(t, ctx, conn, untrackScratchTable) {
			t.Fatalf("fixture did not commit %s into HEAD; the tracked-straggler state is not reproduced", untrackScratchTable)
		}
		if tablePresent(t, ctx, conn, "ignored_schema_migrations") == false {
			t.Fatal("fixture has no live cursor table; the resume path would restore instead of sweeping")
		}
	})

	f.reopenAndClose(t)

	f.withRawConn(t, func(conn *sql.Conn) {
		t.Logf("dolt_ignore rows matching the scratch name: %v", ignoreRowsMatching(t, conn, untrackScratchTable))
		t.Logf("dolt_status after reopen: %v", statusRows(t, conn))
		if tablePresent(t, ctx, conn, untrackScratchTable) {
			t.Errorf("%s is still in the working set after the reopen", untrackScratchTable)
		}
		if trackedAtHead(t, ctx, conn, untrackScratchTable) {
			t.Errorf("%s is still committed at HEAD after the reopen: the drop was never committed", untrackScratchTable)
		}
		for _, row := range statusRows(t, conn) {
			if strings.HasPrefix(row, untrackScratchTable+"|") {
				t.Errorf("dolt_status still lists the scratch table after the reopen: %s", row)
			}
		}
	})

	// A second open must not find anything left to do: a delete delta that a
	// plain DOLT_ADD cannot stage would otherwise persist across every open.
	f.reopenAndClose(t)
	f.withRawConn(t, func(conn *sql.Conn) {
		if trackedAtHead(t, ctx, conn, untrackScratchTable) {
			t.Errorf("%s is still committed at HEAD after a second reopen", untrackScratchTable)
		}
	})
}

// TestIgnoredCursorScratchPendingDeletionIsSweptOnOpen covers the state a
// sweep that dropped the tracked scratch table but never committed the drop
// leaves behind (a crash between the DROP and the forced staging, or any run
// of the pre-force plain-add sweep): the scratch is committed at HEAD and
// already gone from the working set. dolt_status still lists the pending
// deletion, but the reconcile gate must see it too, or nothing ever commits
// it — "__temp__%" keeps it out of every ignore-filtered dirty guard.
func TestIgnoredCursorScratchPendingDeletionIsSweptOnOpen(t *testing.T) {
	ctx := t.Context()
	pristine := newPristineEmbeddedDoltFixture(t, "scratchpending")
	closeEmbeddedDoltStore(t, pristine.store)
	f := &legacyTrackedFixture{
		beadsDir: pristine.beadsDir,
		dataDir:  pristine.dataDir,
		database: "scratchpending",
	}

	f.withRawConn(t, func(conn *sql.Conn) {
		requireTempPatternSeeded(t, conn)
		backupCursorRows(t, ctx, conn)
		mustExecOn(t, ctx, conn, "CALL DOLT_ADD('-f', '"+untrackScratchTable+"')")
		mustExecOn(t, ctx, conn, "CALL DOLT_COMMIT('-m', 'unrelated blanket commit sweeps the scratch into HEAD')")
		mustExecOn(t, ctx, conn, "DROP TABLE "+untrackScratchTable)
		if !trackedAtHead(t, ctx, conn, untrackScratchTable) || tablePresent(t, ctx, conn, untrackScratchTable) {
			t.Fatalf("fixture is not in the pending-deletion state (at HEAD, absent from the working set)")
		}
		if !tablePresent(t, ctx, conn, "ignored_schema_migrations") {
			t.Fatal("fixture has no live cursor table")
		}
	})

	for _, pass := range []string{"first", "second"} {
		f.reopenAndClose(t)
		f.withRawConn(t, func(conn *sql.Conn) {
			if trackedAtHead(t, ctx, conn, untrackScratchTable) {
				t.Errorf("%s reopen: %s is still committed at HEAD; its pending deletion was never committed", pass, untrackScratchTable)
			}
			for _, row := range statusRows(t, conn) {
				if strings.HasPrefix(row, untrackScratchTable+"|") {
					t.Errorf("%s reopen: dolt_status still lists the scratch table: %s", pass, row)
				}
			}
			if !tablePresent(t, ctx, conn, "ignored_schema_migrations") {
				t.Errorf("%s reopen: the live cursor table is gone", pass)
			}
		})
	}
}

// requireTempPatternSeeded fails the test unless dolt_ignore carries the
// exact "__temp__%"=true row, so a scratch-sweep test cannot pass because the
// name it exercises happens not to be ignored.
func requireTempPatternSeeded(t *testing.T, conn *sql.Conn) {
	t.Helper()
	var ignored bool
	err := conn.QueryRowContext(t.Context(),
		"SELECT ignored FROM dolt_ignore WHERE pattern = '__temp__%'").Scan(&ignored)
	if err != nil || !ignored {
		t.Fatalf("dolt_ignore has no \"__temp__%%\"=true row (ignored=%v, err=%v); the scratch name would not be ignored", ignored, err)
	}
}

// statusRows reads dolt_status as "table|staged|status" strings.
func statusRows(t *testing.T, conn *sql.Conn) []string {
	t.Helper()
	rows, err := conn.QueryContext(t.Context(), "SELECT table_name, staged, status FROM dolt_status ORDER BY table_name")
	if err != nil {
		t.Fatalf("reading dolt_status: %v", err)
	}
	defer func() { _ = rows.Close() }()
	var out []string
	for rows.Next() {
		var name, status string
		var staged bool
		if err := rows.Scan(&name, &staged, &status); err != nil {
			t.Fatalf("reading dolt_status: %v", err)
		}
		out = append(out, name+"|"+strconv.FormatBool(staged)+"|"+status)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("reading dolt_status: %v", err)
	}
	return out
}

// ignoreRowsMatching lists the dolt_ignore rows whose pattern matches table,
// using the same LIKE translation the production reader does for % and _
// (close enough for logging; dolt treats _ literally, LIKE does not).
func ignoreRowsMatching(t *testing.T, conn *sql.Conn, table string) []string {
	t.Helper()
	rows, err := conn.QueryContext(t.Context(),
		"SELECT pattern, ignored FROM dolt_ignore WHERE ? LIKE REPLACE(pattern, '*', '%') ORDER BY pattern", table)
	if err != nil {
		t.Fatalf("reading dolt_ignore: %v", err)
	}
	defer func() { _ = rows.Close() }()
	var out []string
	for rows.Next() {
		var pattern string
		var ignored bool
		if err := rows.Scan(&pattern, &ignored); err != nil {
			t.Fatalf("reading dolt_ignore: %v", err)
		}
		out = append(out, pattern+"="+strconv.FormatBool(ignored))
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("reading dolt_ignore: %v", err)
	}
	return out
}
