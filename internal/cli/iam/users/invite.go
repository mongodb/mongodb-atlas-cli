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

package users

import (
	"fmt"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

var inviteTemplate = "The user '{{.Username}}' has been invited.\nInvited users do not have access to the project until they accept the invitation.\n"

type InviteOpts struct {
	cli.OutputOpts
	username    string
	password    string
	country     string
	email       string
	mobile      string
	firstName   string
	lastName    string
	orgRole     []string
	projectRole []string
	store       store.UserCreator
}

func (opts *InviteOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *InviteOpts) Run() error {
	atlasRoles, err := opts.createAtlasRole()
	if err != nil {
		return err
	}

	userRoles, err := opts.createUserRole()
	if err != nil {
		return err
	}

	r, err := opts.store.CreateUser(opts.username, opts.password, opts.firstName, opts.lastName, opts.email, opts.mobile, opts.country, atlasRoles, userRoles)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

const keyParts = 2

func (opts *InviteOpts) createAtlasRole() ([]atlas.AtlasRole, error) {
	if config.Service() == config.CloudService {
		atlasRoles := make([]atlas.AtlasRole, len(opts.orgRole)+len(opts.projectRole))

		i := 0
		for _, role := range opts.orgRole {
			value := strings.Split(role, ":")
			if len(value) != keyParts {
				return nil, fmt.Errorf("unexpected role format: %s", role)
			}
			atlasRoles[i] = atlas.AtlasRole{
				OrgID:    value[0],
				RoleName: strings.ToUpper(value[1]),
			}
			i++
		}

		for _, role := range opts.projectRole {
			value := strings.Split(role, ":")
			if len(value) != keyParts {
				return nil, fmt.Errorf("unexpected role format: %s", role)
			}
			atlasRoles[i] = atlas.AtlasRole{
				GroupID:  value[0],
				RoleName: strings.ToUpper(value[1]),
			}
			i++
		}

		return atlasRoles, nil
	}

	return nil, nil
}

func (opts *InviteOpts) createUserRole() ([]*opsmngr.UserRole, error) {
	if config.Service() != config.CloudService {
		roles := make([]*opsmngr.UserRole, len(opts.orgRole)+len(opts.projectRole))

		i := 0
		for _, role := range opts.orgRole {
			value := strings.Split(role, ":")
			if len(value) != keyParts {
				return nil, fmt.Errorf("unexpected role format: %s", role)
			}
			roles[i] = &opsmngr.UserRole{
				OrgID:    value[0],
				RoleName: strings.ToUpper(value[1]),
			}
			i++
		}

		for _, role := range opts.projectRole {
			value := strings.Split(role, ":")
			if len(value) != keyParts {
				return nil, fmt.Errorf("unexpected role format: %s", role)
			}
			roles[i] = &opsmngr.UserRole{
				GroupID:  value[0],
				RoleName: strings.ToUpper(value[1]),
			}
			i++
		}

		return roles, nil
	}

	return nil, nil
}

// mongocli iam users(s) invite --username username --password password --country country --email email
// --mobile mobile --firstName firstName --lastName lastName --team team1,team2 --orgRole orgID:ROLE_NAME
// --projectRole projectID:ROLE_NAME

func InviteBuilder() *cobra.Command {
	opts := &InviteOpts{}
	opts.Template = inviteTemplate
	cmd := &cobra.Command{
		Use:   "invite",
		Short: inviteUser,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flag.Username, "", usage.Username)
	cmd.Flags().StringVar(&opts.password, flag.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.country, flag.Country, "", usage.Country)
	cmd.Flags().StringVar(&opts.email, flag.Email, "", usage.Email)
	cmd.Flags().StringVar(&opts.mobile, flag.Mobile, "", usage.Mobile)
	cmd.Flags().StringVar(&opts.firstName, flag.FirstName, "", usage.FirstName)
	cmd.Flags().StringVar(&opts.lastName, flag.LastName, "", usage.LastName)
	cmd.Flags().StringSliceVar(&opts.orgRole, flag.OrgRole, []string{}, usage.OrgRole)
	cmd.Flags().StringSliceVar(&opts.projectRole, flag.ProjectRole, []string{}, usage.ProjectRole)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
