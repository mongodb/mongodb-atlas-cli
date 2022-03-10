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

package homebrew

import (
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/spf13/afero"
)

//go:generate mockgen -destination=../mocks/mock_brew_store.go -package=mocks github.com/mongodb/mongocli/internal/homebrew PathStore

type brewPath struct {
	CheckedPathAt  time.Time `yaml:"checked_path_at"`
	ExecutablePath string    `yaml:"path"`
	HomePath       string    `yaml:"home_path"`
}

type LoaderSaver interface {
	LoadBrewPath() (string, string, error)
	SaveBrewPath(execPath, homePath string) error
}

func NewLoaderSaver(fileSystem afero.Fs, t string) LoaderSaver {
	return &loaderSaver{fs: fileSystem, tool: t}
}

type loaderSaver struct {
	fs   afero.Fs
	tool string
}

const (
	brewFileSubPath = "/brew.yaml"
)

// LoadBrewPath will load the latest calculated brew path.
func (s *loaderSaver) LoadBrewPath() (execPath, homePath string, err error) {
	path := new(brewPath)

	filePath, err := config.Path(brewFileSubPath)
	if err != nil {
		return "", "", err
	}

	err = file.Load(s.fs, filePath, path)
	if err != nil {
		return "", "", err
	}

	if path != nil && time.Since(path.CheckedPathAt).Hours() < 24 {
		return path.ExecutablePath, path.HomePath, nil
	}
	return "", "", nil
}

// SaveBrewPath will save the latest calculated brew path.
func (s *loaderSaver) SaveBrewPath(execPath, homePath string) error {
	data := brewPath{CheckedPathAt: time.Now(), ExecutablePath: execPath, HomePath: homePath}

	filePath, err := config.Path(brewFileSubPath)
	if err != nil {
		return err
	}

	return file.Save(s.fs, filePath, data)
}
