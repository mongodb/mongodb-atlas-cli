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
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
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

	mdbContainers, err := opts.GetLocalDeployments(ctx)
	if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
		return err
	}

	var atlasClusters []options.Deployment
	var atlasErr error
	if opts.IsCliAuthenticated() {
		atlasClusters, atlasErr = opts.GetAtlasDeployments(opts.ProjectID)
	}

	err = opts.Print(append(atlasClusters, mdbContainers...))
	if err != nil {
		return err
	}

	if atlasErr != nil {
		return fmt.Errorf(errAtlas, atlasErr.Error())
	}

	return nil
}

func (opts *ListOpts) PostRun(_ context.Context) error {
	return opts.PostRunMessages()
}

// atlas deployments list.
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
			opts.PodmanClient = podman.NewClient(log.IsDebugLevel(), log.Writer())

			if err := opts.PreRunE(
				func() error { return opts.DefaultSetter.InitStore(cmd.Context()) },
				opts.InitStore(cmd.Context(), opts.PodmanClient),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate)); err != nil {
				return err
			}

			opts.DefaultSetter.OutWriter = cmd.OutOrStdout()
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PostRun(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
