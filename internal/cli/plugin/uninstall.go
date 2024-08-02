// Copyright 2024 MongoDB Inc
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

package plugin

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

type UninstallOpts struct {
	cli.OutputOpts
	Opts
	pluginArg string
}

func (opts *UninstallOpts) Run() error {
	pluginToUninstall, err := opts.findPluginWithArg(opts.pluginArg)
	if err != nil {
		return err
	}

	err = pluginToUninstall.Uninstall()
	if err != nil {
		return fmt.Errorf(`plugin %s could not be uninstalled: %w`, pluginToUninstall.Name, err)
	}

	return opts.Print(fmt.Sprintf("Plugin %s uninstalled successfully", pluginToUninstall.Name))
}

func UninstallBuilder(plugins []*plugin.Plugin) *cobra.Command {
	opts := &UninstallOpts{
		Opts: Opts{
			plugins: plugins,
		},
	}

	const use = "uninstall"
	cmd := &cobra.Command{
		Use:     use + " [plugin]",
		Aliases: cli.GenerateAliases(use),
		Annotations: map[string]string{
			"pluginDesc": "Plugin identifier.",
		},
		Short: "Uninstall Atlas CLI plugin.",
		Long: `Uninstall an Atlas CLI plugin.
You can specify a plugin to uninstall using either the "<github-owner>/<github-repository-name>" format or the plugin name.
`,
		Args: require.ExactArgs(1),
		Example: `  # Uninstall a plugin:
  atlas plugin uninstall mongodb/atlas-cli-plugin-example
  atlas plugin uninstall atlas-cli-plugin-example`,
		PreRun: func(_ *cobra.Command, arg []string) {
			opts.pluginArg = arg[0]
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	return cmd
}
