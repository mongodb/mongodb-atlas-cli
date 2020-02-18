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
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/json"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasWhitelistListOpts struct {
	*globalOpts
	pageNum      int
	itemsPerPage int
	store        store.ProjectIPWhitelistLister
}

func (opts *atlasWhitelistListOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasWhitelistListOpts) Run() error {
	listOpts := opts.newListOptions()
	result, err := opts.store.ProjectIPWhitelists(opts.ProjectID(), listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasWhitelistListOpts) newListOptions() *atlas.ListOptions {
	return &atlas.ListOptions{
		PageNum:      opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
	}
}

// mcli atlas whitelist(s) list --projectId projectId [--page N] [--limit N]
func AtlasWhitelistListBuilder() *cobra.Command {
	opts := &atlasWhitelistListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List Atlas whitelist for a project.",
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
