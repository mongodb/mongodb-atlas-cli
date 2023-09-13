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
	"fmt"
	"runtime"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type diagnosticsOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
	podmanClient podman.Client
	podmanDiag   *podman.Diagnostic
	machineDiag  *machineDiagnostic
	podmanLogs   []interface{}
	mongotLogs   []string
	mongodLogs   []string
}

type machineDiagnostic struct {
	OS   string
	Arch string
}

func (opts *diagnosticsOpts) Run(ctx context.Context) error {
	var err error

	// Machine system info
	opts.machineDiag = &machineDiagnostic{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}

	// Podman system info
	opts.podmanDiag = opts.podmanClient.Diagnostics(ctx)

	// Podman logs
	if opts.podmanLogs, err = opts.podmanClient.Logs(ctx); err != nil {
		opts.podmanDiag.Errors = append(opts.podmanDiag.Errors, fmt.Errorf("failed to get podman logs: %w", err).Error())
	}

	if opts.DeploymentName != "" && opts.podmanDiag.MachineInfo.State == podman.PodmanRunningState {
		_, _ = log.Warningf("Fetching logs for deployment %s\n", opts.DeploymentName)
		// ignore error if container does not exist just capture log for that command
		opts.mongotLogs, _ = opts.podmanClient.ContainerLogs(ctx, opts.LocalMongotHostname())
		opts.mongodLogs, _ = opts.podmanClient.ContainerLogs(ctx, opts.LocalMongodHostname())
	}

	diagnosis := map[string]interface{}{
		"Machine": opts.machineDiag,
		"Podman":  opts.podmanDiag,
		"Logs": map[string]interface{}{
			"Podman": opts.podmanLogs,
			"Mongot": opts.mongotLogs,
			"Mongod": opts.mongodLogs,
		},
	}

	return opts.Print(diagnosis)
}

func DiagnosticsBuilder() *cobra.Command {
	opts := &diagnosticsOpts{
		podmanDiag:  &podman.Diagnostic{},
		machineDiag: &machineDiagnostic{},
	}
	cmd := &cobra.Command{
		Use:    "diagnostics [deploymentName]",
		Short:  "Fetch detailed information about all your deployments and system processes.",
		Hidden: true, // always hidden
		Args:   require.MaximumNArgs(1),
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment you want to setup.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}

			w := cmd.OutOrStdout()
			opts.podmanClient = podman.NewClient(log.IsDebugLevel(), w)

			return opts.PreRunE(
				opts.InitOutput(w, ""),
				opts.InitStore(opts.podmanClient),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
