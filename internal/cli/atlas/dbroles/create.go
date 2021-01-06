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

package dbroles

import (
	"errors"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const createTemplate = "Custom Database Role successfully created.\n"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	action         string
	collection     string
	db             string
	roleName       string
	inheritedRoles []string
	cluster        bool
	store          store.DatabaseRoleCreator
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	role := opts.newCustomDBRole()

	r, err := opts.store.CreateDatabaseRole(opts.ConfigProjectID(), role)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCustomDBRole() *atlas.CustomDBRole {
	resource := atlas.Resource{}
	role := &atlas.CustomDBRole{
		RoleName: opts.roleName,
	}
	if opts.action != "" {
		if opts.cluster {
			resource.Cluster = &opts.cluster
		} else {
			resource.Collection = opts.collection
			resource.Db = opts.db
		}

		action := atlas.Action{
			Action:    opts.action,
			Resources: []atlas.Resource{resource},
		}
		role.Actions = []atlas.Action{action}
	}

	if opts.inheritedRoles != nil {
		role.InheritedRoles = convert.BuildAtlasInheritedRoles(opts.inheritedRoles)
	}

	return role
}

func (opts *CreateOpts) validate() error {
	if opts.cluster && (opts.collection != "" || opts.db != "") {
		return errors.New("you can't use --cluster with --db and --collection ")
	}

	if opts.action == "" && opts.inheritedRoles == nil {
		return errors.New("you must provide either actions or inherited roles")
	}

	if opts.action != "" && opts.db == "" && !opts.cluster {
		return errors.New("you must provide --db databaseName with --action")
	}

	return nil
}

// mongocli atlas dbrole(s) create roleName --action actionName --db db --collection collection --inheritedRole role@db
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: createDBRole,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.roleName = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.inheritedRoles, flag.InheritedRole, []string{}, usage.InheritedRoles)
	cmd.Flags().StringVar(&opts.action, flag.Action, "", usage.Action)
	cmd.Flags().StringVar(&opts.db, flag.Database, "", usage.DatabaseCustomRole)
	cmd.Flags().StringVar(&opts.collection, flag.Collection, "", usage.Collection)
	cmd.Flags().BoolVar(&opts.cluster, flag.Cluster, false, usage.CustomDBRoleCluster)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
