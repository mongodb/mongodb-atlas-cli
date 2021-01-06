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

const updateTemplate = "Custom Database Role successfully updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	action         string
	collection     string
	db             string
	roleName       string
	inheritedRoles []string
	cluster        bool
	append         bool
	store          store.DatabaseRoleUpdater
}

func (opts *UpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	var role *atlas.CustomDBRole
	var err error
	if opts.append {
		if role, err = opts.store.DatabaseRole(opts.ConfigProjectID(), opts.roleName); err != nil {
			return err
		}
	}

	out, err := opts.store.UpdateDatabaseRole(opts.ConfigProjectID(), opts.roleName, opts.newCustomDBRole(role))
	if err != nil {
		return err
	}

	return opts.Print(out)
}

func (opts *UpdateOpts) newCustomDBRole(role *atlas.CustomDBRole) *atlas.CustomDBRole {
	out := &atlas.CustomDBRole{}
	resource := atlas.Resource{}
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

		out.Actions = []atlas.Action{action}
	}

	if opts.inheritedRoles != nil {
		out.InheritedRoles = convert.BuildAtlasInheritedRoles(opts.inheritedRoles)
	}

	if opts.append {
		out.Actions = append(out.Actions, role.Actions...)
		out.InheritedRoles = append(out.InheritedRoles, role.InheritedRoles...)
	}

	return out
}

func (opts *UpdateOpts) validate() error {
	if opts.cluster && (opts.collection != "" || opts.db != "") {
		return errors.New("you can't use --cluster with --db and --collection ")
	}

	if opts.action == "" && opts.inheritedRoles == nil {
		return errors.New("you must provide either actions or inherited roles")
	}

	if opts.action != "" && opts.db == "" && !opts.cluster {
		return errors.New("you must provide --db with --action")
	}

	return nil
}

// mongocli atlas dbrole(s) update roleName --action actionName --db db --collection collection --inheritedRole role@db --append
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: updateDBRole,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
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
	cmd.Flags().BoolVar(&opts.append, flag.Append, false, usage.Append)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
