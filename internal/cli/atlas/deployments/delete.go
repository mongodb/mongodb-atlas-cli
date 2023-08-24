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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var errDeploymentNotFound = errors.New("deployment '%s' not found")

type DeleteOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	*cli.DeleteOpts
	options.DeploymentOpts
	podmanClient podman.Client
	debug        bool
}

func (opts *DeleteOpts) checkIfDeploymentExists(podman.Client) error {
	// search locally
	containers, errList := opts.podmanClient.ListContainers(options.MongodHostnamePrefix)
	if errList != nil {
		return errList
	}

	found := false
	for _, c := range containers {
		for _, n := range c.Names {
			if n == opts.LocalMongodHostname() {
				found = true
			}
		}
	}

	if !found {
		return fmt.Errorf(errDeploymentNotFound.Error(), opts.DeploymentName)
	}

	return nil
}

func (opts *DeleteOpts) Run(_ context.Context) error {
	if !opts.Confirm {
		fmt.Println(opts.FailMessage())
		return nil
	}

	if err := opts.checkIfDeploymentExists(opts.podmanClient); err != nil {
		return err
	}

	// remove mongod
	if _, err := opts.podmanClient.RemoveContainers(opts.LocalMongodHostname()); err != nil {
		return err
	}

	// remove mongot
	if _, err := opts.podmanClient.RemoveContainers(opts.LocalMongotHostname()); err != nil {
		return err
	}

	// delete network
	_, err := opts.podmanClient.RemoveNetworks(opts.LocalNetworkName())
	if err != nil {
		return err
	}

	// delete volumes
	if _, err := opts.podmanClient.RemoveVolumes(opts.LocalMongodDataVolume()); err != nil {
		return err
	}

	if _, err := opts.podmanClient.RemoveVolumes(opts.LocalMongotDataVolume()); err != nil {
		return err
	}

	if _, err := opts.podmanClient.RemoveVolumes(opts.LocalMongoMetricsVolume()); err != nil {
		return err
	}

	return nil
}

// atlas deployments delete <clusterName>.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Deployment '%s' deleted", "Deployment not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <clusterName>",
		Short: "Delete a local deployment.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster you want to setup.",
			"output":          opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.InitOutput(cmd.OutOrStdout(), "")); err != nil {
				return err
			}

			opts.podmanClient = podman.NewClient(opts.debug, opts.OutWriter)
			if err := opts.podmanClient.Ready(); err != nil {
				return err
			}

			opts.DeploymentName = args[0]
			opts.Entry = opts.DeploymentName
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)

	return cmd
}
