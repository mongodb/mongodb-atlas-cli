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
	"errors"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/setup"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type StartOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
	store  store.ClusterStarter
	config setup.ProfileReader
}

var (
	ErrDeploymentIsDeleting = errors.New("deployment state is DELETING")
	ErrNoDeploymentName     = errors.New("deployment name is required for Atlas resources")
	ErrNotAuthenticated     = errors.New("you are not authenticated. Please, run atlas auth  login")
	startTemplate           = "Starting deployment '{{.Name}}'.\n"
)

func (opts *StartOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *StartOpts) Run(ctx context.Context) error {
	if err := opts.validateAndPrompt(ctx); err != nil {
		return err
	}

	if strings.EqualFold(opts.DeploymentType, options.LocalCluster) {
		return opts.RunLocal(ctx)
	}

	return opts.RunAtlas()
}

func (opts *StartOpts) RunLocal(ctx context.Context) error {
	localDeployments, err := opts.GetLocalDeployments(ctx)
	if err != nil {
		return err
	}

	for _, deployment := range localDeployments {
		if deployment.Name == opts.DeploymentName {
			if err = opts.startContainer(ctx, deployment); err != nil {
				return err
			}

			return opts.Print(
				admin.AdvancedClusterDescription{
					Name: &opts.DeploymentName,
				})
		}
	}

	return options.ErrDeploymentNotFound
}

func (opts *StartOpts) startContainer(ctx context.Context, deployment options.Deployment) error {
	if deployment.StateName == options.IdleState || deployment.StateName == options.RestartingState {
		return nil
	}

	if deployment.StateName == options.StoppedState {
		if _, err := opts.PodmanClient.StartContainers(ctx, opts.LocalMongodHostname(), opts.LocalMongotHostname()); err != nil {
			return err
		}

		return nil
	}

	if deployment.StateName == options.PausedState {
		if _, err := opts.PodmanClient.UnpauseContainers(ctx, opts.LocalMongodHostname(), opts.LocalMongotHostname()); err != nil {
			return err
		}

		return nil
	}

	return ErrDeploymentIsDeleting
}

func (opts *StartOpts) RunAtlas() error {
	if !opts.IsCliAuthenticated() {
		return ErrNotAuthenticated
	}
	r, err := opts.store.StartCluster(opts.ConfigProjectID(), opts.DeploymentName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *StartOpts) validateAndPrompt(ctx context.Context) error {
	if opts.DeploymentType == "" {
		if err := opts.PromptDeploymentType(); err != nil {
			return err
		}
	}

	if opts.DeploymentType == options.AtlasCluster && opts.DeploymentName == "" {
		return ErrNoDeploymentName
	}

	if opts.DeploymentName == "" {
		if err := opts.DeploymentOpts.Select(ctx); err != nil {
			return err
		}
	}

	return nil
}

func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:     "start <deploymentName>",
		Short:   "Start a deployment",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment.",
			"output":             startTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.config = config.Default()
			opts.CredStore = config.Default()
			log.SetWriter(cmd.OutOrStdout())

			if err := opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(log.Writer(), startTemplate)); err != nil {
				return err
			}

			opts.PodmanClient = podman.NewClient(log.IsDebugLevel(), log.Writer())
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			return opts.Run(cmd.Context())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PostRunMessages()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)

	return cmd
}
