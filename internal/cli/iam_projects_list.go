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
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type iamProjectsListOpts struct {
	*globalOpts
	store store.ProjectLister
}

func (opts *iamProjectsListOpts) init() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *iamProjectsListOpts) Run() error {
	var projects interface{}
	var err error
	if opts.OrgID() != "" && config.Service() == config.OpsManagerService {
		projects, err = opts.store.GetOrgProjects(opts.OrgID())
	} else {
		projects, err = opts.store.GetAllProjects()
	}
	if err != nil {
		return err
	}
	return json.PrettyPrint(projects)
}

// mongocli iam project(s) list [--orgId orgId]
func IAMProjectsListBuilder() *cobra.Command {
	opts := &iamProjectsListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.orgID, flags.OrgID, "", usage.OrgID)

	return cmd
}
