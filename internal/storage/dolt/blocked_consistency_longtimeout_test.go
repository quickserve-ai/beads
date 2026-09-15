package dolt

import (
	"testing"

	mysql "github.com/go-sql-driver/mysql"
)

// TestCountIsBlockedInconsistencies_PinnedLongTimeoutConnection guards the two
// properties that make the doctor "Blocked State" check trustworthy on a
// remote, shared store (ga-fo8w65). Both are silent when broken, which is why
// they are asserted rather than assumed:
//
//  1. BRANCH PIN. A fresh one-shot connection defaults to the default branch.
//     setupTestStore checks out an isolated test branch with a raw
//     CALL DOLT_CHECKOUT on store.db, so an unpinned connection would count on
//     the default branch — schema-only and issue-less — and return 0 with a
//     NIL ERROR. An error check alone cannot catch that; only asserting the
//     count matches the corruption we planted can.
//
//  2. DEDICATED LONG-TIMEOUT CONNECTION. store.db is opened once with a
//     baked-in DSN, so mutating store.connStr afterward cannot affect it. Only
//     a path that re-parses connStr per call (openLongTimeoutConn, via
//     withReadTxLongTimeout) is affected. Breaking connStr must therefore break
//     this call — if it still succeeds, the count is riding the shared pool and
//     its 10s ReadTimeout, which is what killed the check as "invalid
//     connection" against the hub in the first place.
func TestCountIsBlockedInconsistencies_PinnedLongTimeoutConnection(t *testing.T) {
	store, cleanup := setupTestStore(t)
	defer cleanup()
	ctx, cancel := testContext(t)
	defer cancel()

	// Correct graph via the normal write path: bm-w blocked on open bm-x.
	seedBlockedPair(ctx, t, store, true)
	if n := countInconsistencies(ctx, t, store.db); n != 0 {
		t.Fatalf("precondition: consistent graph must count 0, got %d", n)
	}
	if got, err := store.CountIsBlockedInconsistencies(ctx); err != nil || got != 0 {
		t.Fatalf("consistent graph via store: want 0/nil, got %d/%v", got, err)
	}

	// Plant a known inconsistency: clear the flag with no recompute, the shape
	// a merge that bypassed the recompute hook leaves behind.
	if _, err := store.db.ExecContext(ctx, "UPDATE issues SET is_blocked = 0 WHERE id = 'bm-w'"); err != nil {
		t.Fatalf("corrupt is_blocked: %v", err)
	}
	if n := countInconsistencies(ctx, t, store.db); n != 1 {
		t.Fatalf("precondition: corrupted graph must count 1 on the store's branch, got %d", n)
	}

	// ARM 1 — the pin. Must see the planted row, not the default branch's zero.
	got, err := store.CountIsBlockedInconsistencies(ctx)
	if err != nil {
		t.Fatalf("CountIsBlockedInconsistencies failed: %v", err)
	}
	if got != 1 {
		t.Fatalf("want 1 inconsistency, got %d — a 0 here means the one-shot "+
			"connection counted on the DEFAULT branch instead of the store's "+
			"checked-out branch (pinStoreBranch via withReadTxLongTimeout)", got)
	}

	// ARM 2 — the dedicated connection. Break connStr to an address that fails
	// DNS resolution fast and permanently (RFC 2606 .invalid), a clean signal
	// distinct from "connection refused", which the retry layer would treat as
	// transient and spend up to serverRetryMaxElapsed retrying.
	cfg, err := mysql.ParseDSN(store.connStr)
	if err != nil {
		t.Fatalf("failed to parse store.connStr: %v", err)
	}
	cfg.Addr = "ga-fo8w65-test.invalid:3306"
	store.connStr = cfg.FormatDSN()

	if _, err := store.CountIsBlockedInconsistencies(ctx); err == nil {
		t.Fatal("expected the count to fail after store.connStr was broken; " +
			"if it still succeeds it is reading through the shared pool " +
			"(store.db/UnderlyingDB) instead of a fresh connection via " +
			"openLongTimeoutConn/withReadTxLongTimeout — the pooled handle's " +
			"10s ReadTimeout is exactly what blinded this check on the hub")
	}
}
