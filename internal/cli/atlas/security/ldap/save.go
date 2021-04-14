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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type SaveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	hostname              string
	port                  int
	bindUsername          string
	bindPassword          string
	caCertificate         string
	authzQueryTemplate    string
	authenticationEnabled bool
	authorizationEnabled  bool
	store                 store.LDAPConfigurationSaver
}

func (opts *SaveOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.PublicAuthenticatedPreset(config.Default()))
	return err
}

var saveTemplate = `HOSTNAME	PORT	AUTHENTICATION	AUTHORIZATION
{{.LDAP.Hostname}}	{{.LDAP.Port}}	{{.LDAP.AuthenticationEnabled}}	{{.LDAP.AuthorizationEnabled}}
`

func (opts *SaveOpts) Run() error {
	r, err := opts.store.SaveLDAPConfiguration(opts.ConfigProjectID(), opts.newLDAPConfiguration())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *SaveOpts) newLDAPConfiguration() *atlas.LDAPConfiguration {
	return &atlas.LDAPConfiguration{
		LDAP: &atlas.LDAP{
			AuthenticationEnabled: opts.authenticationEnabled,
			AuthorizationEnabled:  opts.authorizationEnabled,
			Hostname:              opts.hostname,
			Port:                  opts.port,
			BindUsername:          opts.bindUsername,
			BindPassword:          opts.bindPassword,
			CaCertificate:         opts.caCertificate,
			AuthzQueryTemplate:    opts.authzQueryTemplate,
		},
	}
}

// mongocli atlas security ldap save --hostname hostname --port port --bindUsername bindUsername --bindPassword bindPassword --caCertificate caCertificate
// --authzQueryTemplate authzQueryTemplate --authenticationEnabled authenticationEnabled --authorizationEnabled authorizationEnabled [--projectId projectId]
func SaveBuilder() *cobra.Command {
	opts := &SaveOpts{}
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save an LDAP configuration.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), saveTemplate))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.hostname, flag.Hostname, "", usage.LDAPHostname)
	cmd.Flags().IntVar(&opts.port, flag.Port, 636, usage.LDAPPort)
	cmd.Flags().StringVar(&opts.bindUsername, flag.BindUsername, "", usage.BindUsername)
	cmd.Flags().StringVar(&opts.bindPassword, flag.BindPassword, "", usage.BindPassword)
	cmd.Flags().StringVar(&opts.caCertificate, flag.CaCertificate, "", usage.CaCertificate)
	cmd.Flags().StringVar(&opts.authzQueryTemplate, flag.AuthzQueryTemplate, "", usage.AuthzQueryTemplate)
	cmd.Flags().BoolVar(&opts.authenticationEnabled, flag.AuthenticationEnabled, false, usage.AuthenticationEnabled)
	cmd.Flags().BoolVar(&opts.authorizationEnabled, flag.AuthorizationEnabled, false, usage.AuthorizationEnabled)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Hostname)
	_ = cmd.MarkFlagRequired(flag.BindUsername)
	_ = cmd.MarkFlagRequired(flag.BindPassword)

	return cmd
}
