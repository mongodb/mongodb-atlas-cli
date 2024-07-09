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
	"fmt"
	"os"
	"reflect"
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
				Name:    "Kubernetes",
				Binary:  "kubernetes",
				Version: "1.3",
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
				Name:    "Kubernetes",
				Version: "1.2.3",
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
					"kubernetes":  {},
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

func generateTestPlugin(directoryName string, binaryName string, manifestContent string) error {
	directory := "./plugins/" + directoryName
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	// Write manifest.yml file
	manifestFile, err := os.Create(directory + "/manifest.yml")
	if err != nil {
		return fmt.Errorf("error creating manifest.yml: %v", err)
	}
	defer manifestFile.Close()

	_, err = manifestFile.WriteString(manifestContent)
	if err != nil {
		return fmt.Errorf("error writing to manifest.yml: %v", err)
	}

	// Create empty binary file
	binaryFile, err := os.Create(directory + "/" + binaryName)
	if err != nil {
		return fmt.Errorf("error creating binary file: %v", err)
	}
	defer binaryFile.Close()

	return nil
}

func deleteTestPlugins() error {
	err := os.RemoveAll("./plugins")
	if err != nil {
		return fmt.Errorf("error deleting plugin directory: %v", err)
	}

	return nil
}

func Test_getManifestsFromPluginDirectory(t *testing.T) { 
	generateTestPlugin("kubernetes","binary",`
name: kubernetes
description: this is the description of the kubernetes plugin
version: 1.2.3
commands:
	kubernetes:
		description: this is the kubernetes command`)
		generateTestPlugin("deployments", "deployments",`
name: deployments
description: this is the description of the deployments plugin
version: 2.1.76
commands:
	deployments:
		description: this is the deployments command
	deployments2:
		description: this is the second description
	command:
		description: here is another command`)

	defer deleteTestPlugins()

	tests := []struct {
		name               string
		pluginDirectoryName string
		expectedManifests  []*Manifest
		expectsError       bool
	}{
		{
			name:               "Valid plugin directory",
			pluginDirectoryName: "./plugins",
			expectedManifests: []*Manifest{
				{
					Name:        "kubernetes",
					Description: "this is the description of the kubernetes plugin",
					Version:     "1.2.3",
					Commands: map[string]struct {
						Description string `yaml:"description,omitempty"`
					}{
						"kubernetes": {Description: "this is the kubernetes command"},
					},
					Binary:     "binary",
					BinaryPath: "./plugins/kubernetes/binary",
				},
				{
					Name:        "deployments",
					Description: "this is the description of the deployments plugin",
					Version:     "2.1.76",
					Commands: map[string]struct {
						Description string `yaml:"description,omitempty"`
					}{
						"deployments": {Description: "this is the deployments command"},
						"deployments2": {Description: "this is the second description"},
						"command": {Description: "here is another command"},
					},
					Binary:     "deployments",
					BinaryPath: "./plugins/deployments/deployments",
				},
				},
			expectsError: false,
		},
		{
			name:               "Empty plugin directory",
			pluginDirectoryName: "./empty_plugins",
			expectedManifests:  []*Manifest{},
			expectsError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manifests, err := getManifestsFromPluginDirectory(tt.pluginDirectoryName)

			if tt.expectsError && err == nil {
				t.Errorf("Expected an error but got none")
			}

			if !tt.expectsError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !reflect.DeepEqual(manifests, tt.expectedManifests) {
				t.Errorf("Expected manifests: %+v, got: %+v", tt.expectedManifests, manifests)
			}
		})
	}
}