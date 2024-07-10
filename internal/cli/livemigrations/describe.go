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

type DescribeOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	liveMigrationID string
	store           store.LiveMigrationDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	LAGTIME	READY FOR CUTOVER	STATUS
{{.Id}}	{{.LagTimeSeconds}}	{{.ReadyForCutover}}	{{.Status}}`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.LiveMigrationDescribe(opts.ConfigProjectID(), opts.liveMigrationID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas liveMigrations|lm describe --liveMigrationId <liveMigrationId> [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   "Return a push live migration job.",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.liveMigrationID, flag.LiveMigrationID, "", usage.LiveMigrationID)

	_ = cmd.MarkFlagRequired(flag.LiveMigrationID)

	return cmd
}
