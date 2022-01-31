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
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	cidrBlock                = "cidrBlock"
	ipAddress                = "ipAddress"
	awsSecurityGroup         = "awsSecurityGroup"
	createTemplate           = "Created new IP access list.\n"
	currentIPAddressNotFound = `not able to find your public IP address. Please providing the desired IP address for this command`
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	entry       string
	entryType   string
	comment     string
	deleteAfter string
	currentIP   bool
	store       store.ProjectIPAccessListCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) Run() error {
	entry := opts.newProjectIPAccessList()
	r, err := opts.store.CreateProjectIPAccessList(entry)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newProjectIPAccessList() []*atlas.ProjectIPAccessList {
	entry := &atlas.ProjectIPAccessList{
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
	return []*atlas.ProjectIPAccessList{entry}
}

func iPAddress() (string, error) {
	if publicIP := store.IPAddress(); publicIP != "" {
		return publicIP, nil
	}

	err := fmt.Errorf(currentIPAddressNotFound)
	return "", err
}

func needsArg(opts *CreateOpts) bool {
	// Unless currentIP flag is enabled and type is ip address, args are required.
	return !(opts.entryType == ipAddress && opts.currentIP)
}

func (opts *CreateOpts) validateCurrentIPFlag(cmd *cobra.Command, args []string) func() error {
	if !needsArg(opts) && len(args) > 0 {
		return func() error {
			return fmt.Errorf(
				"please either provide <entry> or use %s to use current IP Address.\n\nUsage: %s",
				flag.CurrentIP,
				cmd.UseLine(),
			)
		}
	}

	if needsArg(opts) && len(args) == 0 {
		return func() error {
			return fmt.Errorf(
				"%q with entry type %s requires at least %d argument received %d\n\nUsage:  %s",
				cmd.CommandPath(),
				opts.entryType,
				1,
				len(args),
				cmd.UseLine(),
			)
		}
	}
	return func() error { return nil }
}

// mongocli atlas accessList(s) create <entry> --type cidrBlock|ipAddress|awsSecurityGroup [--comment comment] [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <entry>",
		Short: "Create an IP access list for your project.",
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"args":      "entry",
			"entryDesc": "The IP address, CIDR address, or AWS security group ID of the access list entry to create.",
		},
		Example: `  
		  Create IP address access list with the current IP address. Entry is not needed in this case.
		  $ mongocli atlas accessList create --currentIP
		  $ mongocli atlas accessList create --type ipAddress --currentIP
		`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.validateCurrentIPFlag(cmd, args),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if len(args) == 0 {
				opts.entry, err = iPAddress()
			} else {
				opts.entry, err = args[0], nil
			}

			if err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.entryType, flag.Type, ipAddress, usage.AccessListType)
	cmd.Flags().StringVar(&opts.comment, flag.Comment, "", usage.Comment)
	cmd.Flags().StringVar(&opts.deleteAfter, flag.DeleteAfter, "", usage.AccessListsDeleteAfter)
	cmd.Flags().BoolVar(&opts.currentIP, flag.CurrentIP, false, usage.CurrentIP)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
