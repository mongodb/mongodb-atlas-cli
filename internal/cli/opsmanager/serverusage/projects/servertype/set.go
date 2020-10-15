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

package servertype

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type SetOpts struct {
	cli.GlobalOpts
	store      store.ProjectServerTypeUpdater
	serverType string
}

func (opts *SetOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *SetOpts) Run() error {
	err := opts.store.UpdateProjectServerType(opts.ConfigProjectID(), opts.newServerTypeRequest())
	if err != nil {
		return err
	}
	fmt.Printf("Serve type %s correctly set\n", opts.serverType)
	return nil
}

func (opts *SetOpts) newServerTypeRequest() *opsmngr.ServerTypeRequest {
	return &opsmngr.ServerTypeRequest{
		ServerType: &opsmngr.ServerType{
			Name: opts.serverType,
		},
	}
}

// mongocli ops-manager serverUsage projects(s) serverType set type [--projectId projectId]
func SetBuilder() *cobra.Command {
	opts := &SetOpts{}
	cmd := &cobra.Command{
		Use:       "set",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"DEV_SERVER", "TEST_SERVER", "PRODUCTION_SERVER", "RAM_POOL"},
		Short:     get,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.serverType = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
