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

package api

import "testing"

func TestSplitShortAndLongDescription(t *testing.T) {
	tests := []struct {
		name        string
		description string
		wantShort   string
		wantLong    string
	}{
		{
			name:        "empty string",
			description: "",
			wantShort:   "",
			wantLong:    "",
		},
		{
			name:        "single sentence",
			description: "This is a single sentence.",
			wantShort:   "This is a single sentence.",
			wantLong:    "",
		},
		{
			name:        "two sentences",
			description: "First sentence. Second sentence.",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence.",
		},
		{
			name:        "multiple sentences with spaces",
			description: "First sentence.   Second sentence.    Third sentence.",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence.    Third sentence.",
		},
		{
			name:        "multiple sentences with spaces and new lines",
			description: "First sentence.   Second sentence.    Third sentence.  \n Forth sentence.   ",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence.    Third sentence.  \n Forth sentence.",
		},
		{
			name:        "sentence without period",
			description: "This is a sentence without period",
			wantShort:   "This is a sentence without period.",
			wantLong:    "",
		},
		{
			name:        "multiple sentences with extra periods",
			description: "This is version 1.2.3. Second sentence. Third sentence.",
			wantShort:   "This is version 1.2.3.",
			wantLong:    "Second sentence. Third sentence.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShort, gotLong := splitShortAndLongDescription(tt.description)
			if gotShort != tt.wantShort {
				t.Errorf("splitShortAndLongDescription() gotShort = %v, want %v", gotShort, tt.wantShort)
			}
			if gotLong != tt.wantLong {
				t.Errorf("splitShortAndLongDescription() gotLong = %v, want %v", gotLong, tt.wantLong)
			}
		})
	}
}
