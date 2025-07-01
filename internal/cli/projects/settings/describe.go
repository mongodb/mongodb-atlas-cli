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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=settings . ProjectSettingsDescriber

type ProjectSettingsDescriber interface {
	ProjectSettings(string) (*atlasv2.GroupSettings, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store ProjectSettingsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `COLLECT DATABASE SPECIFICS STATISTICS ENABLED	DATA EXPLORER ENABLED	DATA EXPLORER GEN AI FEATURES ENABLED	DATA EXPLORER GEN AI SAMPLE DOCUMENT PASSING ENABLED	PERFORMANCE ADVISOR ENABLED	REALTIME PERFORMANCE PANEL ENABLED	SCHEMA ADVISOR ENABLED
{{.IsCollectDatabaseSpecificsStatisticsEnabled}}	{{.IsDataExplorerEnabled}}	{{.IsDataExplorerGenAIFeaturesEnabled}}	{{.IsDataExplorerGenAISampleDocumentPassingEnabled}}	{{.IsPerformanceAdvisorEnabled}}	{{.IsRealtimePerformancePanelEnabled}}	{{.IsSchemaAdvisorEnabled}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.ProjectSettings(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas projects(s) settings describe [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:         "describe",
		Aliases:     []string{"get"},
		Short:       "Retrieve details for settings to the specified project.",
		Annotations: map[string]string{"output": describeTemplate},
		Example: `  # This example uses the profile named "myprofile" for accessing Atlas.
  atlas projects settings describe -P myprofile --projectId 5e2211c17a3e5a48f5497de3`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
