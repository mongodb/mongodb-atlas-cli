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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const windowsGOOS = "windows"

const matchingMarketplaceJSON = `[{"name":"mongodb-plugins","source":"git","url":"https://github.com/mongodb/agent-skills.git"}]`

const installedAtlasPluginJSON = `[{"id":"mongodb-atlas@mongodb-plugins","version":"1.2.1","scope":"user","enabled":true,"installPath":"/mock/cache/mongodb-atlas/1.2.1","mcpServers":{"mongodb-atlas":{"type":"http","url":"https://mcp.mongodb.com/"}}}]`

func newPluginInstallOpts(runner agentCommandRunner) (*installOpts, *bytes.Buffer) {
	output := new(bytes.Buffer)
	opts := &installOpts{
		agent:  claudeAgent,
		plugin: atlasPluginName,
		runner: runner,
	}
	opts.Output = jsonOutputFormat
	opts.OutWriter = output
	opts.Template = installTemplate
	return opts, output
}

func TestInstallClaudePluginAddsMarketplaceAndInstallsAtUserScope(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", "[]")
	runner.addResponse("plugin marketplace list --json", "[]")
	runner.addResponse("plugin marketplace add mongodb/agent-skills --scope user", "Marketplace added")
	runner.addResponse("plugin marketplace update mongodb-plugins", "Marketplace updated")
	runner.addResponse("plugin marketplace list --json", matchingMarketplaceJSON)
	runner.addResponse("plugin install mongodb-atlas@mongodb-plugins --scope user", "Plugin installed")
	runner.addResponse("plugin list --json", installedAtlasPluginJSON)
	opts, output := newPluginInstallOpts(runner)

	require.NoError(t, opts.Run(context.Background(), nil))
	assert.Equal(t, "claude", runner.lookedUp)
	assert.JSONEq(t, `{"installed":[{"type":"plugin","name":"mongodb-atlas","path":"/mock/cache/mongodb-atlas/1.2.1","version":"1.2.1"}]}`, output.String())
	assert.Equal(t, []string{
		"--version",
		"plugin list --json",
		"plugin marketplace list --json",
		"plugin marketplace add mongodb/agent-skills --scope user",
		"plugin marketplace update mongodb-plugins",
		"plugin marketplace list --json",
		"plugin install mongodb-atlas@mongodb-plugins --scope user",
		"plugin list --json",
	}, runner.calls)
}

func TestInstallBuilderPluginModeUsesNativeClaudeCommands(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", "[]")
	runner.addResponse("plugin marketplace list --json", matchingMarketplaceJSON)
	runner.addResponse("plugin marketplace update mongodb-plugins", "Marketplace updated")
	runner.addResponse("plugin marketplace list --json", matchingMarketplaceJSON)
	runner.addResponse("plugin install mongodb-atlas@mongodb-plugins --scope user", "Plugin installed")
	runner.addResponse("plugin list --json", installedAtlasPluginJSON)
	output := new(bytes.Buffer)
	cmd := buildInstallCommandWithRunner(newInstallTestSource(), afero.NewMemMapFs(), func() (string, error) {
		return filepath.Join(t.TempDir(), "home"), nil
	}, runner)
	cmd.SetOut(output)
	cmd.SetArgs([]string{"--plugin", "mongodb-atlas", "--agent", "claude", "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	assert.JSONEq(t, `{"installed":[{"type":"plugin","name":"mongodb-atlas","path":"/mock/cache/mongodb-atlas/1.2.1","version":"1.2.1"}]}`, output.String())
}

func TestInstallBuilderValidatesPluginSelection(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "unsupported plugin", args: []string{"--plugin", "other-plugin", "--agent", "claude"}, want: "unsupported plugin"},
		{name: "plugin and all", args: []string{"--plugin", "mongodb-atlas", "--all", "--agent", "claude"}, want: "choose exactly one"},
		{name: "plugin and skill name", args: []string{"mongodb-schema-design", "--plugin", "mongodb-atlas", "--agent", "claude"}, want: "choose exactly one"},
		{name: "plugin requires agent", args: []string{"--plugin", "mongodb-atlas"}, want: "specify --agent claude or codex"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := newFakeClaudeCommandRunner()
			cmd := buildInstallCommandWithRunner(newInstallTestSource(), afero.NewMemMapFs(), func() (string, error) {
				return t.TempDir(), nil
			}, runner)
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			require.ErrorContains(t, err, tt.want)
			assert.Empty(t, runner.calls)
		})
	}
}

func TestInstallClaudePluginRefreshesMatchingMarketplaceBeforeInstall(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", "[]")
	runner.addResponse("plugin marketplace list --json", matchingMarketplaceJSON)
	runner.addResponse("plugin marketplace update mongodb-plugins", "Marketplace updated")
	runner.addResponse("plugin marketplace list --json", matchingMarketplaceJSON)
	runner.addResponse("plugin install mongodb-atlas@mongodb-plugins --scope user", "Plugin installed")
	runner.addResponse("plugin list --json", installedAtlasPluginJSON)
	opts, output := newPluginInstallOpts(runner)

	require.NoError(t, opts.Run(context.Background(), nil))
	assert.Contains(t, output.String(), "/mock/cache/mongodb-atlas/1.2.1")
	assert.NotContains(t, runner.calls, "plugin marketplace add mongodb/agent-skills --scope user")
	assert.Equal(t, []string{
		"--version",
		"plugin list --json",
		"plugin marketplace list --json",
		"plugin marketplace update mongodb-plugins",
		"plugin marketplace list --json",
		"plugin install mongodb-atlas@mongodb-plugins --scope user",
		"plugin list --json",
	}, runner.calls)
}

func TestInstallClaudePluginRefusesMarketplaceNameCollision(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", "[]")
	runner.addResponse("plugin marketplace list --json", `[{"name":"mongodb-plugins","source":"github","repo":"someone-else/agent-skills"}]`)
	opts, _ := newPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "expected mongodb/agent-skills")
	assert.NotContains(t, runner.calls, "plugin install mongodb-atlas@mongodb-plugins --scope user")
}

func TestInstallClaudePluginRejectsWrongMarketplaceForExistingUserPlugin(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", installedAtlasPluginJSON)
	runner.addResponse("plugin marketplace list --json", `[{"name":"mongodb-plugins","source":"github","repo":"someone-else/agent-skills"}]`)
	opts, _ := newPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "expected mongodb/agent-skills")
	assert.NotContains(t, runner.calls, "plugin install mongodb-atlas@mongodb-plugins --scope user")
}

func TestInstallClaudePluginRefusesExistingPluginFromDifferentMarketplace(t *testing.T) {
	canonical := `{"id":"mongodb-atlas@mongodb-plugins","scope":"user","enabled":true}`
	other := `{"id":"mongodb-atlas@somewhere-else","scope":"user","enabled":true}`
	for _, pluginsJSON := range []string{"[" + canonical + "," + other + "]", "[" + other + "," + canonical + "]"} {
		runner := newFakeClaudeCommandRunner()
		runner.addResponse("--version", "2.1.281 (Claude Code)")
		runner.addResponse("plugin list --json", pluginsJSON)
		opts, _ := newPluginInstallOpts(runner)

		err := opts.Run(context.Background(), nil)
		require.ErrorContains(t, err, "refusing to install a duplicate")
		assert.NotContains(t, runner.calls, "plugin marketplace add mongodb/agent-skills --scope user")
	}
}

func TestInstallClaudePluginIsIdempotentForEnabledUserPlugin(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", installedAtlasPluginJSON)
	runner.addResponse("plugin marketplace list --json", matchingMarketplaceJSON)
	opts, output := newPluginInstallOpts(runner)

	require.NoError(t, opts.Run(context.Background(), nil))
	assert.JSONEq(t, `{"installed":[{"type":"plugin","name":"mongodb-atlas","path":"/mock/cache/mongodb-atlas/1.2.1","version":"1.2.1","alreadyInstalled":true}]}`, output.String())
	assert.NotContains(t, runner.calls, "plugin install mongodb-atlas@mongodb-plugins --scope user")
	assert.NotContains(t, runner.calls, "plugin marketplace update mongodb-plugins")
}

func TestInstallClaudePluginRejectsDisabledUserPlugin(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.281 (Claude Code)")
	runner.addResponse("plugin list --json", strings.Replace(installedAtlasPluginJSON, `"enabled":true`, `"enabled":false`, 1))
	opts, _ := newPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "installed but disabled")
}

func TestInstallClaudePluginRequiresValidatedVersion(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.addResponse("--version", "2.1.120 (Claude Code)")
	opts, _ := newPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "2.1.281 or later is required")
	assert.Equal(t, []string{"--version"}, runner.calls)
}

func TestParseClaudeVersionRequiresTokenBoundary(t *testing.T) {
	for _, output := range []string{"2.1.281garbage", "2.1.281.0"} {
		_, err := parseClaudeVersion([]byte(output))
		require.Error(t, err)
	}

	version, err := parseClaudeVersion([]byte("v2.1.281 (Claude Code)"))
	require.NoError(t, err)
	assert.Equal(t, "2.1.281", version.String())
}

func TestInstallClaudePluginRequiresClaudeExecutable(t *testing.T) {
	runner := newFakeClaudeCommandRunner()
	runner.lookPathErr = errors.New("not found")
	opts, _ := newPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "claude CLI was not found")
	assert.Empty(t, runner.calls)
}

func TestExecClaudeRunnerRejectsRelativePATHEntries(t *testing.T) {
	if runtime.GOOS == windowsGOOS {
		t.Chdir(t.TempDir())
		if err := os.Mkdir("bin", 0o700); err != nil {
			t.Fatal(err)
		}
		claudePath := filepath.Join("bin", "claude.exe")
		if err := os.WriteFile(claudePath, nil, 0o600); err != nil {
			t.Fatal(err)
		}
	} else {
		t.Chdir(t.TempDir())
		if err := os.Mkdir("bin", 0o700); err != nil {
			t.Fatal(err)
		}
		claudePath := filepath.Join("bin", "claude")
		if err := os.WriteFile(claudePath, []byte("#!/bin/sh\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := os.Chmod(claudePath, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", "bin")

	_, err := (execAgentCommandRunner{}).LookPath("claude")
	require.Error(t, err)
}

func TestMarketplacePointsToAgentSkills(t *testing.T) {
	tests := []struct {
		name        string
		marketplace claudeMarketplace
		want        bool
	}{
		{name: "GitHub shorthand", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "github", Repo: "mongodb/agent-skills"}, want: true},
		{name: "HTTPS git URL", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "git", URL: "https://github.com/mongodb/agent-skills.git"}, want: true},
		{name: "SSH git URL", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "git", URL: "git@github.com:mongodb/agent-skills.git"}, want: true},
		{name: "different repository", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "github", Repo: "someone/else"}, want: false},
		{name: "different marketplace", marketplace: claudeMarketplace{Name: "other-market", Source: "github", Repo: "mongodb/agent-skills"}, want: false},
		{name: "conflicting repo and URL", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "git", Repo: "mongodb/agent-skills", URL: "https://github.com/someone/agent-skills.git"}, want: false},
		{name: "insecure HTTP URL", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "git", URL: "http://github.com/mongodb/agent-skills.git"}, want: false},
		{name: "unsupported source kind", marketplace: claudeMarketplace{Name: agentSkillsMarketplaceName, Source: "directory", Repo: "mongodb/agent-skills"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, marketplacePointsToAgentSkills(tt.marketplace))
		})
	}
}

func TestValidateInstalledAtlasPluginRequiresHTTPMCPServer(t *testing.T) {
	tests := []struct {
		name        string
		server      string
		wantErrText string
	}{
		{name: "HTTP server", server: `{"type":"http","url":"https://mcp.mongodb.com/"}`},
		{name: "wrong server type", server: `{"type":"stdio","url":"https://mcp.mongodb.com/"}`, wantErrText: "expected http"},
		{name: "missing server type", server: `{"url":"https://mcp.mongodb.com/"}`, wantErrText: "expected http"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin := claudeInstalledPlugin{
				ID:          atlasPluginID,
				Version:     "1.2.1",
				Scope:       claudeUserScope,
				Enabled:     true,
				InstallPath: "/mock/cache/mongodb-atlas/1.2.1",
				MCPServers:  map[string]json.RawMessage{expectedAtlasMCPServerName: []byte(tt.server)},
			}
			err := validateInstalledAtlasPlugin(plugin)
			if tt.wantErrText == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tt.wantErrText)
		})
	}
}
