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
package dbusers

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const describeTemplate = `USERNAME	DATABASE
{{.Username}}	{{.DatabaseName}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.WizardOpts
	store    store.DatabaseUserDescriber
	authDB   string
	username string
}

func (opts *DescribeOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.DatabaseUser(opts.authDB, opts.ConfigProjectID(), opts.username)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) newWizardRequiredFlags() ([]*cli.Flag, error) {
	var flags []*cli.Flag
	userNames, err := opts.userNames()
	if err != nil {
		return nil, err
	}
	flags = append(flags,
		&cli.Flag{Name: flag.Username, Usage: usage.DBUsername, Options: userNames})
	return flags, nil
}

func (opts *DescribeOpts) newWizardOptionalFlags() []*cli.Flag {
	var flags []*cli.Flag
	flags = append(flags,
		&cli.Flag{Name: flag.AuthDB, Usage: usage.AuthDB})
	return flags
}

func (opts *DescribeOpts) userNames() ([]string, error) {
	listOpt := &atlas.ListOptions{
		ItemsPerPage: cli.WizardItemsPerPage,
	}
	dbUsers, err := opts.store.DatabaseUsers(opts.ConfigProjectID(), listOpt)

	if err != nil {
		return nil, err
	}

	var userNames []string

	for idx := range dbUsers {
		userNames = append(userNames, dbUsers[idx].Username)
	}
	return userNames, nil
}

func (opts *DescribeOpts) initWizardFlags(answers map[string]string) {
	opts.username = opts.GetAnswer(answers, flag.Username)

	authDB := opts.GetAnswer(answers, flag.AuthDB)
	if authDB != "" {
		opts.authDB = authDB
	}
}

// mongocli atlas dbuser(s) describe <username> --projectId projectId --authDB authDB
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <name>",
		Short: describeDBUser,
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := opts.initStore()
			if err != nil {
				return err
			}

			if len(args) == 0 || opts.Wizard {
				requiredFlags, err := opts.newWizardRequiredFlags()
				if err != nil {
					return err
				}
				answers, err := opts.RunWizard(requiredFlags, opts.newWizardOptionalFlags())
				if err != nil {
					return err
				}
				opts.initWizardFlags(answers)
			} else {
				opts.username = args[0]
			}

			validAuthDBs := []string{convert.AdminDB, convert.ExternalAuthDB}
			if err := validate.FlagInSlice(opts.authDB, flag.AuthDB, validAuthDBs); err != nil {
				return err
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.authDB, flag.AuthDB, convert.AdminDB, usage.AuthDB)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
