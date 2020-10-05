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

package blockstore

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

var createTemplate = "Blockstore configuration '{{.ID}}' created.\n"

type CreateOpts struct {
	cli.OutputOpts
	assignment           bool
	encryptedCredentials bool
	ssl                  bool
	id                   string
	label                []string
	writeConcern         string
	uri                  string
	loadFactor           int64
	maxCapacityGB        int64
	store                store.BlockstoresCreater
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateBlockstore(opts.newBackupStore())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *CreateOpts) newBackupStore() *opsmngr.BackupStore {
	backupStore := &opsmngr.BackupStore{
		AdminBackupConfig: opsmngr.AdminBackupConfig{
			ID:           opts.id,
			URI:          opts.uri,
			WriteConcern: opts.writeConcern,
			Labels:       opts.label,
		},
	}

	if opts.ssl {
		backupStore.SSL = &opts.ssl
	}

	if opts.encryptedCredentials {
		backupStore.EncryptedCredentials = &opts.encryptedCredentials
	}

	if opts.assignment {
		backupStore.AssignmentEnabled = &opts.assignment
	}

	if opts.maxCapacityGB != 0 {
		backupStore.MaxCapacityGB = &opts.maxCapacityGB
	}

	if opts.loadFactor != 0 {
		backupStore.LoadFactor = &opts.loadFactor
	}

	return backupStore
}

// mongocli ops-manager admin backup blockstore(s) create [--assignment][--encryptedCredentials][--id id][
// --label label][--loadFactor loadFactor][--maxCapacityGB maxCapacityGB][--uri uri][--ssl][--writeConcern writeConcern]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:   "create",
		Short: create,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.assignment, flag.Assignment, false, usage.Assignment)
	cmd.Flags().BoolVar(&opts.encryptedCredentials, flag.EncryptedCredentials, false, usage.EncryptedCredentials)
	cmd.Flags().StringVar(&opts.id, flag.ID, "", usage.BlockstoreID)
	cmd.Flags().StringSliceVar(&opts.label, flag.Label, []string{}, usage.Label)
	cmd.Flags().Int64Var(&opts.loadFactor, flag.LoadFactor, 0, usage.LoadFactor)
	cmd.Flags().Int64Var(&opts.maxCapacityGB, flag.MaxCapacityGB, 0, usage.MaxCapacityGB)
	cmd.Flags().StringVar(&opts.uri, flag.URI, "", usage.BlockstoreURI)
	cmd.Flags().BoolVar(&opts.ssl, flag.SSL, false, usage.BlockstoreSSL)
	cmd.Flags().StringVar(&opts.writeConcern, flag.WriteConcern, "", usage.WriteConcern)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.ID)
	_ = cmd.MarkFlagRequired(flag.URI)

	return cmd
}
