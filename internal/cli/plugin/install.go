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
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type InstallOpts struct {
	cli.ProjectOpts
	cli.PreRunOpts
	cli.OutputOpts
	Opts
	ghClient                  *github.Client
	githubAsset               *GithubAsset
	skipSignatureVerification bool
}

func (opts *InstallOpts) checkForDuplicatePlugins() error {
	_, err := opts.findPluginWithGithubValues(opts.githubAsset.owner, opts.githubAsset.name)
	if err != nil {
		return nil
	}

	return fmt.Errorf("a plugin from the repository %s is already installed.\nTo update the plugin run: \n\tatlas plugin update %s", opts.githubAsset.repository(), opts.githubAsset.repository())
}

// checks if the plugin manifest is valid and that the plugin
// doesn't contain any commands that conflict with existing CLI commands.
func (opts *InstallOpts) validatePlugin(pluginDirectoryPath string) error {
	// Get the manifest from the plugin directory
	manifest, err := plugin.GetManifestFromPluginDirectory(pluginDirectoryPath)
	if err != nil {
		return err
	}

	err = validateManifest(manifest)
	if err != nil {
		return err
	}

	// check for duplicate plugin names
	for _, p := range opts.getValidPlugins() {
		if manifest.Name == p.Name {
			return fmt.Errorf("a plugin with the name %s already exists", manifest.Name)
		}
	}

	// Check for duplicate commands
	existingCommandsSet := createExistingCommandsSet(opts.existingCommands)
	if manifest.HasDuplicateCommand(existingCommandsSet) {
		return fmt.Errorf(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, opts.githubAsset.repository())
	}

	return nil
}

func (opts *InstallOpts) Run(ctx context.Context) error {
	// get all plugin assets info from github repository
	assets, err := opts.githubAsset.getReleaseAssets(opts.ghClient)
	if err != nil {
		return err
	}

	// find correct assetID, signatureID and pubKeyID using system requirements
	assetID, signatureID, pubKeyID, err := opts.githubAsset.getIDs(assets)
	if err != nil {
		return err
	}

	// When signatureID and pubKeyID are 0, the signature check is skipped.
	if opts.skipSignatureVerification {
		signatureID = 0
		pubKeyID = 0
	}

	// download plugin asset archive file and save it as ReadCloser
	rc, err := opts.githubAsset.getPluginAssetsAsReadCloser(opts.ghClient, assetID, signatureID, pubKeyID)
	if err != nil {
		return err
	}
	defer rc.Close()

	// use the ReadCloser to save the asset archive file in the default plugin directory
	pluginArchiveFilePath, err := saveReadCloserToPluginAssetArchiveFile(rc)
	if err != nil {
		return err
	}
	defer os.Remove(pluginArchiveFilePath) // delete archive file after install command finishes

	// try to extract content of plugin archive file and save it in default plugin directory
	pluginDirectoryPath, err := extractPluginAssetArchiveFile(ctx, pluginArchiveFilePath, opts.githubAsset.getPluginDirectoryName())
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

	return opts.Print(fmt.Sprintf("Plugin %s successfully installed", opts.githubAsset.repository()))
}

func InstallBuilder(pluginOpts *Opts) *cobra.Command {
	opts := &InstallOpts{
		ghClient: NewAuthenticatedGithubClient(),
	}
	opts.Opts = *pluginOpts

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
			githubAssetRelease, err := parseGithubReleaseValues(args[0])
			if err != nil {
				return err
			}
			opts.githubAsset = githubAssetRelease

			return opts.PreRunE(opts.checkForDuplicatePlugins)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&opts.skipSignatureVerification, flag.SkipSignatureVerification, false, usage.SkipSignatureVerification)

	return cmd
}
