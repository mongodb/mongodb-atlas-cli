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
	"runtime"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/spf13/cobra"
)

type diagnosticsOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
}

type diagnostic struct {
	Machine    machineDiagnostic
	Engine     string
	Version    map[string]any
	Images     []string
	Containers []*containerDiagnostic
	Err        string
}

type machineDiagnostic struct {
	OS     string
	Arch   string
	Memory uint64
	CPU    int
}

type containerDiagnostic struct {
	Inspect *container.InspectData
	Logs    []string
}

func (opts *diagnosticsOpts) Run(ctx context.Context) error {
	d := &diagnostic{
		Machine: machineDiagnostic{
			OS:   runtime.GOOS,
			Arch: runtime.GOARCH,
		},
		Engine: opts.ContainerEngine.Name(),
	}

	errs := []error{}

	cores, err := cpu.Counts(true)
	if err != nil {
		errs = append(errs, err)
		cores = -1
	}
	d.Machine.CPU = cores

	v, err := mem.VirtualMemory()
	if err != nil {
		errs = append(errs, err)
		d.Machine.Memory = 0
	} else {
		d.Machine.Memory = v.Available
	}

	images, err := opts.ContainerEngine.ImageList(ctx, "mongodb/mongodb-atlas-local")
	if err != nil {
		errs = append(errs, err)
	} else {
		for _, image := range images {
			d.Images = append(d.Images, image.Repository+":"+image.Tag)
		}
	}

	d.Version, err = opts.ContainerEngine.Version(ctx)
	if err != nil {
		errs = append(errs, err)
	}

	if opts.LocalMongodHostname() != "" {
		inspectData, err := opts.ContainerEngine.ContainerInspect(ctx, opts.LocalMongodHostname())
		if err != nil {
			errs = append(errs, err)
		}

		logs, err := opts.ContainerEngine.ContainerLogs(ctx, opts.LocalMongodHostname())
		if err != nil {
			errs = append(errs, err)
		}

		d.Containers = append(d.Containers, &containerDiagnostic{
			Inspect: firstOrNil(inspectData),
			Logs:    logs,
		})
	}

	d.Err = errors.Join(errs...).Error()

	return opts.Print(d)
}

func firstOrNil[T any](slice []*T) *T {
	if len(slice) == 0 {
		return nil
	}

	return slice[0]
}

func DiagnosticsBuilder() *cobra.Command {
	opts := &diagnosticsOpts{
		OutputOpts: cli.OutputOpts{
			Output: "json",
		},
	}
	cmd := &cobra.Command{
		Use:     "diagnostics [<deploymentName>]",
		Short:   "Fetch detailed information about all your deployments and system processes.",
		Hidden:  true, // always hidden
		Aliases: []string{"diagnostic", "diag", "diags", "inspect"},
		Args:    require.MaximumNArgs(1),
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
