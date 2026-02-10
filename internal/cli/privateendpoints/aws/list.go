// Copyright 2021 MongoDB Inc
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

package aws

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
)

var listTemplate = `ID	ENDPOINT SERVICE	STATUS	ERROR{{range valueOrEmptySlice .}}
{{.Id}}	{{.EndpointServiceName}}	{{.Status}}	{{.ErrorMessage}}{{end}}
`

const deprecatedFlagMessage = "--page and --limit are not supported by privateendpoint aws list"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=aws . PrivateEndpointLister

type PrivateEndpointLister interface {
	PrivateEndpoints(string, string) ([]atlasv2.EndpointService, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	store PrivateEndpointLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.PrivateEndpoints(opts.ConfigProjectID(), provider)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas privateEndpoint(s)|privateendpoint(s) aws list|ls [--projectId projectId].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Return all AWS private endpoints for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of all AWS private endpoints for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas privateEndpoints aws list --projectId 5e2211c17a3e5a48f5497de3 --output json`,
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

	opts.AddListOptsFlagsWithoutOmitCount(cmd)
	_ = cmd.Flags().MarkDeprecated(flag.Page, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Limit, deprecatedFlagMessage)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
