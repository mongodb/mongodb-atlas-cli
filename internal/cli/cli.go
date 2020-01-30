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
	"github.com/10gen/mcli/internal/config"
)

type globalOpts struct {
	orgID     string
	projectID string
}

// newGlobalOpts returns an globalOpts
func newGlobalOpts() *globalOpts {
	return new(globalOpts)
}

// ProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) ProjectID() string {
	if opts.projectID != "" {
		return opts.projectID
	}
	opts.projectID = config.ProjectID()
	return opts.projectID
}

// OrgID returns the organization id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) OrgID() string {
	if opts.orgID != "" {
		return opts.orgID
	}
	opts.orgID = config.OrgID()
	return opts.orgID
}
