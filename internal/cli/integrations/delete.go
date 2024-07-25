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

package integrations

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.IntegrationDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteIntegration, opts.ConfigProjectID())
}

// atlas integration(s) delete <TYPE> [--force] --projectId projectId.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Integration '%s' deleted\n", "Integration not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <integrationType>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified third-party integration from your project.",
		Long: `Deleting an integration from a project removes that integration configuration only for that project. This does not affect any other project or organization's configured integrations.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.ExactValidArgs(1),
		Example: `  # Remove the Datadog integration for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas integrations delete DATADOG --projectId 5e2211c17a3e5a48f5497de3`,
		ValidArgs: []string{"PAGER_DUTY", "MICROSOFT_TEAMS", "SLACK", "DATADOG", "NEW_RELIC", "OPS_GENIE", "VICTOR_OPS", "WEBHOOK", "PROMETHEUS"},
		Annotations: map[string]string{
			"integrationTypeDesc": "Human-readable label that identifies the service integration to delete. Valid values are PAGER_DUTY, MICROSOFT_TEAMS, SLACK, DATADOG, NEW_RELIC, OPS_GENIE, VICTOR_OPS, WEBHOOK, PROMETHEUS.",
			"output":              opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Entry = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.Prompt,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
