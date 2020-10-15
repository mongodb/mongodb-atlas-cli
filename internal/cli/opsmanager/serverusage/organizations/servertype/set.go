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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

var setTemplate = `NAME	PATH	WT COMPRESSION	MMAPV1 COMPRESSION
{{.ID}}	{{.StorePath}}	{{.WTCompressionSetting}}	{{.MMAPV1CompressionSetting}}
`

type SetOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	store      store.OrganizationServerTypeUpdater
	serverType string
}

func (opts *SetOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *SetOpts) Run() error {
	r, err := opts.store.UpdateOrganizationServerType(opts.ConfigOrgID(), opts.newServerType())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *SetOpts) newServerType() *opsmngr.ServerType {
	return &opsmngr.ServerType{
		Name: opts.serverType,
	}
}

// mongocli ops-manager serverUsage organization(s) serverType set type [--orgId orgId]
func SetBuilder() *cobra.Command {
	opts := &SetOpts{}
	cmd := &cobra.Command{
		Use:       "set",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"DEV_SERVER", "TEST_SERVER", "PRODUCTION_SERVER", "RAM_POOL"},
		Short:     get,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunEOrg(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), setTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.serverType = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
