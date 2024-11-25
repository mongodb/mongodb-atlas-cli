// Copyright 2024 MongoDB Inc
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

package cli

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/cobra"
)

type CompositeOpts struct {
	stdOut  *bytes.Buffer
	stdErr  *bytes.Buffer
	stdIn   io.Reader
	rootCmd *cobra.Command
}

func (opts *CompositeOpts) InitComposite(cmd *cobra.Command) {
	opts.rootCmd = cmd
	if opts.rootCmd == nil {
		return
	}
	for opts.rootCmd.Parent() != nil {
		opts.rootCmd = opts.rootCmd.Parent()
	}
}

func (opts *CompositeOpts) StdOut() io.Reader {
	return opts.stdOut
}

func (opts *CompositeOpts) StdErr() io.Reader {
	return opts.stdErr
}

func (opts *CompositeOpts) SetStdIn(in io.Reader) {
	opts.stdIn = in
}

func (opts *CompositeOpts) RunCommand(ctx context.Context, args ...string) error {
	opts.stdOut = &bytes.Buffer{}
	opts.stdErr = &bytes.Buffer{}
	opts.rootCmd.SetOut(opts.stdOut)
	opts.rootCmd.SetErr(opts.stdErr)
	if opts.stdIn != nil {
		opts.rootCmd.SetIn(opts.stdIn)
	}

	log.Debugf("running command '%s'", "atlas "+strings.Join(args, " "))

	currentArgs := args

	if !strings.EqualFold(config.Default().Name(), "default") {
		currentArgs = append(currentArgs, "--profile", config.Default().Name())
	}

	opts.rootCmd.SetArgs(currentArgs)
	return opts.rootCmd.ExecuteContext(ctx)
}
