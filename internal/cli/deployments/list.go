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
	"os"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
}

const listTemplate = `NAME	TYPE	MDB VER	STATE
{{range valueOrEmptySlice .}}{{.Name}}	{{.Type}}	{{.MongoDBVersion}}	{{.StateName}}
{{end}}`

const errAtlas = "failed to retrieve Atlas deployments with: %w"
const errLocal = "failed to retrieve local deployments with: %w"

func (opts *ListOpts) Run(ctx context.Context) error {
	localDeployments, localErr := opts.runLocal(ctx)
	if localErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, localErr)
	}

	atlasClusters, atlasErr := opts.runAtlas()
	if atlasErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, atlasErr)
	}

	if localErr == nil && atlasErr == nil {
		return opts.Print(append(atlasClusters, localDeployments...))
	}

	if atlasErr != nil && localErr == nil {
		return opts.Print(localDeployments)
	}

	if localErr != nil && atlasErr == nil {
		return opts.Print(atlasClusters)
	}

	return nil
}

func (opts *ListOpts) runLocal(ctx context.Context) ([]options.Deployment, error) {
	if !opts.IsLocalDeploymentType() && !opts.NoDeploymentTypeSet() {
		return nil, nil
	}

	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return nil, fmt.Errorf(errLocal, err)
	}

	mdbContainers, err := opts.GetLocalDeployments(ctx)
	if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
		return nil, fmt.Errorf(errLocal, err)
	}
	return mdbContainers, nil
}

func (opts *ListOpts) runAtlas() ([]options.Deployment, error) {
	if !opts.IsAtlasDeploymentType() && !opts.NoDeploymentTypeSet() {
		return nil, nil
	}

	if !opts.IsCliAuthenticated() {
		return nil, nil
	}

	atlasClusters, err := opts.AtlasDeployments(opts.ProjectID)
	if err != nil {
		return nil, fmt.Errorf(errAtlas, err)
	}

	return atlasClusters, nil
}

func (opts *ListOpts) PostRun() error {
	opts.UpdateDeploymentTelemetry()
	return opts.PostRunMessages()
}

// ListBuilder atlas deployments list.
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all deployments.",
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		GroupID: "all",
		Annotations: map[string]string{
			"output": listTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		PostRunE: func(_ *cobra.Command, _ []string) error {
			return opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)

	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return options.DeploymentTypeOptions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.Flags().MarkHidden(flag.TypeFlag)

	return cmd
}
