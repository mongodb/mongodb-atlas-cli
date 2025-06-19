package api

import (
	"sort"
	"testing"
)

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
