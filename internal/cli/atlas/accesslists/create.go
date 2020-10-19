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

package accesslists

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	cidrBlock        = "cidrBlock"
	ipAddress        = "ipAddress"
	awsSecurityGroup = "awsSecurityGroup"
	createTemplate   = "Created new IP access list.\n"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	entry       string
	entryType   string
	comment     string
	deleteAfter string
	store       store.ProjectIPAccessListCreator
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	entry := opts.newWhitelist()
	r, err := opts.store.CreateProjectIPAccessList(entry)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newWhitelist() *atlas.ProjectIPWhitelist {
	entry := &atlas.ProjectIPWhitelist{
		GroupID:         opts.ConfigProjectID(),
		Comment:         opts.comment,
		DeleteAfterDate: opts.deleteAfter,
	}
	switch opts.entryType {
	case cidrBlock:
		entry.CIDRBlock = opts.entry
	case ipAddress:
		entry.IPAddress = opts.entry
	case awsSecurityGroup:
		entry.AwsSecurityGroup = opts.entry
	}
	return entry
}

// mongocli atlas accessList(s) create <entry> --type cidrBlock|ipAddress|awsSecurityGroup [--comment comment] [--projectId projectId]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <entry>",
		Short: create,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.entry = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.entryType, flag.Type, ipAddress, usage.AccessListType)
	cmd.Flags().StringVar(&opts.comment, flag.Comment, "", usage.Comment)
	cmd.Flags().StringVar(&opts.deleteAfter, flag.DeleteAfter, "", usage.AccessListsDeleteAfter)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
