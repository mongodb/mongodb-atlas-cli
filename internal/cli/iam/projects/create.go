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

package projects

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var createTemplate = "Project '{{.ID}}' created.\n"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.ProjectCreator
}

func (opts *CreateOpts) init() error {
	if opts.ConfigOrgID() == "" {
		return cli.ErrMissingOrgID
	}

	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateProject(opts.name, opts.ConfigOrgID())

	if err != nil {
		return err
	}
	if config.Service() != config.CloudService {
		createTemplate += "Agent API Key: '{{.AgentAPIKey}}'\n"
	}

	return opts.Print(r)
}

// mongocli iam project(s) create <name> [--orgId orgId]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: createProject,
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
