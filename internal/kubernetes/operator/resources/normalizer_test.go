//go:build unit

package resources

import (
	"regexp"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation"
)

func TestNormalizeResourceName(t *testing.T) {
	testCases := []struct {
		name           string
		expectedOutput string
	}{
		{
			name:           "test",
			expectedOutput: "test",
		},
		{
			name:           " .",
			expectedOutput: "dashdot",
		},
		{
			name:           "@",
			expectedOutput: "at",
		},
		{
			name:           "--test--project--",
			expectedOutput: "dash-test--project-dash",
		},
	}

	t.Run("Validate test cases", func(t *testing.T) {
		for _, tc := range testCases {
			if errs := validation.IsDNS1123Label(tc.expectedOutput); len(errs) > 0 {
				t.Errorf("output should be DNS-1123 compliant, got:%s. errors: %v", tc.expectedOutput, errs)
			}
		}
	})

	for _, tc := range testCases {
		got := NormalizeAtlasResourceName(tc.name)
		if got != tc.expectedOutput {
			t.Errorf("NormalizeAtlasResourceName() = %v, want %v", got, tc.expectedOutput)
		}
	}
}

func FuzzNormalizeResourceName(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		// Atlas project name can only contain letters, numbers, spaces, and the following symbols: ( ) @ & + : . _ - ' ,
		var atlasNameRegex = regexp.MustCompile(`[^a-zA-Z0-9.@()&+:,'\-]+`)
		input = atlasNameRegex.ReplaceAllString(input, "")
		if input != "" {
			got := NormalizeAtlasResourceName(input)
			if errs := validation.IsDNS1123Label(got); len(errs) > 0 {
				t.Errorf("output should be DNS-1123 compliant, got:%s. errors: %v", got, errs)
			}
		}
	})
}
