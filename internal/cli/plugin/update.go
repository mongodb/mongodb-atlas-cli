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
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	errTooManyArguments   = errors.New(`either the "--all" flag or the plugin identifier can be provided, but not both`)
	errNotEnoughArguments = errors.New(`either the "--all" flag or the plugin identifier needs to be provided`)
)

type UpdateOpts struct {
	cli.OutputOpts
	Opts
	// assetOpts []*AssetOpts
	UpdateAll bool
	pluginArg string
}

// func printPluginUpdateWarning(p *plugin.Plugin, err error) {
// 	_, _ = log.Warningf("could not update plugin %s because: %v", p.Name, err)
// }

// func (opts *UpdateOpts) updatePlugin(pluginToUpdate *plugin.Plugin) {
// 	if pluginToUpdate.Github != nil {
// 		_, _ = log.Warningf(`plugin "%s" could not be updated because its manifest file does not contain any github values`, pluginToUpdate.Name)
// 		return
// 	}

// 	opts.Print(fmt.Sprintf(`Updating plugin "%s"`, pluginToUpdate.Name))

// 	get all plugin assets info from github repository
// 	assets, err := opts.getPluginAssetInfo()
// 	if err != nil {
// 		printPluginUpdateWarning(plugin, err)
// 		return
// 	}
// }

// func (opts *UpdateOpts) Run() error {
// githubClient := github.NewClient(nil)

// if opts.UpdateAll {

// } else {
// 	fmt.Printf("Plugin %s will be updated\n", opts.pluginArg)
// }

// 	return nil
// }

func UpdateBuilder(plugins []*plugin.Plugin) *cobra.Command {
	opts := &UpdateOpts{
		UpdateAll: false,
		Opts: Opts{
			plugins: plugins,
		},
	}

	const use = "update"
	cmd := &cobra.Command{
		Use:     use + " [plugin]",
		Aliases: cli.GenerateAliases(use),
		Annotations: map[string]string{
			"pluginDesc": "Plugin identifier.",
		},
		Short: "Update Atlas CLI plugin.",
		Long: `Update an Atlas CLI plugin.
You can specify a plugin to update using either the "<github-owner>/<github-repository-name>" format or the plugin name.
Additionally, you can use the "--all" flag to update all plugins.
`,
		Args: require.MaximumNArgs(1),
		Example: `  # Update a plugin:
  atlas plugin update mongodb/atlas-cli-plugin-example
  atlas plugin update atlas-cli-plugin-example
  
  # Update all plugins
  atlas plugin update --all`,
		PreRunE: func(_ *cobra.Command, arg []string) error {
			// make sure either the "--all" flag is set or the plugin identifier but not both
			if opts.UpdateAll && len(arg) >= 1 {
				return errTooManyArguments
			}
			if !opts.UpdateAll && len(arg) != 1 {
				return errNotEnoughArguments
			}
			if !opts.UpdateAll {
				opts.pluginArg = arg[0]
			}

			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			// return opts.Run()
			return nil
		},
	}

	cmd.Flags().BoolVar(&opts.UpdateAll, flag.All, false, usage.UpdateAllPlugins)

	return cmd
}
