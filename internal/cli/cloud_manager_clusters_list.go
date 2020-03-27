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
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type cloudManagerClustersListOpts struct {
	*globalOpts
	store store.CloudManagerClustersLister
}

func (opts *cloudManagerClustersListOpts) init() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *cloudManagerClustersListOpts) Run() error {
	result, err := cloudManagerClustersListRun(opts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func cloudManagerClustersListRun(opts *cloudManagerClustersListOpts) (interface{}, error) {
	var result interface{}
	var err error

	if opts.projectID == "" && config.Service() == config.OpsManagerService {
		result, err = opts.store.ListAllProjectClusters()

	} else {
		var clusterConfigs *om.AutomationConfig
		clusterConfigs, err = opts.store.GetAutomationConfig(opts.ProjectID())
		result = convert.FromAutomationConfig(clusterConfigs)
	}
	return result, err
}

// mongocli cloud-manager cluster(s) list --projectId projectId
func CloudManagerClustersListBuilder() *cobra.Command {
	opts := &cloudManagerClustersListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   description.ListClusters,
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
