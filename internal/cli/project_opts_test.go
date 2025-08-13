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

package cli

import (
	"testing"
)

func TestProjectOpts_ValidateProjectID(t *testing.T) {
	tests := []struct {
		name      string
		projectID string
		wantErr   bool
	}{
		{
			name:      "empty project ID",
			projectID: "",
			wantErr:   true,
		},
		{
			name:      "invalid project ID",
			projectID: "1",
			wantErr:   true,
		},
		{
			name:      "valid project ID",
			projectID: "5e98249d937cfc52efdc2a9f",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		projectID := tt.projectID
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			opts := &ProjectOpts{
				ProjectID: projectID,
			}
			if err := opts.ValidateProjectID(); (err != nil) != wantErr {
				t.Errorf("ValidateProjectID() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}
