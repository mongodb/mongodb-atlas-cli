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

package config

import (
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/config/secure"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/spf13/cobra"
)

type describeOpts struct {
	name string
	cli.OutputOpts
}

var descTemplate = `SETTING	VALUE{{ range $key, $value := . }}
{{$key}}	{{$value}}{{end}}
`

const redacted = "redacted"

// AddSecureProperties adds secure properties to the map with "redacted" value
// if they are available in the config.
func (opts *describeOpts) AddSecureProperties(m map[string]string) (map[string]string, error) {
	// Check if secure storage is available
	configStore, err := config.NewDefaultStore()
	if err != nil {
		return nil, err
	}
	if !configStore.IsSecure() {
		return m, nil
	}

	// We are using a keyring client directly here to avoid printing env vars
	secureKeyring := secure.NewDefaultKeyringClient()
	// Service Account
	if v, err := secureKeyring.Get(opts.name, "client_id"); err == nil && v != "" {
		m["client_id"] = redacted
	}
	if v, err := secureKeyring.Get(opts.name, "client_secret"); err == nil && v != "" {
		m["client_secret"] = redacted
	}
	// API Keys
	if v, err := secureKeyring.Get(opts.name, "public_api_key"); err == nil && v != "" {
		m["public_api_key"] = redacted
	}
	if v, err := secureKeyring.Get(opts.name, "private_api_key"); err == nil && v != "" {
		m["private_api_key"] = redacted
	}
	// User Account
	if v, err := secureKeyring.Get(opts.name, "access_token"); err == nil && v != "" {
		m["access_token"] = redacted
	}
	if v, err := secureKeyring.Get(opts.name, "refresh_token"); err == nil && v != "" {
		m["refresh_token"] = redacted
	}

	return m, nil
}

func (opts *describeOpts) Run() error {
	if !config.Exists(opts.name) {
		return fmt.Errorf("you don't have a profile named '%s'", opts.name)
	}

	if err := config.SetName(opts.name); err != nil {
		return err
	}

	mapConfig, err := opts.AddSecureProperties(config.Map())
	if err != nil {
		return err
	}

	return opts.Print(mapConfig)
}

func DescribeBuilder() *cobra.Command {
	opts := &describeOpts{}
	opts.Template = descTemplate
	cmd := &cobra.Command{
		Use:     "describe <name>",
		Aliases: []string{"get"},
		Short:   "Return the profile you specify.",
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Label that identifies the profile.",
		},
		PreRun: func(cmd *cobra.Command, _ []string) {
			opts.OutWriter = cmd.OutOrStdout()
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	opts.AddOutputOptFlags(cmd)

	return cmd
}
