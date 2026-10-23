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
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	emptyCodexPluginListJSON      = `{"installed":[],"available":[]}`
	matchingCodexMarketplaceJSON  = `{"marketplaces":[{"name":"mongodb-plugins","root":"/mock/marketplaces/mongodb-plugins","marketplaceSource":{"sourceType":"git","source":"https://github.com/mongodb/agent-skills.git"}}]}`
	installedCodexAtlasPluginJSON = `{"installed":[{"pluginId":"mongodb-atlas@mongodb-plugins","name":"mongodb-atlas","marketplaceName":"mongodb-plugins","version":"1.2.1","installed":true,"enabled":true,"source":{"source":"local","path":"/mock/marketplaces/mongodb-plugins/plugins/mongodb-atlas"},"marketplaceSource":{"sourceType":"git","source":"https://github.com/mongodb/agent-skills.git"}}],"available":[]}`
	codexAtlasPluginInstallJSON   = `{"pluginId":"mongodb-atlas@mongodb-plugins","name":"mongodb-atlas","marketplaceName":"mongodb-plugins","version":"1.2.1","installedPath":"/mock/codex/plugins/cache/mongodb-plugins/mongodb-atlas/1.2.1","authPolicy":"ON_INSTALL"}`
)

func newCodexPluginInstallOpts(runner agentCommandRunner) (*installOpts, *bytes.Buffer) {
	output := new(bytes.Buffer)
	opts := &installOpts{
		agent:        codexAgent,
		plugin:       atlasPluginName,
		runner:       runner,
		fs:           afero.NewMemMapFs(),
		codexHomeDir: "/mock/codex",
		homeDir:      func() (string, error) { return "/mock/home", nil },
	}
	opts.Output = jsonOutputFormat
	opts.OutWriter = output
	opts.Template = installTemplate
	return opts, output
}

func TestInstallCodexPluginAddsMarketplaceAndInstalls(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", emptyCodexPluginListJSON)
	runner.addResponse("plugin marketplace list --json", `{"marketplaces":[]}`)
	runner.addResponse("plugin marketplace add mongodb/agent-skills --ref main --json", `{"marketplaceName":"mongodb-plugins"}`)
	runner.addResponse("plugin marketplace upgrade mongodb-plugins --json", `{}`)
	runner.addResponse("plugin marketplace list --json", matchingCodexMarketplaceJSON)
	runner.addResponse("plugin add mongodb-atlas@mongodb-plugins --json", codexAtlasPluginInstallJSON)
	runner.addResponse("plugin list --json", installedCodexAtlasPluginJSON)
	opts, output := newCodexPluginInstallOpts(runner)

	require.NoError(t, opts.Run(context.Background(), nil))
	assert.Equal(t, codexAgent, runner.lookedUp)
	assert.JSONEq(t, `{"installed":[{"type":"plugin","name":"mongodb-atlas","path":"/mock/codex/plugins/cache/mongodb-plugins/mongodb-atlas/1.2.1","version":"1.2.1"}]}`, output.String())
	assert.Equal(t, []string{
		"plugin list --json",
		"plugin marketplace list --json",
		"plugin marketplace add mongodb/agent-skills --ref main --json",
		"plugin marketplace upgrade mongodb-plugins --json",
		"plugin marketplace list --json",
		"plugin add mongodb-atlas@mongodb-plugins --json",
		"plugin list --json",
	}, runner.calls)
}

func TestInstallCodexPluginRefreshesMatchingMarketplaceBeforeInstall(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", emptyCodexPluginListJSON)
	runner.addResponse("plugin marketplace list --json", matchingCodexMarketplaceJSON)
	runner.addResponse("plugin marketplace upgrade mongodb-plugins --json", `{}`)
	runner.addResponse("plugin marketplace list --json", matchingCodexMarketplaceJSON)
	runner.addResponse("plugin add mongodb-atlas@mongodb-plugins --json", codexAtlasPluginInstallJSON)
	runner.addResponse("plugin list --json", installedCodexAtlasPluginJSON)
	opts, output := newCodexPluginInstallOpts(runner)

	require.NoError(t, opts.Run(context.Background(), nil))
	assert.Contains(t, output.String(), "/mock/codex/plugins/cache/mongodb-plugins/mongodb-atlas/1.2.1")
	assert.NotContains(t, runner.calls, "plugin marketplace add mongodb/agent-skills --ref main --json")
	assert.Contains(t, runner.calls, "plugin marketplace upgrade mongodb-plugins --json")
}

func TestInstallCodexPluginRefusesMarketplaceNameCollision(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", emptyCodexPluginListJSON)
	runner.addResponse("plugin marketplace list --json", `{"marketplaces":[{"name":"mongodb-plugins","marketplaceSource":{"sourceType":"git","source":"https://github.com/someone-else/agent-skills.git"}}]}`)
	opts, _ := newCodexPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "expected mongodb/agent-skills")
	assert.NotContains(t, runner.calls, "plugin add mongodb-atlas@mongodb-plugins --json")
}

func TestInstallCodexPluginIsIdempotentForEnabledPlugin(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", installedCodexAtlasPluginJSON)
	runner.addResponse("plugin marketplace list --json", matchingCodexMarketplaceJSON)
	opts, output := newCodexPluginInstallOpts(runner)

	require.NoError(t, opts.Run(context.Background(), nil))
	assert.JSONEq(t, `{"installed":[{"type":"plugin","name":"mongodb-atlas","path":"/mock/codex/plugins/cache/mongodb-plugins/mongodb-atlas/1.2.1","version":"1.2.1","alreadyInstalled":true}]}`, output.String())
	assert.Equal(t, []string{"plugin list --json", "plugin marketplace list --json"}, runner.calls)
}

func TestInstallCodexPluginRejectsWrongMarketplaceForExistingPlugin(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", installedCodexAtlasPluginJSON)
	runner.addResponse("plugin marketplace list --json", strings.ReplaceAll(matchingCodexMarketplaceJSON, "mongodb/agent-skills.git", "someone-else/agent-skills.git"))
	opts, _ := newCodexPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "expected mongodb/agent-skills")
	assert.Equal(t, []string{"plugin list --json", "plugin marketplace list --json"}, runner.calls)
}

func TestInstallCodexPluginRejectsDisabledPlugin(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", strings.Replace(installedCodexAtlasPluginJSON, `"enabled":true`, `"enabled":false`, 1))
	opts, _ := newCodexPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "installed but disabled")
	assert.Len(t, runner.calls, 1)
}

func TestInstallCodexPluginRefusesSameNameFromAnotherMarketplace(t *testing.T) {
	otherMarketplace := strings.ReplaceAll(installedCodexAtlasPluginJSON, "mongodb-atlas@mongodb-plugins", "mongodb-atlas@other-marketplace")
	otherMarketplace = strings.ReplaceAll(otherMarketplace, "mongodb-plugins", "other-marketplace")
	runner := newFakeCodexCommandRunner()
	runner.addResponse("plugin list --json", otherMarketplace)
	opts, _ := newCodexPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "refusing to install a duplicate")
}

func TestInstallCodexPluginRequiresCodexExecutable(t *testing.T) {
	runner := newFakeCodexCommandRunner()
	runner.lookPathErr = errors.New("not found")
	opts, _ := newCodexPluginInstallOpts(runner)

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "codex CLI was not found")
	assert.Empty(t, runner.calls)
}

func TestCodexMarketplaceSourceMustBeSecureAndExpected(t *testing.T) {
	tests := []struct {
		name         string
		marketplace  codexMarketplace
		wantExists   bool
		wantConflict bool
	}{
		{
			name: "expected HTTPS repository",
			marketplace: codexMarketplace{
				Name: agentSkillsMarketplaceName,
				MarketplaceSource: codexMarketplaceSource{
					SourceType: "git",
					Source:     "https://github.com/mongodb/agent-skills.git",
				},
			},
			wantExists: true,
		},
		{
			name: "insecure repository",
			marketplace: codexMarketplace{
				Name: agentSkillsMarketplaceName,
				MarketplaceSource: codexMarketplaceSource{
					SourceType: "git",
					Source:     "http://github.com/mongodb/agent-skills.git",
				},
			},
			wantConflict: true,
		},
		{
			name: "different repository",
			marketplace: codexMarketplace{
				Name: agentSkillsMarketplaceName,
				MarketplaceSource: codexMarketplaceSource{
					SourceType: "git",
					Source:     "https://github.com/someone-else/agent-skills.git",
				},
			},
			wantConflict: true,
		},
		{
			name: "different marketplace",
			marketplace: codexMarketplace{
				Name: "other-marketplace",
				MarketplaceSource: codexMarketplaceSource{
					SourceType: "git",
					Source:     "https://github.com/mongodb/agent-skills.git",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := codexMarketplaceList{Marketplaces: []codexMarketplace{tt.marketplace}}
			exists, err := verifyCodexMarketplace(list)
			assert.Equal(t, tt.wantExists, exists)
			if tt.wantConflict {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestCodexMarketplaceMissingIsNotAnError(t *testing.T) {
	exists, err := verifyCodexMarketplace(codexMarketplaceList{})
	require.NoError(t, err)
	assert.False(t, exists)
}

func TestCodexPluginCachePathUsesCODEXHome(t *testing.T) {
	opts, _ := newCodexPluginInstallOpts(newFakeCodexCommandRunner())
	path, err := opts.codexPluginCachePath("1.2.1")
	require.NoError(t, err)
	assert.Equal(t, filepath.Join("/mock/codex", "plugins", "cache", agentSkillsMarketplaceName, atlasPluginName, "1.2.1"), path)
}
