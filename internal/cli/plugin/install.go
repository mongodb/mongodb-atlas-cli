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

func (opts *InstallOpts) checkForDuplicatePlugins() error {
	for _, p := range opts.plugins {
		if p.Github != nil && p.Github.Equals(opts.githubRelease.owner, opts.githubRelease.name) {
			return fmt.Errorf("a plugin from the repository %s is already installed.\nTo update the plugin run: \n\tatlas plugin update %s", opts.repository(), opts.repository())
		}
	}
	return nil
}

// checks if the plugin manifest is valid and that the plugin
// doesn't contain any commands that conflict with existing CLI commands.
func (opts *InstallOpts) validatePlugin(pluginDirectoryPath string) error {
	// Get the manifest from the plugin directory
	manifest, err := plugin.GetManifestFromPluginDirectory(pluginDirectoryPath)
	if err != nil {
		return err
	}

	// Validate the manifest
	if valid, errorList := manifest.IsValid(); !valid {
		var manifestErrorLog strings.Builder
		manifestErrorLog.WriteString(fmt.Sprintf("plugin in directory \"%s\" could not be loaded due to the following error(s) in the manifest.yaml:\n", pluginDirectoryPath))
		for _, err := range errorList {
			manifestErrorLog.WriteString(fmt.Sprintf("\t- %s\n", err.Error()))
		}

		return errors.New(manifestErrorLog.String())
	}

	// Check for duplicate commands
	existingCommandsMap := make(map[string]bool)
	for _, cmd := range opts.existingCommands {
		existingCommandsMap[cmd.Name()] = true
	}
	if plugin.HasDuplicateCommand(manifest, existingCommandsMap) {
		return fmt.Errorf(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, opts.repository())
	}

	return nil
}

func (opts *InstallOpts) Run() error {
	opts.ghClient = github.NewClient(nil)

	// get all plugin assets info from github repository
	assets, err := opts.getPluginAssetInfo()
	if err != nil {
		return err
	}

	// find correct assetID using system requirements
	assetID, err := opts.getAssetID(assets)
	if err != nil {
		return err
	}

	// download plugin asset zip file and save it as ReadCloser
	rc, err := opts.getPluginAssetAsReadCloser(assetID)
	if err != nil {
		return err
	}

	// use the ReadCloser to save the asset zip file in the default plugin directory
	pluginZipFilePath, err := saveReadCloserToPluginAssetZipFile(rc)
	if err != nil {
		return err
	}
	defer os.Remove(pluginZipFilePath) // delete zip file after install command finishes

	// try to extract content of plugin zip file and save it in default plugin directory
	pluginDirectoryPath, err := opts.extractPluginAssetZipFile(pluginZipFilePath)
	if err != nil {
		return err
	}

	// validate the extracted plugin files
	// if plugin is invalid, delete all of its files
	err = opts.validatePlugin(pluginDirectoryPath)
	if err != nil {
		os.RemoveAll(pluginDirectoryPath)
		return err
	}

	return opts.Print(fmt.Sprintf("Plugin %s successfully installed", opts.repository()))
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
			githubRelease, err := parseGithubReleaseValues(args[0])
			if err != nil {
				return err
			}
			opts.githubRelease = githubRelease

			return opts.PreRunE(opts.checkForDuplicatePlugins)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	return cmd
}
