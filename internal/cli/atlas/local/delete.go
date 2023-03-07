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
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.OutputOpts
	name string
}

func (opts *DeleteOpts) Run(ctx context.Context) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	foundID := ""
	for _, c := range containers {
		for _, n := range c.Names {
			if strings.EqualFold(n, fmt.Sprintf("/%s", opts.name)) {
				foundID = c.ID
			}
		}
	}

	if foundID == "" {
		return fmt.Errorf("%s not found", opts.name)
	}

	return cli.ContainerRemove(ctx, foundID, types.ContainerRemoveOptions{Force: true})
}

// atlas local delete <instanceName>.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{}
	cmd := &cobra.Command{
		Use:     "delete <instanceName>",
		Aliases: []string{"rm"},
		Short:   "Deletes an instance.",
		Args:    require.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run(cmd.Context())
		},
	}

	return cmd
}
