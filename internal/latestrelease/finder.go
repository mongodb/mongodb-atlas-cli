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
	"github.com/google/go-github/v42/github"
	"github.com/mongodb/mongocli/internal/version"
)

const (
	minPageSize = 5
)

type VersionFinder interface {
	NewVersionAvailable(isHomebrew bool) (newVersion string, err error)
	StoredLatestVersionAvailable() (needRefresh bool, foundVersion string, err error)
}

func NewVersionFinder(d version.ReleaseVersionDescriber, s Store, t, c string) VersionFinder {
	return &latestReleaseVersionFinder{
		describer:      d,
		store:          s,
		tool:           t,
		currentVersion: c,
	}
}

type latestReleaseVersionFinder struct {
	describer      version.ReleaseVersionDescriber
	store          Store
	tool           string
	currentVersion string
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

func isAtLeast24HoursPast(t time.Time) bool {
	return !t.IsZero() && time.Since(t) >= time.Hour*24
}

func (f *latestReleaseVersionFinder) searchLatestVersionPerTool(currentVersion *semver.Version) (bool, *version.ReleaseInformation, error) {
	release, err := f.describer.LatestWithCriteria(minPageSize, isValidTagForTool, f.tool)
	if err != nil || release == nil {
		return false, nil, err
	}

	v, err := semver.NewVersion(versionFromTag(release, f.tool))
	if err != nil {
		return false, nil, err
	}

	if currentVersion.Compare(v) < 0 {
		return true, &version.ReleaseInformation{
			Version:     v.Original(),
			PublishedAt: release.GetPublishedAt().Time,
		}, nil
	}

	_ = f.store.SaveLatestVersion(v.Original())
	return false, nil, nil
}

func (f *latestReleaseVersionFinder) StoredLatestVersionAvailable() (needRefresh bool, foundVersion string, err error) {
	latestVersionStored, _ := f.store.LoadLatestVersion()
	v, err := semver.NewVersion(latestVersionStored)
	// if empty or invalid, fetch from GitHub
	if err != nil {
		return true, "", err
	}
	// found a valid store higher latest version, no need to refresh
	currentVersion, err := semver.NewVersion(f.currentVersion)
	if err != nil {
		return false, "", err
	}

	if currentVersion.Compare(v) < 0 {
		return false, v.Original(), nil
	}
	// found a lower or equal latest version, no need to refresh
	return false, "", nil
}

func (f *latestReleaseVersionFinder) NewVersionAvailable(isHomebrew bool) (newVersion string, err error) {
	if f.currentVersion == "" {
		return "", nil
	}

	svCurrentVersion, err := semver.NewVersion(f.currentVersion)
	if err != nil {
		return "", err
	}

	if svCurrentVersion.Prerelease() != "" { // ignoring prerelease for code changes against master
		*svCurrentVersion, err = svCurrentVersion.SetPrerelease("")
		if err != nil {
			return "", err
		}
	}

	newVersionAvailable, newV, err := f.searchLatestVersionPerTool(svCurrentVersion)
	if err != nil {
		return "", err
	}

	if newVersionAvailable && (!isHomebrew || isAtLeast24HoursPast(newV.PublishedAt)) {
		_ = f.store.SaveLatestVersion(newV.Version)
		return newV.Version, nil
	}

	return "", nil
}
