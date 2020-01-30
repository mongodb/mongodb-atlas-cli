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

const (
	cidrBlock = "cidrBlock"
	ipAddress = "ipAddress"
)

type atlasWhitelistCreateOpts struct {
	*globalOpts
	entry     string
	entryType string
	comment   string
	store     store.ProjectIPWhitelistCreator
}

func (opts *atlasWhitelistCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New()

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *atlasWhitelistCreateOpts) Run() error {
	entry := opts.newWhitelist()
	result, err := opts.store.CreateProjectIPWhitelist(entry)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasWhitelistCreateOpts) newWhitelist() *atlas.ProjectIPWhitelist {
	projectIPWhitelist := &atlas.ProjectIPWhitelist{
		GroupID: opts.ProjectID(),
		Comment: opts.comment,
	}
	switch opts.entryType {
	case cidrBlock:
		projectIPWhitelist.CIDRBlock = opts.entry
	case ipAddress:
		projectIPWhitelist.IPAddress = opts.entry
	}
	return projectIPWhitelist
}

// mcli atlas whitelist(s) create value --type cidrBlock|ipAddress [--comment comment] [--projectId projectId]
func AtlasWhitelistCreateBuilder() *cobra.Command {
	opts := &atlasWhitelistCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create [entry]",
		Short: "Create a project IP whitelist.",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.entry = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.entryType, flags.Type, ipAddress, usage.WhitelistType)
	cmd.Flags().StringVar(&opts.comment, flags.Comment, "", usage.Comment)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
