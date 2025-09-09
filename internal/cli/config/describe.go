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
	"slices"

	"github.com/mongodb/atlas-cli-core/config"
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

const (
	redactedSecureText = "[redacted - source: secure storage]"
	redactedConfigText = "[redacted - source: config file]"
	servicePrefix      = "atlascli_"
)

func (*describeOpts) GetConfig(configStore config.Store, profileName string) (map[string]string, error) {
	// Get the profile map, this only contains properties coming from the insecure store
	profileMap := configStore.GetProfileStringMap(profileName)

	redactedText := redactedConfigText
	if configStore.IsSecure() {
		redactedText = redactedSecureText
	}

	// Redact values
	for _, key := range config.SecureProperties {
		if v := configStore.GetProfileValue(profileName, key); v != nil && v != "" {
			profileMap[key] = redactedText
		}
	}

	return profileMap, nil
}

func (opts *describeOpts) Run() error {
	// Create a new config proxy store
	configStore, err := config.NewStoreWithEnvOption(false)
	if err != nil {
		return fmt.Errorf("could not create config store: %w", err)
	}

	profileNames := configStore.GetProfileNames()
	if !slices.Contains(profileNames, opts.name) {
		return fmt.Errorf("you don't have a profile named '%s'", opts.name)
	}

	mapConfig, err := opts.GetConfig(configStore, opts.name)
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
