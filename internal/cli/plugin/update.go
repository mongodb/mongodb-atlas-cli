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
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/set"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	errTooManyArguments        = errors.New(`either the "--all" flag or the plugin identifier can be provided, but not both`)
	errNotEnoughArguments      = errors.New(`either the "--all" flag or the plugin identifier needs to be provided`)
	errPluginHasNoGithubValues = errors.New(`specified plugin does not contain any Github values in its manifest.yaml file. This issue may have occurred because the plugin was added manually instead of using the "plugin install" command`)
	errUpdateArgInvalid        = errors.New(`the format of the plugin indentifier is invalid. You can specify a plugin to update using either the "<github-owner>/<github-repository-name>" format or the plugin name`)
)

type UpdateOpts struct {
	cli.OutputOpts
	Opts
	UpdateAll                 bool
	pluginSpecifier           string
	pluginUpdateVersion       *semver.Version
	ghClient                  *github.Client
	skipSignatureVerification bool
}

func printPluginUpdateWarning(p *plugin.Plugin, err error) {
	_, _ = log.Warningf("could not update plugin %s because: %v\n", p.Name, err)
}

// extract plugin specifier and version given the input argument of the update command.
func extractPluginSpecifierAndVersionFromArg(arg string) (string, *semver.Version, error) {
	regexPattern := `^(?P<pluginValue>[^\s@]+)(@(?P<version>.+))?$`
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return "", nil, fmt.Errorf("error compiling regex: %w", err)
	}

	matches := regex.FindStringSubmatch(arg)
	if matches == nil {
		return "", nil, errUpdateArgInvalid
	}

	names := regex.SubexpNames()
	groupMap := make(map[string]string)
	for i, match := range matches {
		if i == 0 {
			continue
		}
		groupMap[names[i]] = match
	}

	var version *semver.Version

	if versionValue, ok := groupMap["version"]; ok && versionValue != latest && versionValue != "" {
		versionValue := strings.TrimPrefix(versionValue, "v")
		semverVersion, err := semver.NewVersion(versionValue)
		if err != nil {
			return "", nil, fmt.Errorf(`the specified version "%s" is invalid, it needs to follow the rules of Semantic Versioning`, versionValue)
		}
		version = semverVersion
	}

	return groupMap["pluginValue"], version, nil
}

// checks if the plugin manifest is valid and that the plugin
// doesn't contain any commands that conflict with existing CLI commands.
func (opts *UpdateOpts) validatePlugin(pluginDirectoryPath string) error {
	// Get the manifest from the plugin directory
	manifest, err := plugin.GetManifestFromPluginDirectory(pluginDirectoryPath)
	if err != nil {
		return err
	}

	err = validateManifest(manifest)
	if err != nil {
		return err
	}

	// make sure that there is exactly one plugin with the same name
	pluginCount := 0
	for _, p := range opts.getValidPlugins() {
		if manifest.Name == p.Name {
			pluginCount++
		}
	}
	if pluginCount != 1 {
		return fmt.Errorf(`there needs to be exactly 1 plugin with the name "%s", but there are %d`, manifest.Name, pluginCount)
	}

	// Check for duplicate commands
	existingCommandsSet := set.NewSet[string]()
	for _, cmd := range opts.existingCommands {
		// only add command to existing commands map if it is not part of the plugin we want to update
		if sourcePluginName, ok := cmd.Annotations[sourcePluginName]; !ok || sourcePluginName != manifest.Name {
			existingCommandsSet.Add(cmd.Name())
		}
	}
	if manifest.HasDuplicateCommand(existingCommandsSet) {
		return fmt.Errorf(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, manifest.Name)
	}

	return nil
}

func (opts *UpdateOpts) updatePlugin(ctx context.Context, githubAssetRelease *GithubAsset, existingPlugin *plugin.Plugin) error {
	// get all plugin assets info from github repository
	assets, err := githubAssetRelease.getReleaseAssets(opts.ghClient)
	if err != nil {
		return err
	}

	// find correct assetID, signatureID and pubKeyID using system requirements
	assetID, signatureID, pubKeyID, err := githubAssetRelease.getIDs(assets)
	if err != nil {
		return err
	}

	// When signatureID and pubKeyID are 0, the signature check is skipped.
	if opts.skipSignatureVerification {
		signatureID = 0
		pubKeyID = 0
	}

	// download plugin asset archive file and save it as ReadCloser
	rc, err := githubAssetRelease.getPluginAssetsAsReadCloser(opts.ghClient, assetID, signatureID, pubKeyID)
	if err != nil {
		return err
	}
	defer rc.Close()

	// use the ReadCloser to save the asset archive file in the default plugin directory
	pluginArchiveFilePath, err := saveReadCloserToPluginAssetArchiveFile(rc)
	if err != nil {
		return err
	}
	defer os.Remove(pluginArchiveFilePath) // delete archive file after update command finishes

	// try to extract content of plugin archive file and save it in default plugin directory
	tempPluginDirectoryName := githubAssetRelease.getPluginDirectoryName() + "_temp"
	tempPluginDirectoryPath, err := extractPluginAssetArchiveFile(ctx, pluginArchiveFilePath, tempPluginDirectoryName)
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempPluginDirectoryPath)

	// validate the extracted plugin files
	// if plugin is invalid, delete all of its files
	err = opts.validatePlugin(tempPluginDirectoryPath)
	if err != nil {
		return err
	}

	// rename old plugin directory to <plugin-directory>_old so we can rollback in case something goes wrong
	oldPluginDirectoryPath := existingPlugin.PluginDirectoryPath + "_old"
	err = os.Rename(existingPlugin.PluginDirectoryPath, oldPluginDirectoryPath)
	if err != nil {
		return err
	}
	defer os.RemoveAll(oldPluginDirectoryPath)

	// rename temp plugin directory to actual name
	// if anything goes wrong, rollback the old version of the directory
	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		err = os.Rename(oldPluginDirectoryPath, existingPlugin.PluginDirectoryPath)
		return err
	}
	pluginDirectoryPath := path.Join(pluginsDefaultDirectory, githubAssetRelease.getPluginDirectoryName())
	err = os.Rename(tempPluginDirectoryPath, pluginDirectoryPath)
	if err != nil {
		err = os.Rename(oldPluginDirectoryPath, existingPlugin.PluginDirectoryPath)
		return err
	}

	return nil
}

func (opts *UpdateOpts) Run(ctx context.Context) error {
	// if update flag is set, update all plugin, if not update only specified plugin
	if opts.UpdateAll {
		// try to create GithubAssetRelease from each plugin -  when create use it to update the plugin
		for _, p := range opts.getValidPlugins() {
			if !p.HasGithub() {
				continue
			}

			opts.Print(fmt.Sprintf(`Updating plugin "%s"`, p.Name))

			// create GithubAsset and use it to update to update plugin
			githubAsset, err := createGithubAssetFromPlugin(p, nil)
			if err != nil {
				printPluginUpdateWarning(p, err)
				continue
			}

			// update using GithubAsset
			err = opts.updatePlugin(ctx, githubAsset, p)
			if err != nil {
				printPluginUpdateWarning(p, err)
			}
		}
	} else {
		// find existing plugin using plugin args
		existingPlugin, err := opts.findPluginWithArg(opts.pluginSpecifier)
		if err != nil {
			return err
		}

		// make sure the plugin has Github values
		if !existingPlugin.HasGithub() {
			return errPluginHasNoGithubValues
		}

		// create GithubAsset and use it to update to update plugin
		githubAsset, err := createGithubAssetFromPlugin(existingPlugin, opts.pluginUpdateVersion)
		if err != nil {
			return err
		}

		// make sure the specified version is greater than currently installed version
		if githubAsset.version != nil && !githubAsset.version.GreaterThan(existingPlugin.Version) {
			return fmt.Errorf("the specified version %s is not greater than the currently installed version %s", githubAsset.version.String(), existingPlugin.Version.String())
		}

		// update using GithubAsset
		opts.Print(fmt.Sprintf(`Updating plugin "%s"`, existingPlugin.Name))
		err = opts.updatePlugin(ctx, githubAsset, existingPlugin)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateBuilder(pluginOpts *Opts) *cobra.Command {
	opts := &UpdateOpts{
		UpdateAll: false,
		ghClient:  NewAuthenticatedGithubClient(),
	}
	opts.Opts = *pluginOpts

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
				// extract plugin value and version from arg
				pluginSpecifier, version, err := extractPluginSpecifierAndVersionFromArg(arg[0])
				if err != nil {
					return err
				}
				opts.pluginSpecifier = pluginSpecifier
				opts.pluginUpdateVersion = version
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&opts.UpdateAll, flag.All, false, usage.UpdateAllPlugins)
	cmd.Flags().BoolVar(&opts.skipSignatureVerification, flag.SkipSignatureVerification, false, usage.SkipSignatureVerification)

	return cmd
}
