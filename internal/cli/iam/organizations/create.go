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
	"context"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/usage"
	"github.com/spf13/cobra"
)

const createTemplate = "Organization '{{.ID}}' created.\n"

type CreateOpts struct {
	cli.OutputOpts
	name  string
	store store.OrganizationCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateOrganization(opts.name)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// CreateBuilder mongocli iam organization(s) create <name>.
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create an Ops Manager or Cloud Manager organization. This command is unavailable for Atlas.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Label that identifies the organization.",
			"output":   createTemplate,
		},
		Example: `  # Create an Ops Manager organization with the name myOrg:
  mongocli iam organizations create myOrg --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return prerun.ExecuteE(
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
