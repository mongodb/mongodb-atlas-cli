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

package whitelist

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)


const listTemplate = `ID	DESCRIPTION	PUBLIC KEY	PRIVATE KEY{{range .}}
{{.ID}}	{{.Desc}}	{{.PublicKey}}	{{.PrivateKey}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.ListOpts
	id string
	store store.OrganizationAPIKeyLister
}

func (opts *ListOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.OrganizationAPIKeys(opts.ConfigOrgID(), opts.NewListOptions())

	if err != nil {
		return err
	}

	return output.Print(config.Default(), listTemplate, r)
}

// mongocli iam organizations|orgs apiKey(s)|apikey(s) whitelist|ipwhitelist <APIKey_ID> list|ls [--orgId orgId]
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(1),
		Short:   description.ListWhitelistAPIKeys,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	return cmd
}
