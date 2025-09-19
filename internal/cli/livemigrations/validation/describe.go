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

package validation

import (
	"context"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=validation . LiveMigrationValidationsDescriber

type LiveMigrationValidationsDescriber interface {
	GetValidationStatus(string, string) (*atlasv2.LiveImportValidation, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	validationID string
	store        LiveMigrationValidationsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	PROJECT ID	SOURCE PROJECT ID	STATUS
{{.Id}}	{{.GroupId}}	{{.SourceGroupId}}	{{.Status}}`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.GetValidationStatus(opts.ConfigProjectID(), opts.validationID)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas liveMigrations|lm validation(s) describe|get <validationId> [--projectId projectId].

func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:         "describe",
		Aliases:     []string{"get"},
		Short:       "Return one validation job.",
		Annotations: map[string]string{"output": describeTemplate},
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
	cmd.Flags().StringVar(&opts.validationID, flag.LiveMigrationValidationID, "", usage.LiveMigrationValidationID)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
