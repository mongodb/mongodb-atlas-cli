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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
)

const (
	codexGitSourceType = "git"
)

type codexMarketplaceSource struct {
	SourceType string `json:"sourceType"`
	Source     string `json:"source"`
}

type codexMarketplace struct {
	Name              string                 `json:"name"`
	MarketplaceSource codexMarketplaceSource `json:"marketplaceSource"`
}

type codexMarketplaceList struct {
	Marketplaces []codexMarketplace `json:"marketplaces"`
}

type codexPlugin struct {
	ID                string                 `json:"pluginId"`
	Name              string                 `json:"name"`
	MarketplaceName   string                 `json:"marketplaceName"`
	Version           string                 `json:"version"`
	Installed         bool                   `json:"installed"`
	Enabled           bool                   `json:"enabled"`
	MarketplaceSource codexMarketplaceSource `json:"marketplaceSource"`
}

type codexPluginList struct {
	Installed []codexPlugin `json:"installed"`
}

type codexPluginInstallResult struct {
	ID            string `json:"pluginId"`
	Version       string `json:"version"`
	InstalledPath string `json:"installedPath"`
}

func (opts *installOpts) installCodexPlugin(ctx context.Context) error {
	executable, err := opts.codexExecutable()
	if err != nil {
		return err
	}
	existing, err := opts.existingCodexPlugin(ctx, executable)
	if err != nil {
		return err
	}
	if existing != nil {
		if err := validateInstalledCodexPlugin(*existing); err != nil {
			return err
		}
		if err := verifyCodexPluginMarketplace(*existing); err != nil {
			return err
		}
		if err := opts.verifyCodexMarketplaceSource(ctx, executable); err != nil {
			return err
		}
		path, err := opts.codexPluginCachePath(existing.Version)
		if err != nil {
			return err
		}
		return opts.printAgentPlugin(path, existing.Version, true)
	}
	if err := opts.ensureCodexMarketplace(ctx, executable); err != nil {
		return err
	}
	return opts.installCodexPluginFromMarketplace(ctx, executable)
}

func (opts *installOpts) existingCodexPlugin(ctx context.Context, executable string) (*codexPlugin, error) {
	plugins, err := opts.listCodexPlugins(ctx, executable)
	if err != nil {
		return nil, err
	}
	return findCodexAtlasPlugin(plugins.Installed)
}

func (opts *installOpts) installCodexPluginFromMarketplace(ctx context.Context, executable string) error {
	output, err := opts.runner.Run(ctx, executable, "plugin", "add", atlasPluginID, "--json")
	if err != nil {
		return err
	}
	var result codexPluginInstallResult
	if err := json.Unmarshal(output, &result); err != nil {
		return fmt.Errorf("could not parse Codex plugin install result: %w", err)
	}
	if result.ID != atlasPluginID || result.Version == "" || result.InstalledPath == "" {
		return fmt.Errorf("codex did not report %s installed with a version and cache path", atlasPluginID)
	}
	installed, err := opts.existingCodexPlugin(ctx, executable)
	if err != nil {
		return err
	}
	if installed == nil {
		return fmt.Errorf("codex did not report %s installed", atlasPluginID)
	}
	if err := validateInstalledCodexPlugin(*installed); err != nil {
		return err
	}
	if err := verifyCodexPluginMarketplace(*installed); err != nil {
		return err
	}
	if installed.Version != result.Version {
		return fmt.Errorf("codex installed %s version %q but reported %q after installation", atlasPluginID, installed.Version, result.Version)
	}
	return opts.printAgentPlugin(result.InstalledPath, result.Version, false)
}

func (opts *installOpts) codexExecutable() (string, error) {
	if opts.runner == nil {
		return "", errors.New("agent command runner is not configured")
	}
	executable, err := opts.runner.LookPath("codex")
	if err != nil {
		return "", fmt.Errorf("codex CLI was not found; install a version with plugin support to use --plugin: %w", err)
	}
	return executable, nil
}

func (opts *installOpts) listCodexPlugins(ctx context.Context, executable string) (codexPluginList, error) {
	output, err := opts.runner.Run(ctx, executable, "plugin", "list", "--json")
	if err != nil {
		return codexPluginList{}, err
	}
	var plugins codexPluginList
	if err := json.Unmarshal(output, &plugins); err != nil {
		return codexPluginList{}, fmt.Errorf("could not parse codex plugin list: %w", err)
	}
	return plugins, nil
}

func (opts *installOpts) ensureCodexMarketplace(ctx context.Context, executable string) error {
	marketplaces, err := opts.listCodexMarketplaces(ctx, executable)
	if err != nil {
		return err
	}
	exists, err := verifyCodexMarketplace(marketplaces)
	if err != nil {
		return err
	}
	if !exists {
		if _, err := opts.runner.Run(ctx, executable, "plugin", "marketplace", "add", agentSkillsRepository, "--ref", "main", "--json"); err != nil {
			return err
		}
	}
	if _, err := opts.runner.Run(ctx, executable, "plugin", "marketplace", "upgrade", agentSkillsMarketplaceName, "--json"); err != nil {
		return fmt.Errorf("could not refresh codex marketplace %q: %w", agentSkillsMarketplaceName, err)
	}
	return opts.verifyCodexMarketplaceSource(ctx, executable)
}

func (opts *installOpts) verifyCodexMarketplaceSource(ctx context.Context, executable string) error {
	marketplaces, err := opts.listCodexMarketplaces(ctx, executable)
	if err != nil {
		return err
	}
	exists, err := verifyCodexMarketplace(marketplaces)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("codex marketplace %q is not registered to %s", agentSkillsMarketplaceName, agentSkillsRepository)
	}
	return nil
}

func (opts *installOpts) listCodexMarketplaces(ctx context.Context, executable string) (codexMarketplaceList, error) {
	output, err := opts.runner.Run(ctx, executable, "plugin", "marketplace", "list", "--json")
	if err != nil {
		return codexMarketplaceList{}, err
	}
	var marketplaces codexMarketplaceList
	if err := json.Unmarshal(output, &marketplaces); err != nil {
		return codexMarketplaceList{}, fmt.Errorf("could not parse codex marketplace list: %w", err)
	}
	return marketplaces, nil
}

func verifyCodexMarketplace(list codexMarketplaceList) (bool, error) {
	found := false
	for _, marketplace := range list.Marketplaces {
		if marketplace.Name != agentSkillsMarketplaceName {
			continue
		}
		found = true
		if marketplace.MarketplaceSource.SourceType != codexGitSourceType || !isSecureGitHubURL(marketplace.MarketplaceSource.Source) || normalizeGitHubRepository(marketplace.MarketplaceSource.Source) != agentSkillsRepository {
			return false, fmt.Errorf("codex marketplace %q points to %q; expected %s", agentSkillsMarketplaceName, marketplace.MarketplaceSource.Source, agentSkillsRepository)
		}
	}
	return found, nil
}

func findCodexAtlasPlugin(plugins []codexPlugin) (*codexPlugin, error) {
	var atlas *codexPlugin
	for i := range plugins {
		plugin := &plugins[i]
		if plugin.Name != atlasPluginName && pluginName(plugin.ID) != atlasPluginName {
			continue
		}
		if plugin.ID != atlasPluginID || plugin.MarketplaceName != agentSkillsMarketplaceName {
			return nil, fmt.Errorf("codex already has %s installed as %s; refusing to install a duplicate", atlasPluginName, plugin.ID)
		}
		if atlas == nil {
			atlas = plugin
		}
	}
	return atlas, nil
}

func validateInstalledCodexPlugin(plugin codexPlugin) error {
	if !plugin.Installed {
		return fmt.Errorf("codex does not report %s as installed", atlasPluginID)
	}
	if !plugin.Enabled {
		return fmt.Errorf("codex plugin %s is installed but disabled; enable it in the Codex config.toml", atlasPluginID)
	}
	if plugin.Version == "" {
		return fmt.Errorf("codex did not report a version for %s", atlasPluginID)
	}
	return nil
}

func verifyCodexPluginMarketplace(plugin codexPlugin) error {
	if plugin.MarketplaceSource.SourceType != codexGitSourceType || !isSecureGitHubURL(plugin.MarketplaceSource.Source) || normalizeGitHubRepository(plugin.MarketplaceSource.Source) != agentSkillsRepository {
		return fmt.Errorf("codex plugin %s points to marketplace source %q; expected %s", atlasPluginID, plugin.MarketplaceSource.Source, agentSkillsRepository)
	}
	return nil
}

func (opts *installOpts) codexConfigHome() (string, error) {
	configHome := opts.codexHomeDir
	if configHome == "" {
		home, err := opts.homeDir()
		if err != nil {
			return "", fmt.Errorf("could not determine the current user's home directory: %w", err)
		}
		if home == "" {
			return "", errors.New("could not determine the current user's home directory")
		}
		configHome = filepath.Join(home, ".codex")
	}
	return filepath.Abs(configHome)
}

func (opts *installOpts) codexPluginCachePath(version string) (string, error) {
	configHome, err := opts.codexConfigHome()
	if err != nil {
		return "", err
	}
	return filepath.Abs(filepath.Join(configHome, "plugins", "cache", agentSkillsMarketplaceName, atlasPluginName, version))
}
