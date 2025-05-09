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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

var describeTemplate = `ID  PROVIDER  REGION  VENDOR  STATE  INTERFACE_ENDPOINT_ID  SERVICE_ENDPOINT_ID  DNS_DOMAIN  DNS_SUBDOMAIN
{{.Id}}  {{.Provider}}  {{.Region}}  {{.Vendor}}  {{.State}}  {{.InterfaceEndpointId}}  {{.ServiceEndpointId}}  {{.DnsDomain}}  {{.DnsSubDomain}}
`

//go:generate mockgen -typed -destination=describe_mock_test.go -package=privatelink . Describer

type Describer interface {
	DescribePrivateLinkEndpoint(projectID, connectionID string) (*atlasv2.StreamsPrivateLinkConnection, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store        Describer
	connectionID string
}

func (opts *DescribeOpts) Run() error {
	if opts.connectionID == "" {
		return errors.New("connectionID is missing")
	}

	result, err := opts.store.DescribePrivateLinkEndpoint(opts.ConfigProjectID(), opts.connectionID)
	if err != nil {
		return err
	}

	return opts.Print(result)
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams privateLink describe <connectionID>
// Describe a PrivateLink endpoint that can be used as an Atlas Stream Processor connection.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <connectionID>",
		Short: "Describes a PrivateLink endpoint that can be used as an Atlas Stream Processor connection.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, commandRoles),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"connectionIDDesc": "ID of the PrivateLink endpoint.",
			"output":           describeTemplate,
		},
		Example: `# describe a PrivateLink endpoint for Atlas Stream Processing:
  atlas streams privateLink describe 5e2211c17a3e5a48f5497de3
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.connectionID = args[0]
			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
