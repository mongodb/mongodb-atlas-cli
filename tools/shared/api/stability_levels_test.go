package api

import (
	"sort"
	"testing"
)

func TestVersionDateLess(t *testing.T) {
	tests := []struct {
		a, b VersionDate
		want bool
	}{
		// Test year
		{VersionDate{2024, 1, 1}, VersionDate{2024, 1, 1}, false},
		{VersionDate{2024, 1, 1}, VersionDate{2025, 1, 1}, true},
		{VersionDate{2024, 1, 1}, VersionDate{2023, 1, 1}, false},
		// Test month
		{VersionDate{2024, 1, 1}, VersionDate{2024, 2, 1}, true},
		{VersionDate{2024, 2, 1}, VersionDate{2024, 2, 1}, false},
		{VersionDate{2024, 2, 1}, VersionDate{2024, 1, 1}, false},
		// Test day
		{VersionDate{2024, 1, 1}, VersionDate{2024, 1, 2}, true},
		{VersionDate{2024, 1, 2}, VersionDate{2024, 1, 2}, false},
		{VersionDate{2024, 1, 2}, VersionDate{2024, 1, 1}, false},
	}

	for _, test := range tests {
		got := test.a.Less(&test.b)
		if got != test.want {
			t.Errorf("VersionDate.Less(%s, %s) = %t, want %t", test.a, test.b, got, test.want)
		}
	}
}
func TestStabilityLevelSorting(t *testing.T) {
	inputs := []Version{
		NewStableVersion(2025, 1, 1),
		NewPreviewVersion(),
		NewUpcomingVersion(2025, 1, 1),
		NewStableVersion(2024, 1, 1),
		NewUpcomingVersion(2024, 1, 1),
		NewStableVersion(2024, 2, 1),
		NewStableVersion(2024, 2, 3),
		NewStableVersion(2024, 2, 3),
		NewUpcomingVersion(2024, 2, 1),
	}

	// The order of the versions is:
	// 2024-01-01
	// 2024-01-01.upcoming
	// 2024-02-01
	// 2024-02-01.upcoming
	// 2024-02-03
	// 2024-02-03
	// 2025-01-01
	// 2025-01-01.upcoming
	// preview
	want := []Version{
		NewStableVersion(2024, 1, 1),
		NewUpcomingVersion(2024, 1, 1),
		NewStableVersion(2024, 2, 1),
		NewUpcomingVersion(2024, 2, 1),
		NewStableVersion(2024, 2, 3),
		NewStableVersion(2024, 2, 3),
		NewStableVersion(2025, 1, 1),
		NewUpcomingVersion(2025, 1, 1),
		NewPreviewVersion(),
	}

	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i].Less(inputs[j])
	})

	// Check that the inputs are sorted in the correct order
	for i := range inputs {
		if !inputs[i].Equal(want[i]) {
			t.Errorf("got %s, want %s", inputs[i], want[i])
		}
	}
}
