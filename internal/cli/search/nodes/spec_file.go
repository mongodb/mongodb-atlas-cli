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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530004/admin"
)

// Load *atlasv2.ApiSearchDeploymentRequest from a given file.
func loadAPISearchDeploymentSpec(fs afero.Fs, filename string) (*atlasv2.ApiSearchDeploymentRequest, error) {
	spec := new(atlasv2.ApiSearchDeploymentRequest)
	if err := file.Load(fs, filename, spec); err != nil {
		return nil, fmt.Errorf("failed to parse JSON file due to: %w", err)
	}

	return spec, nil
}
