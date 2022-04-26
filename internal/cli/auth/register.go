// Copyright 2022 MongoDB Inc
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

package auth

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/cobra"
)

type registerOpts struct {
	loginOpts
}

func (opts *registerOpts) Run() error {
	_, _ = fmt.Fprintf(opts.OutWriter, "Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.\n")
	// TODO: CLOUDP-120669
	return nil
}

func RegisterBuilder() *cobra.Command {
	opts := &registerOpts{}
	cmd := &cobra.Command{
		Use:    "register",
		Short:  "Register with MongoDB Atlas.",
		Hidden: true,
		Example: fmt.Sprintf(`  To start the interactive setup:
  $ %s auth register
`, config.BinName()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if hasUserProgrammaticKeys() {
				return fmt.Errorf(`you have already set the programmatic keys for this profile. 

Run '%s auth register --profile <profileName>' to use your username and password on a new profile`, config.BinName())
			}

			opts.OutWriter = cmd.OutOrStdout()
			opts.config = config.Default()
			if config.OpsManagerURL() != "" {
				opts.OpsManagerURL = config.OpsManagerURL()
			}
			return opts.initFlow()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.isGov, "gov", false, "Register to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.noBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&opts.skipConfig, "skipConfig", false, "Skip profile configuration.")

	return cmd
}
