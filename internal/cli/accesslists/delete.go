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

package accesslists

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.ProjectIPAccessListDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteProjectIPAccessList, opts.ConfigProjectID())
}

// atlas accessList delete <entry> --force.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Project access list entry '%s' deleted\n", "Project access list entry not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <entry>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified IP access list entry from your project.",
		Long: `The command, when run without the force option, prompts you to confirm the operation.

` + fmt.Sprintf(usage.RequiredRole, "Read Write"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"entryDesc": "The IP address, CIDR address, or AWS security group ID that you want to remove from the access list.",
			"output":    opts.SuccessMessage(),
		},
		Example: `  # Remove the IP address 192.0.2.0 from the access list for the project with the ID 5e2211c17a3e5a48f5497de3 after prompting for a confirmation:
  atlas accessLists delete 192.0.2.0 --projectId 5e2211c17a3e5a48f5497de3
  # Remove the IP address 192.0.2.0 from the access list for the project with the ID 5e2211c17a3e5a48f5497de3 without requiring confirmation:
  atlas accessLists delete 192.0.2.0 --projectId 5e2211c17a3e5a48f5497de3 --force`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
