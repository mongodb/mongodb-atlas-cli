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

package owner

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type CreateOpts struct {
	cli.OutputOpts
	email     string
	password  string
	firstName string
	lastName  string
	accessIps []string
	store     store.OwnerCreator
}

func (opts *CreateOpts) init() error {
	var err error
	opts.store, err = store.NewUnauthenticated(config.Default())
	return err
}

var createTemplate = `Owner successfully created.
{{if .APIKey}}Personal API Key: {{.APIKey}}{{end}}
{{- if .ProgrammaticAPIKey}}
Public API Key: {{.ProgrammaticAPIKey.PublicKey}}
Private API Key: {{.ProgrammaticAPIKey.PrivateKey}}
{{- end}}
`

func (opts *CreateOpts) Run() error {
	user := opts.newOwner()
	r, err := opts.store.CreateOwner(user, opts.accessIps)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newOwner() *opsmngr.User {
	user := &opsmngr.User{
		Username:     opts.email,
		Password:     opts.password,
		FirstName:    opts.firstName,
		LastName:     opts.lastName,
		EmailAddress: opts.email,
		Links:        nil,
	}
	return user
}

func (opts *CreateOpts) Prompt() error {
	if opts.password != "" {
		return nil
	}
	prompt := &survey.Password{
		Message: "Password:",
	}
	return survey.AskOne(prompt, &opts.password)
}

// mongocli ops-manager owner create --email username --password password --firstName firstName --lastName lastName --accessListIp <IP>
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: CreateOwner,
		Long:  CreateLong,
		Args:  cobra.OnlyValidArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			_ = opts.InitOutput(cmd.OutOrStdout(), createTemplate)()
			if err := opts.init(); err != nil {
				return err
			}
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.email, flag.Email, "", usage.Email)
	cmd.Flags().StringVarP(&opts.password, flag.Password, flag.PasswordShort, "", usage.Password)
	cmd.Flags().StringVar(&opts.firstName, flag.FirstName, "", usage.FirstName)
	cmd.Flags().StringVar(&opts.lastName, flag.LastName, "", usage.LastName)
	cmd.Flags().StringSliceVar(&opts.accessIps, flag.AccessListIP, []string{}, usage.AccessListIps)
	_ = cmd.Flags().MarkDeprecated(flag.WhitelistIP, fmt.Sprintf("please use --%s instead", flag.AccessListIP))
	cmd.Flags().StringSliceVar(&opts.accessIps, flag.WhitelistIP, []string{}, usage.AccessListIps)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Email)
	_ = cmd.MarkFlagRequired(flag.FirstName)
	_ = cmd.MarkFlagRequired(flag.LastName)

	return cmd
}
