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
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
)

const (
	claudeUserScope                = "user"
	claudeVersionCaptureGroupCount = 2
	minimumClaudePluginCLIVersion  = "2.1.281"
	expectedAtlasMCPServerName     = "mongodb-atlas"
	expectedAtlasMCPServerURL      = "https://mcp.mongodb.com/"
)

var claudeVersionPattern = regexp.MustCompile(`(?m)^\s*v?([0-9]+\.[0-9]+\.[0-9]+(?:-[0-9A-Za-z.-]+)?(?:\+[0-9A-Za-z.-]+)?)(?:\s|$)`)

type claudeMarketplace struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	Repo   string `json:"repo"`
	URL    string `json:"url"`
	Path   string `json:"path"`
}

type claudeInstalledPlugin struct {
	ID          string                     `json:"id"`
	Version     string                     `json:"version"`
	Scope       string                     `json:"scope"`
	Enabled     bool                       `json:"enabled"`
	InstallPath string                     `json:"installPath"`
	MCPServers  map[string]json.RawMessage `json:"mcpServers"`
}

func (opts *installOpts) installClaudePlugin(ctx context.Context) error {
	executable, err := opts.claudeExecutable(ctx)
	if err != nil {
		return err
	}
	existing, err := opts.existingAtlasPlugin(ctx, executable)
	if err != nil {
		return err
	}
	if existing != nil && existing.Scope == claudeUserScope {
		if err := opts.verifyClaudeMarketplaceSource(ctx, executable); err != nil {
			return err
		}
		return opts.printAgentPlugin(existing.InstallPath, existing.Version, true)
	}
	if err := opts.ensureClaudeMarketplace(ctx, executable); err != nil {
		return err
	}
	return opts.installClaudePluginAtUserScope(ctx, executable)
}

func (opts *installOpts) claudeExecutable(ctx context.Context) (string, error) {
	if opts.runner == nil {
		return "", errors.New("agent command runner is not configured")
	}
	executable, err := opts.runner.LookPath("claude")
	if err != nil {
		return "", fmt.Errorf("claude CLI was not found; install it to use --plugin: %w", err)
	}
	if err := validateClaudePluginVersion(ctx, opts.runner, executable); err != nil {
		return "", err
	}
	return executable, nil
}

func (opts *installOpts) existingAtlasPlugin(ctx context.Context, executable string) (*claudeInstalledPlugin, error) {
	plugins, err := opts.listClaudePlugins(ctx, executable)
	if err != nil {
		return nil, err
	}
	existing, err := findAtlasPlugin(plugins)
	if err != nil {
		return nil, err
	}
	if existing != nil && existing.Scope == claudeUserScope {
		if err := validateInstalledAtlasPlugin(*existing); err != nil {
			return nil, err
		}
	}
	return existing, nil
}

func (opts *installOpts) installClaudePluginAtUserScope(ctx context.Context, executable string) error {
	if _, err := opts.runner.Run(ctx, executable, "plugin", "install", atlasPluginID, "--scope", claudeUserScope); err != nil {
		return err
	}

	plugins, err := opts.listClaudePlugins(ctx, executable)
	if err != nil {
		return err
	}
	installed, err := findAtlasPlugin(plugins)
	if err != nil {
		return err
	}
	if installed == nil {
		return fmt.Errorf("claude Code did not report %s installed at user scope", atlasPluginID)
	}
	if err := validateInstalledAtlasPlugin(*installed); err != nil {
		return err
	}
	return opts.printAgentPlugin(installed.InstallPath, installed.Version, false)
}

func validateClaudePluginVersion(ctx context.Context, runner agentCommandRunner, executable string) error {
	output, err := runner.Run(ctx, executable, "--version")
	if err != nil {
		return fmt.Errorf("could not determine Claude Code version: %w", err)
	}
	version, err := parseClaudeVersion(output)
	if err != nil {
		return err
	}
	minimum, err := semver.NewVersion(minimumClaudePluginCLIVersion)
	if err != nil {
		return fmt.Errorf("invalid minimum Claude Code version %q: %w", minimumClaudePluginCLIVersion, err)
	}
	if version.LessThan(minimum) {
		return fmt.Errorf("claude Code %s or later is required for --plugin; found %s", minimum, version)
	}
	return nil
}

func parseClaudeVersion(output []byte) (*semver.Version, error) {
	match := claudeVersionPattern.FindSubmatch(output)
	if len(match) != claudeVersionCaptureGroupCount {
		return nil, fmt.Errorf("could not parse Claude Code version from %q", strings.TrimSpace(string(output)))
	}
	version, err := semver.NewVersion(string(match[1]))
	if err != nil {
		return nil, fmt.Errorf("could not parse Claude Code version %q: %w", match[1], err)
	}
	return version, nil
}

func (opts *installOpts) listClaudePlugins(ctx context.Context, executable string) ([]claudeInstalledPlugin, error) {
	output, err := opts.runner.Run(ctx, executable, "plugin", "list", "--json")
	if err != nil {
		return nil, err
	}
	var plugins []claudeInstalledPlugin
	if err := json.Unmarshal(output, &plugins); err != nil {
		return nil, fmt.Errorf("could not parse Claude Code plugin list: %w", err)
	}
	return plugins, nil
}

func (opts *installOpts) ensureClaudeMarketplace(ctx context.Context, executable string) error {
	marketplaces, err := opts.listClaudeMarketplaces(ctx, executable)
	if err != nil {
		return err
	}
	exists, err := verifyClaudeMarketplace(marketplaces)
	if err != nil {
		return err
	}
	if !exists {
		if _, err := opts.runner.Run(ctx, executable, "plugin", "marketplace", "add", agentSkillsRepository, "--scope", claudeUserScope); err != nil {
			return err
		}
	}
	if _, err := opts.runner.Run(ctx, executable, "plugin", "marketplace", "update", agentSkillsMarketplaceName); err != nil {
		return fmt.Errorf("could not refresh Claude Code marketplace %q: %w", agentSkillsMarketplaceName, err)
	}
	return opts.verifyClaudeMarketplaceSource(ctx, executable)
}

func (opts *installOpts) verifyClaudeMarketplaceSource(ctx context.Context, executable string) error {
	marketplaces, err := opts.listClaudeMarketplaces(ctx, executable)
	if err != nil {
		return err
	}
	exists, err := verifyClaudeMarketplace(marketplaces)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("claude Code marketplace %q is not registered to %s", agentSkillsMarketplaceName, agentSkillsRepository)
	}
	return nil
}

func (opts *installOpts) listClaudeMarketplaces(ctx context.Context, executable string) ([]claudeMarketplace, error) {
	output, err := opts.runner.Run(ctx, executable, "plugin", "marketplace", "list", "--json")
	if err != nil {
		return nil, err
	}
	var marketplaces []claudeMarketplace
	if err := json.Unmarshal(output, &marketplaces); err != nil {
		return nil, fmt.Errorf("could not parse Claude Code marketplace list: %w", err)
	}
	return marketplaces, nil
}

func verifyClaudeMarketplace(marketplaces []claudeMarketplace) (bool, error) {
	found := false
	for _, marketplace := range marketplaces {
		if marketplace.Name != agentSkillsMarketplaceName {
			continue
		}
		found = true
		if !marketplacePointsToAgentSkills(marketplace) {
			return false, fmt.Errorf("claude marketplace %q points to %q; expected %s", agentSkillsMarketplaceName, marketplaceRepository(marketplace), agentSkillsRepository)
		}
	}
	return found, nil
}

func marketplacePointsToAgentSkills(marketplace claudeMarketplace) bool {
	if marketplace.Name != agentSkillsMarketplaceName {
		return false
	}
	if !strings.EqualFold(marketplace.Source, "github") && !strings.EqualFold(marketplace.Source, "git") {
		return false
	}
	if strings.EqualFold(marketplace.Source, "git") && marketplace.URL == "" {
		return false
	}
	if marketplace.Repo != "" && normalizeGitHubRepository(marketplace.Repo) != agentSkillsRepository {
		return false
	}
	if marketplace.URL != "" && (!isSecureGitHubURL(marketplace.URL) || normalizeGitHubRepository(marketplace.URL) != agentSkillsRepository) {
		return false
	}
	return marketplace.Repo != "" || marketplace.URL != ""
}

func marketplaceRepository(marketplace claudeMarketplace) string {
	if marketplace.Repo != "" {
		return marketplace.Repo
	}
	if marketplace.URL != "" {
		return marketplace.URL
	}
	return marketplace.Path
}

func findAtlasPlugin(plugins []claudeInstalledPlugin) (*claudeInstalledPlugin, error) {
	var userPlugin *claudeInstalledPlugin
	for i := range plugins {
		if pluginName(plugins[i].ID) != atlasPluginName {
			continue
		}
		if plugins[i].ID != atlasPluginID {
			return nil, fmt.Errorf("claude Code already has %s installed as %s; refusing to install a duplicate", atlasPluginName, plugins[i].ID)
		}
		if plugins[i].Scope == claudeUserScope && userPlugin == nil {
			userPlugin = &plugins[i]
		}
	}
	return userPlugin, nil
}

func validateInstalledAtlasPlugin(plugin claudeInstalledPlugin) error {
	if !plugin.Enabled {
		return fmt.Errorf("claude plugin %s is installed but disabled; enable it with `claude plugin enable %s --scope user`", atlasPluginID, atlasPluginID)
	}
	if plugin.InstallPath == "" {
		return fmt.Errorf("claude Code did not report an install path for %s", atlasPluginID)
	}
	server, ok := plugin.MCPServers[expectedAtlasMCPServerName]
	if !ok {
		return fmt.Errorf("claude Code did not report the %s MCP server for %s", expectedAtlasMCPServerName, atlasPluginID)
	}
	var serverConfig struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(server, &serverConfig); err != nil {
		return fmt.Errorf("could not parse %s MCP server metadata: %w", expectedAtlasMCPServerName, err)
	}
	if serverConfig.Type != "http" {
		return fmt.Errorf("claude MCP server %s has type %q; expected http", expectedAtlasMCPServerName, serverConfig.Type)
	}
	if serverConfig.URL != expectedAtlasMCPServerURL {
		return fmt.Errorf("claude Code reports MCP server %s at %q; expected %s", expectedAtlasMCPServerName, serverConfig.URL, expectedAtlasMCPServerURL)
	}
	return nil
}
