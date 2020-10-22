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

package clusters

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/automation"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/atmcfg"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	automation.WatchOpts
	store store.CloudManagerClustersDeleter
}

func (opts *DeleteOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DeleteOpts) Run() error {
	if !opts.Confirm {
		return nil
	}

	// shutdown cluster
	err := opts.shutdownCluster()
	if err != nil {
		return err
	}

	// Remove cluster from automation
	err = opts.removeClusterFromAutomation()
	if err != nil {
		return err
	}

	// Stop monitoring
	err = opts.Delete(opts.store.StopMonitoring, opts.ConfigProjectID())
	if err != nil {
		return err
	}

	fmt.Sprint("Cluster deleted")
	return nil
}

func (opts *DeleteOpts) removeClusterFromAutomation() error {
	current, err := opts.store.GetAutomationConfig(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	atmcfg.RemoveByClusterName(current, opts.Entry)
	if err := opts.store.UpdateAutomationConfig(opts.ConfigProjectID(), current); err != nil {
		return err
	}

	// Wait for changes being deployed on automation
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}
	return nil
}

func (opts *DeleteOpts) shutdownCluster() error {
	current, err := opts.store.GetAutomationConfig(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	if !search.ClusterExists(current, opts.Entry) {
		return fmt.Errorf("cluster '%s' doesn't exist", opts.Entry)
	}

	// Shutdown Cluster
	atmcfg.Shutdown(current, opts.Entry)
	if err := opts.store.UpdateAutomationConfig(opts.ConfigProjectID(), current); err != nil {
		return err
	}

	// Wait for changes being deployed on automation
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return nil
}

func (opts *DeleteOpts) watcher() (bool, error) {
	result, err := opts.store.GetAutomationStatus(opts.ConfigProjectID())
	if err != nil {
		return false, err
	}

	for _, p := range result.Processes {
		if p.LastGoalVersionAchieved != result.GoalVersion {
			return false, nil
		}
	}
	return true, nil
}

// mongocli cloud-manager cluster(s) unmanage <name> --projectId projectId [--force]
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("", "Cluster not deleted\""),
	}
	cmd := &cobra.Command{
		Use:     "unmanage <name>",
		Aliases: []string{"rm", "delete"},
		Short:   DeleteCluster,
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
