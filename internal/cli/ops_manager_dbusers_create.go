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
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/messages"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const scransha1 = "SCRAM-SHA-1"

type opsManagerDBUsersCreateOpts struct {
	*globalOpts
	username   string
	password   string
	authDB     string
	roles      []string
	mechanisms []string
	store      store.AutomationPatcher
}

func (opts *opsManagerDBUsersCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerDBUsersCreateOpts) Run() error {
	current, err := opts.store.GetAutomationConfig(opts.ProjectID())

	if err != nil {
		return err
	}

	convert.AddUser(current, opts.newDBUser())

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), current); err != nil {
		return err
	}

	fmt.Print(messages.DeploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

func (opts *opsManagerDBUsersCreateOpts) newDBUser() *om.MongoDBUser {
	return &om.MongoDBUser{
		Database:                   opts.authDB,
		Username:                   opts.username,
		InitPassword:               opts.password,
		Roles:                      convert.BuildOMRoles(opts.roles),
		AuthenticationRestrictions: []string{},
		Mechanisms:                 opts.mechanisms,
	}
}

func (opts *opsManagerDBUsersCreateOpts) Prompt() error {
	if opts.password != "" {
		return nil
	}
	prompt := &survey.Password{
		Message: "Password:",
	}
	return survey.AskOne(prompt, &opts.password)
}

// mongocli atlas dbuser(s) create --username username --password password --role roleName@dbName [--projectId projectId]
func OpsManagerDBUsersCreateBuilder() *cobra.Command {
	opts := &opsManagerDBUsersCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "create",
		Short:   description.CreateDBUser,
		Example: `  mongocli om dbuser create --username User1 --password passW0rd --role readWriteAnyDatabase,clusterMonitor --mechanism SCRAM-SHA-256 --projectId <>`,
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flags.Username, "", usage.Username)
	cmd.Flags().StringVar(&opts.password, flags.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.authDB, flags.AuthDB, convert.AdminDB, usage.AuthDB)
	cmd.Flags().StringSliceVar(&opts.roles, flags.Role, []string{}, usage.Roles)
	cmd.Flags().StringSliceVar(&opts.mechanisms, flags.Mechanisms, []string{scransha1}, usage.Mechanisms)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.Username)

	return cmd
}
