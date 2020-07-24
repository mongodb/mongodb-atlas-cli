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
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	username string
	password string
	x509Type string
	roles    []string
	store    store.DatabaseUserCreator
}

const (
	X509TypeManaged  = "MANAGED"
	X509TypeCustomer = "CUSTOMER"
	X509TypeNone     = "NONE"
)

var validX509Flags = []string{X509TypeNone, X509TypeManaged, X509TypeCustomer}

func (opts *CreateOpts) isX509Set() bool {
	return opts.x509Type != X509TypeNone
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	user := opts.newDatabaseUser()
	r, err := opts.store.CreateDatabaseUser(user)

	if err != nil {
		return err
	}

	return output.Print(config.Default(), "", r)
}

func (opts *CreateOpts) newDatabaseUser() *atlas.DatabaseUser {
	authDB := "admin"

	if opts.isX509Set() {
		authDB = "$external"
	}

	return &atlas.DatabaseUser{
		Roles:        convert.BuildAtlasRoles(opts.roles),
		GroupID:      opts.ConfigProjectID(),
		Username:     opts.username,
		Password:     opts.password,
		X509Type:     opts.x509Type,
		DatabaseName: authDB,
	}
}

func (opts *CreateOpts) Prompt() error {
	if opts.isX509Set() || opts.password != "" {
		return nil
	}
	prompt := &survey.Password{
		Message: "Password:",
	}
	return survey.AskOne(prompt, &opts.password)
}

func (opts *CreateOpts) validate() error {
	if len(opts.roles) == 0 {
		return errors.New("no role specified for the user")
	}

	if err := validate.FlagInSlice(opts.x509Type, flag.X509Type, validX509Flags); err != nil {
		return err
	}

	if opts.isX509Set() && opts.password != "" {
		return errors.New("cannot supply both x509 auth and password")
	}

	return nil
}

// mongocli atlas dbuser(s) create
//		--username username --password password
//		--role roleName@dbName
//		[--projectId projectId]
//		[--x509Type NONE|MANAGED|CUSTOMER]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateDBUser,
		Example: `  
  Create an Atlas admin user
  $ mongocli atlas dbuser create atlasAdmin --username <username>  --projectId <projectId>

  Create user with read/write access to any database
  $ mongocli atlas dbuser create readWriteAnyDatabase --username <username> --projectId <projectId>

  Create user with multiple roles
  $ mongocli atlas dbuser create --username <username> --role clusterMonitor,backup --projectId <projectId>`,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"atlasAdmin", "readWriteAnyDatabase", "readAnyDatabase", "clusterMonitor", "backup", "dbAdminAnyDatabase", "enableSharding"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.initStore); err != nil {
				return err
			}
			opts.roles = append(opts.roles, args...)

			if err := opts.validate(); err != nil {
				return err
			}

			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.username, flag.Username, flag.UsernameShort, "", usage.Username)
	cmd.Flags().StringVarP(&opts.password, flag.Password, flag.PasswordShort, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.Roles)
	cmd.Flags().StringVar(&opts.x509Type, flag.X509Type, X509TypeNone, usage.X509Type)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Username)

	return cmd
}
