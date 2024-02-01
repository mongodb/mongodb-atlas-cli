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

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/deployments/options"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/podman"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
}

const listTemplate = `NAME	TYPE	MDB VER	STATE
{{range .}}{{.Name}}	{{.Type}}	{{.MongoDBVersion}}	{{.StateName}}
{{end}}`

const errAtlas = "failed to retrieve Atlas deployments with: %s"

func (opts *ListOpts) Run(ctx context.Context) error {
	if err := opts.LocalDeploymentPreRun(ctx); err != nil {
		return err
	}

	var mdbContainers []options.Deployment
	if opts.IsLocalDeploymentType() || opts.NoDeploymentTypeSet() {
		var err error
		mdbContainers, err = opts.GetLocalDeployments(ctx)
		if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
			return err
		}
	}
	var atlasClusters []options.Deployment
	var atlasErr error
	if opts.IsAtlasDeploymentType() || opts.NoDeploymentTypeSet() {
		if opts.IsCliAuthenticated() && cli.TokenRefreshed {
			atlasClusters, atlasErr = opts.AtlasDeployments(opts.ProjectID)
		}
	}
	if err := opts.Print(append(atlasClusters, mdbContainers...)); err != nil {
		return err
	}
	if atlasErr != nil {
		return fmt.Errorf(errAtlas, atlasErr.Error())
	}

	return nil
}

func (opts *ListOpts) PostRun() error {
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitStore(cmd.Context(), cmd.OutOrStdout()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PostRun()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.DeploymentType, flag.TypeFlag, "", usage.DeploymentType)

	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return options.DeploymentTypeOptions, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.Flags().MarkHidden(flag.TypeFlag)

	return cmd
}
