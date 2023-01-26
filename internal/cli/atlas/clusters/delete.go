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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.ClusterDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteCluster, opts.ConfigProjectID())
}

// DeleteBuilder
//
// mongocli atlas cluster(s) delete <clusterName> --projectId projectId [--confirm].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Cluster '%s' deleted\n", "Cluster not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <clusterName>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified cluster from your project.",
		Long: `The command prompts you to confirm the operation when you run the command without the --force option. 
		
Deleting a cluster also deletes any backup snapshots for that cluster.`,
		Example: fmt.Sprintf(`  # Remove a cluster named myCluster after prompting for a confirmation:
  %[1]s clusters delete myCluster
  
  # Remove a cluster named myCluster without requiring confirmation:
  %[1]s clusters delete myCluster --force`, cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to delete.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
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
