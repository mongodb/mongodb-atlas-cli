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
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type clustersListOpts struct {
	*globalOpts
	storeCM store.AutomationGetter
	storeOM store.ListAllClusters
}

func (opts *clustersListOpts) init() error {
	var err error

	if opts.ProjectID() == "" {
		opts.storeOM, err = store.New()
	} else {
		opts.storeCM, err = store.New()
	}

	return err
}

func (opts *clustersListOpts) RunCM() error {
	result, err := opts.storeCM.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	clusterConfigs := convert.FromAutomationConfig(result)

	return json.PrettyPrint(clusterConfigs)
}

func (opts *clustersListOpts) RunOM() error {
	result, err := opts.storeOM.ListAllClustersProjects()
	if err != nil {
		return err
	}
	return json.PrettyPrint(result)
}

// mongocli cloud-manager cluster(s) list --projectId projectId
func CloudManagerClustersListBuilder() *cobra.Command {
	opts := &clustersListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List Cloud Manager clusters.",
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.storeCM != nil {
				return opts.RunCM()
			}
			return opts.RunOM()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
