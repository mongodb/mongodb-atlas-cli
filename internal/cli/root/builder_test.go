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

package root

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/latestrelease"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	versionV1 = "v1.0.0"
	versionV2 = "v2.0.0"
)

func TestOutputOpts_notifyIfApplicable(t *testing.T) {
	f := false
	bareV := versionV2
	tests := []struct {
		currentVersion   string
		expectNewVersion bool
		release          *github.RepositoryRelease
	}{
		{
			// New bare tag format must render the upgrade notice.
			currentVersion:   versionV1,
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &bareV, Prerelease: &f, Draft: &f},
		},
		{
			// Already on the latest version: no notice.
			currentVersion:   versionV2,
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &bareV, Prerelease: &f, Draft: &f},
		},
		{
			// No release found: no notice.
			currentVersion:   versionV1,
			expectNewVersion: false,
			release:          nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v/%v", tt.currentVersion, tt.release), func(t *testing.T) {
			prevVersion := version.Version
			version.Version = tt.currentVersion
			t.Cleanup(func() {
				version.Version = prevVersion
			})

			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)

			mockDescriber.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			bufOut := new(bytes.Buffer)
			fs := afero.NewMemMapFs()
			finder, _ := latestrelease.NewVersionFinder(fs, mockDescriber)

			notifier := &Notifier{
				currentVersion: latestrelease.VersionFromTag(version.Version),
				finder:         finder,
				filesystem:     fs,
				writer:         bufOut,
			}

			if err := notifier.notifyIfApplicable(false); err != nil {
				t.Fatalf("notifyIfApplicable() unexpected error:%v", err)
			}

			v := ""
			if tt.release != nil {
				v = latestrelease.VersionFromTag(tt.release.GetTagName())
			}

			want := ""
			if tt.expectNewVersion {
				want = fmt.Sprintf(`
A new version of atlascli is available %q!
To upgrade, see: https://dochub.mongodb.org/core/install-atlas-cli

To disable this alert, run "atlas config set skip_update_check true"
`, v)
			}

			if got := bufOut.String(); got != want {
				t.Errorf("notifyIfApplicable() got = %v, want %v", got, want)
			}
		})
	}
}

func TestShouldSetService(t *testing.T) {
	build := func(names ...string) *cobra.Command {
		parent := &cobra.Command{Use: atlas}
		for _, n := range names {
			child := &cobra.Command{Use: n}
			parent.AddCommand(child)
			parent = child
		}
		return parent
	}

	tests := []struct {
		name string
		cmd  *cobra.Command
		want bool
	}{
		{name: "regular command", cmd: build("clusters", "list"), want: true},
		{name: "config command", cmd: build("config", "set"), want: false},
		{name: "completion command", cmd: build("completion", "bash"), want: false},
		{name: "auth logout", cmd: build("auth", "logout"), want: false},
		{name: "logout shortcut", cmd: build("logout"), want: false},
		{name: "unrelated command named logout", cmd: build("foo", "logout"), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shouldSetService(tt.cmd))
		})
	}
}

func TestIsAgentSkillsCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: atlas}
	skillsCmd := &cobra.Command{Use: "skills"}
	listCmd := &cobra.Command{Use: "list"}
	otherCmd := &cobra.Command{Use: "skillful"}
	rootCmd.AddCommand(skillsCmd, otherCmd)
	skillsCmd.AddCommand(listCmd)

	assert.True(t, isAgentSkillsCommand(skillsCmd))
	assert.True(t, isAgentSkillsCommand(listCmd))
	assert.False(t, isAgentSkillsCommand(otherCmd))
}

func TestNotifyAgentSkillsIfApplicable(t *testing.T) {
	t.Run("does not warn without an AI caller", func(t *testing.T) {
		home := t.TempDir()
		cacheDir := filepath.Join(home, "cache")
		output := new(bytes.Buffer)
		require.NoError(t, notifyAgentSkillsIfApplicable(output, afero.NewOsFs(), home, "", "", cacheDir, false))
		assert.Empty(t, output.String())
	})

	t.Run("suggests installation when an AI caller has no MongoDB skills", func(t *testing.T) {
		home := t.TempDir()
		cacheDir := filepath.Join(home, "cache")
		fs := afero.NewOsFs()
		output := new(bytes.Buffer)
		require.NoError(t, notifyAgentSkillsIfApplicable(output, fs, home, "", "", cacheDir, true))
		assert.Contains(t, output.String(), "Warning: MongoDB agent skills are not installed for Claude Code or Codex")
		assert.Contains(t, output.String(), "atlas skills install --all --agent claude")
		assert.Contains(t, output.String(), "atlas skills install --all --agent codex")
		assert.Contains(t, output.String(), "atlas skills install --plugin mongodb-atlas --agent claude")
		assert.Contains(t, output.String(), "atlas skills install --plugin mongodb-atlas --agent codex")

		repeatedOutput := new(bytes.Buffer)
		require.NoError(t, notifyAgentSkillsIfApplicable(repeatedOutput, fs, home, "", "", cacheDir, true))
		assert.Empty(t, repeatedOutput.String(), "warning should be rate-limited for repeated AI CLI calls")
	})

	t.Run("does not warn when a standalone MongoDB skill is installed", func(t *testing.T) {
		home := t.TempDir()
		cacheDir := filepath.Join(home, "cache")
		fs := afero.NewOsFs()
		path := filepath.Join(home, ".claude", "skills", "mongodb-connection", "SKILL.md")
		require.NoError(t, fs.MkdirAll(filepath.Dir(path), 0o755))
		require.NoError(t, afero.WriteFile(fs, path, []byte("skill"), 0o600))
		output := new(bytes.Buffer)
		require.NoError(t, notifyAgentSkillsIfApplicable(output, fs, home, "", "", cacheDir, true))
		assert.Empty(t, output.String())
	})

	t.Run("does not warn when a Codex MongoDB skill is installed", func(t *testing.T) {
		home := t.TempDir()
		cacheDir := filepath.Join(home, "cache")
		fs := afero.NewOsFs()
		path := filepath.Join(home, ".codex", "skills", "mongodb-connection", "SKILL.md")
		require.NoError(t, fs.MkdirAll(filepath.Dir(path), 0o755))
		require.NoError(t, afero.WriteFile(fs, path, []byte("skill"), 0o600))
		output := new(bytes.Buffer)
		require.NoError(t, notifyAgentSkillsIfApplicable(output, fs, home, "", "", cacheDir, true))
		assert.Empty(t, output.String())
	})
}

func TestSkillsPostRunWarnsAIDiscoverability(t *testing.T) {
	home := t.TempDir()
	cacheDir := filepath.Join(home, ".cache")
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("XDG_CACHE_HOME", cacheDir)
	t.Setenv("LocalAppData", cacheDir)
	t.Setenv("CLAUDE_CONFIG_DIR", "")
	t.Setenv("CODEX_HOME", "")
	t.Setenv("CLAUDECODE", "1")
	rootCmd := Builder(false)
	rootCmd.SetArgs([]string{"--profile", "profile-that-does-not-exist", "skills", "list"})
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	listCmd, _, err := rootCmd.Find([]string{"skills", "list"})
	require.NoError(t, err)
	listCmd.RunE = func(*cobra.Command, []string) error { return nil }

	require.NoError(t, rootCmd.Execute())
	assert.Contains(t, stderr.String(), "Warning: MongoDB agent skills are not installed for Claude Code or Codex")
	assert.Contains(t, stderr.String(), "atlas skills install --all --agent claude")
	assert.Contains(t, stderr.String(), "atlas skills install --all --agent codex")
}

func TestMaybeWarnAboutAgentSkillsSkipsCobraCompletionCommands(t *testing.T) {
	home := t.TempDir()
	cacheDir := filepath.Join(home, ".cache")
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("XDG_CACHE_HOME", cacheDir)
	t.Setenv("LocalAppData", cacheDir)
	t.Setenv("CLAUDE_CONFIG_DIR", "")
	t.Setenv("CODEX_HOME", "")
	t.Setenv("CLAUDECODE", "1")
	rootCmd := &cobra.Command{Use: atlas}
	completionCmd := &cobra.Command{Use: "__complete"}
	rootCmd.AddCommand(completionCmd)
	output := new(bytes.Buffer)
	completionCmd.SetErr(output)

	maybeWarnAboutAgentSkills(completionCmd)

	assert.Empty(t, output.String())
	_, err := os.Stat(filepath.Join(cacheDir, "mongodb", "atlas-cli", "agent-skills-warning"))
	assert.ErrorIs(t, err, os.ErrNotExist)
}

func TestBuilderRegistersSkillsCommands(t *testing.T) {
	cmd := Builder(false)
	for _, commandPath := range [][]string{{"skills"}, {"skills", "list"}, {"skills", "install"}} {
		command, _, err := cmd.Find(commandPath)
		require.NoError(t, err)
		assert.NotNil(t, command)
		assert.Equal(t, "atlas "+strings.Join(commandPath, " "), command.CommandPath())
		profileFlag := command.LocalFlags().Lookup(flag.Profile)
		require.NotNil(t, profileFlag)
		assert.True(t, profileFlag.Hidden)
		assert.Nil(t, command.InheritedFlags().Lookup(flag.Profile))
	}
}
