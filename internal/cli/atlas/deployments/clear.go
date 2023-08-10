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

package deployments

import (
	"context"
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type ClearOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	debug bool
}

var clearTemplate = `local environment stopped and cleared
`

func (opts *ClearOpts) Run(_ context.Context) error {
	if err := runDockerCompose(opts.debug, "down", "-v"); err != nil {
		return err
	}

	mmsConfigFilename, err := mmsConfigPath()
	if err != nil {
		return err
	}
	_ = os.Remove(mmsConfigFilename)

	return opts.Print(localData)
}

// atlas local clear.
func ClearBuilder() *cobra.Command {
	opts := &ClearOpts{}
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Stops local instance and deletes stored data.",
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), clearTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
