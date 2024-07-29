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

var (
	errGithubParametersInvalid      = errors.New(`github parameter is invalid. It needs to have the format "<github-owner>/<github-repository-name>"`)
	errCreateExtractionDir          = errors.New("failed to create directory while extracting plugin asset")
	errCreateFileForExtracting      = errors.New("failed to create file while extracting plugin asset")
	errCopyFileContentExtraction    = errors.New("failed to copy file content while extracting plugin asset")
	errReadTar                      = errors.New("failed to read tar archive")
	errCreatePluginZipFile          = errors.New("could not create plugin zip file")
	errSaveAssetToPluginDir         = errors.New("failed to save asset to plugin directory")
	errCreateDirToExtractAssetFiles = errors.New("failed to create to plugin directory to extract assets in")
)

type GithubRelease struct {
	owner   string
	name    string
	version *semver.Version
}

type AssetOpts struct {
	existingCommands []*cobra.Command
	ghClient         *github.Client
	githubRelease    *GithubRelease
	pluginAssets     []*github.ReleaseAsset
}

func (opts *AssetOpts) repository() string {
	if opts.githubRelease == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", opts.githubRelease.owner, opts.githubRelease.name)
}

func (opts *AssetOpts) getPluginAssetInfo() ([]*github.ReleaseAsset, error) {
	var err error
	var release *github.RepositoryRelease

	// download latest release if version is not specified
	if opts.githubRelease.version == nil {
		release, _, err = opts.ghClient.Repositories.GetLatestRelease(context.Background(), opts.githubRelease.owner, opts.githubRelease.name)

		if err != nil {
			return nil, fmt.Errorf("could not find latest release for %s", opts.repository())
		}
	} else {
		// try to find the release with the version tag with v prefix, if it does not exist try again without the prefix
		release, _, err = opts.ghClient.Repositories.GetReleaseByTag(context.Background(), opts.githubRelease.owner, opts.githubRelease.name, "v"+opts.githubRelease.version.String())

		if release == nil || err != nil {
			release, _, err = opts.ghClient.Repositories.GetReleaseByTag(context.Background(), opts.githubRelease.owner, opts.githubRelease.name, opts.githubRelease.version.String())
		}

		if err != nil {
			return nil, fmt.Errorf("could not find the release %s release for %s", opts.githubRelease.name, opts.repository())
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

	return 0, fmt.Errorf("could not find an asset to download from %s for %s %s", opts.repository(), operatingSystem, architecture)
}

func (opts *AssetOpts) getPluginAssetAsReadCloser() (io.ReadCloser, error) {
	assetID, err := opts.getAssetID()

	if err != nil {
		return nil, err
	}

	rc, _, err := opts.ghClient.Repositories.DownloadReleaseAsset(context.Background(), opts.githubRelease.owner, opts.githubRelease.name, assetID, http.DefaultClient)

	if err != nil {
		return nil, fmt.Errorf("could not download asset with ID %d from %s", assetID, opts.repository())
	}

	return rc, nil
}

func parseGithubReleaseValues(arg string) (*GithubRelease, error) {
	parts := strings.Split(arg, "@")

	owner, name, err := parseGithubRepoValues(parts[0])
	if err != nil {
		return nil, err
	}

	var version *semver.Version
	versionPartsCount := 2
	if len(parts) == versionPartsCount {
		version, err = parseGithubReleaseVersion(parts[1])
		if err != nil {
			return nil, err
		}
	}

	return &GithubRelease{owner: owner, name: name, version: version}, nil
}

func parseGithubRepoValues(arg string) (string, string, error) {
	arg = strings.TrimSuffix(arg, "/")

	parts := strings.Split(arg, "/")

	minParts := 2
	if len(parts) < minParts {
		return "", "", errGithubParametersInvalid
	}

	owner := parts[len(parts)-2]
	name := parts[len(parts)-1]
	if owner == "" || name == "" {
		return "", "", errGithubParametersInvalid
	}

	return owner, name, nil
}

func parseGithubReleaseVersion(arg string) (*semver.Version, error) {
	if arg == "" || arg == "latest" {
		return nil, nil
	}

	version, err := semver.NewVersion(arg)
	if err != nil {
		return nil, fmt.Errorf(`the specified version "%s" is invalid, it needs to follow the rules of Semantic Versioning`, arg)
	}

	return version, nil
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
			return errReadTar
		}

		fileName := filepath.Clean(header.Name)
		if strings.HasPrefix(fileName, "..") {
			return fmt.Errorf("illegal file path for extracted plugin asset file: %s", fileName)
		}

		filePath := filepath.Join(dest, fileName)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filePath, os.FileMode(header.Mode)); err != nil {
				return errCreateExtractionDir
			}
		case tar.TypeReg:
			outFile, err := os.Create(filePath)
			if err != nil {
				return errCreateFileForExtracting
			}
			for {
				_, err := io.CopyN(outFile, tarReader, 1024) //nolint:mnd // 1k each write to avoid compression bomb
				if err != nil {
					if err == io.EOF {
						break
					}
					outFile.Close()
					return errCopyFileContentExtraction
				}
			}
			outFile.Close()
		default:
			return errReadTar
		}
	}

	return nil
}

func saveReadCloserToPluginAssetZipFile(rc io.ReadCloser) (string, error) {
	defer rc.Close()

	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", err
	}

	pluginZipFilePath := filepath.Join(pluginsDefaultDirectory, "zippedPluginFiles.tar.gz")
	pluginTarGzFile, err := os.Create(pluginZipFilePath)
	if err != nil {
		return "", errCreatePluginZipFile
	}
	defer pluginTarGzFile.Close()

	_, err = io.Copy(pluginTarGzFile, rc)
	if err != nil {
		os.Remove(pluginZipFilePath)
		return "", errSaveAssetToPluginDir
	}

	return pluginZipFilePath, nil
}

func (opts *AssetOpts) extractPluginAssetZipFile(pluginZipFilePath string) (string, error) {
	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", err
	}

	pluginDirectoryPath := filepath.Join(pluginsDefaultDirectory, fmt.Sprintf("%s@%s", opts.githubRelease.owner, opts.githubRelease.name))
	err = os.MkdirAll(pluginDirectoryPath, os.ModePerm)
	if err != nil {
		return "", errCreateDirToExtractAssetFiles
	}

	if err = extractTarGz(pluginZipFilePath, pluginDirectoryPath); err != nil {
		os.RemoveAll(pluginDirectoryPath)
		return pluginDirectoryPath, err
	}

	return pluginDirectoryPath, nil
}
