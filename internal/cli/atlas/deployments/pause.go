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
	"os/exec"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/setup"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115006/admin"
)

type PauseOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
	store  store.ClusterPauser
	config setup.ProfileReader
}

const (
	pauseTemplate                     = "Pausing deployment '{{.Name}}'.\n"
	podmanContainerTerminatedExitCode = 137
)

var (
	errDeploymentIsNotIDLE = errors.New("deployment state is not IDLE")
)

func (opts *PauseOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *PauseOpts) Run(ctx context.Context) error {
	deployment, err := opts.SelectDeployments(ctx, opts.ConfigProjectID())
	if err != nil {
		return err
	}

	if opts.IsLocalDeploymentType() {
		return opts.RunLocal(ctx, deployment)
	}

	return opts.RunAtlas()
}

func (opts *PauseOpts) RunLocal(ctx context.Context, deployment options.Deployment) error {
	if err := opts.pauseContainer(ctx, deployment); err != nil {
		return err
	}

	return opts.Print(
		admin.AdvancedClusterDescription{
			Name: &opts.DeploymentName,
		})
}

func (opts *PauseOpts) pauseContainer(ctx context.Context, deployment options.Deployment) error {
	if deployment.StateName == options.PausedState || deployment.StateName == options.StoppedState {
		return nil
	}

	if deployment.StateName != options.IdleState {
		return errDeploymentIsNotIDLE
	}
	opts.StartSpinner()
	defer opts.StopSpinner()

	// search for a process "java ... mongot", extract the process ID, kill the process with SIGTERM (15)
	const stopMongot = "grep -l java.*mongot /proc/*/cmdline | awk -F'/' '{print $3; exit}' | while read -r pid; do kill -15 $pid; done"
	if err := opts.PodmanClient.Exec(ctx, opts.LocalMongotHostname(), "/bin/sh", "-c", stopMongot); err != nil {
		return err
	}

	err := opts.PodmanClient.Exec(ctx, opts.LocalMongodHostname(), "mongod", "--shutdown")
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) && exitErr.ExitCode() == podmanContainerTerminatedExitCode {
		return nil
	}
	return err
}

func (opts *PauseOpts) RunAtlas() error {
	opts.StartSpinner()
	defer opts.StopSpinner()

	r, err := opts.store.PauseCluster(opts.ConfigProjectID(), opts.DeploymentName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *PauseOpts) StopMongoD(ctx context.Context, names string) error {
	return opts.PodmanClient.Exec(ctx, "-d", names, "mongod", "--shutdown")
}

func PauseBuilder() *cobra.Command {
	opts := &PauseOpts{}
	cmd := &cobra.Command{
		Use:     "pause <deploymentName>",
		Aliases: []string{"stop"},
		Short:   "Pause a deployment.",
		Args:    require.MaximumNArgs(1),
		GroupID: "all",
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment.",
			"output":             pauseTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.config = config.Default()
			opts.CredStore = config.Default()

			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitOutput(log.Writer(), pauseTemplate))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			return opts.Run(cmd.Context())
		},

		PostRunE: func(_ *cobra.Command, _ []string) error {
			return opts.PostRunMessages()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)

	return cmd
}
