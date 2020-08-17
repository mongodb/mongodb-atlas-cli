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
package peering

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type ListOpts struct {
	cli.GlobalOpts
	cli.ListOpts
	provider string
	store    store.PeeringConnectionLister
}

func (opts *ListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var listTemplate = `ID	STATUS	CONTAINER ID{{range .}}
{{.ID}}	{{if .AWSAccountID}}{{.StatusName}}{{else}}{{.Status}}{{end}}	{{.ContainerID}}{{end}}
`

func (opts *ListOpts) Run() error {
	var r []atlas.Peer
	var err error
	r, err = opts.store.PeeringConnections(opts.ConfigProjectID(), opts.newContainerListOptions())
	if err != nil {
		return err
	}
	return output.Print(config.Default(), listTemplate, r)
}

func (opts *ListOpts) newContainerListOptions() *atlas.ContainersListOptions {
	return &atlas.ContainersListOptions{
		ListOptions:  *opts.NewListOptions(),
		ProviderName: opts.provider,
	}
}

// mongocli atlas networking peering list [--provider provider] [--projectId projectId] [--page N] [--limit N]
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   listPeering,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.initStore()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.Provider)
	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
