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

package alerts

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	status string
	store  store.AlertLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	TYPE   	STATUS{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.EventTypeName}}	{{.Status}}{{end}}
`

func (opts *ListOpts) Run() error {
	params := &atlasv2.ListAlertsApiParams{
		GroupId:      opts.ConfigProjectID(),
		ItemsPerPage: &opts.ItemsPerPage,
		PageNum:      &opts.PageNum,
	}

	if opts.OmitCount {
		params.IncludeCount = pointer.Get(false)
	}

	if opts.status != "" {
		params.Status = &opts.status
	}

	r, err := opts.store.Alerts(params)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas alerts list [--status status] [--projectId projectId] [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return all alerts for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # Return a JSON-formatted list of all alerts for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas alerts list --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		Aliases: []string{"ls"},
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Args: require.NoArgs,
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
	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().BoolVar(&opts.OmitCount, flag.OmitCount, false, usage.OmitCount)

	cmd.Flags().StringVar(&opts.status, flag.Status, "", usage.Status)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	return cmd
}
