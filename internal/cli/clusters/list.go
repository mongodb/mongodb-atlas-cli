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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=clusters . ClusterLister

type ClusterLister interface {
	ProjectClusters(string, *store.ListOptions) (*atlasClustersPinned.PaginatedAdvancedClusterDescription, error)
	ListFlexClusters(*atlasv2.ListFlexClustersApiParams) (*atlasv2.PaginatedFlexClusters20241113, error)
	LatestProjectClusters(string, *store.ListOptions) (*atlasv2.PaginatedClusterDescription20240805, error)
	GetClusterAutoScalingConfig(string, string) (*atlasv2.ClusterDescriptionAutoScalingModeConfiguration, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	tier            string
	autoScalingMode string
	store           ClusterLister
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
	if isIndependentShardScaling(opts.autoScalingMode) {
		r, err := opts.store.LatestProjectClusters(opts.ConfigProjectID(), listOpts)
		if err != nil {
			return err
		}
		r, err = opts.filterClustersByAutoScalingMode(r)
		if err != nil {
			return err
		}
		return opts.Print(r)
	}

	r, err := opts.store.ProjectClusters(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) filterClustersByAutoScalingMode(clusters *atlasv2.PaginatedClusterDescription20240805) (*atlasv2.PaginatedClusterDescription20240805, error) {
	filteredClusters := make([]atlasv2.ClusterDescription20240805, 0)
	for _, cluster := range clusters.GetResults() {
		clusterAutoScalingConfig, err := opts.store.GetClusterAutoScalingConfig(opts.ConfigProjectID(), *cluster.Name)
		if err != nil {
			return nil, err
		}
		if isIndependentShardScaling(clusterAutoScalingConfig.GetAutoScalingMode()) {
			filteredClusters = append(filteredClusters, cluster)
		}
	}

	clusters.Results = &filteredClusters
	return clusters, nil
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
  atlas clusters list --projectId 5e2211c17a3e5a48f5497de3 --output json
 
  # Return a JSON-formatted list of all clusters for the project with ID 5e2211c17a3e5a48f5497de3 and with independent shard scaling mode:
  atlas clusters list --projectId 5e2211c17a3e5a48f5497de3 --autoScalingMode independentShardScaling --output json
  `,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				validate.AutoScalingMode(opts.autoScalingMode),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)
	cmd.Flags().StringVar(&opts.autoScalingMode, flag.AutoScalingMode, clusterWideScalingFlag, usage.AutoScalingMode)
	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
