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

package settings

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312003/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=settings . AlertConfigurationLister

type AlertConfigurationLister interface {
	AlertConfigurations(*atlasv2.ListAlertConfigurationsApiParams) (*atlasv2.PaginatedAlertConfig, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	CompactResponse bool
	store           AlertConfigurationLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var settingsListTemplate = `ID	TYPE	ENABLED{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.EventTypeName}}	{{.Enabled}}{{end}}
`

func (opts *ListOpts) Run() error {
	params := &atlasv2.ListAlertConfigurationsApiParams{
		GroupId: opts.ConfigProjectID(),
		PageNum: &opts.PageNum,
	}
	if opts.ItemsPerPage > 0 {
		params.ItemsPerPage = &opts.ItemsPerPage
	}
	r, err := opts.store.AlertConfigurations(params)
	if err != nil {
		return err
	}

	if opts.CompactResponse {
		return opts.PrintForCompactResultsResponse(r)
	}

	return opts.Print(r)
}

// atlas alerts config(s) list --projectId projectId [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Returns all alert configurations for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # Return a JSON-formatted list of all alert configurations for the project with the ID 5df90590f10fab5e33de2305:
  atlas alerts settings list --projectId 5df90590f10fab5e33de2305 --output json`,
		Annotations: map[string]string{},
		Aliases:     []string{"ls"},
		Args:        require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), settingsListTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddListOptsFlagsWithoutOmitCount(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)
	cmd.Flags().BoolVarP(&opts.CompactResponse, flag.CompactResponse, flag.CompactResponseShort, false, usage.CompactResponse)

	return cmd
}
