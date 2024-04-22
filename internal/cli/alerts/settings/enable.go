// Copyright 2023 MongoDB Inc
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

package settings

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type EnableOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	alertID string
	store   store.AlertConfigurationEnabler
}

func (opts *EnableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var enableTemplate = "Alert configuration '{{.Id}}' enabled\n"

func (opts *EnableOpts) Run() error {
	r, err := opts.store.EnableAlertConfiguration(opts.ConfigProjectID(), opts.alertID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas alerts enable <ID> --projectId projectId.
func EnableBuilder() *cobra.Command {
	opts := new(EnableOpts)
	cmd := &cobra.Command{
		Use:   "enable <alertConfigId>",
		Short: "Enables one alert configuration for the specified project.",
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		Example: `  # Enable the alert configuration with the ID 5d1113b25a115342acc2d1aa in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas alerts settings enable 5d1113b25a115342acc2d1aa --projectId 5e2211c17a3e5a48f5497de3`,
		Annotations: map[string]string{
			"alertConfigIdDesc": "ID of the alert you want to enable.",
			"output":            enableTemplate,
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}
	cmd.OutOrStdout()
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
