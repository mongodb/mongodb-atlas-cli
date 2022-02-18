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

package latest

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
	"github.com/google/go-github/v42/github"
	"github.com/mongodb/mongocli/internal/version"
)

const (
	minPageSize = 5
)

type VersionFinder interface {
	PrintNewVersionAvailable(w io.Writer, v, tool, bin string) error
}

func NewVersionFinder(ctx context.Context, re version.ReleaseVersionDescriber) VersionFinder {
	return &latestReleaseVersionFinder{c: ctx, r: re}
}

type latestReleaseVersionFinder struct {
	c context.Context
	r version.ReleaseVersionDescriber
}

func versionFromTag(release *github.RepositoryRelease, toolName string) string {
	if strings.Contains(release.GetTagName(), toolName) {
		return strings.ReplaceAll(release.GetTagName(), toolName+"/", "")
	}
	return release.GetTagName()
}

func isValidTagForTool(tag, tool string) bool {
	if tool == version.MongoCLI {
		return !strings.Contains(tag, version.AtlasCLI)
	}
	return strings.Contains(tag, tool)
}

func (s *latestReleaseVersionFinder) searchLatestVersionPerTool(currentVersion *semver.Version, toolName string) (bool, *version.ReleaseInformation, error) {
	release, err := s.r.LatestWithCriteria(minPageSize, isValidTagForTool, toolName)

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

func (s *latestReleaseVersionFinder) hasNewVersionAvailable(v, tool string) (newVersionAvailable bool, newVersion string, err error) {
	if v == "" {
		return false, "", nil
	}
	svCurrentVersion, err := semver.NewVersion(v)
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
	newVersionAvailable, latestVersion, err := s.hasNewVersionAvailable(v, tool)

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
