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
	username     string
	password     string
	country      string
	email        string
	mobile       string
	firstName    string
	lastName     string
	orgRoles     []string
	projectRoles []string
	store        store.UserCreator
}

func (opts *InviteOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *InviteOpts) createUserView() (*store.UserView, error) {
	atlasRoles, err := opts.createAtlasRole()
	if err != nil {
		return nil, err
	}

	userRoles, err := opts.createUserRole()
	if err != nil {
		return nil, err
	}
	user := &store.UserView{
		AtlasRoles:   atlasRoles,
		Country:      opts.country,
		MobileNumber: opts.mobile,
		User: opsmngr.User{
			Username:     opts.username,
			Password:     opts.password,
			FirstName:    opts.firstName,
			LastName:     opts.lastName,
			EmailAddress: opts.email,
			Roles:        userRoles,
		},
	}

	return user, nil
}

func (opts *InviteOpts) Run() error {
	user, err := opts.createUserView()
	if err != nil {
		return err
	}

	r, err := opts.store.CreateUser(user)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

const keyParts = 2

func (opts *InviteOpts) createAtlasRole() ([]atlas.AtlasRole, error) {
	if config.Service() != config.CloudService {
		return nil, nil
	}

	atlasRoles := make([]atlas.AtlasRole, len(opts.orgRoles)+len(opts.projectRoles))

	i := 0
	for _, role := range opts.orgRoles {
		value := strings.Split(role, ":")
		if len(value) != keyParts {
			return nil, fmt.Errorf("unexpected role format: %s", role)
		}
		atlasRoles[i] = newAtlasOrgRole(value)
		i++
	}

	for _, role := range opts.projectRoles {
		value := strings.Split(role, ":")
		if len(value) != keyParts {
			return nil, fmt.Errorf("unexpected role format: %s", role)
		}
		atlasRoles[i] = newAtlasProjectRole(value)
		i++
	}

	return atlasRoles, nil
}

func (opts *InviteOpts) createUserRole() ([]*opsmngr.UserRole, error) {
	if config.Service() == config.CloudService {
		return nil, nil
	}

	if config.Service() != config.CloudService {
		roles := make([]*opsmngr.UserRole, len(opts.orgRoles)+len(opts.projectRoles))

		i := 0
		for _, role := range opts.orgRoles {
			value := strings.Split(role, ":")
			if len(value) != keyParts {
				return nil, fmt.Errorf("unexpected role format: %s", role)
			}
			roles[i] = newUserOrgRole(value)
			i++
		}

		for _, role := range opts.projectRoles {
			value := strings.Split(role, ":")
			if len(value) != keyParts {
				return nil, fmt.Errorf("unexpected role format: %s", role)
			}
			roles[i] = newUserProjectRole(value)
			i++
		}

		return roles, nil
	}

	return nil, nil
}

func newUserOrgRole(role []string) *opsmngr.UserRole {
	return &opsmngr.UserRole{
		OrgID:    role[0],
		RoleName: strings.ToUpper(role[1]),
	}
}

func newAtlasProjectRole(role []string) atlas.AtlasRole {
	return atlas.AtlasRole{
		GroupID:  role[0],
		RoleName: strings.ToUpper(role[1]),
	}
}

func newAtlasOrgRole(role []string) atlas.AtlasRole {
	return atlas.AtlasRole{
		OrgID:    role[0],
		RoleName: strings.ToUpper(role[1]),
	}
}

func newUserProjectRole(role []string) *opsmngr.UserRole {
	return &opsmngr.UserRole{
		GroupID:  role[0],
		RoleName: strings.ToUpper(role[1]),
	}
}

// mongocli iam users(s) invite --username username --password password --country country --email email
// --mobile mobile --firstName firstName --lastName lastName --team team1,team2 --orgRoles orgID:ROLE_NAME
// --projectRoles projectID:ROLE_NAME

func InviteBuilder() *cobra.Command {
	opts := &InviteOpts{}
	opts.Template = inviteTemplate
	cmd := &cobra.Command{
		Use:   "invite",
		Short: inviteUser,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			if config.Service() == config.CloudService {
				_ = cmd.MarkFlagRequired(flag.Country)
			}
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
	cmd.Flags().StringSliceVar(&opts.orgRoles, flag.OrgRole, []string{}, usage.OrgRole)
	cmd.Flags().StringSliceVar(&opts.projectRoles, flag.ProjectRole, []string{}, usage.ProjectRole)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Username)
	_ = cmd.MarkFlagRequired(flag.Password)
	_ = cmd.MarkFlagRequired(flag.Email)
	_ = cmd.MarkFlagRequired(flag.FirstName)
	_ = cmd.MarkFlagRequired(flag.LastName)

	return cmd
}
