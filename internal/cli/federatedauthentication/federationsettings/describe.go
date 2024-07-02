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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const describeTemplate = `ID	IDENTITY PROVIDER ID	IDENTITY PROVIDER STATUS
{{.Id}}	{{.IdentityProviderId}}	{{.IdentityProviderStatus}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.FederationSettingsDescriber
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
		OrgId: opts.ConfigOrgID(),
	}
	r, err := opts.store.FederationSetting(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings describe --orgId orgId.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Args:    require.NoArgs,
		Short:   "Return the Federation Settings details for the specified organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization Owner"),
		Annotations: map[string]string{
			"output": describeTemplate,
		},
		Example: `  # Return the JSON-formatted Federation Settings details:
  atlas federatedAuthentication federationSettings describe --orgId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
				opts.ValidateOrgID,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
