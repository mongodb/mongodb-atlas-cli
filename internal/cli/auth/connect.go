// Copyright 2026 MongoDB Inc
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
	"context"
	"errors"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/spf13/cobra"
)

type ConnectOpts struct {
	config SetSaver
}

func (opts *ConnectOpts) Run(ctx context.Context) error {
	opts.config.SetAuthType(config.UserDelegation)

	if err := opts.config.Save(); err != nil {
		return err
	}

	return errors.New("connect is not yet implemented")
}

func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}

	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect to MongoDB Atlas using the dedicated authorization server.",
		Long:  `This command authenticates with MongoDB Atlas via the dedicated OAuth Authorization Server at authorize.mongodb.com.`,
		Example: `  # Connect to your MongoDB Atlas account:
  atlas auth connect
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.config = config.Default()
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
		Args: require.NoArgs,
	}

	return cmd
}
