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

package cli

import (
	"fmt"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasClustersWatchOpts struct {
	globalOpts
	name  string
	store store.ClusterDescriber
}

func (opts *atlasClustersWatchOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *atlasClustersWatchOpts) Run() error {
	for {
		result, err := opts.store.Cluster(opts.ProjectID(), opts.name)
		if err != nil {
			return err
		}
		if result.StateName == "IDLE" {
			fmt.Printf("\nCluster available at: %s\n", result.MongoURIWithOptions)
			break
		}
		fmt.Print(".")
		time.Sleep(4 * time.Second)
	}

	return nil
}

// mongocli atlas cluster(s) watch [name] [--projectId projectId]
func AtlasClustersWatchBuilder() *cobra.Command {
	opts := &atlasClustersWatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch [name]",
		Short: description.WatchCluster,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
