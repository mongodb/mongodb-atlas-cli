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
	"os"
	"os/exec"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
)

type StartOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
	store store.ClusterStarter
}

var (
	ErrDeploymentIsDeleting = errors.New("deployment state is DELETING")
	startTemplate           = "\nStarting deployment '{{.Name}}'.\n"
)

func (opts *StartOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *StartOpts) Run(ctx context.Context) error {
	deployment, err := opts.SelectDeployments(ctx, opts.ConfigProjectID())
	if err != nil {
		return err
	}

	if opts.IsLocalDeploymentType() {
		return opts.RunLocal(ctx, deployment)
	}

	return opts.RunAtlas()
}

func (opts *StartOpts) RunLocal(ctx context.Context, deployment options.Deployment) error {
	opts.StartSpinner()
	defer opts.StopSpinner()

	if err := opts.startContainer(ctx, deployment); err != nil {
		return err
	}

	return opts.Print(
		admin.AdvancedClusterDescription{
			Name: &opts.DeploymentName,
		})
}

func (opts *StartOpts) startContainer(ctx context.Context, deployment options.Deployment) error {
	if deployment.StateName == options.IdleState || deployment.StateName == options.RestartingState {
		return nil
	}

	if deployment.StateName == options.StoppedState {
		buf, err := options.ComposeDefinition(&options.ComposeDefinitionOptions{
			Name:          opts.DeploymentName,
			Port:          "27017",
			MongodVersion: "7.0",
			BindIp:        "127.0.0.1",
		})
		if err != nil {
			return err
		}
		cmd := exec.Command("docker", "compose", "-f", "/dev/stdin", "start")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = buf
		cmd.Env = append(os.Environ(), "KEY_FILE=keyfile")
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if deployment.StateName == options.PausedState {
		buf, err := options.ComposeDefinition(&options.ComposeDefinitionOptions{
			Name:          opts.DeploymentName,
			Port:          "27017",
			MongodVersion: "7.0",
			BindIp:        "127.0.0.1",
		})
		if err != nil {
			return err
		}
		cmd := exec.Command("docker", "compose", "-f", "/dev/stdin", "unpause")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = buf
		cmd.Env = append(os.Environ(), "KEY_FILE=keyfile")
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return ErrDeploymentIsDeleting
}

func (opts *StartOpts) RunAtlas() error {
	opts.StartSpinner()
	defer opts.StopSpinner()

	r, err := opts.store.StartCluster(opts.ConfigProjectID(), opts.DeploymentName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:     "start <deploymentName>",
		Short:   "Start a deployment.",
		Long:    "After you stop a machine, it goes into sleep mode, or restarts.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment.",
			"output":             startTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitOutput(cmd.OutOrStdout(), startTemplate))
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
