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
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/unicode/norm"
)

const (
	claudeAgent          = "claude"
	codexAgent           = "codex"
	maxInstallFileCount  = 1024
	maxInstallTotalBytes = 64 * 1024 * 1024
	parentPathSegment    = ".."
	agentSkillsDirMode   = os.FileMode(0o755)
)

const installTemplate = `{{range .Installed}}{{if .AlreadyInstalled}}Already installed{{else}}Installed{{end}} {{.Type}} {{.Name}}{{if .Version}} version {{.Version}}{{end}} at {{.Path}}
{{end}}`

type installItem struct {
	Type             string `json:"type"`
	Name             string `json:"name"`
	Path             string `json:"path"`
	Version          string `json:"version,omitempty"`
	AlreadyInstalled bool   `json:"alreadyInstalled,omitempty"`
}

type installResult struct {
	Installed []installItem `json:"installed"`
}

type skillInstallPlan struct {
	name  string
	dir   string
	files []skillFile
}

type installOpts struct {
	cli.OutputOpts
	source          skillSource
	fs              afero.Fs
	homeDir         func() (string, error)
	claudeConfigDir string
	codexHomeDir    string
	agent           string
	all             bool
	plugin          string
	runner          agentCommandRunner
}

func (opts *installOpts) validateSelection(args []string) error {
	if opts.agent == "" {
		return fmt.Errorf("specify --%s %s or %s", flag.Agent, claudeAgent, codexAgent)
	}
	if opts.agent != claudeAgent && opts.agent != codexAgent {
		return fmt.Errorf("unsupported agent %q: supported agents are %s and %s", opts.agent, claudeAgent, codexAgent)
	}
	if opts.plugin != "" {
		if opts.all || len(args) > 0 {
			return fmt.Errorf("choose exactly one of a skill name, --%s, or --%s", flag.All, flag.Plugin)
		}
		if opts.plugin != atlasPluginName {
			return fmt.Errorf("unsupported plugin %q: only %s is supported", opts.plugin, atlasPluginName)
		}
		return nil
	}
	if opts.all && len(args) > 0 {
		return fmt.Errorf("specify a skill name or --%s, not both", flag.All)
	}
	if opts.all {
		return nil
	}
	if len(args) != 1 {
		return fmt.Errorf("specify one skill name or use --%s", flag.All)
	}
	if !isSafePathComponent(args[0]) {
		return fmt.Errorf("invalid skill name %q", args[0])
	}
	return nil
}

func (opts *installOpts) Run(ctx context.Context, args []string) error {
	if opts.plugin != "" {
		if opts.agent == codexAgent {
			return opts.installCodexPlugin(ctx)
		}
		return opts.installClaudePlugin(ctx)
	}

	snapshot, err := opts.source.Snapshot(ctx)
	if err != nil {
		return err
	}
	names, err := opts.selectedSkillNames(snapshot, args)
	if err != nil {
		return err
	}
	skillsRoot, err := opts.skillsRoot()
	if err != nil {
		return err
	}
	plans, err := opts.createPlans(snapshot, skillsRoot, names)
	if err != nil {
		return err
	}
	installed, err := opts.installStandalone(ctx, snapshot, skillsRoot, plans)
	if err != nil {
		return err
	}
	return opts.Print(installResult{Installed: installed})
}

func (opts *installOpts) selectedSkillNames(snapshot *skillSnapshot, args []string) ([]string, error) {
	names := make([]string, 0, len(snapshot.skills))
	if opts.all {
		for _, available := range snapshot.skills {
			names = append(names, available.Name)
		}
	} else {
		names = append(names, args[0])
	}
	if len(names) == 0 {
		return nil, errors.New("no standalone skills are available in the upstream repository")
	}
	return names, nil
}

func (opts *installOpts) installStandalone(ctx context.Context, snapshot *skillSnapshot, skillsRoot string, plans []skillInstallPlan) ([]installItem, error) {
	if err := opts.fs.MkdirAll(skillsRoot, agentSkillsDirMode); err != nil {
		return nil, fmt.Errorf("could not create %s skills directory %s: %w", opts.agent, skillsRoot, err)
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	stageRoot, err := afero.TempDir(opts.fs, skillsRoot, ".atlas-skills-install-")
	if err != nil {
		return nil, fmt.Errorf("could not create temporary skill installation directory: %w", err)
	}
	defer func() {
		_ = opts.fs.RemoveAll(stageRoot)
	}()

	if err := opts.stagePlans(ctx, snapshot, plans, stageRoot); err != nil {
		return nil, err
	}
	return opts.commitPlans(ctx, plans, stageRoot)
}

func (opts *installOpts) stagePlans(ctx context.Context, snapshot *skillSnapshot, plans []skillInstallPlan, stageRoot string) error {
	var downloadedBytes int64
	for _, plan := range plans {
		if err := ctx.Err(); err != nil {
			return err
		}
		stagedDirectory, err := pathWithinRoot(stageRoot, plan.name)
		if err != nil {
			return err
		}
		if err := opts.fs.Mkdir(stagedDirectory, agentSkillsDirMode); err != nil {
			return fmt.Errorf("could not create temporary directory for skill %s: %w", plan.name, err)
		}
		for _, file := range plan.files {
			if err := opts.stageSkillFile(ctx, snapshot.commitSHA, stagedDirectory, file, &downloadedBytes); err != nil {
				return err
			}
		}
	}
	return nil
}

func (opts *installOpts) stageSkillFile(ctx context.Context, commitSHA, stagedDirectory string, file skillFile, downloadedBytes *int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	content, err := opts.source.ReadFile(ctx, commitSHA, file.repositoryPath)
	if err != nil {
		return fmt.Errorf("could not download %s: %w", file.repositoryPath, err)
	}
	if len(content) > maxSkillFileSize {
		return fmt.Errorf("upstream file %s exceeds the %d-byte size limit", file.repositoryPath, maxSkillFileSize)
	}
	if file.size > 0 && int64(len(content)) != file.size {
		return fmt.Errorf("upstream file %s changed size during download", file.repositoryPath)
	}
	*downloadedBytes += int64(len(content))
	if *downloadedBytes > maxInstallTotalBytes {
		return fmt.Errorf("install exceeds the %d-byte total size limit", maxInstallTotalBytes)
	}

	destination, err := pathWithinRoot(stagedDirectory, file.relativePath)
	if err != nil {
		return err
	}
	if err := opts.fs.MkdirAll(filepath.Dir(destination), agentSkillsDirMode); err != nil {
		return fmt.Errorf("could not create directory for %s: %w", destination, err)
	}
	if err := afero.WriteFile(opts.fs, destination, content, file.mode.Perm()); err != nil {
		return fmt.Errorf("could not write temporary skill file %s: %w", destination, err)
	}
	return ctx.Err()
}

func (opts *installOpts) commitPlans(ctx context.Context, plans []skillInstallPlan, stageRoot string) ([]installItem, error) {
	installed := make([]installItem, 0, len(plans))
	createdDirectories := make([]string, 0, len(plans))
	for _, plan := range plans {
		if err := ctx.Err(); err != nil {
			return nil, rollbackSkillInstall(opts.fs, createdDirectories, err)
		}
		exists, err := pathExistsIncludingSymlink(opts.fs, plan.dir)
		if err != nil {
			return nil, rollbackSkillInstall(opts.fs, createdDirectories, fmt.Errorf("could not inspect skill destination %s: %w", plan.dir, err))
		}
		if exists {
			return nil, rollbackSkillInstall(opts.fs, createdDirectories, fmt.Errorf("skill %q already exists at %s; refusing to overwrite it", plan.name, plan.dir))
		}
		stagedDirectory, err := pathWithinRoot(stageRoot, plan.name)
		if err != nil {
			return nil, rollbackSkillInstall(opts.fs, createdDirectories, err)
		}
		if err := opts.fs.Rename(stagedDirectory, plan.dir); err != nil {
			return nil, rollbackSkillInstall(opts.fs, createdDirectories, fmt.Errorf("could not install skill directory %s: %w", plan.dir, err))
		}
		createdDirectories = append(createdDirectories, plan.dir)
		installed = append(installed, installItem{Type: "skill", Name: plan.name, Path: plan.dir})
	}

	if err := ctx.Err(); err != nil {
		return nil, rollbackSkillInstall(opts.fs, createdDirectories, err)
	}
	return installed, nil
}

func (opts *installOpts) skillsRoot() (string, error) {
	if opts.agent == codexAgent {
		configHome, err := opts.codexConfigHome()
		if err != nil {
			return "", err
		}
		return filepath.Join(configHome, "skills"), nil
	}
	configHome := opts.claudeConfigDir
	if configHome == "" {
		home, err := opts.homeDir()
		if err != nil {
			return "", fmt.Errorf("could not determine the current user's home directory: %w", err)
		}
		if home == "" {
			return "", errors.New("could not determine the current user's home directory")
		}
		configHome = filepath.Join(home, ".claude")
	}
	root, err := filepath.Abs(filepath.Join(configHome, "skills"))
	if err != nil {
		return "", fmt.Errorf("could not resolve %s skills directory: %w", opts.agent, err)
	}
	return root, nil
}

func (opts *installOpts) createPlans(snapshot *skillSnapshot, skillsRoot string, names []string) ([]skillInstallPlan, error) {
	plans := make([]skillInstallPlan, 0, len(names))
	seenNames := make(map[string]string, len(names))
	for _, name := range names {
		nameKey := filesystemPathKey(name)
		if priorName, ok := seenNames[nameKey]; ok {
			return nil, fmt.Errorf("skill names %q and %q collide on case-insensitive filesystems", priorName, name)
		}
		seenNames[nameKey] = name

		destination, err := pathWithinRoot(skillsRoot, name)
		if err != nil {
			return nil, fmt.Errorf("invalid skill name %q: %w", name, err)
		}

		exists, err := pathExistsIncludingSymlink(opts.fs, destination)
		if err != nil {
			return nil, fmt.Errorf("could not inspect skill destination %s: %w", destination, err)
		}
		if exists {
			return nil, fmt.Errorf("skill %q already exists at %s; refusing to overwrite it", name, destination)
		}

		files, err := snapshot.filesForSkill(name)
		if err != nil {
			return nil, err
		}
		if err := validateCaseInsensitiveFilePaths(name, files); err != nil {
			return nil, err
		}

		plan := skillInstallPlan{name: name, dir: destination, files: files}
		hasSkillFile := false
		for _, file := range plan.files {
			if _, err := pathWithinRoot(destination, file.relativePath); err != nil {
				return nil, err
			}
			if file.relativePath == skillDefinitionFile {
				hasSkillFile = true
			}
		}
		if !hasSkillFile {
			return nil, fmt.Errorf("upstream skill %q does not contain SKILL.md", name)
		}
		plans = append(plans, plan)
	}
	if err := validateInstallLimits(plans); err != nil {
		return nil, err
	}

	return plans, nil
}

func validateInstallLimits(plans []skillInstallPlan) error {
	fileCount := 0
	for _, plan := range plans {
		fileCount += len(plan.files)
	}
	if fileCount > maxInstallFileCount {
		return fmt.Errorf("install contains more than %d files", maxInstallFileCount)
	}

	var totalBytes int64
	for _, plan := range plans {
		for _, file := range plan.files {
			if file.size < 0 || file.size > maxSkillFileSize {
				return fmt.Errorf("upstream file %s exceeds the %d-byte size limit", file.repositoryPath, maxSkillFileSize)
			}
			totalBytes += file.size
			if totalBytes > maxInstallTotalBytes {
				return fmt.Errorf("install exceeds the %d-byte total size limit", maxInstallTotalBytes)
			}
		}
	}
	return nil
}

func validateCaseInsensitiveFilePaths(skillName string, files []skillFile) error {
	seenFiles := make(map[string]string, len(files))
	seenDirectories := make(map[string]string, len(files))
	for _, file := range files {
		filePath := filepath.ToSlash(filepath.Clean(filepath.FromSlash(file.relativePath)))
		segments := strings.Split(filePath, "/")
		for i := 1; i < len(segments); i++ {
			directory := strings.Join(segments[:i], "/")
			directoryKey := filesystemPathKey(directory)
			if priorFile, ok := seenFiles[directoryKey]; ok {
				return fmt.Errorf("skill %q contains file path %q and directory path %q that collide on case-insensitive filesystems", skillName, priorFile, directory)
			}
			if priorDirectory, ok := seenDirectories[directoryKey]; ok && priorDirectory != directory {
				return fmt.Errorf("skill %q contains directory paths %q and %q that collide on case-insensitive filesystems", skillName, priorDirectory, directory)
			}
			seenDirectories[directoryKey] = directory
		}

		key := filesystemPathKey(filePath)
		if priorFile, ok := seenFiles[key]; ok {
			return fmt.Errorf("skill %q contains paths %q and %q that collide on case-insensitive filesystems", skillName, priorFile, file.relativePath)
		}
		if priorDirectory, ok := seenDirectories[key]; ok {
			return fmt.Errorf("skill %q contains file path %q and directory path %q that collide on case-insensitive filesystems", skillName, file.relativePath, priorDirectory)
		}
		seenFiles[key] = file.relativePath
	}
	return nil
}

func filesystemPathKey(value string) string {
	return cases.Fold().String(norm.NFD.String(value))
}

func pathWithinRoot(root, relativePath string) (string, error) {
	if !isSafeRelativePath(relativePath) {
		return "", fmt.Errorf("unsafe relative path %q", relativePath)
	}

	destination := filepath.Join(root, filepath.FromSlash(relativePath))
	relativeToRoot, err := filepath.Rel(root, destination)
	if err != nil {
		return "", fmt.Errorf("could not verify destination path %s: %w", destination, err)
	}
	if filepath.IsAbs(relativeToRoot) || relativeToRoot == parentPathSegment || strings.HasPrefix(relativeToRoot, parentPathSegment+string(filepath.Separator)) {
		return "", fmt.Errorf("destination path %s escapes %s", destination, root)
	}

	return destination, nil
}

func pathExistsIncludingSymlink(fs afero.Fs, path string) (bool, error) {
	if lstatFS, ok := fs.(afero.Lstater); ok {
		info, _, err := lstatFS.LstatIfPossible(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return false, nil
			}
			return false, err
		}
		return info != nil, nil
	}

	return afero.Exists(fs, path)
}

func rollbackSkillInstall(fs afero.Fs, directories []string, cause error) error {
	errorsToReturn := []error{cause}
	for i := len(directories) - 1; i >= 0; i-- {
		if err := fs.RemoveAll(directories[i]); err != nil {
			errorsToReturn = append(errorsToReturn, fmt.Errorf("could not remove partial skill directory %s: %w", directories[i], err))
		}
	}
	return errors.Join(errorsToReturn...)
}

// InstallBuilder builds the standalone skill install command.
func InstallBuilder() *cobra.Command {
	return buildInstallCommand(newGitHubSource(), afero.NewOsFs(), os.UserHomeDir)
}

func buildInstallCommand(source skillSource, filesystem afero.Fs, homeDir func() (string, error)) *cobra.Command {
	return buildInstallCommandWithRunner(source, filesystem, homeDir, execAgentCommandRunner{})
}

func buildInstallCommandWithRunner(source skillSource, filesystem afero.Fs, homeDir func() (string, error), runner agentCommandRunner) *cobra.Command {
	opts := &installOpts{
		source:          source,
		fs:              filesystem,
		homeDir:         homeDir,
		claudeConfigDir: os.Getenv("CLAUDE_CONFIG_DIR"),
		codexHomeDir:    os.Getenv("CODEX_HOME"),
		runner:          runner,
	}
	cmd := &cobra.Command{
		Use:   "install [skill-name]",
		Short: "Install MongoDB agent skills for Claude Code or Codex.",
		Long: `Install a standalone MongoDB agent skill to the selected agent's personal skills directory.
Use --all to install every standalone skill available in the upstream repository.
Use --plugin mongodb-atlas to install the native Atlas plugin for Claude Code or Codex.
Codex plugin/app installation is separate from its direct MCP server registry.
Use a named skill or --all to install standalone Codex skills.`,
		Example: `  # Install one skill for Claude Code:
  atlas skills install mongodb-schema-design --agent claude

  # Install one skill for Codex:
  atlas skills install mongodb-schema-design --agent codex

  # Install all standalone skills for Codex:
  atlas skills install --all --agent codex

  # Install the Atlas plugin and its MCP server configuration:
  atlas skills install --plugin mongodb-atlas --agent claude

  # Install the Atlas plugin for Codex:
  atlas skills install --plugin mongodb-atlas --agent codex`,
		Annotations: map[string]string{
			"skill-nameDesc": "Name of the standalone MongoDB agent skill to install.",
		},
		Args: require.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.validateSelection(args); err != nil {
				return err
			}
			return opts.InitOutput(cmd.OutOrStdout(), installTemplate)()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
			defer stop()
			return opts.Run(ctx, args)
		},
	}
	hideAtlasProfileFlag(cmd)

	cmd.Flags().StringVar(&opts.agent, flag.Agent, "", usage.Agent)
	cmd.Flags().BoolVar(&opts.all, flag.All, false, usage.AllSkills)
	cmd.Flags().StringVar(&opts.plugin, flag.Plugin, "", usage.AgentPlugin)
	_ = cmd.MarkFlagRequired(flag.Agent)
	opts.AddOutputOptFlags(cmd)
	return cmd
}
