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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	DefaultPage      = 1
	DefaultPageLimit = 100
)

type ListOpts struct {
	PageNum      int
	ItemsPerPage int
	OmitCount    bool
}

func (opts *ListOpts) NewAtlasListOptions() *store.ListOptions {
	return &store.ListOptions{
		PageNum:      opts.PageNum,
		ItemsPerPage: opts.ItemsPerPage,
		IncludeCount: !opts.OmitCount,
	}
}

func (opts *ListOpts) AddListOptsFlags(cmd *cobra.Command) {
	opts.AddListOptsFlagsWithoutOmitCount(cmd)
	cmd.Flags().BoolVar(&opts.OmitCount, flag.OmitCount, false, usage.OmitCount)
}

func (opts *ListOpts) AddListOptsFlagsWithoutOmitCount(cmd *cobra.Command) {
	cmd.Flags().IntVar(&opts.PageNum, flag.Page, DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, DefaultPageLimit, usage.Limit)
}
