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
	"github.com/mongodb/mongocli/internal/homebrew"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
)

type Checker interface {
	CheckAvailable() error
}

func NewChecker(cv, t string, p Printer) Checker {
	return &checker{
		currentVersion: versionFromTag(cv, t),
		tool:           t,
		printer:        p,
		finder:         nil,
		filesystem:     afero.NewOsFs(),
	}
}

type checker struct {
	currentVersion string
	tool           string
	printer        Printer
	finder         VersionFinder
	filesystem     afero.Fs
}

func newCheckerForTest(cv, t string, p Printer, f VersionFinder, fs afero.Fs) Checker {
	return &checker{
		currentVersion: versionFromTag(cv, t),
		tool:           t,
		printer:        p,
		finder:         f,
		filesystem:     fs,
	}
}

func (c *checker) latestVersion() (isHomebrew bool, latestVersion string, err error) {
	releaseStore := NewStore(c.filesystem, c.tool)

	if c.finder == nil {
		c.finder = NewVersionFinder(
			version.NewReleaseVersionDescriber(),
			releaseStore,
			c.tool,
			c.currentVersion,
		)
	}

	needRefresh, latestVersion, err := c.finder.StoredLatestVersionAvailable()
	if !needRefresh && latestVersion == "" {
		return false, "", err
	}

	brewStore := homebrew.NewPathStore(c.filesystem, c.tool)
	isHomebrew = homebrew.IsHomebrew(c.tool, brewStore)

	if !needRefresh {
		return isHomebrew, latestVersion, nil
	}

	latestVersion, err = c.finder.NewVersionAvailable(isHomebrew)

	return isHomebrew, latestVersion, err
}

// CheckAvailable latest release version and returns it if found.
func (c *checker) CheckAvailable() error {
	cmd := ""
	isHomebrew, latestVersion, err := c.latestVersion()
	if err != nil || latestVersion == "" {
		return err
	}

	if isHomebrew {
		cmd = homebrew.Command(c.tool)
	}

	_ = c.printer.PrintNewVersionAvailable(latestVersion, cmd)

	return nil
}
