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

package search

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=delete_mock_test.go -package=search . Deleter

type Deleter interface {
	DeleteSearchIndex(string, string, string) error
}

type DeleteOpts struct {
	cli.ProjectOpts
	*cli.DeleteOpts
	clusterName string
	store       Deleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteSearchIndex, opts.ConfigProjectID(), opts.clusterName)
}

// atlas cluster(s) search(s) index(es) delete <id> [--clusterName name][--projectId projectId][--force].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Index '%s' deleted\n", "Index not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <indexId>",
		Aliases: []string{"rm"},
		Short:   "Delete the specified search index from the specified cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"indexIdDesc": "ID of the index.",
			"output":      opts.SuccessMessage(),
		},
		Example: `  # Delete the search index with the ID 5f2099cd683fc55fbb30bef6 for the cluster named myCluster without requiring confirmation:
  atlas clusters search indexes delete 5f2099cd683fc55fbb30bef6 --clusterName myCluster --force`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			if err := validate.ObjectID(args[0]); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	opts.AddProjectOptsFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
