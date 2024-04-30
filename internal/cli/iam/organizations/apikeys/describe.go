// Copyright 2020 MongoDB Inc
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

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/store"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	store store.OrganizationAPIKeyDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

const describeTemplate = `ID	DESCRIPTION	PUBLIC KEY	PRIVATE KEY
{{.ID}}	{{.Desc}}	{{.PublicKey}}	{{.PrivateKey}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.OrganizationAPIKey(opts.ConfigOrgID(), opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli iam organizations(s) apiKey(s)|apikey(s) describe <ID> --orgID.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Aliases: []string{"show"},
		Args:    require.ExactArgs(1),
		Short:   "Return the details for the specified API key for your organization.",
		Long: fmt.Sprintf(`To view possible values for the ID argument, run %s organizations apiKeys list.

%s`, cli.ExampleAtlasEntryPoint(), fmt.Sprintf(usage.RequiredRole, "Organization Member")),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies your API key.",
			"output": describeTemplate,
		},
		Example: fmt.Sprintf(`  # Return the JSON-formatted details for the organization API key with the ID 5f24084d8dbffa3ad3f21234 for the organization with the ID 5a1b39eec902201990f12345:
  %s organizations apiKeys describe 5f24084d8dbffa3ad3f21234 --orgId 5a1b39eec902201990f12345 -output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
