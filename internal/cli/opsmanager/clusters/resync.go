// Copyright 2023 MongoDB Inc
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
	"context"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/atmcfg"
)

type resyncOpts struct {
	cli.GlobalOpts
	confirm     bool
	clusterName string
	timestamp   string
	processes   []string
	store       store.AutomationPatcher
}

func (opts *resyncOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *resyncOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	if err = atmcfg.StartInitialSyncAtForProcessesByClusterName(current, opts.clusterName, opts.timestamp, opts.processes); err != nil {
		return err
	}

	if err = opts.store.UpdateAutomationConfig(opts.ConfigProjectID(), current); err != nil {
		return err
	}

	fmt.Print(cli.DeploymentStatus(config.OpsManagerURL(), opts.ConfigProjectID()))

	return nil
}

func (opts *resyncOpts) Confirm() error {
	if opts.confirm {
		return nil
	}

	process := opts.clusterName

	if len(opts.processes) > 0 {
		process = fmt.Sprintf("%s (%s)", opts.clusterName, strings.Join(opts.processes, ", "))
	}

	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to reclaim free space for: %s", process),
	}
	return survey.AskOne(prompt, &opts.confirm)
}

// ResyncBuilder mongocli cloud-manager cluster(s) resync <clusterName> [--processName process1,process2...][--timestamp timestamp] [--force].
func ResyncBuilder() *cobra.Command {
	opts := &resyncOpts{}
	cmd := &cobra.Command{
		Use:   "resync <clusterName>",
		Short: "Start an initial sync for a cluster or process.",
		Long: `The MongoDB Agent checks whether the specified timestamp is later than the time of the last resync, and if confirmed, the MongoDB Agent:

1. Starts the initial sync on the secondary nodes in a rolling fashion
2. Waits until you ask the primary node to become the secondary with the rs.stepDown() method 
3. Starts the initial sync on the primary node

Warning: Use this method with caution. During initial sync, Automation removes the entire contents of the node’s dbPath directory.

To learn more, see: https://docs.mongodb.com/manual/tutorial/resync-replica-set-member/`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster for which you want to start a resync.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.clusterName = args[0]
			return opts.Confirm()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.processes, flag.ProcessName, []string{}, usage.ProcessName)
	cmd.Flags().BoolVar(&opts.confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.timestamp, flag.Timestamp, "", usage.ResyncTimestamp)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
