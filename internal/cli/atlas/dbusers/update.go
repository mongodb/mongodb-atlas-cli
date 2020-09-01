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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const updateTemplate = "Successfully updated database user '{{.Username}}'.\n"

type UpdateOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	username string
	password string
	roles    []string
	store    store.DatabaseUserUpdater
}

func (opts *UpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	current := new(atlas.DatabaseUser)
	opts.update(current)
	r, err := opts.store.UpdateDatabaseUser(current)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) update(out *atlas.DatabaseUser) {
	out.GroupID = opts.ConfigProjectID()
	out.Username = opts.username
	if opts.password != "" {
		out.Password = opts.password
	}

	out.Roles = convert.BuildAtlasRoles(opts.roles)
}

// mongocli atlas dbuser(s) update <username> [--password password] [--role roleName@dbName] [--projectId projectId]
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <username>",
		Short: updateDBUser,
		Example: `
  Update roles for a user
  $ mongocli atlas dbuser update <username> --role readWriteAnyDatabase --projectId <projectId>`,
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.username = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.username, flag.Username, flag.UsernameShort, "", usage.DBUsername)
	cmd.Flags().StringVarP(&opts.password, flag.Password, flag.PasswordShort, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.Roles)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
