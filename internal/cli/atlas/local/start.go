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

package local

import (
	"context"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

const speed = 100 * time.Millisecond

type StartOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	s *spinner.Spinner
}

var startTemplate = `local environment started at {{.ConnectionString}}
`

func (opts *StartOpts) Run(ctx context.Context) error {
	if opts.s != nil {
		opts.s.Start()
	}

	defer func() {
		if opts.s != nil {
			opts.s.Stop()
		}
	}()

	mongotHome, err := mongotHome()
	if err != nil {
		return err
	}
	cmd := exec.Command("make", "docker.up")
	cmd.Dir = mongotHome
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return err
	}

	if opts.s != nil {
		opts.s.Stop()
	}
	return opts.Print(localData)
}

// atlas local start.
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	if opts.IsTerminal() {
		opts.s = spinner.New(spinner.CharSets[9], speed)
	}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a local instance.",
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), startTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
