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

package dbusers

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	authDB string
	store  store.DatabaseUserDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteDatabaseUser, opts.authDB, opts.ConfigProjectID())
}

// atlas dbuser(s) delete <username> --force.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("DB user '%s' deleted\n", "DB user not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <username>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified database user from your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"usernameDesc": "Username to delete from the MongoDB database. The format of the username depends on the user's method of authentication.",
			"output":       opts.SuccessMessage(),
		},
		Example: `  # Delete the SCRAM SHA-authenticating database user named dylan for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers delete dylan --projectId 5e2211c17a3e5a48f5497de3

  # Delete the AWS IAM-authenticating database user with the ARN arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs for the project with ID 5e2211c17a3e5a48f5497de3. Prepend $external with \ to escape the special-use character:
  atlas dbusers delete arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs --authDB \$external --projectId 5e2211c17a3e5a48f5497de3
			
  # Delete the xLDAP-authenticating database user with the RFC 2253 Distinguished Name CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM for the project with ID 5e2211c17a3e5a48f5497de3. Prepend $external with \ to escape the special-use character:
  atlas dbusers delete CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM --authDB \$external --projectId 5e2211c17a3e5a48f5497de3`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.authDB, flag.AuthDB, convert.AdminDB, usage.AtlasAuthDB)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
