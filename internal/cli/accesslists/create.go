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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	cidrBlock        = "cidrBlock"
	ipAddress        = "ipAddress"
	awsSecurityGroup = "awsSecurityGroup"
	createTemplate   = "Created a new IP access list.\n"
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
	entry, err := opts.newProjectIPAccessList()

	if err != nil {
		return err
	}

	r, err := opts.store.CreateProjectIPAccessList(entry)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newProjectIPAccessList() ([]*atlasv2.NetworkPermissionEntry, error) {
	entry := &atlasv2.NetworkPermissionEntry{
		GroupId: pointer.Get(opts.ConfigProjectID()),
		Comment: &opts.comment,
	}
	if opts.deleteAfter != "" {
		deleteAfterDate, err := convert.ParseTimestamp(opts.deleteAfter)
		if err != nil {
			return nil, err
		}
		entry.DeleteAfterDate = &deleteAfterDate
	}
	switch opts.entryType {
	case cidrBlock:
		entry.CidrBlock = &opts.entry
	case ipAddress:
		entry.IpAddress = &opts.entry
	case awsSecurityGroup:
		entry.AwsSecurityGroup = &opts.entry
	}
	return []*atlasv2.NetworkPermissionEntry{entry}, nil
}

func IPAddress() (string, error) {
	if publicIP := store.IPAddress(); publicIP != "" {
		return publicIP, nil
	}

	return "", errors.New("unable to find your public IP address. Specify the public IP address for this command")
}

func (opts *CreateOpts) needsArg() bool {
	// Unless currentIP flag is enabled and type is ip address, args are required.
	return !(opts.entryType == ipAddress && opts.currentIP)
}

func (opts *CreateOpts) validateCurrentIPFlag(cmd *cobra.Command, args []string) func() error {
	return func() error {
		if !opts.needsArg() && len(args) > 0 {
			return fmt.Errorf(
				"please either provide [entry] or use %s to use your current IP Address.\n\nUsage: %s",
				flag.CurrentIP,
				cmd.UseLine(),
			)
		}

		if opts.needsArg() && len(args) == 0 {
			return fmt.Errorf(
				"%q with entry type %s requires at least 1 argument, received %d\n\nUsage:  %s",
				cmd.CommandPath(),
				opts.entryType,
				len(args),
				cmd.UseLine(),
			)
		}

		if opts.entryType != ipAddress && opts.currentIP {
			return fmt.Errorf("%q with entry type %s does not support %s flag.\n\n Usage: %s",
				cmd.CommandPath(),
				opts.entryType,
				flag.CurrentIP,
				cmd.UseLine(),
			)
		}
		return nil
	}
}

// atlas accessList(s) create <entry> --type cidrBlock|ipAddress|awsSecurityGroup [--comment comment] [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create [entry]",
		Short: "Create an IP access list entry for your project.",
		Long: `The access list can contain trusted IP addresses, AWS security group IDs, and entries in Classless Inter-Domain Routing (CIDR) notation. You can add only one access list entry at a time. You can create one access list per project. 
		
The command doesn't overwrite existing entries in the access list. Instead, it adds the new entries to the list of entries.

` + fmt.Sprintf(usage.RequiredRole, "Read Write"),
		Args: require.MaximumNArgs(1),
		Annotations: map[string]string{
			"entryDesc": "IP address, CIDR address, or AWS security group ID that you want to add to the access list.",
			"output":    createTemplate,
		},
		Example: `  # Create an IP access list entry using the current IP address:
  atlas accessList create --currentIp
  
  # Create an access list entry for the IP address 192.0.2.15 in the project with ID 5e2211c17a3e5a48f5497de3:
  atlas accessList create 192.0.2.15 --type ipAddress --projectId 5e2211c17a3e5a48f5497de3 --comment "IP address for app server 2" --output json
  
  # Create an access list entry in CIDR notation for 73.231.201.205/24 in the project with ID 5e2211c17a3e5a48f5497de3:
  atlas accessList create 73.231.201.205/24 --type cidrBlock --projectId 5e2211c17a3e5a48f5497de3 --output json --comment "CIDR block for servers C - F"
  
  # Create an access list entry for the AWS security group sg-903004f8 in the project with ID 5e2211c17a3e5a48f5497de3:
  atlas accessList create sg-903004f8 --type awsSecurityGroup
  --projectId 5e2211c17a3e5a48f5497de3 --output json --comment "AWS Security Group"`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.validateCurrentIPFlag(cmd, args),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			var err error
			if len(args) == 0 {
				opts.entry, err = IPAddress()
			} else {
				opts.entry, err = args[0], nil
			}

			if err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.entryType, flag.TypeFlag, ipAddress, usage.AccessListType)
	cmd.Flags().StringVar(&opts.comment, flag.Comment, "", usage.Comment)
	cmd.Flags().StringVar(&opts.deleteAfter, flag.DeleteAfter, "", usage.AccessListsDeleteAfter)
	cmd.Flags().BoolVar(&opts.currentIP, flag.CurrentIP, false, usage.CurrentIP)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
