package issueops

import (
	"testing"

	"github.com/steveyegge/beads/internal/types"
	publicops "github.com/steveyegge/beads/issueops"
)

// A gate with no await_type is invisible to every notifier (ga-49tby1: all
// nine live human-authored gates were null-typed because `bd create -t gate`
// never set one). PrepareIssueForInsert is the single seam every create path
// flows through, so the default belongs there, not in one CLI command.
func TestPrepareIssueForInsertDefaultsGateAwaitTypeToHuman(t *testing.T) {
	gate := &types.Issue{
		ID:        "bd-gate-default",
		Title:     "gate",
		Status:    types.StatusOpen,
		Priority:  2,
		IssueType: types.TypeGate,
	}
	if err := PrepareIssueForInsert(gate, nil, nil); err != nil {
		t.Fatalf("PrepareIssueForInsert() error = %v", err)
	}
	if gate.AwaitType != "human" {
		t.Fatalf("gate with no await_type must default to human, got %q", gate.AwaitType)
	}
}

func TestPrepareIssueForInsertKeepsExplicitGateAwaitType(t *testing.T) {
	gate := &types.Issue{
		ID:        "bd-gate-timer",
		Title:     "gate",
		Status:    types.StatusOpen,
		Priority:  2,
		IssueType: types.TypeGate,
		AwaitType: "timer",
	}
	if err := PrepareIssueForInsert(gate, nil, nil); err != nil {
		t.Fatalf("PrepareIssueForInsert() error = %v", err)
	}
	if gate.AwaitType != "timer" {
		t.Fatalf("explicit await_type must be preserved, got %q", gate.AwaitType)
	}
}

func TestPrepareIssueForInsertLeavesNonGateAwaitTypeEmpty(t *testing.T) {
	task := &types.Issue{
		ID:        "bd-task-plain",
		Title:     "task",
		Status:    types.StatusOpen,
		Priority:  2,
		IssueType: types.TypeTask,
	}
	if err := PrepareIssueForInsert(task, nil, nil); err != nil {
		t.Fatalf("PrepareIssueForInsert() error = %v", err)
	}
	if task.AwaitType != "" {
		t.Fatalf("non-gate issue must not receive an await_type, got %q", task.AwaitType)
	}
}

// The 9 live null-typed gates cannot be fixed at create time; the update path
// must carry await_type so they can be re-typed in place instead of re-minted
// under new ids that grooming prose already cites.
func TestUpdateFieldsCarriesAwaitType(t *testing.T) {
	patch := publicops.IssuePatch{}
	patch.AwaitType.Set = true
	patch.AwaitType.Value = "human"

	updates := UpdateFields(patch)
	got, ok := updates["await_type"]
	if !ok {
		t.Fatalf("UpdateFields must include await_type when set; got %v", updates)
	}
	if got != "human" {
		t.Fatalf("await_type = %v, want human", got)
	}
	if !IsAllowedUpdateField("await_type") {
		t.Fatalf("await_type must be an allowed update field")
	}
}
