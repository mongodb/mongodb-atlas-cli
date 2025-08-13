// Copyright 2020 MongoDB Inc
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

package convert

import (
	"testing"
	"time"
)

func Test_ParseTimestamp(t *testing.T) {
	testCases := map[string]struct {
		timestamp string
		want      time.Time
		wantErr   bool
	}{
		"Valid timestamps in RFC3339 format": {
			timestamp: "2023-06-17T00:00:00Z",
			want:      time.Date(2023, 6, 17, 0, 0, 0, 0, time.UTC),
			wantErr:   false,
		},
		"Valid timestamp with custom layout": {
			timestamp: "2023-06-17T00:00:00-0000",
			want:      time.Date(2023, 6, 17, 0, 0, 0, 0, time.UTC),
			wantErr:   false,
		},
		"Invalid timestamp": {
			timestamp: "2023-06-15T18:25:55",
			want:      time.Time{},
			wantErr:   true,
		},
	}

	for name, tc := range testCases {
		input := tc.timestamp
		expected := tc.want
		wantErr := tc.wantErr
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ts, err := ParseTimestamp(input)
			if (err != nil) != wantErr {
				t.Fatalf("parseTimestamp() unexpected error: %v\n", err)
			}
			if !ts.Equal(expected) {
				t.Errorf("parseTimestamp() expected: %s but got: %s", expected, ts)
			}
		})
	}
}
