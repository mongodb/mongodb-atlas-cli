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

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=disable_mock_test.go -package=settings . AlertConfigurationDisabler

type AlertConfigurationDisabler interface {
	DisableAlertConfiguration(string, string) (*atlasv2.GroupAlertsConfig, error)
}

type DisableOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	alertID string
	store   AlertConfigurationDisabler
}

func (opts *DisableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var disableTemplate = "Alert configuration '{{.ID}}' disabled\n"

func (opts *DisableOpts) Run() error {
	r, err := opts.store.DisableAlertConfiguration(opts.ConfigProjectID(), opts.alertID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas alerts disable <ID> --projectId projectId.
func DisableBuilder() *cobra.Command {
	opts := new(DisableOpts)
	cmd := &cobra.Command{
		Use:   "disable <alertConfigId>",
		Short: "Disables one alert configuration for the specified project.",
		Args:  require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), disableTemplate),
			)
		},
		Annotations: map[string]string{
			"alertConfigIdDesc": "ID of the alert configuration you want to disable.",
			"output":            disableTemplate,
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}
	cmd.OutOrStdout()
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
