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

package cli

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasDBUsersCreateOpts struct {
	globalOpts
	username string
	password string
	authDB   string
	roles    []string
	store    store.DatabaseUserCreator
}

func (opts *atlasDBUsersCreateOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasDBUsersCreateOpts) Run() error {
	user := opts.newDatabaseUser()
	result, err := opts.store.CreateDatabaseUser(user)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasDBUsersCreateOpts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		DatabaseName: opts.authDB,
		Roles:        convert.BuildAtlasRoles(opts.roles),
		GroupID:      opts.ProjectID(),
		Username:     opts.username,
		Password:     opts.password,
	}
}

func (opts *atlasDBUsersCreateOpts) Prompt() error {
	if opts.password != "" {
		return nil
	}
	prompt := &survey.Password{
		Message: "Password:",
	}
	return survey.AskOne(prompt, &opts.password)
}

// mongocli atlas dbuser(s) create --username username --password password --role roleName@dbName [--projectId projectId]
func AtlasDBUsersCreateBuilder() *cobra.Command {
	opts := &atlasDBUsersCreateOpts{}
	cmd := &cobra.Command{
		Use:       "create",
		Short:     description.CreateDBUser,
		Example:   `  mongocli atlas dbuser create --username User1 --password passW0rd --role atlasAdmin,readWriteAnyDatabase --projectId <>`,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"atlasAdmin", "readWriteAnyDatabase", "readAnyDatabase", "clusterMonitor", "backup", "dbAdminAnyDatabase", "enableSharding"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.initStore); err != nil {
				return err
			}
			if len(args) == 0 && len(opts.roles) == 0 {
				return errors.New("no role specified for the user")
			}
			opts.roles = append(opts.roles, args...)
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.username, flags.Username, flags.UsernameShort, "", usage.Username)
	cmd.Flags().StringVarP(&opts.password, flags.Password, flags.PasswordShort, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flags.Role, []string{}, usage.Roles)
	cmd.Flags().StringVar(&opts.authDB, flags.AuthDB, convert.AdminDB, usage.AuthDB)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.Username)

	return cmd
}
