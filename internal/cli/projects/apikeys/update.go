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
package apikeys

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type UpdateOpts struct {
	cli.GlobalOpts
	id    string
	roles []string
	store store.ProjectAPIKeyUpdater
}

func (opts *UpdateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) newAssignAPIKey() *atlas.AssignAPIKey {
	return &atlas.AssignAPIKey{
		Roles: opts.roles,
	}
}

const updateTemplate = "Successfully updated APIKey.\n"

func (opts *UpdateOpts) Run() error {
	err := opts.store.UpdateProjectAPIKey(opts.ConfigOrgID(), opts.id, opts.newAssignAPIKey())
	if err != nil {
		return err
	}

	return output.Print(config.Default(), updateTemplate, "")
}

// mongocli iam project(s) apiKey(s)|apikey(s) update <ID> [--role role][--projectId projectId]
func UpdateBuilder() *cobra.Command {
	opts := new(UpdateOpts)
	cmd := &cobra.Command{
		Use:     "update <ID>",
		Aliases: []string{"updates"},
		Args:    cobra.ExactArgs(1),
		Short:   description.ProjectOrganizationsAPIKey,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.init)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.APIKeyRoles)

	cmd.Flags().StringVar(&opts.OrgID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
