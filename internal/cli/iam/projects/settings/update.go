// Copyright 2022 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const updateTemplate = "Project settings updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	projectSettings                           *atlas.ProjectSettings
	store                                     store.ProjectSettingsGetterUpdater
	enableCollectDatabaseSpecificsStatistics  bool
	disableCollectDatabaseSpecificsStatistics bool
	enableDataExplorer                        bool
	disableDataExplorer                       bool
	enablePerformanceAdvisor                  bool
	disablePerformanceAdvisor                 bool
	enableSchemaAdvisor                       bool
	disableSchemaAdvisor                      bool
	enableRealtimePerformancePanel            bool
	disableRealtimePerformancePanel           bool
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateProjectSettings(opts.ConfigProjectID(), opts.projectSettings)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) GetCurrentProjectSettings() *atlas.ProjectSettings {
	r, err := opts.store.ProjectSettings(opts.ConfigProjectID())
	if err != nil {
		return nil
	}
	return r
}

func (opts *UpdateOpts) newProjectSettings(currentSettings *atlas.ProjectSettings) {
	opts.projectSettings = &atlas.ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: returnValueForSetting(opts.enableCollectDatabaseSpecificsStatistics, opts.disableCollectDatabaseSpecificsStatistics, *currentSettings.IsCollectDatabaseSpecificsStatisticsEnabled),
		IsDataExplorerEnabled:                       returnValueForSetting(opts.enableDataExplorer, opts.disableDataExplorer, *currentSettings.IsDataExplorerEnabled),
		IsPerformanceAdvisorEnabled:                 returnValueForSetting(opts.enablePerformanceAdvisor, opts.disablePerformanceAdvisor, *currentSettings.IsPerformanceAdvisorEnabled),
		IsRealtimePerformancePanelEnabled:           returnValueForSetting(opts.enableRealtimePerformancePanel, opts.disableRealtimePerformancePanel, *currentSettings.IsSchemaAdvisorEnabled),
		IsSchemaAdvisorEnabled:                      returnValueForSetting(opts.enableSchemaAdvisor, opts.disableSchemaAdvisor, *currentSettings.IsRealtimePerformancePanelEnabled),
	}
}

func returnValueForSetting(enableFlag, disableFlag, previousValue bool) *bool {
	var valueToSet bool
	if enableFlag {
		valueToSet = true
		return &valueToSet
	}
	if disableFlag {
		valueToSet = false
		return &valueToSet
	}
	return &previousValue
}

func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"updates"},
		Short:   "Updates settings of a given project.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			preRun := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
			opts.newProjectSettings(opts.GetCurrentProjectSettings())
			return preRun
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	cmd.Flags().BoolVarP(&opts.enableCollectDatabaseSpecificsStatistics, flag.EnableCollectDatabaseSpecificsStatistics, "", false, usage.EnableCollectDatabaseSpecificsStatistics)
	cmd.Flags().BoolVarP(&opts.disableCollectDatabaseSpecificsStatistics, flag.DisableCollectDatabaseSpecificsStatistics, "", false, usage.DisableCollectDatabaseSpecificsStatistics)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableCollectDatabaseSpecificsStatistics, flag.DisableCollectDatabaseSpecificsStatistics)

	cmd.Flags().BoolVarP(&opts.enableDataExplorer, flag.EnableDataExplorer, "", false, usage.EnableDataExplorer)
	cmd.Flags().BoolVarP(&opts.disableDataExplorer, flag.DisableDataExplorer, "", false, usage.DisableDataExplorer)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableDataExplorer, flag.DisableDataExplorer)

	cmd.Flags().BoolVarP(&opts.enablePerformanceAdvisor, flag.EnablePerformanceAdvisor, "", false, usage.EnablePerformanceAdvisor)
	cmd.Flags().BoolVarP(&opts.disablePerformanceAdvisor, flag.DisablePerformanceAdvisor, "", false, usage.DisablePerformanceAdvisor)
	cmd.MarkFlagsMutuallyExclusive(flag.EnablePerformanceAdvisor, flag.DisablePerformanceAdvisor)

	cmd.Flags().BoolVarP(&opts.enableSchemaAdvisor, flag.EnableSchemaAdvisor, "", false, usage.EnableSchemaAdvisor)
	cmd.Flags().BoolVarP(&opts.disableSchemaAdvisor, flag.DisableSchemaAdvisor, "", false, usage.DisableSchemaAdvisor)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableSchemaAdvisor, flag.DisableSchemaAdvisor)

	cmd.Flags().BoolVarP(&opts.enableRealtimePerformancePanel, flag.EnableRealtimePerformancePanel, "", false, usage.EnableRealtimePerformancePanel)
	cmd.Flags().BoolVarP(&opts.disableRealtimePerformancePanel, flag.DisableRealtimePerformancePanel, "", false, usage.DisableRealtimePerformancePanel)
	cmd.MarkFlagsMutuallyExclusive(flag.EnableRealtimePerformancePanel, flag.DisableRealtimePerformancePanel)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
