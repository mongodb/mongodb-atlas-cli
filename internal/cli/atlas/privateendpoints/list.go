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

package privateendpoints

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var listTemplate = `ID	ENDPOINT SERVICE	STATUS	ERROR{{range .}}
{{.ID}}	{{.EndpointServiceName}}	{{.Status}}	{{.ErrorMessage}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	store store.PrivateEndpointLister
}

func (opts *ListOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.PrivateEndpoints(opts.ConfigProjectID(), opts.NewListOptions())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas privateEndpoint(s)|privateendpoint(s) list|ls [--projectId projectId]
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   listPrivateEndpoints,
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.init,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
