// Copyright 2021 MongoDB Inc
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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const updateTemplate = "Successfully updated database user '{{.Username}}'.\n"

type UpdateOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	username        string
	currentUsername string
	password        string
	authDB          string
	roles           []string
	scopes          []string
	x509Type        string
	store           store.DatabaseUserUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	current := new(admin.CloudDatabaseUser)
	opts.update(current)

	params := &admin.UpdateDatabaseUserApiParams{
		GroupId:           current.GroupId,
		DatabaseName:      current.DatabaseName,
		Username:          opts.currentUsername,
		CloudDatabaseUser: current,
	}
	r, err := opts.store.UpdateDatabaseUser(params)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) update(out *admin.CloudDatabaseUser) {
	out.GroupId = opts.ConfigProjectID()
	out.Username = opts.username
	if opts.username == "" {
		out.Username = opts.currentUsername
	}
	if opts.password != "" {
		out.Password = &opts.password
	}

	roles := convert.BuildAtlasRoles(opts.roles)
	out.Roles = &roles
	scopes := convert.BuildAtlasScopes(opts.scopes)
	out.Scopes = &scopes
	out.DatabaseName = opts.authDB
	if opts.authDB == "" {
		out.DatabaseName = convert.GetAuthDB(out)
	}
	if opts.x509Type != "" {
		out.X509Type = &opts.x509Type
	}
}

func (opts *UpdateOpts) validateAuthDB() error {
	if opts.authDB == "" {
		return nil
	}
	validAuthDBs := []string{convert.AdminDB, convert.ExternalAuthDB}
	return validate.FlagInSlice(opts.authDB, flag.AuthDB, validAuthDBs)
}

// atlas dbuser(s) update <username> [--password password] [--role roleName@dbName] [--projectId projectId] [--authDB authDB].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <username>",
		Short: "Modify the details of a database user in your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # Update roles for a database user named myUser for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas dbuser update myUser --role readWriteAnyDatabase --projectId 5e2211c17a3e5a48f5497de3

  # Update scopes for a database user named myUser for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas dbuser update myUser --scope resourceName:resourceType --projectId 5e2211c17a3e5a48f5497de3`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"usernameDesc": "Username to update in the MongoDB database.",
			"output":       updateTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.validateAuthDB,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.currentUsername = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.username, flag.Username, flag.UsernameShort, "", usage.DBUsername)
	cmd.Flags().StringVarP(&opts.password, flag.Password, flag.PasswordShort, "", usage.DBUserPassword)
	cmd.Flags().StringVar(&opts.authDB, flag.AuthDB, "", usage.AtlasAuthDB)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.Roles+usage.UpdateWarning)
	cmd.Flags().StringSliceVar(&opts.scopes, flag.Scope, []string{}, usage.Scopes+usage.UpdateWarning)
	cmd.Flags().StringVar(&opts.x509Type, flag.X509Type, none, usage.X509Type)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
