// Copyright 2021 MongoDB Inc
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

package livemigrations

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

var cutoverTemplate = "Cutover process successfully started.\n"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=cutover_mock_test.go -package=livemigrations . LiveMigrationCutoverCreator

type LiveMigrationCutoverCreator interface {
	CreateLiveMigrationCutover(string, string) error
}

type CutoverOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store           LiveMigrationCutoverCreator
	liveMigrationID string
}

func (opts *CutoverOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CutoverOpts) Run() error {
	err := opts.store.CreateLiveMigrationCutover(opts.ConfigProjectID(), opts.liveMigrationID)
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

// atlas liveMigrations|lm cutover [--liveMigrationID liveMigrationId] [--projectId projectId].
func CutoverBuilder() *cobra.Command {
	opts := &CutoverOpts{}
	cmd := &cobra.Command{
		Use:   "cutover",
		Short: "Start the cutover for a push live migration and confirm when the cutover completes. When the cutover completes, the application completes the live migration process and stops synchronizing with the source cluster.",
		Long:  `To migrate using scripts, use mongomirror instead of the Atlas CLI. To learn more about mongomirror, see https://www.mongodb.com/docs/atlas/reference/mongomirror/.`,
		Annotations: map[string]string{
			"output": cutoverTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), cutoverTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}
	opts.AddProjectOptsFlags(cmd)
	cmd.Flags().StringVar(&opts.liveMigrationID, flag.LiveMigrationID, "", usage.LiveMigrationID)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.LiveMigrationID)

	return cmd
}
