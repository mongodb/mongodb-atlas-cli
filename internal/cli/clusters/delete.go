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
	"context"
	"fmt"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/watchers"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=delete_mock_test.go -package=clusters . ClusterDeleter

type ClusterDeleter interface {
	DeleteCluster(string, string) error
	DeleteFlexCluster(string, string) error
}

type DeleteOpts struct {
	cli.ProjectOpts
	cli.WatchOpts
	*cli.DeleteOpts
	store         ClusterDeleter
	isFlexCluster bool
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	if err := opts.Delete(opts.store.DeleteCluster, opts.ConfigProjectID()); err != nil {
		return opts.RunFlexCluster(err)
	}

	opts.isFlexCluster = false
	return nil
}

func (opts *DeleteOpts) RunFlexCluster(err error) error {
	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return err
	}

	if apiError.ErrorCode != cannotUseFlexWithClusterApisErrorCode {
		return err
	}

	opts.isFlexCluster = true
	return opts.Delete(opts.store.DeleteFlexCluster, opts.ConfigProjectID())
}

func (opts *DeleteOpts) PostRun() error {
	if !opts.EnableWatch {
		return nil
	}

	if opts.isFlexCluster {
		return opts.PostRunFlexCluster()
	}

	watcher := watchers.NewWatcher(
		*watchers.ClusterDeleted,
		watchers.NewAtlasClusterStateDescriber(
			opts.store.(store.ClusterDescriber),
			opts.ConfigProjectID(),
			opts.Entry,
		),
	)

	watcher.Timeout = time.Duration(opts.Timeout)
	if err := opts.WatchWatcher(watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

func (opts *DeleteOpts) PostRunFlexCluster() error {
	watcher := watchers.NewWatcherWithDefaultWait(
		*watchers.ClusterDeleted,
		watchers.NewAtlasFlexClusterStateDescriber(
			opts.store.(store.ClusterDescriber),
			opts.ConfigProjectID(),
			opts.Entry,
		),
		opts.GetDefaultWait(),
	)

	watcher.Timeout = time.Duration(opts.Timeout)
	if err := opts.WatchWatcher(watcher); err != nil {
		return err
	}

	return opts.Print(flexCluster)
}

// DeleteBuilder
//
// atlas cluster(s) delete <clusterName> --projectId projectId [--confirm].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Deleting cluster '%s'", "Cluster not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <clusterName>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified cluster from your project.",
		Long: `The command prompts you to confirm the operation when you run the command without the --force option. 
		
Deleting a cluster also deletes any backup snapshots for that cluster.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # Remove a cluster named myCluster after prompting for a confirmation:
  atlas clusters delete myCluster
  
  # Remove a cluster named myCluster without requiring confirmation:
  atlas clusters delete myCluster --force`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to delete.",
			"output":          opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), "Cluster deleted\n")); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.PromptWithMessage("This operation will delete the cluster and all of its data. Confirm your backup settings before terminating your cluster. This action cannot be undone.\nAre you sure you want to terminate %s?")
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		PostRunE: func(_ *cobra.Command, _ []string) error {
			return opts.PostRun()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatch)
	cmd.Flags().Int64Var(&opts.Timeout, flag.WatchTimeout, 0, usage.WatchTimeout)

	opts.AddProjectOptsFlags(cmd)

	return cmd
}
