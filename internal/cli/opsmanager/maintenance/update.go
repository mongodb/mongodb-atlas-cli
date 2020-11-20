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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	alertType   []string
	startDate   string
	endDate     string
	description string
	id          string
	store       store.OpsManagerMaintenanceWindowUpdater
}

func (opts *UpdateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var updateTemplate = "Maintenance window '{{.ID}}' successfully updated.\n"

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateOpsManagerMaintenanceWindow(opts.ConfigProjectID(), opts.newMaintenanceWindow())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *UpdateOpts) newMaintenanceWindow() *opsmngr.MaintenanceWindow {
	return &opsmngr.MaintenanceWindow{
		ID:             opts.id,
		StartDate:      opts.startDate,
		EndDate:        opts.endDate,
		AlertTypeNames: opts.alertType,
		Description:    opts.description,
	}
}

// mongocli ops-manager maintenanceWindows update --startDate startDate --endDate endDate --alertType alertType --desc desc [--projectId projectId]
func UpdateBuilder() *cobra.Command {
	opts := new(UpdateOpts)
	cmd := &cobra.Command{
		Use:   "update <id>",
		Short: updateMaintenanceWindow,
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.init,
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.startDate, flag.StartDate, "", usage.StartDate)
	cmd.Flags().StringVar(&opts.endDate, flag.EndDate, "", usage.EndDate)
	cmd.Flags().StringSliceVar(&opts.alertType, flag.AlertType, []string{}, usage.AlertType)
	cmd.Flags().StringVar(&opts.description, flag.Description, "", usage.MaintenanceDescription)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
