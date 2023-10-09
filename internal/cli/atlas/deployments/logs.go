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
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DownloadOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
}

func (opts *DownloadOpts) Run(ctx context.Context) error {
	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	if err := opts.DetectLocalDeploymentName(ctx); err != nil {
		return err
	}

	logs, err := opts.PodmanClient.ContainerLogs(ctx, opts.LocalMongodHostname())
	if err != nil {
		return err
	}

	// format log entries into lines
	if opts.IsJSONOutput() {
		return opts.Print(logs)
	}
	return opts.Print(strings.Join(logs, "\n"))
}

// atlas deployments logs.
func LogsBuilder() *cobra.Command {
	opts := &DownloadOpts{}
	cmd := &cobra.Command{
		Use:     "logs",
		Short:   "Get deployments logs.",
		Aliases: []string{"log"},
		Args:    require.NoArgs,
		GroupID: "local",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			log.SetWriter(w)

			return opts.PreRunE(
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitOutput(w, ""))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVar(&opts.DeploymentName, flag.DeploymentName, "", usage.DeploymentName)

	return cmd
}
