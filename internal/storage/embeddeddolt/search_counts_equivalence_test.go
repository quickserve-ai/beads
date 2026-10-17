//go:build cgo

package embeddeddolt_test

import (
	"database/sql"
	"testing"

	"github.com/steveyegge/beads/internal/storage/embeddeddolt"
	"github.com/steveyegge/beads/internal/storage/searchcountsequiv"
)

// TestSearchCountsIDsFirstMatchesPredicateForm_Embedded runs the counts-search
// equivalence scenarios (internal/storage/searchcountsequiv) on the embedded
// engine: a bounded page answered ids-first, and the by-IDs form on its own,
// must return the predicate form's rows (#7635).
func TestSearchCountsIDsFirstMatchesPredicateForm_Embedded(t *testing.T) {
	skipUnlessEmbeddedDolt(t)
	searchcountsequiv.Run(t, func(t *testing.T) *sql.DB {
		t.Helper()
		fixture := newPristineEmbeddedDoltFixture(t, "eq")
		// Fleet adaptation (v1.3.4 base): upstream opens a counting-wrapped
		// raw connection (openCountedConn, sqlcount; #7051's harness, absent
		// here). The counts were never read, so a plain OpenSQL handle
		// serves. The embedded engine is exclusive, so the fixture's store is
		// closed first: a raw handle opened beside a live store retries its
		// open forever.
		closeEmbeddedDoltStore(t, fixture.store)
		db, cleanup, err := embeddeddolt.OpenSQL(t.Context(), fixture.dataDir, fixture.database, "main")
		if err != nil {
			t.Fatalf("OpenSQL: %v", err)
		}
		t.Cleanup(func() { _ = cleanup() })
		return db
	})
}
