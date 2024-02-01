// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package globalapikeys

import (
	"context"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/store"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type UpdateOpts struct {
	cli.OutputOpts
	id    string
	desc  string
	roles []string
	store store.GlobalAPIKeyUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) newAPIKeyInput() *opsmngr.APIKeyInput {
	return &opsmngr.APIKeyInput{
		Desc:  opts.desc,
		Roles: opts.roles,
	}
}

const updateTemplate = "API Key '{{.ID}}' successfully updated.\n"

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateGlobalAPIKey(opts.id, opts.newAPIKeyInput())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli iam globalApiKey(s) update <ID> [--role role][--desc description].
func UpdateBuilder() *cobra.Command {
	opts := new(UpdateOpts)
	opts.Template = updateTemplate
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Args:  require.ExactArgs(1),
		Short: "Modify the roles and description for a global API key.",
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies the global API key.",
		},
		Example: `  # Modify the roles and description for the global API key with the ID 5f5bad7a57aef32b04ed0210:
  mongocli iam globalApiKeys update 5f5bad7a57aef32b04ed0210 --desc "My Sample Global API Key" --role GLOBAL_MONITORING_ADMIN --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.GlobalAPIKeyRoles+usage.UpdateWarning)
	cmd.Flags().StringVar(&opts.desc, flag.Description, "", usage.APIKeyDescription)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Description)
	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
