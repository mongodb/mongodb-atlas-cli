// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"strings"

	"github.com/10gen/mcli/internal/usage"

	"github.com/10gen/mcli/internal/utils"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

const (
	adminDB = "admin"
	roleSep = "@"
)

type atlasDBUsersCreateOpts struct {
	*globalOpts
	username string
	password string
	roles    []string
	store    store.DatabaseUserCreator
}

func (opts *atlasDBUsersCreateOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New(opts.Config)

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *atlasDBUsersCreateOpts) Run() error {
	user := opts.newDatabaseUser()
	result, err := opts.store.CreateDatabaseUser(user)

	if err != nil {
		return err
	}

	return utils.PrettyJSON(result)
}

func (opts *atlasDBUsersCreateOpts) newDatabaseUser() *atlas.DatabaseUser {
	return &atlas.DatabaseUser{
		DatabaseName: adminDB,
		Roles:        opts.buildRoles(),
		GroupID:      opts.ProjectID(),
		Username:     opts.username,
		Password:     opts.password,
	}
}

func (opts *atlasDBUsersCreateOpts) buildRoles() []atlas.Role {
	rolesLen := len(opts.roles)
	roles := make([]atlas.Role, rolesLen)
	for i, roleP := range opts.roles {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := adminDB
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlas.Role{
			RoleName:     roleName,
			DatabaseName: databaseName,
		}
	}
	return roles
}

// mcli atlas dbuser(s) create --username username --password password --role roleName@dbName [--projectId projectId]
func AtlasDBUsersCreateBuilder() *cobra.Command {
	opts := &atlasDBUsersCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a database user for a project.",
		Args:  cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flags.Username, "", usage.Username)
	cmd.Flags().StringVar(&opts.password, flags.Password, "", usage.Password)
	cmd.Flags().StringSliceVar(&opts.roles, flags.Role, []string{}, usage.Roles)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.profile, flags.Profile, flags.ProfileShort, config.DefaultProfile, usage.Profile)

	_ = cmd.MarkFlagRequired(flags.Username)
	_ = cmd.MarkFlagRequired(flags.Password)
	_ = cmd.MarkFlagRequired(flags.Role)

	return cmd
}
