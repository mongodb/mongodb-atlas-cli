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

package ldap

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type VerifyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	hostname           string
	port               int
	bindUsername       string
	bindPassword       string
	caCertificate      string
	authzQueryTemplate string
	store              store.LDAPConfigurationVerifier
}

func (opts *VerifyOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var verifyTemplate = `REQUEST ID	PROJECT ID	STATUS
{{.RequestId}}	{{.GroupId}}	{{.Status}}
`

func (opts *VerifyOpts) Run() error {
	r, err := opts.store.VerifyLDAPConfiguration(opts.ConfigProjectID(), opts.newLDAP())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *VerifyOpts) Prompt() error {
	if opts.bindPassword != "" {
		return nil
	}

	if !opts.IsTerminalInput() {
		_, err := fmt.Fscanln(opts.InReader, &opts.bindPassword)
		return err
	}

	prompt := &survey.Password{
		Message: "Password:",
	}

	if err := telemetry.TrackAskOne(prompt, &opts.bindPassword); err != nil {
		return err
	}

	if opts.bindPassword == "" {
		return errors.New("no password provided")
	}

	return nil
}

func (opts *VerifyOpts) newLDAP() *atlasv2.LDAPVerifyConnectivityJobRequestParams {
	return &atlasv2.LDAPVerifyConnectivityJobRequestParams{
		Hostname:           opts.hostname,
		Port:               opts.port,
		BindUsername:       opts.bindUsername,
		BindPassword:       opts.bindPassword,
		CaCertificate:      &opts.caCertificate,
		AuthzQueryTemplate: &opts.authzQueryTemplate,
	}
}

// VerifyBuilder mongocli security atlas ldap verify --hostname hostname --port port --bindUsername bindUsername --bindPassword bindPassword --caCertificate caCertificate --authzQueryTemplate authzQueryTemplate [--projectId projectId].
func VerifyBuilder() *cobra.Command {
	opts := &VerifyOpts{}
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Request verification of an LDAP configuration for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": verifyTemplate,
			"toc":    "true",
		},
		Example: `  # Request the JSON-formatted verification of the LDAP configuration for the atlas-ldaps-01.ldap.myteam.com host in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas security ldap verify --hostname atlas-ldaps-01.ldap.myteam.com --bindUsername "CN=Administrator,CN=Users,DC=atlas-ldaps-01,DC=myteam,DC=com" --bindPassword changeMe --projectId 5e2211c17a3e5a48f5497de3 --output json
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), verifyTemplate),
				opts.InitInput(cmd.InOrStdin())); err != nil {
				return err
			}
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.hostname, flag.Hostname, "", usage.LDAPHostname)
	cmd.Flags().IntVar(&opts.port, flag.Port, defaultLDAPPort, usage.LDAPPort)
	cmd.Flags().StringVar(&opts.bindUsername, flag.BindUsername, "", usage.BindUsername)
	cmd.Flags().StringVar(&opts.bindPassword, flag.BindPassword, "", usage.BindPassword)
	cmd.Flags().StringVar(&opts.caCertificate, flag.CaCertificate, "", usage.CaCertificate)
	cmd.Flags().StringVar(&opts.authzQueryTemplate, flag.AuthzQueryTemplate, "", usage.AuthzQueryTemplate)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Hostname)
	_ = cmd.MarkFlagRequired(flag.BindUsername)

	cmd.AddCommand(
		StatusBuilder(),
	)

	return cmd
}
