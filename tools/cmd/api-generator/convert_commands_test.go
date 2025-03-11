// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"
)

func TestExtractVersionAndContentType(t *testing.T) {
	tests := []struct {
		input           string
		wantVersion     string
		wantContentType string
	}{
		{
			"application/vnd.atlas.2025-01-01+json",
			"2025-01-01",
			"json",
		},
		{
			"application/vnd.atlas.2024-08-05+json",
			"2024-08-05",
			"json",
		},
		{
			"application/vnd.atlas.2023-01-01+csv",
			"2023-01-01",
			"csv",
		},
		{
			"application/vnd.atlas.preview+json",
			"preview",
			"json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotVersion, gotContentType, err := extractVersionAndContentType(tt.input)
			if err != nil {
				t.Fatalf("Error = %v", err)
			}
			if gotVersion != tt.wantVersion {
				t.Errorf("Expected: %s. Got: %s", tt.wantVersion, gotVersion)
			}
			if gotContentType != tt.wantContentType {
				t.Errorf("Expected: %s Got: %s,", tt.wantContentType, gotContentType)
			}
		})
	}
}
