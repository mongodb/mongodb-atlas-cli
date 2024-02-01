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

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/store"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type FailoverOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.ClusterTester
}

func (opts *FailoverOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *FailoverOpts) Run() error {
	return opts.Delete(opts.store.TestClusterFailover, opts.ConfigProjectID())
}

// FailoverBuilder
//
// atlas cluster(s) failover <clusterName> --projectId projectId [--force].
func FailoverBuilder() *cobra.Command {
	opts := &FailoverOpts{
		DeleteOpts: cli.NewDeleteOpts("Failover test for '%s' started\n", "Failover test not started"),
	}
	cmd := &cobra.Command{
		Use:   "failover <clusterName>",
		Short: "Starts a failover test for the specified cluster in the specified project.",
		Long:  `Clusters contain a group of hosts that maintain the same data set. A failover test checks how MongoDB Cloud handles the failure of the cluster's primary node. During the test, MongoDB Cloud shuts down the primary node and elects a new primary.`,
		Example: fmt.Sprintf(`  # Test failover for a cluster named myCluster:
  %s clusters failover myCluster`, cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Human-readable label that identifies the cluster to start a failover test for.",
			"output":          opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.PromptWithMessage("Are you sure you want to start a failover test for %q")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
