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
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mongodb/mongocli/internal/config"
)

const (
	atlasToolFullName = "mongodb-atlas"
)

// Command get homebrew suitable command for a given tool.
func Command(tool string) string {
	if strings.Contains(tool, config.AtlasCLI) {
		return atlasToolFullName
	}
	return tool
}

// IsHomebrew checks if the cli was installed with homebrew.
func IsHomebrew(store LoaderSaver) bool {
	executablePath, brewFormulaPath, err := store.LoadBrewPath()
	// If one of the values was not found previously it is still a valid case - rely on the file.
	if (executablePath != "" || brewFormulaPath != "") && err == nil {
		return strings.HasPrefix(executablePath, brewFormulaPath)
	}

	executablePath, err = executableCurrentPath()
	if err != nil {
		_ = store.SaveBrewPath(executablePath, brewFormulaPath)
		return false
	}

	brewFormulaPath, err = homebrewFormulaPath()
	if err != nil {
		_ = store.SaveBrewPath(executablePath, brewFormulaPath)
		return false
	}

	_ = store.SaveBrewPath(executablePath, brewFormulaPath)
	return strings.HasPrefix(executablePath, brewFormulaPath)
}

func homebrewFormulaPath() (string, error) {
	formula := config.BinName()
	brewFormulaPathBytes, err := exec.Command("brew", "--prefix", "--installed", formula).Output()
	if err != nil {
		return "", err
	}

	brewFormulaPath := strings.TrimSpace(string(brewFormulaPathBytes))

	return filepath.EvalSymlinks(brewFormulaPath)
}

func executableCurrentPath() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(executablePath)
}
