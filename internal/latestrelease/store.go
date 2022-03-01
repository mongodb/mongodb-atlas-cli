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
	"bytes"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
)

//go:generate mockgen -destination=../mocks/mock_version_store.go -package=mocks github.com/mongodb/mongocli/internal/latestrelease Store

type State struct {
	CheckedForUpdateAt   time.Time `yaml:"checked_for_update_at"`
	LatestReleaseVersion string    `yaml:"latest_release"`
}

type BrewPath struct {
	CheckedPathAt  time.Time `yaml:"checked_path_at"`
	ExecutablePath string    `yaml:"path"`
	HomePath       string    `yaml:"home_path"`
}

type Store interface {
	LoadLatestVersion(tool string) (string, error)
	SaveLatestVersion(tool, ver string) error
	LoadBrewPath(tool string) (string, string, error)
	SaveBrewPath(tool, execPath, homePath string) error
}

func NewStore(fileSystem afero.Fs) Store {
	return &store{fs: fileSystem}
}

type store struct {
	fs afero.Fs
}

const (
	stateFileSubPath = "/rstate.yaml"
	brewFileSubPath  = "/brew.yaml"
)

// LoadLatestVersion will load the latest checked version if it was retrieved in the last 24 hours.
func (opts *store) LoadLatestVersion(tool string) (string, error) {
	latestReleaseState := new(State)
	err := opts.loadWithFileName(stateFileSubPath, tool, latestReleaseState)
	if err != nil {
		return "", err
	}

	if latestReleaseState != nil && time.Since(latestReleaseState.CheckedForUpdateAt).Hours() < 24 {
		return latestReleaseState.LatestReleaseVersion, nil
	}
	return "", nil
}

// SaveLatestVersion will save the latest retrieved version and date it was retrieved.
func (opts *store) SaveLatestVersion(tool, ver string) error {
	data := State{CheckedForUpdateAt: time.Now(), LatestReleaseVersion: ver}
	return opts.saveWithFileName(stateFileSubPath, tool, data)
}

// LoadBrewPath will load the latest calculated brew path.
func (opts *store) LoadBrewPath(tool string) (execPath, homePath string, err error) {
	path := new(BrewPath)
	err = opts.loadWithFileName(brewFileSubPath, tool, path)
	if err != nil {
		return "", "", err
	}

	if path != nil && time.Since(path.CheckedPathAt).Hours() < 24 {
		return path.ExecutablePath, path.HomePath, nil
	}
	return "", "", nil
}

// SaveBrewPath will save the latest calculated brew path.
func (opts *store) SaveBrewPath(tool, execPath, homePath string) error {
	data := BrewPath{CheckedPathAt: time.Now(), ExecutablePath: execPath, HomePath: homePath}
	return opts.saveWithFileName(brewFileSubPath, tool, data)
}

func (opts *store) loadWithFileName(fileName, tool string, data interface{}) error {
	filePath, err := filePath(tool, fileName)
	if err != nil {
		return err
	}

	err = file.Load(opts.fs, filePath, data)
	if err != nil {
		return err
	}
	return nil
}

func (opts *store) saveWithFileName(fileName, tool string, data interface{}) error {
	filePath, err := filePath(tool, fileName)
	if err != nil {
		return err
	}

	err = file.Save(opts.fs, filePath, data)
	if err != nil {
		return err
	}
	return nil
}

func filePath(tool, fileName string) (string, error) {
	var path bytes.Buffer
	var home string
	var err error

	if tool == version.AtlasCLI {
		home, err = config.AtlasCLIConfigHome()
	} else {
		home, err = config.MongoCLIConfigHome()
	}
	if err != nil {
		return "", err
	}

	path.WriteString(home)
	path.WriteString(fileName)
	return path.String(), nil
}
