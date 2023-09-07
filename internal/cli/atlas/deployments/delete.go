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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	*cli.DeleteOpts
	options.DeploymentOpts
	podmanClient podman.Client
}

func (opts *DeleteOpts) Run(ctx context.Context) (err error) {
	if err := opts.podmanClient.Ready(ctx); err != nil {
		return err
	}

	if opts.DeploymentName != "" {
		if err := opts.DeploymentOpts.CheckIfDeploymentExists(ctx); err != nil {
			return err
		}
	} else {
		if err = opts.DeploymentOpts.Select(ctx); err != nil {
			return err
		}
	}

	opts.Entry = opts.DeploymentName
	if err := opts.Prompt(); err != nil {
		return err
	}

	return opts.Delete(func() error {
		log.Warningln("deleting deployment...")
		opts.StartSpinner()
		defer opts.StopSpinner()
		return opts.DeploymentOpts.Remove(ctx)
	})
}

// atlas deployments delete <clusterName>.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Deployment '%s' deleted", "Deployment not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <clusterName>",
		Short: "Delete a local deployment.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster you want to delete.",
			"output":          opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}
			w := cmd.OutOrStdout()
			opts.podmanClient = podman.NewClient(log.IsDebugLevel(), w)
			return opts.PreRunE(opts.InitOutput(w, ""), opts.InitStore(opts.podmanClient))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
