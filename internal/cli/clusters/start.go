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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=start_mock_test.go -package=clusters . ClusterStarter

type ClusterStarter interface {
	StartCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	StartClusterLatest(string, string) (*atlasv2.ClusterDescription20240805, error)
}

type StartOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name            string
	autoScalingMode string
	store           ClusterStarter
}

func (opts *StartOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var startTmpl = "Starting cluster '{{.Name}}'.\n"

func (opts *StartOpts) Run() error {
	if isIndependentShardScaling(opts.autoScalingMode) {
		r, err := opts.store.StartClusterLatest(opts.ConfigProjectID(), opts.name)
		if err != nil {
			return err
		}
		return opts.Print(r)
	}

	r, err := opts.store.StartCluster(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas cluster(s) start <clusterName> [--projectId projectId].
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:   "start <clusterName>",
		Short: "Start the specified paused MongoDB cluster.",
		Long:  fmt.Sprintf("%s\n%s", fmt.Sprintf(usage.RequiredRole, "Project Cluster Manager"), "Atlas supports this command only for M10+ clusters."),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster to start.",
			"output":          startTmpl,
		},
		Example: `  # Start a cluster named myCluster for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas clusters start myCluster --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), startTmpl),
				validate.AutoScalingMode(opts.autoScalingMode),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.autoScalingMode, flag.AutoScalingMode, clusterWideScalingFlag, usage.AutoScalingMode)
	_ = cmd.RegisterFlagCompletionFunc(flag.AutoScalingMode, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{clusterWideScalingFlag, independentShardScalingFlag}, cobra.ShellCompDirectiveDefault
	})

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
