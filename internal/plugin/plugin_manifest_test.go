// Copyright 2024 MongoDB Inc
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

package plugin

import (
	"testing"
)

func TestManifest_IsValid(t *testing.T) {
	tests := []struct {
		name         string
		manifest     Manifest
		wantValid    bool
		wantErrCount int
	}{
		{
			name: "Valid manifest",
			manifest: Manifest{
				Name:        "Kubernetes",
				Description: "Kubernetes plugin",
				Binary:      "kubernetes",
				Version:     "1.2.3",
				Commands: map[string]struct {
					Description string `yaml:"description,omitempty"`
				}{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    true,
			wantErrCount: 0,
		},
		{
			name: "Missing name",
			manifest: Manifest{
				Description: "Kubernetes plugin",
				Binary:      "kubernetes",
				Version:     "1.2.3",
				Commands: map[string]struct {
					Description string `yaml:"description,omitempty"`
				}{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    false,
			wantErrCount: 1,
		},
		{
			name: "Invalid version and missing description",
			manifest: Manifest{
				Name:        "Kubernetes",
				Binary:      "kubernetes",
				Version:     "1.3",
				Commands: map[string]struct {
					Description string `yaml:"description,omitempty"`
				}{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    false,
			wantErrCount: 2,
		},
		{
			name: "Missing commands and missing version",
			manifest: Manifest{
				Name:        "MyApp",
				Description: "An example application",
				Binary:      "myapp",
			},
			wantValid:    false,
			wantErrCount: 2,
		},
		{
			name: "Command without description and missing binary",
			manifest: Manifest{
				Name:        "Kubernetes",
				Version:     "1.2.3",
				Commands: map[string]struct {
					Description string `yaml:"description,omitempty"`
				}{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    false,
			wantErrCount: 2,
		},
		{
			name: "Two command descriptions missing and missing name",
			manifest: Manifest{
				Description: "Kubernetes plugin",
				Binary:      "kubernetes",
				Version:     "1.2.3",
				Commands: map[string]struct {
					Description string `yaml:"description,omitempty"`
				}{
					"kubernetes": {},
					"kubernetes2": {},
				},
			},
			wantValid:    false,
			wantErrCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotErrs := tt.manifest.IsValid()

			if gotValid != tt.wantValid {
				t.Errorf("IsValid() gotValid = %v, want %v", gotValid, tt.wantValid)
			}

			if len(gotErrs) != tt.wantErrCount {
				t.Errorf("IsValid() got %d errors, want %d errors", len(gotErrs), tt.wantErrCount)
			}
		})
	}
}	
