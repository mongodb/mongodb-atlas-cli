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
	"go.mongodb.org/atlas/mongodbatlas"
)

var createTemplate = "New API Key '{{.ID}}' created for project.\n"

type CreateOpts struct {
	cli.GlobalOpts
	store       store.ProjectAPIKeyCreator
	description string
	roles       []string
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	apiKeyInput := &mongodbatlas.APIKeyInput{
		Desc:  opts.description,
		Roles: opts.roles,
	}

	r, err := opts.store.CreateProjectAPIKey(opts.ProjectID, apiKeyInput)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTemplate, r)
}

// mongocli iam project apiKey create [--projectId projectId] [--roles roles] [--description description]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateProjectAPIKey,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}

			return opts.PreRunE()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringArrayVar(&opts.roles, flag.Role, nil, usage.Roles)
	cmd.Flags().StringVar(&opts.description, flag.Description, "", usage.APIKeyDescription)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Description)
	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
