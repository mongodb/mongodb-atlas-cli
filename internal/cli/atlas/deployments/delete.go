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

package deployments

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/internal/watchers"
	"github.com/spf13/cobra"
)

const (
	deleteSuccessMessage = "Deployment '%s' deleted\n"
	deleteFailMessage    = "Deployment not deleted"
)

type DeleteOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	*cli.DeleteOpts
	cli.WatchOpts
	options.DeploymentOpts
	atlasStore store.ClusterDeleter
}

func (opts *DeleteOpts) initAtlasStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.atlasStore, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run(ctx context.Context) error {
	if err := opts.validateAndPrompt(ctx); err != nil {
		return err
	}
	opts.Entry = opts.DeploymentName

	if err := opts.Prompt(); err != nil {
		return err
	}

	if opts.IsAtlasDeploymentType() {
		return opts.runAtlas()
	}
	return opts.runLocal(ctx)
}

func (opts *DeleteOpts) validateAndPrompt(ctx context.Context) error {
	if err := opts.ValidateAndPromptDeploymentType(); err != nil {
		return err
	}
	telemetry.AppendOption(telemetry.WithDeploymentType(options.LocalCluster))

	if opts.IsAtlasDeploymentType() {
		return opts.validateAndPromptAtlas()
	}
	return opts.validateAndPromptLocal(ctx)
}

func (opts *DeleteOpts) validateAndPromptAtlas() error {
	if opts.DeploymentName == "" {
		return ErrNoDeploymentName
	}

	return opts.ValidateProjectID()
}

func (opts *DeleteOpts) validateAndPromptLocal(ctx context.Context) error {
	if err := opts.PodmanClient.Ready(ctx); err != nil {
		return err
	}

	if opts.DeploymentName == "" {
		if err := opts.DeploymentOpts.SelectLocal(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (opts *DeleteOpts) runAtlas() error {
	return opts.Delete(opts.atlasStore.DeleteCluster, opts.ConfigProjectID())
}

func (opts *DeleteOpts) runLocal(ctx context.Context) error {
	err := opts.LocalDeploymentPreRun(ctx)
	if err != nil {
		return err
	}
	return opts.Delete(func() error {
		_, _ = log.Warningln("deleting deployment...")
		opts.StartSpinner()
		defer opts.StopSpinner()
		return opts.DeploymentOpts.Remove(ctx)
	})
}

func (opts *DeleteOpts) PostRun() error {
	if !opts.EnableWatch || !opts.IsAtlasDeploymentType() {
		return nil
	}

	watcher := watchers.NewWatcher(
		*watchers.ClusterDeleted,
		watchers.NewAtlasClusterStateDescriber(
			opts.atlasStore.(store.AtlasClusterDescriber),
			opts.ProjectID,
			opts.Entry,
		),
	)

	watcher.Timeout = time.Duration(opts.Timeout)
	if err := opts.WatchWatcher(watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// atlas deployments delete <clusterName>.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts(deleteSuccessMessage, deleteFailMessage),
	}
	cmd := &cobra.Command{
		Use:   "delete [deploymentName]",
		Short: "Delete a deployment.",
		Long: `The command prompts you to confirm the operation when you run the command without the --force option. 
		
Deleting an Atlas deployment also deletes any backup snapshots for that cluster.
Deleting a Local deployment also deletes any local data volumes.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  # Remove an Atlas deployment named myDeployment after prompting for a confirmation:
  %[1]s deployments delete myDeployment --type ATLAS
  
  # Remove an Atlas deployment named myDeployment without requiring confirmation:
  %[1]s deployments delete myDeployment --type ATLAS --force

  # Remove an Local deployment named myDeployment without requiring confirmation:
  %[1]s deployments delete myDeployment --type LOCAL --force`, cli.ExampleAtlasEntryPoint()),
		Aliases: []string{"rm"},
		GroupID: "all",
		Args:    require.MaximumNArgs(1),
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment that you want to delete.",
			"output":             opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			return opts.PreRunE(
				opts.initAtlasStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatch)
	cmd.Flags().UintVar(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
