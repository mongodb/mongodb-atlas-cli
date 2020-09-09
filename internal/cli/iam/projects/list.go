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

package projects

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const listTemplate = `id	NAME{{range .Results}}
{{.id}}	{{.Name}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	store store.ProjectLister
}

func (opts *ListOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ListOpts) Run() error {
	var r interface{}
	var err error
	listOptions := opts.NewListOptions()
	if opts.ConfigOrgID() != "" && config.Service() == config.OpsManagerService {
		r, err = opts.store.GetOrgProjects(opts.ConfigOrgID(), listOptions)
	} else {
		r, err = opts.store.Projects(listOptions)
	}
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// mongocli iam project(s) list [--orgId orgId]
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	opts.Template = listTemplate
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   listProjects,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
