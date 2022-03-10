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
	"github.com/mongodb/mongocli/internal/version"
)

const (
	minPageSize = 5
)

type VersionFinder interface {
	NewVersionAvailable() (newVersion string, err error)
	StoredLatestVersionAvailable() (needRefresh bool, foundVersion string, err error)
}

func NewVersionFinder(d version.ReleaseVersionDescriber, s LoaderSaver, t, c string) VersionFinder {
	return &latestReleaseVersionFinder{
		describer:      d,
		store:          s,
		tool:           t,
		currentVersion: versionFromTag(c, t),
	}
}

type latestReleaseVersionFinder struct {
	describer      version.ReleaseVersionDescriber
	store          LoaderSaver
	tool           string
	currentVersion string
}

func versionFromTag(ver, toolName string) string {
	if prefix := toolName + "/"; strings.HasPrefix(ver, prefix) {
		return strings.ReplaceAll(ver, prefix, "")
	}
	return ver
}

const (
	mongoCLI = "mongocli"
	atlasCLI = "atlascli"
)

func isValidTagForTool(tag, tool string) bool {
	if tool == mongoCLI {
		return !strings.Contains(tag, atlasCLI)
	}
	return strings.Contains(tag, tool)
}

func isAtLeast24HoursPast(t time.Time) bool {
	return !t.IsZero() && time.Since(t) >= time.Hour*24
}

// ReleaseInformation Release information.
type ReleaseInformation struct {
	Version     string
	PublishedAt time.Time
}

func (f *latestReleaseVersionFinder) searchLatestVersionPerTool(currentVersion *semver.Version) (bool, *ReleaseInformation, error) {
	release, err := f.describer.LatestWithCriteria(minPageSize, isValidTagForTool, f.tool)
	if err != nil || release == nil {
		return false, nil, err
	}

	v, err := semver.NewVersion(versionFromTag(release.GetTagName(), f.tool))
	if err != nil {
		return false, nil, err
	}

	if currentVersion.Compare(v) < 0 {
		return true, &ReleaseInformation{
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
	// found a valid version that is higher than latest version, no need to refresh
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

func (f *latestReleaseVersionFinder) NewVersionAvailable() (newVersion string, err error) {
	if f.currentVersion == "" {
		return "", nil
	}

	f.currentVersion = versionFromTag(f.currentVersion, f.tool)
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

	if newVersionAvailable {
		_ = f.store.SaveLatestVersion(newV.Version)
		return newV.Version, nil
	}

	return "", nil
}
