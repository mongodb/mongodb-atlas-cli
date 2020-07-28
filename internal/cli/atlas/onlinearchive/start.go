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

package onlinearchive

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type StartOpts struct {
	cli.GlobalOpts
	id          string
	clusterName string
	store       store.OnlineArchiveUpdater
}

func (opts *StartOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var startTemplate = "Online archive '{{.ID}}' started.\n"

func (opts *StartOpts) Run() error {
	paused := false
	archive := &atlas.OnlineArchive{
		ID:     opts.id,
		Paused: &paused,
	}
	r, err := opts.store.UpdateOnlineArchive(opts.ConfigProjectID(), opts.clusterName, archive)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), startTemplate, r)
}

// mongocli atlas cluster(s) onlineArchive(s) start <ID> [--clusterName name][--projectId projectId]
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:   "start <ID>",
		Short: description.StartOnlineArchive,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
