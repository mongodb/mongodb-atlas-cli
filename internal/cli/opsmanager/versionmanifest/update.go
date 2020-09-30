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

package versionmanifest

import (
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const updateTemplate = "Version manifest updated.\n"

type UpdateOpts struct {
	cli.OutputOpts
	versionManifest string
	store           store.VersionManifestUpdater
	storeStaticPath store.VersionManifestGetter
}

func (opts *UpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	if err != nil {
		return err
	}
	opts.storeStaticPath, err = store.NewStaticPath(config.Default())
	return err
}

func (opts *UpdateOpts) Run() error {
	versionManifest, err := opts.storeStaticPath.GetVersionManifest(opts.version())
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateVersionManifest(versionManifest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) version() string {
	if strings.Contains(opts.versionManifest, ".json") {
		return opts.versionManifest
	}
	return opts.versionManifest + ".json"
}

// mongocli om versionManifest(s) update <version>
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	opts.Template = updateTemplate
	cmd := &cobra.Command{
		Use:   "update <version>",
		Short: update,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.versionManifest = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
