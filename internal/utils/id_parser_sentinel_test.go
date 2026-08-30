package utils

import (
	"context"
	"strings"
	"testing"
)

// The sentinel guard must fire before any storage access, so these tests run
// with a nil store: reaching the "storage is nil" error proves the guard was
// NOT triggered for that input. See ga-emfu6w — a jq-produced literal "null"
// substring-matched and silently mutated an unrelated issue.
func TestResolvePartialIDRefusesSentinelTokens(t *testing.T) {
	ctx := context.Background()

	refused := []string{
		"null", "NULL", "Null",
		"undefined", "UNDEFINED",
		"", "  ", " null ",
	}
	for _, input := range refused {
		if _, err := ResolvePartialID(ctx, nil, input); err == nil {
			t.Errorf("ResolvePartialID(%q) = nil error, want sentinel refusal", input)
		} else if strings.Contains(err.Error(), "storage is nil") {
			t.Errorf("ResolvePartialID(%q) reached the storage check; the sentinel guard did not fire: %v", input, err)
		}
	}

	// Non-sentinel inputs must pass the guard and proceed (here: hit the
	// nil-store check, proving the guard did not over-block them).
	passed := []string{
		"bd-a3f8e9",      // full ID
		"a3f8",           // partial hash
		"nullos",         // contains a sentinel as substring — not a bare token
		"ga-null",        // prefixed form of a hash containing the token
		"bd-null.1",      // hierarchical form
		"undefined-behavior", // prefixed-looking, not the bare token
	}
	for _, input := range passed {
		_, err := ResolvePartialID(ctx, nil, input)
		if err == nil || !strings.Contains(err.Error(), "storage is nil") {
			t.Errorf("ResolvePartialID(%q) = %v, want to pass the sentinel guard and hit the nil-store check", input, err)
		}
	}
}
