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
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=fields_type_mock_test.go -package=settings . MatcherFieldsLister

type MatcherFieldsLister interface {
	MatcherFields() ([]string, error)
}

type FieldsTypeOpts struct {
	cli.OutputOpts
	store MatcherFieldsLister
}

func (opts *FieldsTypeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var matcherFieldsTemplate = "{{range valueOrEmptySlice .}}{{.}}\n{{end}}"

func (opts *FieldsTypeOpts) Run() error {
	r, err := opts.store.MatcherFields()
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas alerts config(s) fields type.
func FieldsTypeBuilder() *cobra.Command {
	opts := &FieldsTypeOpts{}
	opts.Template = matcherFieldsTemplate
	cmd := &cobra.Command{
		Use:     "type",
		Short:   "Return all available field types that the matcherFieldName option accepts when you create or update an alert configuration.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"types"},
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of accepted field types for the matchersFieldName option:
  atlas alerts settings fields type --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddOutputOptFlags(cmd)

	return cmd
}
