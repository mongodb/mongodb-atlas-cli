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
	"runtime"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
	"github.com/spf13/cobra"
)

type diagnosticsOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
}

type diagnosticLogs struct {
	MongoD []string
	MongoT []string
}
type diagnostic struct {
	Machine    machineDiagnostic
	Diagnostic *podman.Diagnostic
	Containers []*podman.InspectContainerData
	Logs       diagnosticLogs
	Errors     []error
}
type machineDiagnostic struct {
	OS   string
	Arch string
}

func (opts *diagnosticsOpts) Run(ctx context.Context) error {
	d := &diagnostic{
		Machine: machineDiagnostic{
			OS:   runtime.GOOS,
			Arch: runtime.GOARCH,
		},
		Diagnostic: opts.PodmanClient.Diagnostics(ctx),
	}

	var err error
	d.Containers, err = opts.PodmanClient.ContainerInspect(ctx, opts.LocalMongodHostname(), opts.LocalMongotHostname())
	if err != nil {
		d.Errors = append(d.Errors, err)
	}

	if d.Logs.MongoT, err = opts.PodmanClient.ContainerLogs(ctx, opts.LocalMongotHostname()); err != nil {
		d.Errors = append(d.Errors, err)
	}
	if d.Logs.MongoD, err = opts.PodmanClient.ContainerLogs(ctx, opts.LocalMongodHostname()); err != nil {
		d.Errors = append(d.Errors, err)
	}

	return opts.Print(d)
}

func DiagnosticsBuilder() *cobra.Command {
	opts := &diagnosticsOpts{
		OutputOpts: cli.OutputOpts{
			Output: "json",
		},
	}
	cmd := &cobra.Command{
		Use:     "diagnostics <deploymentName>",
		Short:   "Fetch detailed information about all your deployments and system processes.",
		Hidden:  true, // always hidden
		Aliases: []string{"diagnostic", "diag", "diags"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"deploymentNameDesc": "Name of the deployment you want to setup.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.DeploymentName = args[0]
			}

			w := cmd.OutOrStdout()

			return opts.PreRunE(
				opts.InitOutput(w, ""),
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
	}

	return cmd
}
