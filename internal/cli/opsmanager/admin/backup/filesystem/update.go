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

package filesystem

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/opsmanager/admin/backupstore"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var updateTemplate = "File System configuration '{{.ID}}' updated.\n"

type UpdateOpts struct {
	cli.OutputOpts
	backupstore.AdminOpts
	store store.FileSystemsUpdater
}

func (opts *UpdateOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateFileSystems(opts.NewFileSystemConfiguration())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// mongocli ops-manager admin backup fileSystem(s) update <ID> [--assignment][--encryptedCredentials]
// [--label label][--loadFactor loadFactor][--id ID][--storePath storePath][--mmapv1CompressionSetting mmapv1CompressionSetting]
// [--wtCompressionSetting wtCompressionSetting]
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	opts.Template = updateTemplate
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: update,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.ID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Assignment, flag.Assignment, false, usage.Assignment)
	cmd.Flags().BoolVar(&opts.EncryptedCredentials, flag.EncryptedCredentials, false, usage.EncryptedCredentials)
	cmd.Flags().StringVar(&opts.ID, flag.ID, "", usage.BlockstoreID)
	cmd.Flags().StringSliceVar(&opts.Label, flag.Label, []string{}, usage.Label)
	cmd.Flags().Int64Var(&opts.LoadFactor, flag.LoadFactor, 0, usage.LoadFactor)
	cmd.Flags().StringVar(&opts.MMAPV1CompressionSetting, flag.MMAPV1CompressionSetting, "", usage.MMAPV1CompressionSetting)
	cmd.Flags().StringVar(&opts.WTCompressionSetting, flag.WTCompressionSetting, "", usage.WTCompressionSetting)
	cmd.Flags().StringVar(&opts.StorePath, flag.StorePath, "", usage.StorePath)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
