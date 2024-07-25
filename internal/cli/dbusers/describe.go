// Copyright 2021 MongoDB Inc
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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

const describeTemplate = `USERNAME	DATABASE
{{.Username}}	{{.DatabaseName}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store    store.DatabaseUserDescriber
	authDB   string
	username string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.DatabaseUser(opts.authDB, opts.ConfigProjectID(), opts.username)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas dbuser(s) describe <username> --projectId projectId --authDB authDB.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <username>",
		Short:   "Return the details for the specified database user for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.ExactArgs(1),
		Aliases: []string{"get"},
		Annotations: map[string]string{
			"usernameDesc": "Username to retrieve from the MongoDB database. The format of the username depends on the user's method of authentication.",
			"output":       describeTemplate,
		},
		Example: `  # Return the details for the SCRAM SHA-authenticating database user named myDbUser:
  atlas dbuser describe myDbUser --authDB admin --output json

  # Return the details for the AWS IAM-authenticating database user with the ARN arn:aws:iam::772401394250:user/my-test-user. Prepend $external with \ to escape the special-use character:
  atlas dbuser describe arn:aws:iam::772401394250:user/my-test-user --authDB \$external --output json

  # Return the details for the X.509-authenticating database user with the RFC 2253 Distinguished Name CN=ellen@example.com,OU=users,DC=example,DC=com. Prepend $external with \ to escape the special-use character:
  atlas dbuser describe CN=ellen@example.com,OU=users,DC=example,DC=com --authDB \$external --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.username = args[0]

			validAuthDBs := []string{convert.AdminDB, convert.ExternalAuthDB}
			if err := validate.FlagInSlice(opts.authDB, flag.AuthDB, validAuthDBs); err != nil {
				return err
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.authDB, flag.AuthDB, convert.AdminDB, usage.AtlasAuthDB)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
