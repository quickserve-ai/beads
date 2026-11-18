package workapi

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/steveyegge/beads/internal/types"
	"github.com/steveyegge/beads/issueops"
)

// TestListExcludeStatusAppendsAndValidates pins ListRequest.ExcludeStatus in
// the builder: it is appended to the default exclusions (never in their place),
// stands alone under AllFlag, splits commas, and an unknown name is
// ErrValidation. Fleet carry (ga-yuiof4), from upstream #7647's builder test.
func TestListExcludeStatusAppendsAndValidates(t *testing.T) {
	var cfg ListConfig
	f, err := BuildListFilter(issueops.ListRequest{ExcludeStatus: []string{"blocked, in_progress"}}, cfg)
	if err != nil {
		t.Fatalf("BuildListFilter: %v", err)
	}
	want := []types.Status{types.StatusClosed, types.StatusPinned, types.StatusBlocked, types.StatusInProgress}
	if !slices.Equal(f.ExcludeStatus, want) {
		t.Errorf("default listing ExcludeStatus = %v, want %v", f.ExcludeStatus, want)
	}
	f, err = BuildListFilter(issueops.ListRequest{AllFlag: true, ExcludeStatus: []string{"closed"}}, cfg)
	if err != nil {
		t.Fatalf("BuildListFilter: %v", err)
	}
	if want := []types.Status{types.StatusClosed}; !slices.Equal(f.ExcludeStatus, want) {
		t.Errorf("AllFlag ExcludeStatus = %v, want %v", f.ExcludeStatus, want)
	}
	f, err = BuildListFilter(issueops.ListRequest{AllFlag: true}, cfg)
	if err != nil {
		t.Fatalf("BuildListFilter: %v", err)
	}
	if len(f.ExcludeStatus) != 0 {
		t.Errorf("AllFlag with no ExcludeStatus = %v, want none (the carry must not change an existing request)", f.ExcludeStatus)
	}
	_, err = BuildListFilter(issueops.ListRequest{ExcludeStatus: []string{"clsoed"}}, cfg)
	if !errors.Is(err, issueops.ErrValidation) || !strings.HasPrefix(err.Error(), "invalid exclude-status ") {
		t.Errorf("unknown exclude status: err = %v, want ErrValidation starting \"invalid exclude-status \"", err)
	}
	// Blank entries and blank comma parts are skipped, not refused and not
	// read as a status.
	f, err = BuildListFilter(issueops.ListRequest{AllFlag: true, ExcludeStatus: []string{"", " , closed ,", "  "}}, cfg)
	if err != nil {
		t.Fatalf("BuildListFilter with blank entries: %v", err)
	}
	if want := []types.Status{types.StatusClosed}; !slices.Equal(f.ExcludeStatus, want) {
		t.Errorf("blank entries: ExcludeStatus = %v, want %v", f.ExcludeStatus, want)
	}
	// A workspace custom status is a valid name to exclude; the same name in a
	// workspace that does not define it is refused.
	custom := ListConfig{CustomStatuses: []types.CustomStatus{{Name: "review", Category: types.CategoryWIP}}}
	f, err = BuildListFilter(issueops.ListRequest{AllFlag: true, ExcludeStatus: []string{"review"}}, custom)
	if err != nil {
		t.Fatalf("BuildListFilter excluding a custom status: %v", err)
	}
	if want := []types.Status{"review"}; !slices.Equal(f.ExcludeStatus, want) {
		t.Errorf("custom status: ExcludeStatus = %v, want %v", f.ExcludeStatus, want)
	}
	if _, err = BuildListFilter(issueops.ListRequest{AllFlag: true, ExcludeStatus: []string{"review"}}, cfg); !errors.Is(err, issueops.ErrValidation) {
		t.Errorf("custom status absent from the workspace: err = %v, want ErrValidation", err)
	}
	_, err = BuildListFilter(issueops.ListRequest{ReadyFlag: true, ExcludeStatus: []string{"closed"}}, cfg)
	if !errors.Is(err, issueops.ErrValidation) {
		t.Errorf("ReadyFlag with ExcludeStatus: err = %v, want a refusal wrapping ErrValidation (the ready query cannot carry it)", err)
	}
}
