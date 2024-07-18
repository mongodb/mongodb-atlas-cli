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
	"net/http"
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

type InstallOpts struct {
	cli.OutputOpts
	ghClient        *github.Client
	repositoryOwner string
	repositoryName  string
	releaseVersion  string
	plugins         []*plugin.Plugin
	pluginAssets    []*github.ReleaseAsset
}

func (opts *InstallOpts) fullRepositoryDefinition() string {
	return fmt.Sprintf("%s/%s", opts.repositoryOwner, opts.repositoryName)
}

func parseGithubValues(arg string) (string, string, error) {
	arg = strings.Split(arg, "@")[0]
	arg = strings.TrimSuffix(arg, "/")

	parts := strings.Split(arg, "/")

	err := errors.New(`github parameter is invalid. It needs to have the format "<github-owner>/<github-repository-name>"`)

	minParts := 2
	if len(parts) < minParts {
		return "", "", err
	}
	owner := parts[len(parts)-2]
	name := parts[len(parts)-1]

	if owner == "" || name == "" {
		return "", "", err
	}

	return owner, name, nil
}

func parseReleaseVersion(arg string) (string, error) {
	parts := strings.Split(arg, "@")

	minParts := 2
	if len(parts) < minParts {
		return "", nil
	}

	version := parts[1]

	if version == "" || version == "latest" {
		return "", nil
	}

	parsedVersion, err := semver.NewVersion(version)

	if err != nil {
		return "", errors.New(`the specified version is invalid, it needs to follow the rules of Semantic Versioning`)
	}

	return parsedVersion.String(), nil
}

func (opts *InstallOpts) getPluginAssets() ([]*github.ReleaseAsset, error) {
	var err error
	var release *github.RepositoryRelease

	// download latest release if version is not specified
	if opts.releaseVersion == "" {
		release, _, err = opts.ghClient.Repositories.GetLatestRelease(context.Background(), opts.repositoryOwner, opts.repositoryName)

		if err != nil {
			return nil, fmt.Errorf("could not find latest release for %s", opts.fullRepositoryDefinition())
		}
	} else {
		// try to find the release with the version tag with v prefix, if it does not exist try again without the prefix
		release, _, err = opts.ghClient.Repositories.GetReleaseByTag(context.Background(), opts.repositoryOwner, opts.repositoryName, "v"+opts.releaseVersion)

		if release == nil || err != nil {
			release, _, err = opts.ghClient.Repositories.GetReleaseByTag(context.Background(), opts.repositoryOwner, opts.repositoryName, opts.releaseVersion)
		}

		if err != nil {
			return nil, fmt.Errorf("could not find the release %s release for %s", opts.releaseVersion, opts.fullRepositoryDefinition())
		}
	}

	return release.Assets, nil
}

func (opts *InstallOpts) getAssetID() (int64, error) {
	operatingSystem, architecture := runtime.GOOS, runtime.GOARCH
	for _, asset := range opts.pluginAssets {
		if *asset.ContentType != "application/gzip" {
			continue
		}
		fmt.Printf("%s | type: %s\n", *asset.Name, *asset.ContentType)
		name := *asset.Name

		if strings.Contains(name, operatingSystem) && strings.Contains(name, architecture) {
			return *asset.ID, nil
		}
	}

	return 0, fmt.Errorf("could not find an asset to download from %s for %s %s", opts.fullRepositoryDefinition(), operatingSystem, architecture)
}

func (opts *InstallOpts) Run() error {
	assetID, err := opts.getAssetID()

	if err != nil {
		return err
	}

	rc, _, err := opts.ghClient.Repositories.DownloadReleaseAsset(context.Background(), opts.repositoryOwner, opts.repositoryName, assetID, http.DefaultClient)

	if err != nil {
		return fmt.Errorf("could not download asset with ID %d from %s", assetID, opts.fullRepositoryDefinition())
	}
	defer rc.Close()

	return opts.Print(fmt.Sprintf("Plugin %s successfully installed", opts.fullRepositoryDefinition()))
}

func InstallBuilder(plugins []*plugin.Plugin) *cobra.Command {
	opts := &InstallOpts{plugins: plugins}
	const use = "install"
	cmd := &cobra.Command{
		Use:     use + " <github-owner>/<github-repository-name>",
		Aliases: cli.GenerateAliases(use),
		Short:   "Install Atlas CLI plugin from a GitHub repository.",
		Long: `Install an Atlas CLI plugin from a GitHub repository.
The GitHub repository can be specified using either the "<github-owner>/<github-repository-name>" format or a full URL.
By default, the latest release on GitHub will be used to install the plugin.
If a specific version is needed, it can be specified using the --version flag.

An example plugin can be found here: https://github.com/mongodb/atlas-cli-plugin-example
`,
		Args: require.ExactArgs(1),
		Example: `  # Install latest version of plugin:
  atlas plugin install mongodb/atlas-cli-plugin-example
  atlas plugin install https://github.com/mongodb/atlas-cli-plugin-example
  
  # Install a specific version of plugin:
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

			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Printf("Operating System: %s\n", runtime.GOOS)
			fmt.Printf("Architecture: %s\n", runtime.GOARCH)
			return opts.Run()
		},
	}

	return cmd
}
