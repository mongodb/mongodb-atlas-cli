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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	dayOfWeek int
	hourOfDay int
	startASAP bool
	store     store.MaintenanceWindowUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Maintenance window updated.\n"

func (opts *UpdateOpts) Run() error {
	err := opts.store.UpdateMaintenanceWindow(opts.ConfigProjectID(), opts.newMaintenanceWindow())
	if err != nil {
		return err
	}
	return opts.Print(nil)
}

func (opts *UpdateOpts) newMaintenanceWindow() *atlas.MaintenanceWindow {
	return &atlas.MaintenanceWindow{
		DayOfWeek: opts.dayOfWeek,
		HourOfDay: &opts.hourOfDay,
		StartASAP: &opts.startASAP,
	}
}

// mongocli atlas maintenanceWindow(s) update(s) --dayOfWeek dayOfWeek --hourOfDay hourOfDay --startASAP [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Modify the maintenance window for your project.",
		Long:  `To learn more about maintenance windows, see https://www.mongodb.com/docs/atlas/tutorial/cluster-maintenance-window/.`,
		Example: fmt.Sprintf(`  # Update the maintenance window to midnight on Saturdays for the project with the ID 5e2211c17a3e5a48f5497de3:
  %s maintenanceWindows update --dayOfWeek 7 --hourOfDay 0 --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !opts.startASAP {
				_ = cmd.MarkFlagRequired(flag.DayOfWeek)
				_ = cmd.MarkFlagRequired(flag.HourOfDay)
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.dayOfWeek, flag.DayOfWeek, 0, usage.DayOfWeek)
	cmd.Flags().IntVar(&opts.hourOfDay, flag.HourOfDay, 0, usage.HourOfDay)
	cmd.Flags().BoolVar(&opts.startASAP, flag.StartASAP, false, usage.StartASAP)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
