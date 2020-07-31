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
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/validate"
)

type GlobalOpts struct {
	OrgID     string
	ProjectID string
}

// ConfigProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *GlobalOpts) ConfigProjectID() string {
	if opts.ProjectID != "" {
		return opts.ProjectID
	}
	opts.ProjectID = config.ProjectID()
	return opts.ProjectID
}

// ConfigOrgID returns the organization id.
// If the id is empty, it caches it after querying config.
func (opts *GlobalOpts) ConfigOrgID() string {
	if opts.OrgID != "" {
		return opts.OrgID
	}
	opts.OrgID = config.OrgID()
	return opts.OrgID
}

type cmdOpt func() error

// PreRunE is a function to call before running the command,
// this will validate the project ID and call any additional function pass as a callback
func (opts *GlobalOpts) PreRunE(cbs ...cmdOpt) error {
	if opts.ConfigProjectID() == "" {
		return errMissingProjectID
	}
	if err := validate.ObjectID(opts.ConfigProjectID()); err != nil {
		return err
	}
	for _, f := range cbs {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

// PreRunEOrg is a function to call before running the command,
// this will validate the org ID and call any additional function pass as a callback
func (opts *GlobalOpts) PreRunEOrg(cbs ...cmdOpt) error {
	if opts.ConfigOrgID() == "" && opts.OrgID == "" {
		return ErrMissingOrgID
	}
	if err := validate.ObjectID(opts.ConfigOrgID()); err != nil {
		return err
	}
	for _, f := range cbs {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func DeploymentStatus(baseURL, projectID string) string {
	return fmt.Sprintf("Changes are being applied, please check %sv2/%s#deployment/topology for status\n", baseURL, projectID)
}
