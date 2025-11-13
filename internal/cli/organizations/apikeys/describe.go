// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apikeys

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=apikeys . OrganizationAPIKeyDescriber

type OrganizationAPIKeyDescriber interface {
	OrganizationAPIKey(string, string) (*atlasv2.ApiKeyUserDetails, error)
}

type DescribeOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	id    string
	store OrganizationAPIKeyDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

const describeTemplate = `ID	DESCRIPTION	PUBLIC KEY	PRIVATE KEY
{{.Id}}	{{.Desc}}	{{.PublicKey}}	{{.PrivateKey}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.OrganizationAPIKey(opts.ConfigOrgID(), opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas organizations(s) apiKey(s)|apikey(s) describe <ID> --orgID.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Aliases: []string{"show"},
		Args:    require.ExactArgs(1),
		Short:   "Return the details for the specified API key for your organization.",
		Long:    longDesc + fmt.Sprintf(usage.RequiredRole, "Organization Member"),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies your API key.",
			"output": describeTemplate,
		},
		Example: `  # Return the JSON-formatted details for the organization API key with the ID 5f24084d8dbffa3ad3f21234 for the organization with the ID 5a1b39eec902201990f12345:
  atlas organizations apiKeys describe 5f24084d8dbffa3ad3f21234 --orgId 5a1b39eec902201990f12345 -output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
