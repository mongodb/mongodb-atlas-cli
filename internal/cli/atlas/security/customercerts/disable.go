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

package customercerts

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DisableOpts struct {
	cli.GlobalOpts
	store   store.X509CertificateConfDisabler
	confirm bool
}

func (opts *DisableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DisableOpts) Run() error {
	if !opts.confirm {
		fmt.Printf("X.509 configuration was not disabled.\n")
		return nil
	}

	if err := opts.store.DisableX509Configuration(opts.ConfigProjectID()); err != nil {
		return err
	}

	fmt.Printf("X.509 configuration for project %s was deleted.\n", opts.ConfigProjectID())

	return nil
}

func (opts *DisableOpts) Prompt() error {
	prompt := &survey.Confirm{
		Message: "Are you sure you want to delete the X509 configuration for this project?",
	}

	return telemetry.TrackAskOne(prompt, &opts.confirm)
}

// mongocli atlas security certs disable --projectId projectId.
func DisableBuilder() *cobra.Command {
	opts := &DisableOpts{}
	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Clear customer-managed X.509 settings on a project, including the uploaded Certificate Authority, and disable self-managed X.509.",
		Long:  "Disabling customer-managed X.509 triggers a rolling restart.",
		Args:  require.NoArgs,
		Example: fmt.Sprintf(`  # Disable the customer-managed X.509 configuration in the project with the ID 5e2211c17a3e5a48f5497de3:
  %s security customerCerts disable --projectId 5e2211c17a3e5a48f5497de3`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}

			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
