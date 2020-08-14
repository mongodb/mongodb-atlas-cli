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

package organizations

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/spf13/cobra"
)

const createTemplate = "Organization '{{.ID}}' created.\n"

type CreateOpts struct {
	name  string
	store store.OrganizationCreator
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	p, err := opts.store.CreateOrganization(opts.name)

	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTemplate, p)
}

// mongocli iam organization(s) create <name>
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: createOrganization,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}

	return cmd
}
