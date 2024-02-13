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

package nodes

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

type SpecFile []SpecFileEntry

type SpecFileEntry struct {
	InstanceSize string
	NodeCount    int
}

// Load []atlasv2.ApiSearchDeploymentSpec from a given file.
func LoadAPISearchDeploymentSpec(fs afero.Fs, filename string) ([]atlasv2.ApiSearchDeploymentSpec, error) {
	spec := new(SpecFile)
	if err := file.Load(fs, filename, spec); err != nil {
		return nil, fmt.Errorf("failed to parse JSON file due to: %w", err)
	}

	return asAPISearchDeploymentSpec(*spec), nil
}

// Convert SpecFile into []atlasv2.ApiSearchDeploymentSpec.
func asAPISearchDeploymentSpec(s SpecFile) []atlasv2.ApiSearchDeploymentSpec {
	out := make([]atlasv2.ApiSearchDeploymentSpec, len(s))

	for i, entry := range s {
		out[i] = atlasv2.ApiSearchDeploymentSpec{
			InstanceSize: entry.InstanceSize,
			NodeCount:    entry.NodeCount,
		}
	}

	return out
}
