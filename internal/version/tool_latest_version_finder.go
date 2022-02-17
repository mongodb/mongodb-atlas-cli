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

package version

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v38/github"
)

const (
	minPageSize = 5
)

type LatestVersionFinder interface {
	PrintNewVersionAvailable(w io.Writer, v, tool, bin string) error
	IsValidTagForTool(tag, tool string) bool
}

func NewLatestVersionFinder(re ReleaseVersionDescriber) LatestVersionFinder {
	return &latestReleaseVersionFinder{ctx: context.Background(), r: re}
}

type latestReleaseVersionFinder struct {
	ctx context.Context
	r   ReleaseVersionDescriber
}

func (s *latestReleaseVersionFinder) versionFromTag(release *github.RepositoryRelease, toolName string) (version string) {
	if strings.Contains(release.GetTagName(), toolName) {
		return strings.ReplaceAll(release.GetTagName(), toolName+"/", "")
	}
	return release.GetTagName()
}

func (s *latestReleaseVersionFinder) IsValidTagForTool(tag, tool string) bool {
	if tool == mongoCLI {
		return !strings.Contains(tag, atlasCLI)
	}
	return strings.Contains(tag, tool)
}

func (s *latestReleaseVersionFinder) searchLatestVersionPerTool(currentVersion *semver.Version, toolName string) (bool, *ReleaseInformation, error) {
	release, err := s.r.LatestWithCriteria(minPageSize, s.IsValidTagForTool, toolName)

	if err != nil || release == nil {
		return false, nil, err
	}

	v, err := semver.NewVersion(s.versionFromTag(release, toolName))

	if err != nil {
		return false, nil, err
	}

	if currentVersion.Compare(v) < 0 {
		return true, &ReleaseInformation{
			Version:     v.Original(),
			PublishedAt: release.GetPublishedAt().Time,
		}, nil
	}
	return false, nil, nil
}

func isAtLeast24HoursPast(t time.Time) bool {
	return !t.IsZero() && time.Since(t) >= time.Hour*24
}

func isHomebrew(tool string) bool {
	brewFormulaPath, err := homebrewFormulaPath(tool)
	if err != nil {
		return false
	}

	executablePath, err := executableCurrentPath()
	if err != nil {
		return false
	}

	return strings.HasPrefix(executablePath, brewFormulaPath)
}

func homebrewFormulaPath(tool string) (string, error) {
	formula := tool
	brewFormulaPathBytes, err := exec.Command("brew", "--prefix", "--installed", formula).Output()
	if err != nil {
		return "", err
	}

	brewFormulaPath := strings.TrimSpace(string(brewFormulaPathBytes))

	brewFormulaPath, err = filepath.EvalSymlinks(brewFormulaPath)
	if err != nil {
		return "", err
	}

	return brewFormulaPath, nil
}

func executableCurrentPath() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executablePath, err = filepath.EvalSymlinks(executablePath)
	if err != nil {
		return "", err
	}

	return executablePath, nil
}

func (s *latestReleaseVersionFinder) HasNewVersionAvailable(version, tool string) (newVersionAvailable bool, newVersion string, err error) {
	if version == "" {
		return false, "", nil
	}
	svCurrentVersion, err := semver.NewVersion(version)
	if err != nil {
		return false, "", err
	}

	if svCurrentVersion.Prerelease() != "" { // ignoring prerelease for code changes against master
		*svCurrentVersion, err = svCurrentVersion.SetPrerelease("")
		if err != nil {
			return false, "", err
		}
	}

	newVersionAvailable, newV, err := s.searchLatestVersionPerTool(svCurrentVersion, tool)
	if err != nil {
		return false, "", err
	}

	if newVersionAvailable && (!isHomebrew(tool) || isAtLeast24HoursPast(newV.PublishedAt)) {
		return newVersionAvailable, newV.Version, nil
	}

	return false, "", nil
}

func (s *latestReleaseVersionFinder) PrintNewVersionAvailable(w io.Writer, v, tool, bin string) error {
	newVersionAvailable, latestVersion, err := s.HasNewVersionAvailable(v, tool)

	if err != nil {
		return err
	}
	if newVersionAvailable {
		var upgradeInstructions string
		if isHomebrew(tool) {
			upgradeInstructions = fmt.Sprintf(`To upgrade, run "brew update && brew upgrade %s".`, bin)
		} else {
			upgradeInstructions = fmt.Sprintf(`To upgrade, see: https://dochub.mongodb.org/core/%s-install.`, tool)
		}

		newVersionTemplate := `
A new version of %s is available '%s'!
%s

To disable this alert, run "%s config set skip_update_check true".
`
		_, err = fmt.Fprintf(w, newVersionTemplate, tool, latestVersion, upgradeInstructions, bin)
		return err
	}
	return nil
}
