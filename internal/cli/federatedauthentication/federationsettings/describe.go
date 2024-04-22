// Copyright 2024 MongoDB Inc
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

package federationsettings

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

const describeTemplate = `ID	IDENTITY PROVIDER ID	IDENTITY PROVIDER STATUS
{{.Id}}	{{.IdentityProviderId}}	{{.IdentityProviderStatus}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	store store.FederationSettignsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	params := &atlasv2.GetFederationSettingsApiParams{
		OrgId: opts.id,
	}
	r, err := opts.store.FederationSetting(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings describe <orgID>.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <orgID>",
		Aliases: []string{"show"},
		Args:    require.ExactArgs(1),
		Short:   "Return the federation Settings details for the specified organizations.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization Owner"),
		Annotations: map[string]string{
			"orgIDDesc": "Unique 24-digit string that identifies the organization.",
			"output":    describeTemplate,
		},
		Example: `  # Return the JSON-formatted Federation Settings details for the organization with the ID 5e2211c17a3e5a48f5497de3:
  atlas federatedAuthentication federationSettings describe 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
