// Copyright 2020 MongoDB Inc
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

package ldap

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

var getTemplate = `HOSTNAME	PORT	AUTHENTICATION	AUTHORIZATION
{{.Ldap.Hostname}}	{{.Ldap.Port}}	{{.Ldap.AuthenticationEnabled}}	{{.Ldap.AuthorizationEnabled}}
`

type GetOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.LDAPConfigurationGetter
}

func (opts *GetOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GetOpts) Run() error {
	r, err := opts.store.GetLDAPConfiguration(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas security ldap get --projectId projectId.
func GetBuilder() *cobra.Command {
	opts := &GetOpts{}
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Return the current LDAP configuration details for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": getTemplate,
		},
		Example: `  # Return the JSON-formatted details of the current LDAP configuration in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas security ldap get --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), getTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
