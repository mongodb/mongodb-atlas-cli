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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=autoscalingconfig_mock_test.go -package=clusters . ClusterAutoscalingConfigGetter

type ClusterAutoscalingConfigGetter interface {
	GetClusterAutoScalingConfig(string, string) (*atlasv2.ClusterDescriptionAutoScalingModeConfiguration, error)
}

type GetAutoscalingConfigOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store ClusterAutoscalingConfigGetter
	name  string
}

func (opts *GetAutoscalingConfigOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GetAutoscalingConfigOpts) Run() error {
	r, err := opts.store.GetClusterAutoScalingConfig(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// GetAutoscalingConfigBuilder
//
// atlas cluster(s) autoScalingConfig <clusterName> --projectId projectId.
func GetAutoscalingConfigBuilder() *cobra.Command {
	opts := &GetAutoscalingConfigOpts{}
	cmd := &cobra.Command{
		Use:   "autoScalingConfig <clusterName>",
		Short: "Get the autoscaling config for the specified cluster.",

		Example: `  # Get the autoscaling config for a cluster named myCluster:
  atlas clusters autoScalingConfig myCluster`,
		Args:   require.ExactArgs(1),
		Hidden: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), "Autoscaling config for cluster is {{.AutoScalingMode}}\n"),
			); err != nil {
				return err
			}
			opts.name = args[0]
			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
