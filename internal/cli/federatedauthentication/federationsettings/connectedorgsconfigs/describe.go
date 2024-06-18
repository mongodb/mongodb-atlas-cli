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

package connectedorgsconfigs

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	*DescribeOrgConfigsOpts
	federationSettingsID string
}

const describeTemplate = `ORG ID	DENTITY PROVIDER ID	DATA ACCESS IDENTITY PRODIVER IDs
{{.OrgId}}	{{if .IdentityProviderId }}	{{ .IdentityProviderId }}{{else}}N/A{{end}}	{{if and .DataAccessIdentityProviderIds (gt (len .DataAccessIdentityProviderIds) 0)}}{{range $index, $element := .DataAccessIdentityProviderIds}}{{if $index}}, {{end}}{{$element}}{{end}}{{else}}N/A{{end}}`

func (opts *DescribeOpts) Run() error {
	r, err := opts.GetConnectedOrgConfig(opts.federationSettingsID, opts.ConfigOrgID())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings connectedOrgConfigsConfig describe --identityProviderId identityProviderId --federationSettingsId federationSettingsId [-o/--output output].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{
		DescribeOrgConfigsOpts: &DescribeOrgConfigsOpts{},
	}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a Connected Org Config.",
		Args:  cobra.NoArgs,
		Example: `  # Describe a connected org config from the current profile org and federationSettingsId 5d1113b25a115342acc2d1aa 
			atlas federatedAuthentication federationSettings connectedOrgConfigsConfig describe --federationSettingsId 5d1113b25a115342acc2d1aa 
			# Describe a connected org config from the org with ID 7d1113b25a115342acc2d1aa and federationSettingsId 5d1113b25a115342acc2d1aa 
			atlas federatedAuthentication federationSettings connectedOrgConfigs describe --orgId 7d1113b25a115342acc2d1aa --federationSettingsId 5d1113b25a115342acc2d1aa 
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitDescribeStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.IdentityProviderID)

	return cmd
}
