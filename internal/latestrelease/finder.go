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
	"context"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v42/github"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
)

type VersionFinder interface {
	HasNewVersionAvailable(v, tool string) (newVersionAvailable bool, newVersion string, err error)
}

func NewVersionFinder(ctx context.Context, re version.ReleaseVersionDescriber) VersionFinder {
	return &latestReleaseVersionFinder{c: ctx, r: re, s: NewStore(afero.NewOsFs())}
}

func NewVersionFinderWithStore(ctx context.Context, re version.ReleaseVersionDescriber, store Store) VersionFinder {
	return &latestReleaseVersionFinder{c: ctx, r: re, s: store}
}

type latestReleaseVersionFinder struct {
	c context.Context
	r version.ReleaseVersionDescriber
	s Store
}

func versionFromTag(release *github.RepositoryRelease, toolName string) string {
	if prefix := toolName + "/"; strings.HasPrefix(release.GetTagName(), prefix) {
		return strings.ReplaceAll(release.GetTagName(), prefix, "")
	}
	return release.GetTagName()
}

func isValidTagForTool(tag, tool string) bool {
	if tool == version.MongoCLI {
		return !strings.Contains(tag, version.AtlasCLI)
	}
	return strings.Contains(tag, tool)
}

func (f *latestReleaseVersionFinder) searchLatestVersionPerTool(currentVersion *semver.Version, toolName string) (bool, *version.ReleaseInformation, error) {
	release, err := f.r.LatestWithCriteria(minPageSize, isValidTagForTool, toolName)

	if err != nil || release == nil {
		return false, nil, err
	}

	v, err := semver.NewVersion(versionFromTag(release, toolName))
	if err != nil {
		return false, nil, err
	}

	if currentVersion.Compare(v) < 0 {
		return true, &version.ReleaseInformation{
			Version:     v.Original(),
			PublishedAt: release.GetPublishedAt().Time,
		}, nil
	}
	return false, nil, nil
}

func (f *latestReleaseVersionFinder) storedLatestVersionAvailable(tool string, currentVersion *semver.Version) (needRefresh bool, foundVersion string, err error) {
	latestVersionStored, _ := f.s.LoadLatestVersion(tool)
	// no valid store version, need to fetch from GitHub
	if latestVersionStored == "" {
		return true, "", nil
	}
	// retrieved an invalid version, need to fetch from GitHub
	v, err := semver.NewVersion(latestVersionStored)
	if err != nil {
		return true, "", err
	}
	// found a valid store higher latest version, no need to refresh
	if currentVersion.Compare(v) < 0 {
		return false, v.Original(), nil
	}
	// found a lower or equal latest version, no need to refresh
	return false, "", nil
}

func (f *latestReleaseVersionFinder) HasNewVersionAvailable(currentV, tool string) (newVersionAvailable bool, newVersion string, err error) {
	if currentV == "" {
		return false, "", nil
	}

	svCurrentVersion, err := semver.NewVersion(currentV)
	if err != nil {
		return false, "", err
	}

	if svCurrentVersion.Prerelease() != "" { // ignoring prerelease for code changes against master
		*svCurrentVersion, err = svCurrentVersion.SetPrerelease("")
		if err != nil {
			return false, "", err
		}
	}

	needRefresh, storedLatestVersion, err := f.storedLatestVersionAvailable(tool, svCurrentVersion)
	if !needRefresh {
		return storedLatestVersion != "", storedLatestVersion, err
	}

	newVersionAvailable, newV, err := f.searchLatestVersionPerTool(svCurrentVersion, tool)
	if err != nil {
		return false, "", err
	}

	if newVersionAvailable && (!isHomebrew(tool, f.s) || isAtLeast24HoursPast(newV.PublishedAt)) {
		_ = f.s.SaveLatestVersion(tool, newV.Version)
		return newVersionAvailable, newV.Version, nil
	}

	return false, "", nil
}
