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
	"github.com/AlecAivazis/survey/v2"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerOwnerCreateOpts struct {
	email        string
	password     string
	firstName    string
	lastName     string
	whitelistIps []string
	store        store.OwnerCreator
}

func (opts *opsManagerOwnerCreateOpts) init() error {
	var err error
	opts.store, err = store.NewUnauthenticated()
	return err
}

func (opts *opsManagerOwnerCreateOpts) Run() error {
	user := opts.newOwner()
	result, err := opts.store.CreateOwner(user, opts.whitelistIps)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *opsManagerOwnerCreateOpts) newOwner() *om.User {
	user := &om.User{
		Username:     opts.email,
		Password:     opts.password,
		FirstName:    opts.firstName,
		LastName:     opts.lastName,
		EmailAddress: opts.email,
		Links:        nil,
	}
	return user
}

func (opts *opsManagerOwnerCreateOpts) Prompt() error {
	if opts.password != "" {
		return nil
	}
	prompt := &survey.Password{
		Message: "Password:",
	}
	return survey.AskOne(prompt, &opts.password)
}

// mongocli ops-manager owner create --email username --password password --firstName firstName --lastName lastName --whitelistIps whitelistIp
func OpsManagerOwnerCreateBuilder() *cobra.Command {
	opts := new(opsManagerOwnerCreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: description.CreateOwner,
		Args:  cobra.OnlyValidArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.email, flags.Email, "", usage.Email)
	cmd.Flags().StringVar(&opts.password, flags.Password, "", usage.Password)
	cmd.Flags().StringVar(&opts.firstName, flags.FirstName, "", usage.FirstName)
	cmd.Flags().StringVar(&opts.lastName, flags.LastName, "", usage.LastName)
	cmd.Flags().StringSliceVar(&opts.whitelistIps, flags.WhitelistIP, []string{}, usage.WhitelistIps)

	_ = cmd.MarkFlagRequired(flags.Username)
	_ = cmd.MarkFlagRequired(flags.FirstName)
	_ = cmd.MarkFlagRequired(flags.LastName)

	return cmd
}
