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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/setup"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

type ListOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	options.DeploymentOpts
	defaultSetter cli.DefaultSetterOpts
	store         store.ClusterLister
	config        setup.ProfileReader
}

const listTemplate = `NAME	TYPE	MDB VER	STATE
{{range .}}{{.Name}}	{{.Type}}	{{.MongoDBVersion}}	{{.StateName}}
{{end}}`

const MaxItemsPerPage = 500
const errAtlas = "failed to retrieve Atlas deployments with: %s"

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) getAtlasDeployments() ([]options.Deployment, error) {
	if !opts.IsCliAuthenticated() {
		return nil, nil
	}

	if opts.ProjectID == "" {
		opts.ProjectID = opts.config.ProjectID()
	}

	if opts.ProjectID == "" {
		if err := opts.defaultSetter.AskProject(); err != nil {
			return nil, err
		}
		opts.ProjectID = opts.defaultSetter.ProjectID
	}

	listOpts := &mongodbatlas.ListOptions{
		PageNum:      cli.DefaultPage,
		ItemsPerPage: MaxItemsPerPage,
	}

	projectClusters, err := opts.store.ProjectClusters(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return nil, err
	}
	atlasClusters := projectClusters.(*admin.PaginatedAdvancedClusterDescription)

	deployments := make([]options.Deployment, len(atlasClusters.Results))
	for i, c := range atlasClusters.Results {
		stateName := *c.StateName
		if *c.Paused {
			// for paused clusters, Atlas returns stateName IDLE and Paused=true
			stateName = options.PausedState
		}
		deployments[i] = options.Deployment{
			Type:           "ATLAS",
			Name:           *c.Name,
			MongoDBVersion: *c.MongoDBVersion,
			StateName:      stateName,
		}
	}

	return deployments, nil
}

func (opts *ListOpts) Run(ctx context.Context) error {
	mdbContainers, err := opts.GetLocalDeployments(ctx)
	if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
		return err
	}

	atlasClusters, atlasErr := opts.getAtlasDeployments()

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
			opts.config = config.Default()
			opts.CredStore = config.Default()
			log.SetWriter(cmd.OutOrStdout())

			if err := opts.PreRunE(
				opts.initStore(cmd.Context()),
				func() error { return opts.defaultSetter.InitStore(cmd.Context()) },
				opts.InitOutput(log.Writer(), listTemplate)); err != nil {
				return err
			}

			opts.defaultSetter.OutWriter = cmd.OutOrStdout()

			opts.PodmanClient = podman.NewClient(log.IsDebugLevel(), log.Writer())
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
