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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

func Run(ctx context.Context, opts *options.ConnectOpts) error {
	return opts.Connect(ctx)
}

// atlas deployments connect [clusterName].
func ConnectBuilder() *cobra.Command {
	opts := &options.ConnectOpts{}
	cmd := &cobra.Command{
		Use:     "connect [deploymentName]",
		Short:   "Connect to a deployment.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment that you want to connect to.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), ""),
				opts.InitInput(cmd.InOrStdin()),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitAtlasStore(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.ConnectWith, flag.ConnectWith, "", usage.ConnectWith)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)
	cmd.Flags().StringVar(&opts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.ConnectionStringType, flag.ConnectionStringType, options.ConnectionStringTypeStandard, usage.ConnectionStringType)

	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return options.ConnectWithOptions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return options.DeploymentTypeOptions, cobra.ShellCompDirectiveDefault
	})

	return cmd
}
