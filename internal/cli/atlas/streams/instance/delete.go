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

package instance

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.StreamsDeleter
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteStream, opts.ConfigProjectID(), opts.Entry)
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// DeleteBuilder
// atlas streams instance delete [name].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Atlas Streams processor instance '%s' deleted\n", "Atlas Streams processor instance not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete an Atlas Stream Processor Instance",
		Long: `The command prompts you to confirm the operation when you run the command without the --force option.

An Atlas Streams processor instance with running processors cannot be deleted without stopping the processes first.
` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  # Remove an Atlas Streams processor instance after prompting for a confirmation:
  %[1]s streams delete myProcessorInstance

  # Remove an Atlas Streams processor instance named myProcessorInstance without requiring confirmation:
  %[1]s streams delete myProcessorInstance --force`, cli.ExampleAtlasEntryPoint()),
		Aliases: []string{"rm"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the Atlas Streams instance.",
			"output":   opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("the Atlas Streams Processor instance name is missing")
			}

			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
