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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/compass"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type ConnectOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
	podmanClient podman.Client
	connectWith  string
}

func (opts *ConnectOpts) validateFlags() error {
	if opts.DeploymentName != "" {
		if err := options.ValidateDeploymentName(opts.DeploymentName); err != nil {
			return err
		}
	}

	if opts.connectWith != "" {
		if err := options.ValidateConnectWith(opts.connectWith); err != nil {
			return err
		}
	}

	return nil
}

func (opts *ConnectOpts) Run(ctx context.Context) error {
	if err := opts.podmanClient.Ready(ctx); err != nil {
		return err
	}

	if opts.DeploymentName != "" {
		if err := opts.DeploymentOpts.CheckIfDeploymentExists(ctx); err != nil {
			return err
		}
	} else {
		if err := opts.DeploymentOpts.Select(ctx); err != nil {
			return err
		}
	}

	if opts.connectWith == "" {
		var err error
		if opts.connectWith, err = opts.DeploymentOpts.PromptConnectWith(); err != nil {
			return err
		}
	}

	connectionString, err := opts.ConnectionString(ctx)
	if err != nil {
		return err
	}

	switch opts.connectWith {
	case options.ConnectWithConnectionString:
		opts.Print(connectionString)
	case compassConnect:
		if !compass.Detect() {
			return errCompassNotInstalled
		}
		if _, err := log.Warningln("Launching MongoDB Compass..."); err != nil {
			return err
		}
		return compass.Run("", "", connectionString)
	case mongoshConnect:
		if !mongosh.Detect() {
			return errMongoshNotInstalled
		}
		return mongosh.Run("", "", connectionString)
	}

	return nil
}

// atlas deployments connect [clusterName].
func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}
	cmd := &cobra.Command{
		Use:   "connect [deploymentName]",
		Short: "Connect to a deployment.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment that you want to connect to.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			opts.podmanClient = podman.NewClient(log.IsDebugLevel(), log.Writer())
			return opts.PreRunE(opts.InitOutput(cmd.OutOrStdout(), ""), opts.InitStore(opts.podmanClient), opts.validateFlags)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.connectWith, flag.ConnectWith, "", usage.ConnectWith)

	_ = cmd.RegisterFlagCompletionFunc(flag.ConnectWith, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return options.ConnectWithOptions, cobra.ShellCompDirectiveDefault
	})

	return cmd
}
