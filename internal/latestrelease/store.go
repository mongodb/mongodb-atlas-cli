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
	"time"

	"github.com/mongodb/mongocli/internal/file"
	"github.com/spf13/afero"
)

//go:generate mockgen -destination=../mocks/mock_version_store.go -package=mocks github.com/mongodb/mongocli/internal/latestrelease Store

const (
	stateFileSubPath = "/rstate.yaml"
)

type state struct {
	CheckedForUpdateAt   time.Time `yaml:"checked_for_update_at"`
	LatestReleaseVersion string    `yaml:"latest_release"`
}

type Store interface {
	LoadLatestVersion() (string, error)
	SaveLatestVersion(ver string) error
}

func NewStore(fileSystem afero.Fs, t string) Store {
	return &store{fs: fileSystem, tool: t}
}

type store struct {
	fs   afero.Fs
	tool string
}

// LoadLatestVersion will load the latest checked version if it was retrieved in the last 24 hours.
func (s *store) LoadLatestVersion() (string, error) {
	latestReleaseState := new(state)
	filePath, err := file.Path(s.tool, stateFileSubPath)
	if err != nil {
		return "", err
	}

	err = file.Load(s.fs, filePath, latestReleaseState)
	if err != nil {
		return "", err
	}

	if latestReleaseState != nil && time.Since(latestReleaseState.CheckedForUpdateAt).Hours() < 24 {
		return latestReleaseState.LatestReleaseVersion, nil
	}
	return "", nil
}

// SaveLatestVersion will save the latest retrieved version and date it was retrieved.
func (s *store) SaveLatestVersion(ver string) error {
	data := state{CheckedForUpdateAt: time.Now(), LatestReleaseVersion: ver}

	filePath, err := file.Path(s.tool, stateFileSubPath)
	if err != nil {
		return err
	}

	return file.Save(s.fs, filePath, data)
}
