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

package maintenance

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const longDesc = `To learn more about maintenance windows, see https://www.mongodb.com/docs/atlas/tutorial/cluster-maintenance-window/.

`

type ClearOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	Confirm bool
	store   store.MaintenanceWindowClearer
}

func (opts *ClearOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var clearTemplate = "Maintenance window removed.\n"

func (opts *ClearOpts) Run() error {
	err := opts.store.ClearMaintenanceWindow(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

// Prompt confirms that the resource should be deleted.
func (opts *ClearOpts) Prompt() error {
	if opts.Confirm {
		return nil
	}

	p := prompt.NewDeleteConfirm("maintenance window")
	return telemetry.TrackAskOne(p, &opts.Confirm)
}

// atlas maintenanceWindow(s) clear [--projectId projectId].
func ClearBuilder() *cobra.Command {
	opts := &ClearOpts{}
	cmd := &cobra.Command{
		Use:     "clear",
		Short:   "Clear the current maintenance window setting for your project.",
		Long:    longDesc + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Aliases: []string{"rm", "delete"},
		Annotations: map[string]string{
			"output": clearTemplate,
		},
		Example: `  # Clear the current maintenance window for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas maintenanceWindows clear --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), clearTemplate),
				opts.Prompt,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
