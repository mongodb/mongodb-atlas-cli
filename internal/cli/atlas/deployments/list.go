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
	"sort"
	"strings"

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

type Deployment struct {
	Type           string
	Name           string
	MongoDBVersion string
	StateName      string
}

const listTemplate = `NAME	TYPE	MDB VER	STATE
{{range .}}{{.Name}}	{{.Type}}	{{.MongoDBVersion}}	{{.StateName}}
{{end}}`

const MaxItemsPerPage = 500

var localStateMap = map[string]string{
	"running":  "IDLE",
	"removing": "DELETING",
	// a "created" container is ready to be started but is currently stopped,
	// which for a local deployment is equivalent to being paused.
	"created":    "PAUSED",
	"paused":     "PAUSED",
	"restarting": "RESTARTING",
	"exited":     "STOPPED",
	"dead":       "STOPPED",
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) getAtlasDeployments() ([]Deployment, error) {
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

	deployments := make([]Deployment, len(atlasClusters.Results))
	for i, c := range atlasClusters.Results {
		deployments[i] = Deployment{
			Type:           "ATLAS",
			Name:           *c.Name,
			MongoDBVersion: *c.MongoDBVersion,
			StateName:      *c.StateName,
		}
	}

	return deployments, nil
}

func (opts *ListOpts) getLocalDeployments(ctx context.Context) ([]Deployment, error) {
	if err := opts.PodmanClient.Ready(ctx); err != nil {
		return nil, err
	}

	mdbContainers, err := opts.PodmanClient.ListContainers(ctx, options.MongodHostnamePrefix)
	if err != nil {
		return nil, err
	}
	sort.Slice(mdbContainers, func(i, j int) bool {
		return mdbContainers[i].Names[0] < mdbContainers[j].Names[0]
	})

	deployments := make([]Deployment, len(mdbContainers))
	for i, c := range mdbContainers {
		stateName, found := localStateMap[c.State]
		if !found {
			stateName = strings.ToUpper(c.State)
		}
		name := strings.TrimPrefix(c.Names[0], options.MongodHostnamePrefix+"-")
		deployments[i] = Deployment{
			Type:           "LOCAL",
			Name:           name,
			MongoDBVersion: c.Labels["version"],
			StateName:      stateName,
		}
	}

	return deployments, nil
}

func (opts *ListOpts) Run(ctx context.Context) error {
	atlasClusters, err := opts.getAtlasDeployments()
	if err != nil {
		return err
	}

	mdbContainers, err := opts.getLocalDeployments(ctx)
	if err != nil && !errors.Is(err, podman.ErrPodmanNotFound) {
		return err
	}

	err = opts.Print(append(atlasClusters, mdbContainers...))
	if err != nil {
		return err
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
