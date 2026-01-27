// Copyright 2024 MongoDB Inc
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

package nodes

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=nodes . SearchNodesLister

type SearchNodesLister interface {
	SearchNodes(string, string) (*atlasv2.ApiSearchDeploymentResponse, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	clusterName string
	store       SearchNodesLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID							STATE		INSTANCE SIZE		NODE COUNT{{ $id:=.Id}}{{ $state:=.StateName }}{{range valueOrEmptySlice .Specs}}
{{$id}}	{{$state}}	{{.InstanceSize}}	{{.NodeCount}}{{end}}
`

func (opts *ListOpts) Run() error {
	r, err := opts.store.SearchNodes(opts.ConfigProjectID(), opts.clusterName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// ListBuilder represents atlas clusters search nodes list [--projectId projectId] [--clusterName name].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Atlas Search nodes for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return the JSON-formatted list of Atlas search nodes in the cluster named myCluster:
  atlas clusters search nodes list --clusterName myCluster --output json`,
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

	// Command specific flags
	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	_ = cmd.MarkFlagRequired(flag.ClusterName)

	// Global flags
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
