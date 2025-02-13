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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"github.com/google/go-github/v61/github"
	"github.com/mholt/archives"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
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
	latest    = "latest"
	publicKey = "signature.asc"
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
		// download the 100 latest releases
		const MaxPerPage = 100
		releases, _, err := g.ghClient.Repositories.ListReleases(context.Background(), g.owner, g.name, &github.ListOptions{
			Page:    0,
			PerPage: MaxPerPage,
		})

		if err != nil {
			return nil, fmt.Errorf("could not fetch releases for %s, access to GitHub is required, %w", g.repository(), err)
		}

		// get the latest release that doesn't have prerelease info or metadata in the version tag
		release = getLatestStableRelease(releases)
		if release == nil {
			return nil, fmt.Errorf("could not find latest stable release for %s", g.repository())
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

func getLatestStableRelease(releases []*github.RepositoryRelease) *github.RepositoryRelease {
	var latestStableVersion *semver.Version
	var latestStableRelease *github.RepositoryRelease

	for _, release := range releases {
		version, err := semver.NewVersion(*release.TagName)

		// if we can't parse the version tag, skip this release
		if err != nil {
			continue
		}

		// if the version has pre-release info or metadata, skip this version
		if version.Prerelease() != "" || version.Metadata() != "" {
			continue
		}

		if latestStableVersion == nil || version.GreaterThan(latestStableVersion) {
			latestStableVersion = version
			latestStableRelease = release
		}
	}

	return latestStableRelease
}

var architectureAliases = map[string][]string{
	"amd64": {"x86_64"},
	"arm64": {"aarch64"},
	"386":   {"i386", "x86"},
}

//nolint:mnd
var contentTypePriority = map[string]int{
	"application/gzip":   0, // tar.gz
	"application/x-gtar": 1, // tar.gz
	"application/x-gzip": 2, // tar.gz
	"application/zip":    3, // zip
}

func (g *GithubAsset) getIDs(assets []*github.ReleaseAsset) (int64, int64, int64, error) {
	return g.getIDsForOSArch(assets, runtime.GOOS, runtime.GOARCH)
}

func (g *GithubAsset) getIDsForOSArch(assets []*github.ReleaseAsset, goos, goarch string) (int64, int64, int64, error) {
	// Get all possible architecture names for the current architecture
	archNames := []string{goarch}
	if aliases, ok := architectureAliases[goarch]; ok {
		archNames = append(archNames, aliases...)
	}

	var archiveAssets []*github.ReleaseAsset
	for _, asset := range assets {
		if asset.ContentType == nil || asset.Name == nil {
			continue
		}

		if _, ok := contentTypePriority[*asset.ContentType]; !ok {
			continue
		}

		name := strings.ToLower(*asset.Name)
		if !strings.Contains(name, goos) {
			continue
		}

		// Check if any of the architecture names match
		for _, arch := range archNames {
			if strings.Contains(name, arch) {
				archiveAssets = append(archiveAssets, asset)
				break
			}
		}
	}

	if len(archiveAssets) == 0 {
		return 0, 0, 0, fmt.Errorf("no compatible asset found in %s for OS=%s, arch=%s (including aliases: %v)",
			g.repository(), goos, goarch, archNames[1:])
	}

	// Sort by content type priority
	slices.SortFunc(archiveAssets, func(a, b *github.ReleaseAsset) int {
		return contentTypePriority[*a.ContentType] - contentTypePriority[*b.ContentType]
	})
	name := strings.ToLower(*archiveAssets[0].Name)
	signatureID, pubKeyID, err := getSignatureAssetandKeyID(name, assets)
	if err != nil {
		return 0, 0, 0, err
	}
	return *archiveAssets[0].ID, signatureID, pubKeyID, nil
}

func getSignatureAssetandKeyID(name string, assets []*github.ReleaseAsset) (int64, int64, error) {
	var signatureAsset *github.ReleaseAsset
	var pubKeyAsset *github.ReleaseAsset

	for _, asset := range assets {
		assetName := strings.ToLower(*asset.Name)

		// Check if asset is a public key
		if strings.Compare(assetName, publicKey) == 0 {
			pubKeyAsset = asset
			continue
		}

		// Check if asset is a signature
		if strings.Contains(assetName, name+".sig") {
			signatureAsset = asset
		}
	}

	// If no signature package is found, provide warning
	if signatureAsset == nil {
		_, _ = log.Warningf("-- plugin warning: no corresponding signature asset found for package %s\n", name)
		return 0, 0, nil
	}

	// If signature package exists but public key does not, return error
	if pubKeyAsset == nil {
		return 0, 0, fmt.Errorf("-- plugin warning: no public key '%s' found for signature verification", publicKey)
	}

	return *signatureAsset.ID, *pubKeyAsset.ID, nil
}

func (g *GithubAsset) getPluginAssetsAsReadCloser(assetID, sigAssetID, pubKeyAssetID int64) (io.ReadCloser, error) {
	rc, _, err := g.ghClient.Repositories.DownloadReleaseAsset(context.Background(), g.owner, g.name, assetID, http.DefaultClient)
	if err != nil {
		return nil, fmt.Errorf("could not download asset with ID %d from %s", assetID, g.repository())
	}

	asset, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("could not convert reader to bytes")
	}

	// Only do verification if IDs are not 0, i.e. when there is signature package available
	if sigAssetID != 0 && pubKeyAssetID != 0 {
		err = g.verifyAssetSignature(asset, sigAssetID, pubKeyAssetID)
		if err != nil {
			return nil, err
		}
	}

	return io.NopCloser(bytes.NewReader(asset)), nil
}

// verifyAssetSignature verifies the asset signature.
// Returns nil if signature check is successful.
func (g *GithubAsset) verifyAssetSignature(asset []byte, sigAssetID, pubKeyAssetID int64) error {
	sigRc, _, err := g.ghClient.Repositories.DownloadReleaseAsset(context.Background(), g.owner, g.name, sigAssetID, http.DefaultClient)
	if err != nil {
		return fmt.Errorf("could not download signature asset with ID %d from %s", sigAssetID, g.repository())
	}
	defer sigRc.Close()

	keyRc, _, err := g.ghClient.Repositories.DownloadReleaseAsset(context.Background(), g.owner, g.name, pubKeyAssetID, http.DefaultClient)
	if err != nil {
		return fmt.Errorf("could not download public key asset with ID %d from %s", pubKeyAssetID, g.repository())
	}
	defer keyRc.Close()

	key, err := openpgp.ReadArmoredKeyRing(keyRc)
	if err != nil {
		return err
	}

	config := &packet.Config{}
	_, err = openpgp.CheckArmoredDetachedSignature(key, bytes.NewReader(asset), sigRc, config)
	if err != nil {
		return fmt.Errorf("signature verification unsuccessful: %w", err)
	}

	fmt.Println("PGP signature verification successful!")
	return nil
}

func parseGithubReleaseValues(arg string) (*GithubAsset, error) {
	regexPattern := `^((https?://(www\.)?)?github\.com/)?(?P<owner>[\w.\-]+)/(?P<name>[\w.\-]+)/?(@(?P<version>.+))?$`
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
		version := strings.TrimPrefix(version, "v")
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
	// Strip prefix
	prefix, err := getArchivePrefix(ctx, pluginArchivePath)
	if err != nil {
		return fmt.Errorf("failed to determine archive prefix: %w", err)
	}

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
		destPath := filepath.Join(pluginDirectoryName, strings.TrimPrefix(fileInfo.NameInArchive, prefix))

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

func getArchivePrefix(ctx context.Context, pluginArchivePath string) (string, error) {
	fsys, err := archives.FileSystem(ctx, pluginArchivePath, nil)
	if err != nil {
		return "", fmt.Errorf("failed to open archive file: %w", err)
	}

	// Read the contents of the archive root
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return "", fmt.Errorf("failed to read root directory of archive: %w", err)
	}

	// Strip prefix
	prefix := ""
	if len(entries) == 1 {
		entry := entries[0]
		if entry.IsDir() {
			prefix = entry.Name()
		}
	}

	return prefix, nil
}
