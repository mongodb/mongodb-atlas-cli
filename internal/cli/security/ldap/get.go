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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

var getTemplate = `HOSTNAME	PORT	AUTHENTICATION	AUTHORIZATION
{{.Ldap.Hostname}}	{{.Ldap.Port}}	{{.Ldap.AuthenticationEnabled}}	{{.Ldap.AuthorizationEnabled}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=get_mock_test.go -package=ldap . Getter

type Getter interface {
	GetLDAPConfiguration(string) (*atlasv2.UserSecurity, error)
}

type GetOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store Getter
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

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
