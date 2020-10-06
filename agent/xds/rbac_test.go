package xds

import (
	"fmt"
	"path/filepath"
	"sort"
	"testing"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeRBACNetworkFilter(t *testing.T) {
	testIntention := func(t *testing.T, src, dst string, action structs.IntentionAction) *structs.Intention {
		t.Helper()
		ixn := structs.TestIntention(t)
		ixn.SourceName = src
		ixn.DestinationName = dst
		ixn.Action = action
		//nolint:staticcheck
		ixn.UpdatePrecedence()
		return ixn
	}
	testSourceIntention := func(src string, action structs.IntentionAction) *structs.Intention {
		return testIntention(t, src, "api", action)
	}
	sorted := func(ixns ...*structs.Intention) structs.Intentions {
		sort.SliceStable(ixns, func(i, j int) bool {
			return ixns[j].Precedence < ixns[i].Precedence
		})
		return structs.Intentions(ixns)
	}

	tests := map[string]struct {
		intentionDefaultAllow bool
		intentions            structs.Intentions
	}{
		"default-deny-mixed-precedence": {
			intentionDefaultAllow: false,
			intentions: sorted(
				testIntention(t, "web", "api", structs.IntentionActionAllow),
				testIntention(t, "*", "api", structs.IntentionActionDeny),
				testIntention(t, "web", "*", structs.IntentionActionDeny),
			),
		},
		"default-deny-service-wildcard-allow": {
			intentionDefaultAllow: false,
			intentions: sorted(
				testSourceIntention("*", structs.IntentionActionAllow),
			),
		},
		"default-allow-service-wildcard-deny": {
			intentionDefaultAllow: true,
			intentions: sorted(
				testSourceIntention("*", structs.IntentionActionDeny),
			),
		},
		"default-deny-one-allow": {
			intentionDefaultAllow: false,
			intentions: sorted(
				testSourceIntention("web", structs.IntentionActionAllow),
			),
		},
		"default-allow-one-deny": {
			intentionDefaultAllow: true,
			intentions: sorted(
				testSourceIntention("web", structs.IntentionActionDeny),
			),
		},
		"default-deny-allow-deny": {
			intentionDefaultAllow: false,
			intentions: sorted(
				testSourceIntention("web", structs.IntentionActionDeny),
				testSourceIntention("*", structs.IntentionActionAllow),
			),
		},
		"default-deny-kitchen-sink": {
			intentionDefaultAllow: false,
			intentions: sorted(
				// (double exact)
				testSourceIntention("web", structs.IntentionActionAllow),
				testSourceIntention("unsafe", structs.IntentionActionDeny),
				testSourceIntention("cron", structs.IntentionActionAllow),
				// and we invert the default-ness of the whole thing
				testSourceIntention("*", structs.IntentionActionAllow),
			),
		},
		"default-allow-kitchen-sink": {
			intentionDefaultAllow: true,
			intentions: sorted(
				// (double exact)
				testSourceIntention("web", structs.IntentionActionDeny),
				testSourceIntention("unsafe", structs.IntentionActionAllow),
				testSourceIntention("cron", structs.IntentionActionDeny),
				// and we invert the default-ness of the whole thing
				testSourceIntention("*", structs.IntentionActionDeny),
			),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			filter, err := makeRBACNetworkFilter(tt.intentions, tt.intentionDefaultAllow)
			require.NoError(t, err)

			gotJSON := protoToJSON(t, filter)

			require.JSONEq(t, golden(t, filepath.Join("rbac", name), "", gotJSON), gotJSON)
		})
	}
}

func TestRemoveSameSourceIntentions(t *testing.T) {
	testIntention := func(t *testing.T, src, dst string) *structs.Intention {
		t.Helper()
		ixn := structs.TestIntention(t)
		ixn.SourceName = src
		ixn.DestinationName = dst
		//nolint:staticcheck
		ixn.UpdatePrecedence()
		return ixn
	}
	sorted := func(ixns ...*structs.Intention) structs.Intentions {
		sort.SliceStable(ixns, func(i, j int) bool {
			return ixns[j].Precedence < ixns[i].Precedence
		})
		return structs.Intentions(ixns)
	}
	tests := map[string]struct {
		in     structs.Intentions
		expect structs.Intentions
	}{
		"empty": {},
		"one": {
			in: sorted(
				testIntention(t, "*", "*"),
			),
			expect: sorted(
				testIntention(t, "*", "*"),
			),
		},
		"two with no match": {
			in: sorted(
				testIntention(t, "*", "foo"),
				testIntention(t, "bar", "*"),
			),
			expect: sorted(
				testIntention(t, "*", "foo"),
				testIntention(t, "bar", "*"),
			),
		},
		"two with match, exact": {
			in: sorted(
				testIntention(t, "bar", "foo"),
				testIntention(t, "bar", "*"),
			),
			expect: sorted(
				testIntention(t, "bar", "foo"),
			),
		},
		"two with match, wildcard": {
			in: sorted(
				testIntention(t, "*", "foo"),
				testIntention(t, "*", "*"),
			),
			expect: sorted(
				testIntention(t, "*", "foo"),
			),
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := removeSameSourceIntentions(tc.in)
			require.Equal(t, tc.expect, got)
		})
	}
}

func TestSimplifyNotSourceSlice(t *testing.T) {
	tests := map[string]struct {
		in     []string
		expect []string
	}{
		"empty": {},
		"one": {
			[]string{"bar"},
			[]string{"bar"},
		},
		"two with no match": {
			[]string{"foo", "bar"},
			[]string{"foo", "bar"},
		},
		"two with match": {
			[]string{"*", "bar"},
			[]string{"*"},
		},
		"three with two matches down to one": {
			[]string{"*", "foo", "bar"},
			[]string{"*"},
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := simplifyNotSourceSlice(makeServiceNameSlice(tc.in))
			require.Equal(t, makeServiceNameSlice(tc.expect), got)
		})
	}
}

func TestIxnSourceMatches(t *testing.T) {
	tests := []struct {
		tester, against string
		matches         bool
	}{
		// identical precedence
		{"web", "api", false},
		{"*", "*", false},
		// backwards precedence
		{"*", "web", false},
		// name wildcards
		{"web", "*", true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s cmp %s", tc.tester, tc.against), func(t *testing.T) {
			matches := ixnSourceMatches(
				structs.ServiceNameFromString(tc.tester),
				structs.ServiceNameFromString(tc.against),
			)
			assert.Equal(t, tc.matches, matches)
		})
	}
}

func makeServiceNameSlice(slice []string) []structs.ServiceName {
	if len(slice) == 0 {
		return nil
	}
	var out []structs.ServiceName
	for _, src := range slice {
		out = append(out, structs.ServiceNameFromString(src))
	}
	return out
}

func unmakeServiceNameSlice(slice []structs.ServiceName) []string {
	if len(slice) == 0 {
		return nil
	}
	var out []string
	for _, src := range slice {
		out = append(out, src.String())
	}
	return out
}
