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
	"path/filepath"
	"strings"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/spf13/afero"
	exec "golang.org/x/sys/execabs"
)

const (
	atlasFormulaName = "mongodb-atlas-cli"
	brewFileSubPath  = "/brew.yaml"
)

type Checker struct {
	path string
	fs   afero.Fs
}

func NewChecker(fileSystem afero.Fs) (*Checker, error) {
	filePath, err := config.Path(brewFileSubPath)
	if err != nil {
		return nil, err
	}
	return &Checker{fs: fileSystem, path: filePath}, nil
}

func FormulaName(tool string) string {
	if strings.Contains(tool, "atlas") {
		return atlasFormulaName
	}
	return tool
}

// IsHomebrew checks if the cli was installed with homebrew.
func (s Checker) IsHomebrew() bool {
	// Load from cache
	h, err := s.load()
	if h != nil && h.ExecutablePath != "" && h.FormulaPath != "" && err == nil {
		return strings.HasPrefix(h.ExecutablePath, h.FormulaPath)
	}

	formula := FormulaName(config.BinName())
	cmdResult := new(bytes.Buffer)
	cmd := exec.Command("brew", "--prefix", formula)

	if err = cmd.Start(); err != nil {
		return false
	}

	h = new(homebrew)
	h.ExecutablePath, err = executableCurrentPath()
	if err != nil {
		return false
	}

	if err = cmd.Wait(); err != nil {
		return false
	}

	h.FormulaPath, err = filepath.EvalSymlinks(strings.TrimSpace(cmdResult.String()))
	if err != nil {
		return false
	}
	_ = s.save(h)
	return strings.HasPrefix(h.ExecutablePath, h.FormulaPath)
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

func (s *Checker) load() (*homebrew, error) {
	path := new(homebrew)
	if err := file.Load(s.fs, s.path, path); err != nil {
		return nil, err
	}

	if path != nil && time.Since(path.CheckedAt).Hours() < 24 {
		return path, nil
	}
	return nil, nil
}

func (s *Checker) save(h *homebrew) error {
	h.CheckedAt = time.Now()
	return file.Save(s.fs, s.path, h)
}
