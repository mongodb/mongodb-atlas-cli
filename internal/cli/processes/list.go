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

package processes

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const listTemplate = `ID	REPLICA SET NAME	SHARD NAME	VERSION{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.ReplicaSetName}}	{{.ShardName}}	{{.Version}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=processes . ProcessLister

type ProcessLister interface {
	Processes(*atlasv2.ListAtlasProcessesApiParams) (*atlasv2.PaginatedHostViewAtlas, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	CompactResponse bool
	store           ProcessLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	listParams := opts.newProcessesListParams()
	r, err := opts.store.Processes(listParams)
	if err != nil {
		return err
	}

	if opts.CompactResponse {
		return opts.PrintForCompactResultsResponse(r)
	}

	return opts.Print(r)
}

func (opts *ListOpts) newProcessesListParams() *atlasv2.ListAtlasProcessesApiParams {
	listOpts := opts.NewAtlasListOptions()
	processesList := &atlasv2.ListAtlasProcessesApiParams{
		GroupId: opts.ConfigProjectID(),
	}
	if listOpts.PageNum > 0 {
		processesList.PageNum = &listOpts.PageNum
	}
	if listOpts.ItemsPerPage > 0 {
		processesList.ItemsPerPage = &listOpts.ItemsPerPage
	}

	if listOpts.IncludeCount {
		processesList.IncludeCount = &listOpts.IncludeCount
	}
	return processesList
}

// ListBuilder atlas process(es) list --projectId projectId [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all MongoDB processes for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of all MongoDB processes in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas processes list --projectId 5e2211c17a3e5a48f5497de3 --output json`,
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

	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)
	cmd.Flags().BoolVarP(&opts.CompactResponse, flag.CompactResponse, flag.CompactResponseShort, false, usage.CompactResponse)

	return cmd
}
