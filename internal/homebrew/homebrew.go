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
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/spf13/afero"
)

//go:generate mockgen -destination=../mocks/mock_brew_store.go -package=mocks github.com/mongodb/mongocli/internal/homebrew PathStore

const atlasFormulaName = "mongodb-atlas"

// formulaName get homebrew suitable command for a given tool.
func formulaName(tool string) string {
	if strings.Contains(tool, "atlas") {
		return atlasFormulaName
	}
	return tool
}

// IsHomebrew checks if the cli was installed with homebrew.
func (s Checker) IsHomebrew() bool {
	h, err := s.Load()
	// If one of the values was not found previously it is still a valid case - rely on the file.
	if (h.ExecutablePath != "" || h.FormulaPath != "") && err == nil {
		return strings.HasPrefix(h.ExecutablePath, h.FormulaPath)
	}
	formula := formulaName(config.BinName())
	buf := new(bytes.Buffer)
	cmd := exec.Command("brew", "--prefix", "--installed", formula)

	cmd.Stdout = buf
	if err := cmd.Start(); err != nil {
		return false
	}

	executablePath, err := executableCurrentPath()
	if err != nil {
		return false
	}
	h.ExecutablePath = executablePath

	if err := cmd.Wait(); err != nil {
		return false
	}
	brewFormulaPath, err := filepath.EvalSymlinks(strings.TrimSpace(buf.String()))
	if err != nil {
		return false
	}
	h.FormulaPath = brewFormulaPath
	_ = s.Save(h)
	return strings.HasPrefix(executablePath, brewFormulaPath)
}

func executableCurrentPath() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(executablePath)
}

type homebrew struct {
	CheckedAt      time.Time `yaml:"checked_at"`
	ExecutablePath string    `yaml:"executable_path"`
	FormulaPath    string    `yaml:"formula_path"`
}

type Loader interface {
	Load() (*homebrew, error)
}

type LoaderSaver interface {
	Loader
	Save(*homebrew) error
}

func NewChecker(fileSystem afero.Fs) (*Checker, error) {
	filePath, err := config.Path(brewFileSubPath)
	if err != nil {
		return nil, err
	}
	return &Checker{fs: fileSystem, path: filePath}, nil
}

type Checker struct {
	path string
	fs   afero.Fs
}

const brewFileSubPath = "/brew.yaml"

// Load will load the latest calculated brew path.
func (s *Checker) Load() (*homebrew, error) {
	path := new(homebrew)
	if err := file.Load(s.fs, s.path, path); err != nil {
		return nil, err
	}

	if path != nil && time.Since(path.CheckedAt).Hours() < 24 {
		return path, nil
	}
	return nil, nil
}

// Save will save the latest calculated brew path.
func (s *Checker) Save(h *homebrew) error {
	return file.Save(s.fs, s.path, h)
}
