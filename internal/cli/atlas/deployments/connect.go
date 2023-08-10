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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/spf13/cobra"
)

type ConnectOpts struct{}

func (*ConnectOpts) Run(_ context.Context) error {
	return mongosh.Run(localUser, localPassword, localURI)
}

// atlas local connect.
func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Connects to the local instance.",
		Args:  require.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	return cmd
}
