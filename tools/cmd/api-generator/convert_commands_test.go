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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

func TestExtractVersionAndContentType(t *testing.T) {
	tests := []struct {
		input           string
		wantVersion     api.Version
		wantContentType string
	}{
		{
			"application/vnd.atlas.2025-01-01+json",
			api.NewStableVersion(2025, 1, 1),
			"json",
		},
		{
			"application/vnd.atlas.2024-08-05+json",
			api.NewStableVersion(2024, 8, 5),
			"json",
		},
		{
			"application/vnd.atlas.2023-01-01+csv",
			api.NewStableVersion(2023, 1, 1),
			"csv",
		},
		{
			"application/vnd.atlas.preview+json",
			api.NewPreviewVersion(),
			"json",
		},
		{
			"application/vnd.atlas.preview+csv",
			api.NewPreviewVersion(),
			"csv",
		},
		{
			"application/vnd.atlas.2024-08-05.upcoming+json",
			api.NewUpcomingVersion(2024, 8, 5),
			"json",
		},
		{
			"application/vnd.atlas.2023-01-01.upcoming+csv",
			api.NewUpcomingVersion(2023, 1, 1),
			"csv",
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
