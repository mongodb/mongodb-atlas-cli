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
	"fmt"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	atlasPluginName            = "mongodb-atlas"
	agentSkillsMarketplaceName = "mongodb-plugins"
	agentSkillsRepository      = upstreamOwner + "/" + upstreamRepository
	atlasPluginID              = atlasPluginName + "@" + agentSkillsMarketplaceName
	httpsScheme                = "https"
	sshScheme                  = "ssh"
)

type agentCommandRunner interface {
	LookPath(string) (string, error)
	Run(context.Context, string, ...string) ([]byte, error)
}

type execAgentCommandRunner struct{}

func (execAgentCommandRunner) LookPath(command string) (string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return "", err
	}
	if !filepath.IsAbs(path) {
		return "", fmt.Errorf("agent executable resolved to relative path %q; refusing to execute it", path)
	}
	return path, nil
}

func (execAgentCommandRunner) Run(ctx context.Context, executable string, args ...string) ([]byte, error) {
	command := exec.CommandContext(ctx, executable, args...)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	command.Stderr = &stderr
	command.Stdout = &stdout
	if err := command.Run(); err != nil {
		commandName := filepath.Base(executable)
		if len(args) > 0 {
			commandName += " " + strings.Join(args, " ")
		}
		if stderrText := strings.TrimSpace(stderr.String()); stderrText != "" {
			return nil, fmt.Errorf("agent command %q failed: %w: %s", commandName, err, stderrText)
		}
		return nil, fmt.Errorf("agent command %q failed: %w", commandName, err)
	}
	return stdout.Bytes(), nil
}

func isSecureGitHubURL(value string) bool {
	if strings.HasPrefix(strings.ToLower(value), "git@github.com:") {
		return true
	}
	parsed, err := url.Parse(value)
	if err != nil || (!strings.EqualFold(parsed.Host, "github.com") && !strings.EqualFold(parsed.Host, "www.github.com")) {
		return false
	}
	return parsed.Scheme == httpsScheme || parsed.Scheme == sshScheme
}

func normalizeGitHubRepository(repository string) string {
	repository = strings.TrimSpace(repository)
	if strings.HasPrefix(strings.ToLower(repository), "git@github.com:") {
		repository = repository[len("git@github.com:"):]
	}
	parsed, err := url.Parse(repository)
	if err == nil && parsed.Host != "" {
		if !strings.EqualFold(parsed.Host, "github.com") && !strings.EqualFold(parsed.Host, "www.github.com") {
			return ""
		}
		if parsed.Scheme != httpsScheme && parsed.Scheme != sshScheme {
			return ""
		}
		if parsed.RawQuery != "" || parsed.Fragment != "" {
			return ""
		}
		repository = parsed.Path
	}
	repository = strings.Trim(strings.TrimSuffix(repository, ".git"), "/")
	return strings.ToLower(repository)
}

func pluginName(id string) string {
	name, _, _ := strings.Cut(id, "@")
	return name
}

func (opts *installOpts) printAgentPlugin(path, version string, alreadyInstalled bool) error {
	return opts.Print(installResult{Installed: []installItem{{
		Type:             "plugin",
		Name:             atlasPluginName,
		Path:             path,
		Version:          version,
		AlreadyInstalled: alreadyInstalled,
	}}})
}
