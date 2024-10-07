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
	"errors"
	"fmt"
	"io"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/cobra"
)

type whoOpts struct {
	OutWriter   io.Writer
	authSubject string
	authType    string
}

func (opts *whoOpts) Run() error {
	_, _ = fmt.Fprintf(opts.OutWriter, "Logged in as %s %s\n", opts.authSubject, opts.authType)

	return nil
}

var ErrUnauthenticated = errors.New("not logged in with an Atlas account or API key")

func AccountWithAccessToken() (string, error) {
	if config.AccessToken() == "" {
		return "", ErrUnauthenticated
	}

	return config.AccessTokenSubject()
}

func authTypeAndSubject() (string, string, error) {
	if config.PublicAPIKey() != "" {
		return "key", config.PublicAPIKey(), nil
	}

	if subject, err := AccountWithAccessToken(); err == nil {
		return "account", subject, nil
	}

	return "", "", ErrUnauthenticated
}

func WhoAmIBuilder() *cobra.Command {
	opts := &whoOpts{}

	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Verifies and displays information about your authentication state.",
		Example: `  # See the logged account:
  atlas auth whoami
`,
		PreRun: func(cmd *cobra.Command, _ []string) {
			opts.OutWriter = cmd.OutOrStdout()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			var err error
			if opts.authType, opts.authSubject, err = authTypeAndSubject(); err != nil {
				return err
			}

			return opts.Run()
		},
		Args: require.NoArgs,
	}

	return cmd
}
