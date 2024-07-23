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
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

type AssetOpts struct {
	existingCommands []*cobra.Command
	ghClient         *github.Client
	repositoryOwner  string
	repositoryName   string
	releaseVersion   string
	pluginAssets     []*github.ReleaseAsset
}

func (opts *AssetOpts) fullRepositoryDefinition() string {
	return fmt.Sprintf("%s/%s", opts.repositoryOwner, opts.repositoryName)
}

func (opts *AssetOpts) getPluginAssets() ([]*github.ReleaseAsset, error) {
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

func (opts *AssetOpts) getAssetID() (int64, error) {
	operatingSystem, architecture := runtime.GOOS, runtime.GOARCH
	for _, asset := range opts.pluginAssets {
		if *asset.ContentType != "application/gzip" {
			continue
		}
		name := *asset.Name

		if strings.Contains(name, operatingSystem) && strings.Contains(name, architecture) {
			return *asset.ID, nil
		}
	}

	return 0, fmt.Errorf("could not find an asset to download from %s for %s %s", opts.fullRepositoryDefinition(), operatingSystem, architecture)
}

func (opts *AssetOpts) getPluginAssetAsReadCloser() (io.ReadCloser, error) {
	assetID, err := opts.getAssetID()

	if err != nil {
		return nil, err
	}

	rc, _, err := opts.ghClient.Repositories.DownloadReleaseAsset(context.Background(), opts.repositoryOwner, opts.repositoryName, assetID, http.DefaultClient)

	if err != nil {
		return nil, fmt.Errorf("could not download asset with ID %d from %s", assetID, opts.fullRepositoryDefinition())
	}

	return rc, nil
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

func extractTarGz(src string, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open .tar.gz file: %w", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New("failed to read tar archive")
		}

		fileName := filepath.Clean(header.Name)
		if strings.HasPrefix(fileName, "..") {
			return fmt.Errorf("illegal file path for extracted plugin asset file: %s", fileName)
		}

		filePath := filepath.Join(dest, fileName)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filePath, os.FileMode(header.Mode)); err != nil {
				return errors.New("failed to create directory while extracting plugin asset")
			}
		case tar.TypeReg:
			outFile, err := os.Create(filePath)
			if err != nil {
				return errors.New("failed to create file while extracting plugin asset")
			}
			defer outFile.Close()

			// each extracted file can be a maximum of 0.5GB in size to avoid G110 DoS decompression bomb
			maxBytesToCopy := 500000000
			if _, err := io.CopyN(outFile, tarReader, int64(maxBytesToCopy)); err != nil && err != io.EOF {
				return errors.New("failed to copy file content while extracting plugin asset")
			}
		default:
			return errors.New("failed to read tar archive")
		}
	}

	return nil
}

func saveReadCloserToPluginAssetZipFile(rc io.ReadCloser) (string, error) {
	defer rc.Close()

	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", errors.New("could not find plugin directory")
	}

	pluginZipFilePath := filepath.Join(pluginsDefaultDirectory, "zippedPluginFiles.tar.gz")
	pluginTarGzFile, err := os.Create(pluginZipFilePath)
	if err != nil {
		return "", errors.New("could not open plugins directory")
	}
	defer pluginTarGzFile.Close()

	_, err = io.Copy(pluginTarGzFile, rc)
	if err != nil {
		return "", errors.New("failed to save asset to plugin directory: " + err.Error())
	}

	return pluginZipFilePath, nil
}

func (opts *AssetOpts) extractPluginAssetZipFile(pluginZipFilePath string) (string, error) {
	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", errors.New("could not find plugin directory")
	}

	pluginDirectoryPath := filepath.Join(pluginsDefaultDirectory, fmt.Sprintf("%s@%s", opts.repositoryOwner, opts.repositoryName))
	err = os.MkdirAll(pluginDirectoryPath, os.ModePerm)
	if err != nil {
		return "", errors.New("could not create to plugin directory to extract assets in")
	}

	if err = extractTarGz(pluginZipFilePath, pluginDirectoryPath); err != nil {
		return pluginDirectoryPath, err
	}

	return pluginDirectoryPath, nil
}
