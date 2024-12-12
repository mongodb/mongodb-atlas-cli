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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

type OrgOpts struct {
	PreRunOpts
	OrgID string
}

// ConfigOrgID returns the organization id.
// If the id is empty, it caches it after querying config.
func (opts *OrgOpts) ConfigOrgID() string {
	if opts.OrgID != "" {
		return opts.OrgID
	}
	opts.OrgID = config.OrgID()
	return opts.OrgID
}

// ValidateOrgID validates orgID.
func (opts *OrgOpts) ValidateOrgID() error {
	if opts.ConfigOrgID() == "" {
		return ErrMissingOrgID
	}
	return validate.ObjectID(opts.ConfigOrgID())
}

func (opts *OrgOpts) AddOrgOptFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
}
