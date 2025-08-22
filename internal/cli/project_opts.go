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

package cli

import (
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

type ProjectOpts struct {
	PreRunOpts
	ProjectID string
}

// ConfigProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *ProjectOpts) ConfigProjectID() string {
	if opts.ProjectID != "" {
		return opts.ProjectID
	}
	opts.ProjectID = config.ProjectID()
	return opts.ProjectID
}

// ValidateProjectID validates projectID.
func (opts *ProjectOpts) ValidateProjectID() error {
	if opts.ConfigProjectID() == "" {
		return errMissingProjectID
	}
	return validate.ObjectID(opts.ConfigProjectID())
}

func (opts *ProjectOpts) AddProjectOptsFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
}
