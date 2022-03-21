// Copyright 2022 MongoDB Inc
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

package latestrelease

import (
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
)

const (
	stateFileSubPath = "/rstate.yaml"
	minPageSize      = 5
)

// ReleaseInformation Release information.
type ReleaseInformation struct {
	CheckedAt   time.Time `yaml:"saved_at"`
	PublishedAt time.Time `yaml:"released_at"`
	Version     string    `yaml:"latest_version"`
}

type VersionFinder interface {
	Find() (releaseInfo *ReleaseInformation, err error)
}

func NewVersionFinder(fs afero.Fs, d version.ReleaseVersionDescriber) (VersionFinder, error) {
	filePath, err := config.Path(stateFileSubPath)
	if err != nil {
		return nil, err
	}

	return &finder{
		describer:      d,
		filesystem:     fs,
		tool:           config.ToolName,
		currentVersion: VersionFromTag(version.Version, config.ToolName),
		path:           filePath,
	}, nil
}

type finder struct {
	describer      version.ReleaseVersionDescriber
	filesystem     afero.Fs
	tool           string
	currentVersion string
	path           string
}

func VersionFromTag(ver, toolName string) string {
	if prefix := toolName + "/"; strings.HasPrefix(ver, prefix) {
		return strings.ReplaceAll(ver, prefix, "")
	}
	return ver
}

func isValidTagForTool(tag, tool string) bool {
	if tool == config.MongoCLI {
		return !strings.Contains(tag, config.AtlasCLI)
	}
	return strings.Contains(tag, tool)
}

func (f *finder) find() (*ReleaseInformation, error) {
	release, err := f.describer.LatestWithCriteria(minPageSize, isValidTagForTool, f.tool)
	if err != nil || release == nil {
		return nil, err
	}

	latestFoundRelease := &ReleaseInformation{
		Version:     VersionFromTag(release.GetTagName(), f.tool),
		PublishedAt: release.GetPublishedAt().Time,
		CheckedAt:   time.Now(),
	}
	_ = f.save(latestFoundRelease)

	return latestFoundRelease, nil
}

func (f *finder) loadOrGet() (*ReleaseInformation, *semver.Version, error) {
	if newestRelease, err := f.load(); newestRelease != nil && err == nil {
		ver, err := semver.NewVersion(newestRelease.Version)
		if err == nil {
			return newestRelease, ver, nil
		}
	}

	newestRelease, err := f.find()
	if err != nil || newestRelease == nil {
		return nil, nil, err
	}

	ver, err := semver.NewVersion(newestRelease.Version)
	return newestRelease, ver, err
}

func (f *finder) Find() (releaseInfo *ReleaseInformation, err error) {
	svCurrentVersion, err := semver.NewVersion(f.currentVersion)
	if err != nil {
		return nil, err
	}

	// ignoring prerelease for code changes against master
	if svCurrentVersion.Prerelease() != "" {
		*svCurrentVersion, err = svCurrentVersion.SetPrerelease("")
		if err != nil {
			return nil, err
		}
	}

	releaseInfo, newestVersion, err := f.loadOrGet()
	if err != nil || releaseInfo == nil || svCurrentVersion.Compare(newestVersion) >= 0 {
		return nil, err
	}
	return releaseInfo, nil
}

func (f *finder) load() (*ReleaseInformation, error) {
	latestReleaseState := new(ReleaseInformation)
	err := file.Load(f.filesystem, f.path, latestReleaseState)
	if err != nil {
		return nil, err
	}

	if latestReleaseState != nil && time.Since(latestReleaseState.CheckedAt).Hours() < 24 {
		return latestReleaseState, nil
	}
	return nil, nil
}

func (f *finder) save(ver *ReleaseInformation) error {
	return file.Save(f.filesystem, f.path, ver)
}
