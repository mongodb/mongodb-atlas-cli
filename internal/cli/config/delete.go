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
	"context"
	"fmt"
	"os"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/auth"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/workflows"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
}

func (opts *DeleteOpts) Run(ctx context.Context) error {
	logout := auth.LogoutBuilder()

	var newArgs []string
	_, _ = log.Debugf("Removing flags and args from original args %s\n", os.Args)

	newArgs, err := workflows.RemoveFlagsAndArgs(nil, map[string]bool{opts.Entry: true}, os.Args)
	if err != nil {
		return err
	}

	logout.SetArgs(newArgs)

	// Send profile as a context value to the logout command
	if err := config.SetName(opts.Entry); err != nil {
		return err
	}

	ctx = config.WithProfile(ctx, config.Default())

	_, _ = log.Debugf("Executing logout with args '%s' and profile '%s'", newArgs, opts.Entry)
	_, err = logout.ExecuteContextC(ctx)
	return err
}

func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Profile '%s' deleted\n", "Profile not deleted"),
	}
	cmd := &cobra.Command{
		Use:        "delete <name>",
		Aliases:    []string{"rm"},
		Short:      "Delete a profile.",
		Args:       require.ExactArgs(1),
		Deprecated: "Please use the 'atlas auth logout' command instead.",
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

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
