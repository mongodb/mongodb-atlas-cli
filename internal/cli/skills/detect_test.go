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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasInstalledMongoDBSkills(t *testing.T) {
	const home = "/home/test"
	fs := afero.NewMemMapFs()
	installed, err := HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.False(t, installed)

	writeDetectionTestFile(t, fs, filepath.Join(home, ".claude", "skills", "mongodb-connection", "SKILL.md"), "---\nname: mongodb-connection\n---")
	installed, err = HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsIgnoresUnrelatedSkills(t *testing.T) {
	const home = "/home/test"
	fs := afero.NewMemMapFs()
	writeDetectionTestFile(t, fs, filepath.Join(home, ".claude", "skills", "other-skill", "SKILL.md"), "---\nname: other-skill\n---")

	installed, err := HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.False(t, installed)
}

func TestHasInstalledMongoDBSkillsRecognizesSymlinkedSkill(t *testing.T) {
	home := t.TempDir()
	filesystem := afero.NewOsFs()
	targetDir := filepath.Join(home, "shared", "mongodb-connection")
	require.NoError(t, os.MkdirAll(targetDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(targetDir, skillDefinitionFile), []byte("skill"), 0o600))
	skillsDir := filepath.Join(home, ".claude", "skills")
	require.NoError(t, os.MkdirAll(skillsDir, 0o755))
	linkPath := filepath.Join(skillsDir, "mongodb-connection")
	if err := os.Symlink(targetDir, linkPath); err != nil {
		t.Skipf("directory symlinks are unavailable: %v", err)
	}

	installed, err := HasInstalledMongoDBSkills(filesystem, home, "", "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsRecognizesUserAtlasPlugin(t *testing.T) {
	const home = "/home/test"
	claudeConfigDir := filepath.Join(home, ".claude")
	fs := afero.NewMemMapFs()
	pluginPath := filepath.Join(claudeConfigDir, "plugins", "cache", "mongodb-plugins", "mongodb-atlas", "1.2.1")
	require.NoError(t, fs.MkdirAll(pluginPath, 0o755))
	installedPluginsPath := filepath.Join(claudeConfigDir, "plugins", "installed_plugins.json")
	writeDetectionTestFile(t, fs, installedPluginsPath, fmt.Sprintf(`{"plugins":{"%s":[{"scope":"user","installPath":%q}]}}`, installedAtlasClaudePluginID, pluginPath))

	installed, err := HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsReadsConfiguredClaudeDirectory(t *testing.T) {
	const home = "/home/test"
	claudeConfigDir := "/tmp/custom-claude-config"
	fs := afero.NewMemMapFs()
	pluginPath := filepath.Join(claudeConfigDir, "plugins", "cache", "mongodb-plugins", "mongodb-atlas", "1.2.1")
	require.NoError(t, fs.MkdirAll(pluginPath, 0o755))
	installedPluginsPath := filepath.Join(claudeConfigDir, "plugins", "installed_plugins.json")
	writeDetectionTestFile(t, fs, installedPluginsPath, fmt.Sprintf(`{"plugins":{"%s":[{"scope":"user","installPath":%q}]}}`, installedAtlasClaudePluginID, pluginPath))

	installed, err := HasInstalledMongoDBSkills(fs, home, claudeConfigDir, "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsUsesActiveClaudeConfigDirectory(t *testing.T) {
	const home = "/home/test"
	fs := afero.NewMemMapFs()
	writeDetectionTestFile(t, fs, filepath.Join(home, ".claude", "skills", "mongodb-connection", "SKILL.md"), "skill")

	installed, err := HasInstalledMongoDBSkills(fs, home, "/tmp/custom-claude-config", "")
	require.NoError(t, err)
	assert.False(t, installed)
}

func TestHasInstalledMongoDBSkillsRecognizesCodexStandaloneSkill(t *testing.T) {
	const home = "/home/test"
	fs := afero.NewMemMapFs()
	writeDetectionTestFile(t, fs, filepath.Join(home, ".codex", "skills", "mongodb-connection", "SKILL.md"), "skill")

	installed, err := HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsUsesConfiguredCodexHome(t *testing.T) {
	const home = "/home/test"
	const codexHome = "/tmp/custom-codex-home"
	fs := afero.NewMemMapFs()
	writeDetectionTestFile(t, fs, filepath.Join(codexHome, "skills", "mongodb-connection", "SKILL.md"), "skill")

	installed, err := HasInstalledMongoDBSkills(fs, home, "", codexHome)
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsRecognizesCodexAtlasPlugin(t *testing.T) {
	const home = "/home/test"
	const codexHome = "/home/test/.codex"
	fs := afero.NewMemMapFs()
	pluginPath := filepath.Join(codexHome, "plugins", "cache", "mongodb-plugins", "mongodb-atlas", "1.2.1")
	writeDetectionTestFile(t, fs, filepath.Join(pluginPath, ".codex-plugin", "plugin.json"), `{}`)
	writeDetectionTestFile(t, fs, filepath.Join(codexHome, "config.toml"), "[plugins.\"mongodb-atlas@mongodb-plugins\"]\nenabled = true\n")

	installed, err := HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func TestHasInstalledMongoDBSkillsRecognizesDisabledCodexAtlasPlugin(t *testing.T) {
	const home = "/home/test"
	codexHome := filepath.Join(home, ".codex")
	fs := afero.NewMemMapFs()
	pluginPath := filepath.Join(codexHome, "plugins", "cache", "mongodb-plugins", "mongodb-atlas", "1.2.1")
	writeDetectionTestFile(t, fs, filepath.Join(pluginPath, ".codex-plugin", "plugin.json"), `{}`)
	writeDetectionTestFile(t, fs, filepath.Join(codexHome, "config.toml"), "[plugins.\"mongodb-atlas@mongodb-plugins\"]\nenabled = false\n")

	installed, err := HasInstalledMongoDBSkills(fs, home, "", "")
	require.NoError(t, err)
	assert.True(t, installed)
}

func writeDetectionTestFile(t *testing.T, fs afero.Fs, path, content string) {
	t.Helper()
	require.NoError(t, fs.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, afero.WriteFile(fs, path, []byte(content), 0o600))
}
