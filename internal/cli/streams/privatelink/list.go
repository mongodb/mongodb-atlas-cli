// Copyright 2025 MongoDB Inc
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

package privatelink

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312006/admin"
)

var listTemplate = `ID  PROVIDER  REGION  VENDOR  STATE  INTERFACE_ENDPOINT_ID  SERVICE_ENDPOINT_ID  DNS_DOMAIN  DNS_SUBDOMAIN{{range valueOrEmptySlice .Results}}
{{.Id}}  {{.Provider}}  {{.Region}}  {{.Vendor}}  {{.State}}  {{.InterfaceEndpointId}}  {{.ServiceEndpointId}}  {{.DnsDomain}}  {{.DnsSubDomain}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=privatelink . Lister

type Lister interface {
	ListPrivateLinkEndpoints(projectID string) (*atlasv2.PaginatedApiStreamsPrivateLink, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store Lister
}

func (opts *ListOpts) Run() error {
	result, err := opts.store.ListPrivateLinkEndpoints(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(result)
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams privateLink list
// List the PrivateLink endpoints in the project that can be used as Atlas Stream Processor connections.
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists the PrivateLink endpoints in the project that can be used as Atlas Stream Processor connections.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, commandRoles),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Example: `# list PrivateLink endpoints for Atlas Stream Processing:
  atlas streams privateLink list
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
