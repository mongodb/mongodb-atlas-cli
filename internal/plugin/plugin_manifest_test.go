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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/set"
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
				Commands: map[string]ManifestCommand{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    true,
			wantErrCount: 0,
		},
		{
			name: "Valid manifest with alias",
			manifest: Manifest{
				Name:        "Kubernetes",
				Description: "Kubernetes plugin",
				Binary:      "kubernetes",
				Version:     "1.2.3",
				Commands: map[string]ManifestCommand{
					"kubernetes": {Description: "the kubernetes command", Aliases: []string{"k8s"}},
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
				Commands: map[string]ManifestCommand{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    false,
			wantErrCount: 1,
		},
		{
			name: "Invalid version and missing description",
			manifest: Manifest{
				Name:    "Kubernetes",
				Binary:  "kubernetes",
				Version: "version1.3",
				Commands: map[string]ManifestCommand{
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
				Name:    "Kubernetes",
				Version: "1.2.3",
				Commands: map[string]ManifestCommand{
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
				Commands: map[string]ManifestCommand{
					"kubernetes":  {},
					"kubernetes2": {},
				},
			},
			wantValid:    false,
			wantErrCount: 3,
		},
		{
			name: "Github defined but owner missing",
			manifest: Manifest{
				Name:        "Kubernetes",
				Description: "Kubernetes plugin",
				Binary:      "kubernetes",
				Version:     "1.2.3",
				Github:      &ManifestGithubValues{Name: "repo-name"},
				Commands: map[string]ManifestCommand{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    false,
			wantErrCount: 1,
		},
		{
			name: "Github defined but owner and name missing",
			manifest: Manifest{
				Name:        "Kubernetes",
				Description: "Kubernetes plugin",
				Binary:      "kubernetes",
				Version:     "1.2.3",
				Github:      &ManifestGithubValues{},
				Commands: map[string]ManifestCommand{
					"kubernetes": {Description: "the kubernetes command"},
				},
			},
			wantValid:    false,
			wantErrCount: 2,
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

func getExistingCommandsSet() set.Set[string] {
	existingCommandsSet := set.NewSet[string]()
	existingCommandsSet.Add("existingCmd1")
	existingCommandsSet.Add("existingCmd2")
	return existingCommandsSet
}

func Test_getUniqueManifests(t *testing.T) {
	existingCommandsSet := getExistingCommandsSet()

	manifests := []*Manifest{
		{
			Name:        "deployments",
			Description: "this is the description of the deployments plugin",
			Version:     "2.1.76",
			Binary:      "deployments",
			Commands: map[string]ManifestCommand{
				"deployments":  {Description: "this is the deployments command"},
				"deployments2": {Description: "this is the second description"},
				"command":      {Description: "here is another command"},
			},
		},
		{
			Name:        "Plugin2",
			Description: "Another test plugin",
			Binary:      "plugin2",
			Version:     "1.0.0",
			Commands: map[string]ManifestCommand{
				"existingCmd1": {Description: "A duplicate command"},
				"newCommand":   {Description: "This command does not exist yet"},
			},
		},
		{
			Name:        "kubernetes",
			Description: "this is the description of the kubernetes plugin",
			Version:     "1.2.3",
			Binary:      "binary",
			Commands: map[string]ManifestCommand{
				"kubernetes": {Description: "this is the kubernetes command"},
			},
		},
	}

	uniqueManifests, duplicateManifests := getUniqueManifests(manifests, existingCommandsSet)

	if len(uniqueManifests) != 2 {
		t.Errorf("expected 2 unique manifests, got %d", len(uniqueManifests))
	}

	if len(duplicateManifests) != 1 {
		t.Errorf("expected 1 duplicate manifest, got %d", len(duplicateManifests))
	}
}

func Test_removeManifestsWithDuplicateNames(t *testing.T) {
	tests := []struct {
		name                        string
		manifests                   []*Manifest
		expectedUniqueManifestCount int
		expectedDuplicateNamesCount int
	}{
		{
			name: "No duplicate manifests",
			manifests: []*Manifest{
				{
					Name: "manifestName1",
				},
				{
					Name: "manifestName2",
				},
				{
					Name: "manifestName3",
				},
			},
			expectedUniqueManifestCount: 3,
			expectedDuplicateNamesCount: 0,
		},
		{
			name: "Two duplicate manifests",
			manifests: []*Manifest{
				{
					Name: "manifestName1",
				},
				{
					Name: "manifestName2",
				},
				{
					Name: "manifestName2",
				},
			},
			expectedUniqueManifestCount: 1,
			expectedDuplicateNamesCount: 1,
		},
		{
			name: "Multiple uplicate manifests",
			manifests: []*Manifest{
				{
					Name: "manifestName1",
				},
				{
					Name: "manifestName2",
				},
				{
					Name: "manifestName2",
				},
				{
					Name: "manifestName3",
				},
				{
					Name: "manifestName4",
				},
				{
					Name: "manifestName5",
				},
				{
					Name: "manifestName5",
				},
				{
					Name: "manifestName5",
				},
				{
					Name: "manifestName5",
				},
			},
			expectedUniqueManifestCount: 3,
			expectedDuplicateNamesCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniqueManifests, duplicateNames := removeManifestsWithDuplicateNames(tt.manifests)
			if len(uniqueManifests) != tt.expectedUniqueManifestCount {
				t.Errorf("Expected %d unique manifests, got %d", tt.expectedUniqueManifestCount, len(uniqueManifests))
			}
			if len(duplicateNames) != tt.expectedDuplicateNamesCount {
				t.Errorf("Expected %d duplicate names, got %d", tt.expectedDuplicateNamesCount, len(duplicateNames))
			}
		})
	}
}

func Test_hasDuplicateCommand(t *testing.T) {
	existingCommandsSet := getExistingCommandsSet()

	tests := []struct {
		name           string
		manifest       *Manifest
		expectedResult bool
	}{
		{
			name: "Manifest without duplicate commands",
			manifest: &Manifest{
				Name:        "deployments",
				Description: "this is the description of the deployments plugin",
				Version:     "2.1.76",
				Binary:      "deployments",
				Commands: map[string]ManifestCommand{
					"deployments":  {Description: "this is the deployments command"},
					"deployments2": {Description: "this is the second description"},
					"command":      {Description: "here is another command"},
				},
			},
			expectedResult: false,
		},
		{
			name: "Manifest with duplicate commands",
			manifest: &Manifest{
				Name:        "kubernetes",
				Description: "this is the description of the kubernetes plugin",
				Version:     "1.2.3",
				Binary:      "binary",
				Commands: map[string]ManifestCommand{
					"kubernetes":   {Description: "this is the kubernetes command"},
					"existingCmd2": {Description: "this command already exsists"},
				},
			},
			expectedResult: true,
		},
		{
			name: "Manifest with duplicate alias",
			manifest: &Manifest{
				Name:        "deployments",
				Description: "this is the description of the deployments plugin",
				Version:     "2.1.76",
				Binary:      "deployments",
				Commands: map[string]ManifestCommand{
					"deployments":  {Description: "this is the deployments command"},
					"deployments2": {Description: "this is the second description", Aliases: []string{"existingCmd2"}},
					"command":      {Description: "here is another command"},
				},
			},
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.manifest.HasDuplicateCommand(existingCommandsSet)

			if result != tt.expectedResult {
				t.Errorf("expected result %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
