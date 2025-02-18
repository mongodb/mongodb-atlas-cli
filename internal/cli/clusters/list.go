// Copyright 2020 MongoDB Inc
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

package clusters

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	tier  string
	store store.ClusterLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	NAME	MDB VER	STATE{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Name}}	{{.MongoDBVersion}}	{{.StateName}}{{end}}
`

func (opts *ListOpts) Run() error {
	if opts.tier == atlasFlex {
		return opts.RunFlexCluster()
	}

	return opts.RunDedicatedCluster()
}

func (opts *ListOpts) RunDedicatedCluster() error {
	listOpts := opts.NewAtlasListOptions()
	r, err := opts.store.ProjectClusters(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) RunFlexCluster() error {
	r, err := opts.store.ListFlexClusters(opts.newListFlexClustersAPIParams())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newListFlexClustersAPIParams() *atlasv2.ListFlexClustersApiParams {
	includeCount := !opts.OmitCount
	return &atlasv2.ListFlexClustersApiParams{
		GroupId:      opts.ConfigProjectID(),
		IncludeCount: &includeCount,
		ItemsPerPage: &opts.ItemsPerPage,
		PageNum:      &opts.PageNum,
	}
}

// ListBuilder builds a cobra.Command that can run as:
// atlas cluster(s) list --projectId projectId [--page N] [--limit N] [--tier tier].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all clusters for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Example: `  # Return a JSON-formatted list of all clusters for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas clusters list --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)
	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
