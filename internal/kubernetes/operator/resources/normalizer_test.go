// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package resources

import (
	"regexp"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation"
)

func TestNormalizeAtlasName(t *testing.T) {
	d := AtlasNameToKubernetesName()
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
		got := NormalizeAtlasName(tc.name, d)
		if got != tc.expectedOutput {
			t.Errorf("NormalizeAtlasName() = %v, want %v", got, tc.expectedOutput)
		}
	}
}

func FuzzNormalizeAtlasName(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		d := AtlasNameToKubernetesName()
		// Atlas project name can only contain letters, numbers, spaces, and the following symbols: ( ) @ & + : . _ - ' ,
		var atlasNameRegex = regexp.MustCompile(`[^a-zA-Z0-9_.@()&+:,'\-]+`)
		input = atlasNameRegex.ReplaceAllString(input, "")
		if input != "" {
			got := NormalizeAtlasName(input, d)
			if errs := validation.IsDNS1123Label(got); len(errs) > 0 {
				t.Errorf("output should be DNS-1123 compliant, got:%s. errors: %v", got, errs)
			}
		}
	})
}
