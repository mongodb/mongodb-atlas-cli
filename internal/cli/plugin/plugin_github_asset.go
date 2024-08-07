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
	"regexp"
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
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
	errCreatePluginAssetFromPlugin  = errors.New("failed to create plugin asset from plugin")
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

func createGithubAssetFromPlugin(p *plugin.Plugin) (*GithubAsset, error) {
	if !p.HasGithub() {
		return nil, errCreatePluginAssetFromPlugin
	}

	return &GithubAsset{
		owner: p.Github.Owner,
		name:  p.Github.Name,
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
			return nil, fmt.Errorf("could not find the release %s release for %s", g.name, g.repository())
		}
	}

	return release.Assets, nil
}

func (g *GithubAsset) getID(assets []*github.ReleaseAsset) (int64, error) {
	operatingSystem, architecture := runtime.GOOS, runtime.GOARCH
	for _, asset := range assets {
		if *asset.ContentType != "application/gzip" {
			continue
		}
		name := *asset.Name

		if strings.Contains(name, operatingSystem) && strings.Contains(name, architecture) {
			return *asset.ID, nil
		}
	}

	return 0, fmt.Errorf("could not find an asset to download from %s for %s %s", g.repository(), operatingSystem, architecture)
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

	if version, ok := groupMap["version"]; ok && version != "latest" && version != "" {
		semverVersion, err := semver.NewVersion(version)
		if err != nil {
			return nil, fmt.Errorf(`the specified version "%s" is invalid, it needs to follow the rules of Semantic Versioning`, version)
		}
		githubRelease.version = semverVersion
	}

	return githubRelease, nil
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

func extractPluginAssetZipFile(pluginZipFilePath string, pluginDirectoryName string) (string, error) {
	pluginsDefaultDirectory, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", err
	}

	pluginDirectoryPath := filepath.Join(pluginsDefaultDirectory, pluginDirectoryName)
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
