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
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

type InstallOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	AssetOpts
	plugins []*plugin.Plugin
}

func (opts *InstallOpts) validateForExistingPlugins() error {
	for _, plugin := range opts.plugins {
		if plugin.Github != nil && plugin.Github.Equals(opts.repositoryOwner, opts.repositoryName) {
			return fmt.Errorf("a plugin from the repository %s is already installed.\nTo update the plugin run: \n\tatlas plugin update %s", opts.fullRepositoryDefinition(), opts.fullRepositoryDefinition())
		}
	}
	return nil
}

func (opts *InstallOpts) validatePlugin(pluginDirectoryPath string) error {
	manifest, err := plugin.GetManifestFromPluginDirectory(pluginDirectoryPath)
	if err != nil {
		return err
	}

	if valid, errorList := manifest.IsValid(); !valid {
		var manifestErrorLog strings.Builder
		manifestErrorLog.WriteString(fmt.Sprintf("plugin in directory \"%s\" could not be loaded due to the following error(s) in the manifest.yaml:\n", pluginDirectoryPath))
		for _, err := range errorList {
			manifestErrorLog.WriteString(fmt.Sprintf("\t- %s\n", err.Error()))
		}

		return errors.New(manifestErrorLog.String())
	}

	existingCommandsMap := make(map[string]bool)
	for _, cmd := range opts.existingCommands {
		existingCommandsMap[cmd.Name()] = true
	}
	if plugin.HasDuplicateCommand(manifest, existingCommandsMap) {
		return fmt.Errorf(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, opts.fullRepositoryDefinition())
	}

	return nil
}

func (opts *InstallOpts) Run() error {
	rc, err := opts.getPluginAssetAsReadCloser()
	if err != nil {
		return err
	}

	pluginZipFilePath, err := saveReadCloserToPluginAssetZipFile(rc)
	defer os.Remove(pluginZipFilePath)
	if err != nil {
		return err
	}

	pluginDirectoryPath, err := opts.extractPluginAssetZipFile(pluginZipFilePath)
	if err != nil {
		if pluginDirectoryPath != "" {
			os.RemoveAll(pluginDirectoryPath)
		}

		return err
	}

	err = opts.validatePlugin(pluginDirectoryPath)
	if err != nil {
		os.RemoveAll(pluginDirectoryPath)
		return err
	}

	return opts.Print(fmt.Sprintf("Plugin %s successfully installed", opts.fullRepositoryDefinition()))
}

func InstallBuilder(plugins []*plugin.Plugin, existingCommands []*cobra.Command) *cobra.Command {
	opts := &InstallOpts{
		plugins: plugins,
		AssetOpts: AssetOpts{
			existingCommands: existingCommands,
		},
	}

	const use = "install"
	cmd := &cobra.Command{
		Use:     use + " [<github-owner>/<github-repository-name>]",
		Aliases: cli.GenerateAliases(use),
		Annotations: map[string]string{
			"<github-owner>/<github-repository-name>Desc": "Repository identifier.",
		},
		Short: "Install Atlas CLI plugin from a GitHub repository.",
		Long: `Install an Atlas CLI plugin from a GitHub repository.
You can specify a GitHub repository using either the "<github-owner>/<github-repository-name>" format or a full URL.
When you install the plugin, its latest release on GitHub is used by default.
To install a specific version of the plugin, append the version number directly to the plugin name using the @ symbol.

MongoDB provides an example plugin: https://github.com/mongodb/atlas-cli-plugin-example
`,
		Args: require.ExactArgs(1),
		Example: `  # Install the latest version of the plugin:
  atlas plugin install mongodb/atlas-cli-plugin-example
  atlas plugin install https://github.com/mongodb/atlas-cli-plugin-example
  
  # Install a specific version of the plugin:
  atlas plugin install mongodb/atlas-cli-plugin-example@1.0.4
  atlas plugin install https://github.com/mongodb/atlas-cli-plugin-example/@v1.2.3`,
		PreRunE: func(_ *cobra.Command, args []string) error {
			repositoryOwner, repositoryName, err := parseGithubValues(args[0])
			if err != nil {
				return err
			}
			opts.repositoryOwner, opts.repositoryName = repositoryOwner, repositoryName

			version, err := parseReleaseVersion(args[0])
			if err != nil {
				return err
			}
			opts.releaseVersion = version

			opts.ghClient = github.NewClient(nil)

			assets, err := opts.getPluginAssets()
			if err != nil {
				return err
			}
			opts.pluginAssets = assets

			return opts.PreRunE(opts.validateForExistingPlugins)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	return cmd
}
