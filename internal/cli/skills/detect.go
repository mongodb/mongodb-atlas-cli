// Copyright 2026 MongoDB Inc
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

package skills

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
)

const (
	installedAtlasClaudePluginID = "mongodb-atlas@mongodb-plugins"
	installedAtlasCodexPluginID  = installedAtlasClaudePluginID
)

// HasInstalledMongoDBSkills reports whether Claude Code or Codex has a
// MongoDB standalone skill or Atlas plugin installed in its active config home.
func HasInstalledMongoDBSkills(filesystem afero.Fs, homeDir, claudeConfigDir, codexHomeDir string) (bool, error) {
	if homeDir == "" {
		return false, errors.New("home directory is empty")
	}
	if claudeConfigDir == "" {
		claudeConfigDir = filepath.Join(homeDir, ".claude")
	}
	installed, err := hasMongoDBStandaloneSkill(filesystem, filepath.Join(claudeConfigDir, "skills"))
	if err != nil || installed {
		return installed, err
	}
	installed, err = hasUserAtlasClaudePlugin(filesystem, claudeConfigDir)
	if err != nil || installed {
		return installed, err
	}
	if codexHomeDir == "" {
		codexHomeDir = filepath.Join(homeDir, ".codex")
	}
	installed, err = hasMongoDBStandaloneSkill(filesystem, filepath.Join(codexHomeDir, "skills"))
	if err != nil || installed {
		return installed, err
	}
	return hasUserAtlasCodexPlugin(filesystem, codexHomeDir)
}

func hasMongoDBStandaloneSkill(filesystem afero.Fs, skillsDir string) (bool, error) {
	entries, err := afero.ReadDir(filesystem, skillsDir)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), "mongodb-") {
			continue
		}
		skillDir := filepath.Join(skillsDir, entry.Name())
		dirInfo, err := filesystem.Stat(skillDir)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return false, err
		}
		if !dirInfo.IsDir() {
			continue
		}
		info, err := filesystem.Stat(filepath.Join(skillDir, skillDefinitionFile))
		if err == nil && !info.IsDir() {
			return true, nil
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
	}
	return false, nil
}

func hasUserAtlasClaudePlugin(filesystem afero.Fs, claudeConfigDir string) (bool, error) {
	installedPluginsPath := filepath.Join(claudeConfigDir, "plugins", "installed_plugins.json")
	data, err := afero.ReadFile(filesystem, installedPluginsPath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	var installedPlugins struct {
		Plugins map[string][]struct {
			Scope       string `json:"scope"`
			InstallPath string `json:"installPath"`
		} `json:"plugins"`
	}
	if err := json.Unmarshal(data, &installedPlugins); err != nil {
		return false, fmt.Errorf("could not read Claude plugin install state: %w", err)
	}
	for _, plugin := range installedPlugins.Plugins[installedAtlasClaudePluginID] {
		if plugin.Scope != claudeUserScope || plugin.InstallPath == "" {
			continue
		}
		info, err := filesystem.Stat(plugin.InstallPath)
		if err == nil && info.IsDir() {
			return true, nil
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
	}
	return false, nil
}

func hasUserAtlasCodexPlugin(filesystem afero.Fs, codexHomeDir string) (bool, error) {
	configPath := filepath.Join(codexHomeDir, "config.toml")
	data, err := afero.ReadFile(filesystem, configPath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	var config struct {
		Plugins map[string]any `toml:"plugins"`
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		return false, fmt.Errorf("could not read Codex plugin install state: %w", err)
	}
	if _, ok := config.Plugins[installedAtlasCodexPluginID]; !ok {
		return false, nil
	}
	cacheRoot := filepath.Join(codexHomeDir, "plugins", "cache", "mongodb-plugins", "mongodb-atlas")
	versions, err := afero.ReadDir(filesystem, cacheRoot)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	for _, version := range versions {
		if !version.IsDir() {
			continue
		}
		manifestPath := filepath.Join(cacheRoot, version.Name(), ".codex-plugin", "plugin.json")
		info, err := filesystem.Stat(manifestPath)
		if err == nil && !info.IsDir() {
			return true, nil
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
	}
	return false, nil
}
