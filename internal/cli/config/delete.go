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

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
}

func (opts *DeleteOpts) Run() error {
	config.SetName(opts.Entry)
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
		Use:     "delete <profileName>",
		Aliases: []string{"rm"},
		Short:   "Delete a profile.",
		Long:    `This command removes the specified profile from the MongoDB CLI configuration file. 
For more information on the configuration file, see the documentation: https://docs.mongodb.com/mongocli/stable/configure/configuration-file/ `,
		Example:  `Delete a profile called myProfile:
$ mongocli config delete myProfile`,
		Annotations: map[string]string{
			"args":            "profileName",
			"profileNameDesc": "Name of the profile to delete. Specify 'default' to delete the default profile.",
		},
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Entry = args[0]
			if !config.Exists(opts.Entry) {
				return fmt.Errorf("profile %v does not exist", opts.Entry)
			}

			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
