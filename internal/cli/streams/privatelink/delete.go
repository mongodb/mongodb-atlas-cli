// Copyright 2025 MongoDB Inc
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

package privatelink

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	successDeleteTemplate = "Atlas Stream Processing PrivateLink endpoint '%s' deleted.\n"
	failDeleteTemplate    = "Atlas Stream Processing PrivateLink endpoint not deleted"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=delete_mock_test.go -package=privatelink . Deleter

type Deleter interface {
	DeletePrivateLinkEndpoint(projectID, connectionID string) error
}

type DeleteOpts struct {
	cli.ProjectOpts
	*cli.DeleteOpts
	store Deleter
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeletePrivateLinkEndpoint, opts.ConfigProjectID())
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams privateLink deleted <connectionID>
// Deletes a PrivateLink endpoint.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts(successDeleteTemplate, failDeleteTemplate),
	}
	cmd := &cobra.Command{
		Use:     "delete <connectionID>",
		Aliases: []string{"rm"},
		Short:   "Deletes an Atlas Stream Processing PrivateLink endpoint.",
		Long:    fmt.Sprintf(usage.RequiredOneOfRoles, commandRoles),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"connectionIDDesc": "ID of the PrivateLink endpoint.",
			"output":           successDeleteTemplate,
		},
		Example: `# delete an Atlas Stream Processing PrivateLink endpoint:
  atlas streams privateLink delete 5e2211c17a3e5a48f5497de3
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	opts.AddProjectOptsFlags(cmd)

	return cmd
}
