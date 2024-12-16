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
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/mholt/archives"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
)

var (
	errGithubParametersInvalid      = errors.New(`github parameter is invalid. It needs to have the format "<github-owner>/<github-repository-name>"`)
	errCreatePluginArchiveFile      = errors.New("could not create plugin archive file")
	errSaveAssetToPluginDir         = errors.New("failed to save asset to plugin directory")
	errCreateDirToExtractAssetFiles = errors.New("failed to create to plugin directory to extract assets in")
	errCreatePluginAssetFromPlugin  = errors.New("failed to create plugin asset from plugin")
)

const (
	latest = "latest"
)

type GithubAsset struct {
	ghClient *github.Client
	owner    string
	name     string
	version  *semver.Version
}

func (g *GithubAsset) repository() string {
	return fmt.Sprintf("%s/%s", g.owner, g.name)
}

func createGithubAssetFromPlugin(p *plugin.Plugin, version *semver.Version) (*GithubAsset, error) {
	if !p.HasGithub() {
		return nil, errCreatePluginAssetFromPlugin
	}

	return &GithubAsset{
		owner:   p.Github.Owner,
		name:    p.Github.Name,
		version: version,
	}, nil
}

func (g *GithubAsset) getPluginDirectoryName() string {
	return fmt.Sprintf("%s@%s", g.owner, g.name)
}

func (g *GithubAsset) getReleaseAssets() ([]*github.ReleaseAsset, error) {
	var err error
	var release *github.RepositoryRelease

	// download latest release if version is not specified
	if g.version == nil {
		release, _, err = g.ghClient.Repositories.GetLatestRelease(context.Background(), g.owner, g.name)

		if err != nil {
			return nil, fmt.Errorf("could not find latest release for %s", g.repository())
		}
	} else {
		// try to find the release with the version tag with v prefix, if it does not exist try again without the prefix
		release, _, err = g.ghClient.Repositories.GetReleaseByTag(context.Background(), g.owner, g.name, "v"+g.version.String())

		if release == nil || err != nil {
			release, _, err = g.ghClient.Repositories.GetReleaseByTag(context.Background(), g.owner, g.name, g.version.String())
		}

		if err != nil {
			return nil, fmt.Errorf("could not find the release %s for %s", g.version, g.repository())
		}
	}

	return release.Assets, nil
}

func (g *GithubAsset) getID(assets []*github.ReleaseAsset) (int64, error) {
	operatingSystem, architecture := runtime.GOOS, runtime.GOARCH
	archiveContentTypesAndPriority := map[string]int{
		"application/gzip": 0,
		"application/zip":  1,
	}

	archiveAssets := make([]*github.ReleaseAsset, 0)

	for _, asset := range assets {
		if _, ok := archiveContentTypesAndPriority[*asset.ContentType]; !ok {
			continue
		}
		name := *asset.Name

		if strings.Contains(name, operatingSystem) && strings.Contains(name, architecture) {
			archiveAssets = append(archiveAssets, asset)
		}
	}

	if len(archiveAssets) == 0 {
		return 0, fmt.Errorf("could not find an asset to download from %s for %s %s", g.repository(), operatingSystem, architecture)
	}

	slices.SortFunc(archiveAssets, func(a, b *github.ReleaseAsset) int {
		priortyA := archiveContentTypesAndPriority[*a.ContentType]
		priortyB := archiveContentTypesAndPriority[*b.ContentType]
		return priortyA - priortyB
	})

	return *archiveAssets[0].ID, nil
}

func (g *GithubAsset) getPluginAssetAsReadCloser(assetID int64) (io.ReadCloser, error) {
	rc, _, err := g.ghClient.Repositories.DownloadReleaseAsset(context.Background(), g.owner, g.name, assetID, http.DefaultClient)

	if err != nil {
		return nil, fmt.Errorf("could not download asset with ID %d from %s", assetID, g.repository())
	}

	return rc, nil
}

func parseGithubReleaseValues(arg string) (*GithubAsset, error) {
	regexPattern := `^((https?://(www\.)?)?github\.com/)?(?P<owner>[\w.\-]+)/(?P<name>[\w.\-]+)/?(@(?P<version>v?(\d+)(\.\d+)?(\.\d+)?|latest))?$`
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, fmt.Errorf("error compiling regex: %w", err)
	}

	matches := regex.FindStringSubmatch(arg)
	if matches == nil {
		return nil, errGithubParametersInvalid
	}

	names := regex.SubexpNames()
	groupMap := make(map[string]string)
	for i, match := range matches {
		if i == 0 {
			continue
		}
		groupMap[names[i]] = match
	}

	githubRelease := &GithubAsset{owner: groupMap["owner"], name: groupMap["name"]}

	if version, ok := groupMap["version"]; ok && version != latest && version != "" {
		semverVersion, err := semver.NewVersion(version)
		if err != nil {
			return nil, fmt.Errorf(`the specified version "%s" is invalid, it needs to follow the rules of Semantic Versioning`, version)
		}
		githubRelease.version = semverVersion
	}

	return githubRelease, nil
}

func saveReadCloserToPluginAssetArchiveFile(rc io.ReadCloser) (string, error) {
	defer rc.Close()

	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", err
	}

	pluginArchiveFilePath := filepath.Join(pluginsDefaultDirectory, "plugin.partial")
	pluginTarGzFile, err := os.Create(pluginArchiveFilePath)
	if err != nil {
		return "", errCreatePluginArchiveFile
	}
	defer pluginTarGzFile.Close()

	_, err = io.Copy(pluginTarGzFile, rc)
	if err != nil {
		os.Remove(pluginArchiveFilePath)
		return "", errSaveAssetToPluginDir
	}

	return pluginArchiveFilePath, nil
}

func extractPluginAssetArchiveFile(ctx context.Context, pluginArchivePath string, pluginDirectoryName string) (string, error) {
	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", err
	}

	pluginDirectoryPath := filepath.Join(pluginsDefaultDirectory, pluginDirectoryName)
	err = os.MkdirAll(pluginDirectoryPath, os.ModePerm)
	if err != nil {
		return "", errCreateDirToExtractAssetFiles
	}

	if err = extractArchive(ctx, pluginArchivePath, pluginDirectoryPath); err != nil {
		os.RemoveAll(pluginDirectoryPath)
		return pluginDirectoryPath, err
	}

	return pluginDirectoryPath, nil
}

func extractArchive(ctx context.Context, pluginArchivePath string, pluginDirectoryName string) error {
	archiveFile, err := os.Open(pluginArchivePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer archiveFile.Close()

	// Identify the archive format
	// The library we're using supports: zip, .tar, .tar.gz, .rar, .7z
	format, _, err := archives.Identify(ctx, pluginArchivePath, archiveFile)
	if err != nil {
		return fmt.Errorf("failed to identify archive format: %w", err)
	}

	// Try to get an extractor for the format
	ex, ok := format.(archives.Extractor)
	if !ok {
		return fmt.Errorf("%s is not supported", format.MediaType())
	}

	// Extract the archive
	if err := ex.Extract(ctx, archiveFile, func(_ context.Context, fileInfo archives.FileInfo) error {
		// Get the destination path
		destPath := filepath.Join(pluginDirectoryName, fileInfo.NameInArchive)

		// Handle directories
		if fileInfo.IsDir() {
			return os.MkdirAll(destPath, fileInfo.Mode())
		}

		// Only handle regular files
		if !fileInfo.Mode().IsRegular() {
			return fmt.Errorf("plugin archive should only contain directoreis and regular files, encountered: %s", fileInfo.Mode())
		}

		// Create parent directories if they don't exist
		if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create parent directory: %w", err)
		}

		// Open file in archive
		file, err := fileInfo.Open()
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		// Create the file
		destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileInfo.Mode())
		if err != nil {
			return fmt.Errorf("failed to create destination file: %w", err)
		}
		defer destFile.Close()

		// Copy file contents
		if _, err := io.Copy(destFile, file); err != nil {
			return fmt.Errorf("failed to copy contents to destination file: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to extract archive: %w", err)
	}

	return nil
}
