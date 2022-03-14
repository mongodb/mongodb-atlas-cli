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

package logs

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type KeyProviderListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	file string
}

func (opts *KeyProviderListOpts) initStore() func() error {
	// Keeping for now, not sure if needed
	return func() error {
		return nil
	}
}

func (opts *KeyProviderListOpts) Run() error {
	// Run the command. For now printing the file path.
	return opts.Print(opts.file)
}

// atlas logs listKeyProvider --file <encryptedLogFile>.
func KeyProvidersListBuilder() *cobra.Command {
	opts := &KeyProviderListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Lists all key provider configurations found in the encrypted log file.",
		Example: `
  $ atlas listKeyProvider --file log.gz`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore(),
				opts.InitOutput(cmd.OutOrStdout(), ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.file, flag.File, flag.FileShort, "", usage.EncryptedLogFile)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.File)
	_ = cmd.MarkFlagFilename(flag.File)
	return cmd
}
