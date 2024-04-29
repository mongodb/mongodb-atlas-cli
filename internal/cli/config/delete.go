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

package config

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
}

func (opts *DeleteOpts) Run() error {
	if !opts.Confirm {
		return nil
	}

	if err := config.SetName(opts.Entry); err != nil {
		return err
	}

	if err := config.Delete(); err != nil {
		return err
	}
	fmt.Printf(opts.SuccessMessage(), opts.Entry)
	return nil
}

func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Profile '%s' deleted\n", "Profile not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <name>",
		Aliases: []string{"rm"},
		Short:   "Delete a profile.",
		Args:    require.ExactArgs(1),
		Example: `  # Delete the default profile configuration:
  atlas config delete default

  # Skip the confirmation question and delete the default profile configuration:
  atlas config delete default --force`,
		Annotations: map[string]string{
			"nameDesc": "Name of the profile.",
			"output":   opts.SuccessMessage(),
		},
		PreRunE: func(_ *cobra.Command, args []string) error {
			opts.Entry = args[0]
			if !config.Exists(opts.Entry) {
				return fmt.Errorf("profile %v does not exist", opts.Entry)
			}

			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
