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
	"os"
	"os/exec"
	"syscall"

	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	mongo = "mongo"
)

type atlasClustersConnectOpts struct {
	*globalOpts
	name     string
	username string
	store    store.ClusterDescriber
}

func (opts *atlasClustersConnectOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasClustersConnectOpts) Run() error {
	result, err := opts.store.Cluster(opts.ProjectID(), opts.name)

	if err != nil {
		return err
	}
	binary, err := exec.LookPath(mongo)
	if err != nil {
		return err
	}
	args := []string{mongo, result.MongoURIWithOptions, "-u", opts.username}
	env := os.Environ()
	err = syscall.Exec(binary, args, env)
	if err != nil {
		return err
	}

	return nil
}

// mcli atlas cluster(s) connect [name] --username <username> --password <password> --projectId <projectId>
func AtlasClustersConnectBuilder() *cobra.Command {
	opts := &atlasClustersConnectOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "connect [name]",
		Short: "Connect to an Atlas cluster.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.username, flags.Username, "", usage.Username)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
