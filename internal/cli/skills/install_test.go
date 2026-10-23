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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const betaSkillMarkdownPath = "skills/beta/SKILL.md"

type installTestSource struct {
	snapshot   *skillSnapshot
	files      map[string][]byte
	failReadAt string
}

func (source *installTestSource) Snapshot(context.Context) (*skillSnapshot, error) {
	return source.snapshot, nil
}

func (source *installTestSource) ReadFile(_ context.Context, _ string, path string) ([]byte, error) {
	if path == source.failReadAt {
		return nil, errors.New("simulated upstream read failure")
	}
	content, ok := source.files[path]
	if !ok {
		return nil, fmt.Errorf("file %s not found in test source", path)
	}
	return content, nil
}

func newInstallTestSource() *installTestSource {
	return &installTestSource{
		snapshot: &skillSnapshot{
			commitSHA: "test-commit",
			skills:    []Skill{{Name: "alpha"}, {Name: "beta"}},
			filesByName: map[string][]skillFile{
				"alpha": {
					{repositoryPath: "skills/alpha/SKILL.md", relativePath: "SKILL.md", mode: 0o644},
					{repositoryPath: "skills/alpha/references/guide.md", relativePath: "references/guide.md", mode: 0o644},
					{repositoryPath: "skills/alpha/scripts/run.py", relativePath: "scripts/run.py", mode: 0o755},
				},
				"beta": {
					{repositoryPath: betaSkillMarkdownPath, relativePath: "SKILL.md", mode: 0o644},
				},
			},
		},
		files: map[string][]byte{
			"skills/alpha/SKILL.md":            []byte("---\ndescription: Alpha\n---\n"),
			"skills/alpha/references/guide.md": []byte("Reference content"),
			"skills/alpha/scripts/run.py":      []byte("print('hello')"),
			betaSkillMarkdownPath:              []byte("---\ndescription: Beta\n---\n"),
		},
	}
}

func TestInstallBuilderInstallsNamedSkillAndSupportFiles(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	output := new(bytes.Buffer)
	cmd := buildInstallCommand(newInstallTestSource(), fs, func() (string, error) { return home, nil })
	cmd.SetOut(output)
	cmd.SetArgs([]string{"alpha", "--agent", "claude", "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	assert.JSONEq(t, fmt.Sprintf(`{"installed":[{"type":"skill","name":"alpha","path":%q}]}`, filepath.Join(home, ".claude", "skills", "alpha")), output.String())

	for path, expected := range map[string]string{
		filepath.Join(home, ".claude", "skills", "alpha", "SKILL.md"):               "---\ndescription: Alpha\n---\n",
		filepath.Join(home, ".claude", "skills", "alpha", "references", "guide.md"): "Reference content",
		filepath.Join(home, ".claude", "skills", "alpha", "scripts", "run.py"):      "print('hello')",
	} {
		content, err := afero.ReadFile(fs, path)
		require.NoError(t, err)
		assert.Equal(t, expected, string(content))
	}

	skillInfo, err := fs.Stat(filepath.Join(home, ".claude", "skills", "alpha", "SKILL.md"))
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o644), skillInfo.Mode().Perm())
	scriptInfo, err := fs.Stat(filepath.Join(home, ".claude", "skills", "alpha", "scripts", "run.py"))
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o755), scriptInfo.Mode().Perm())
}

func TestInstallBuilderInstallsCodexSkillUnderCODEXHome(t *testing.T) {
	fs := afero.NewMemMapFs()
	codexHome := filepath.Join(t.TempDir(), "codex-home")
	home := filepath.Join(t.TempDir(), "home")
	output := new(bytes.Buffer)
	t.Setenv("CODEX_HOME", codexHome)
	cmd := buildInstallCommand(newInstallTestSource(), fs, func() (string, error) { return home, nil })
	cmd.SetOut(output)
	cmd.SetArgs([]string{"alpha", "--agent", codexAgent, "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	wanted := filepath.Join(codexHome, "skills", "alpha")
	assert.JSONEq(t, fmt.Sprintf(`{"installed":[{"type":"skill","name":"alpha","path":%q}]}`, wanted), output.String())
	assertFSExists(t, fs, filepath.Join(wanted, "SKILL.md"), true)
	assertFSExists(t, fs, filepath.Join(wanted, "references", "guide.md"), true)
}

func TestInstallBuilderDefaultsCodexSkillToUserCodexHome(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	output := new(bytes.Buffer)
	t.Setenv("CODEX_HOME", "")
	cmd := buildInstallCommand(newInstallTestSource(), fs, func() (string, error) { return home, nil })
	cmd.SetOut(output)
	cmd.SetArgs([]string{"alpha", "--agent", codexAgent, "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	wanted := filepath.Join(home, ".codex", "skills", "alpha")
	assert.JSONEq(t, fmt.Sprintf(`{"installed":[{"type":"skill","name":"alpha","path":%q}]}`, wanted), output.String())
	assertFSExists(t, fs, filepath.Join(wanted, "SKILL.md"), true)
}

func TestInstallBuilderUsesConfiguredClaudeDirectory(t *testing.T) {
	fs := afero.NewMemMapFs()
	claudeConfigDir := filepath.Join(t.TempDir(), "claude-config")
	home := filepath.Join(t.TempDir(), "home")
	output := new(bytes.Buffer)
	t.Setenv("CLAUDE_CONFIG_DIR", claudeConfigDir)
	cmd := buildInstallCommand(newInstallTestSource(), fs, func() (string, error) { return home, nil })
	cmd.SetOut(output)
	cmd.SetArgs([]string{"alpha", "--agent", claudeAgent, "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	assert.JSONEq(t, fmt.Sprintf(`{"installed":[{"type":"skill","name":"alpha","path":%q}]}`, filepath.Join(claudeConfigDir, "skills", "alpha")), output.String())
}

func TestInstallNamedSkillDoesNotReadOtherSkillMetadata(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	source := newInstallTestSource()
	source.failReadAt = betaSkillMarkdownPath
	opts := installOpts{
		source:  source,
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}
	opts.Output = jsonOutputFormat
	opts.OutWriter = new(bytes.Buffer)
	opts.Template = installTemplate

	require.NoError(t, opts.Run(context.Background(), []string{"alpha"}))
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "alpha", "SKILL.md"), true)
}

func TestInstallBuilderInstallsAllSkills(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	output := new(bytes.Buffer)
	cmd := buildInstallCommand(newInstallTestSource(), fs, func() (string, error) { return home, nil })
	cmd.SetOut(output)
	cmd.SetArgs([]string{"--all", "--agent", "claude", "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	assert.JSONEq(t, fmt.Sprintf(`{"installed":[{"type":"skill","name":"alpha","path":%q},{"type":"skill","name":"beta","path":%q}]}`,
		filepath.Join(home, ".claude", "skills", "alpha"),
		filepath.Join(home, ".claude", "skills", "beta")), output.String())
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "alpha", "SKILL.md"), true)
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "beta", "SKILL.md"), true)
}

func TestInstallBuilderInstallsAllSkillsForCodex(t *testing.T) {
	fs := afero.NewMemMapFs()
	codexHome := filepath.Join(t.TempDir(), "codex-home")
	output := new(bytes.Buffer)
	t.Setenv("CODEX_HOME", codexHome)
	cmd := buildInstallCommand(newInstallTestSource(), fs, func() (string, error) { return t.TempDir(), nil })
	cmd.SetOut(output)
	cmd.SetArgs([]string{"--all", "--agent", codexAgent, "--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	assert.JSONEq(t, fmt.Sprintf(`{"installed":[{"type":"skill","name":"alpha","path":%q},{"type":"skill","name":"beta","path":%q}]}`,
		filepath.Join(codexHome, "skills", "alpha"),
		filepath.Join(codexHome, "skills", "beta")), output.String())
	assertFSExists(t, fs, filepath.Join(codexHome, "skills", "alpha", "SKILL.md"), true)
	assertFSExists(t, fs, filepath.Join(codexHome, "skills", "beta", "SKILL.md"), true)
}

func TestInstallOptsRefusesExistingDestination(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	destination := filepath.Join(home, ".claude", "skills", "alpha")
	require.NoError(t, fs.MkdirAll(destination, 0o755))
	require.NoError(t, afero.WriteFile(fs, filepath.Join(destination, "user.md"), []byte("keep"), 0o644))
	opts := installOpts{
		source:  newInstallTestSource(),
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}
	opts.Output = jsonOutputFormat

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "refusing to overwrite")
	content, err := afero.ReadFile(fs, filepath.Join(destination, "user.md"))
	require.NoError(t, err)
	assert.Equal(t, "keep", string(content))
}

func TestInstallOptsRefusesDanglingDestinationSymlink(t *testing.T) {
	fs := afero.NewOsFs()
	home := t.TempDir()
	skillsRoot := filepath.Join(home, ".claude", "skills")
	require.NoError(t, os.MkdirAll(skillsRoot, 0o755))
	destination := filepath.Join(skillsRoot, "alpha")
	if err := os.Symlink(filepath.Join(home, "missing-target"), destination); err != nil {
		t.Skipf("creating symlinks is not supported: %v", err)
	}
	opts := installOpts{
		source:  newInstallTestSource(),
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "refusing to overwrite")
	info, err := os.Lstat(destination)
	require.NoError(t, err)
	assert.NotZero(t, info.Mode()&os.ModeSymlink)
}

func TestInstallOptsRejectsDestinationTraversal(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	source := newInstallTestSource()
	source.snapshot.filesByName["alpha"] = append(source.snapshot.filesByName["alpha"], skillFile{
		repositoryPath: "skills/alpha/escape.md",
		relativePath:   "../outside.md",
		mode:           0o644,
	})
	source.files["skills/alpha/escape.md"] = []byte("outside")
	opts := installOpts{
		source:  source,
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "unsafe relative path")
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "outside.md"), false)
}

func TestInstallOptsRejectsCaseInsensitiveFileCollisions(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	source := newInstallTestSource()
	source.snapshot.filesByName["alpha"] = append(source.snapshot.filesByName["alpha"], skillFile{
		repositoryPath: "skills/alpha/references/guide.md",
		relativePath:   "references/Guide.md",
		mode:           0o644,
	})
	source.files["skills/alpha/references/Guide.md"] = []byte("case collision")
	opts := installOpts{
		source:  source,
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "collide on case-insensitive filesystems")
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "alpha"), false)
}

func TestInstallOptsRejectsUnicodeNormalizationCollisions(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	source := newInstallTestSource()
	composed := "references/caf\u00e9.md"
	decomposed := "references/cafe\u0301.md"
	source.snapshot.filesByName["alpha"] = append(source.snapshot.filesByName["alpha"],
		skillFile{repositoryPath: "skills/alpha/" + composed, relativePath: composed, mode: 0o644},
		skillFile{repositoryPath: "skills/alpha/" + decomposed, relativePath: decomposed, mode: 0o644},
	)
	source.files["skills/alpha/"+composed] = []byte("composed")
	source.files["skills/alpha/"+decomposed] = []byte("decomposed")
	opts := installOpts{
		source:  source,
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "collide on case-insensitive filesystems")
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "alpha"), false)
}

func TestValidateInstallLimits(t *testing.T) {
	tests := []struct {
		name    string
		files   []skillFile
		wantErr string
	}{
		{
			name:  "per-file boundary",
			files: []skillFile{{repositoryPath: "skills/alpha/SKILL.md", size: int64(maxSkillFileSize)}},
		},
		{
			name:    "per-file limit",
			files:   []skillFile{{repositoryPath: "skills/alpha/SKILL.md", size: int64(maxSkillFileSize) + 1}},
			wantErr: "size limit",
		},
		{
			name:  "aggregate byte boundary",
			files: installLimitFiles(8, int64(maxSkillFileSize)),
		},
		{
			name:    "aggregate byte limit",
			files:   installLimitFiles(9, int64(maxSkillFileSize)),
			wantErr: "total size limit",
		},
		{
			name:    "file count limit",
			files:   installLimitFiles(maxInstallFileCount+1, 0),
			wantErr: "more than 1024 files",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInstallLimits([]skillInstallPlan{{files: tt.files}})
			if tt.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tt.wantErr)
		})
	}
}

func installLimitFiles(count int, size int64) []skillFile {
	files := make([]skillFile, count)
	for i := range files {
		path := fmt.Sprintf("references/file-%04d.md", i)
		files[i] = skillFile{
			repositoryPath: "skills/alpha/" + path,
			relativePath:   path,
			mode:           0o644,
			size:           size,
		}
	}
	return files
}

func TestInstallRejectsActualFileOverLimit(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	source := newInstallTestSource()
	source.files["skills/alpha/SKILL.md"] = bytes.Repeat([]byte("x"), maxSkillFileSize+1)
	opts := installOpts{
		source:  source,
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "size limit")
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "alpha"), false)
}

func TestInstallOptsRejectsFileDirectoryCollisions(t *testing.T) {
	cases := [][]skillFile{
		{
			{relativePath: "README"},
			{relativePath: "readme/guide.md"},
		},
		{
			{relativePath: "readme/guide.md"},
			{relativePath: "README"},
		},
	}
	for i, files := range cases {
		t.Run(fmt.Sprintf("order-%d", i), func(t *testing.T) {
			require.ErrorContains(t, validateCaseInsensitiveFilePaths("alpha", files), "collide on case-insensitive filesystems")
		})
	}
}

func TestInstallAllDoesNotLeavePartialResultsOnDownloadFailure(t *testing.T) {
	fs := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	source := newInstallTestSource()
	source.failReadAt = betaSkillMarkdownPath
	opts := installOpts{
		source:  source,
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
		all:     true,
	}

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "simulated upstream read failure")
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "alpha"), false)
	assertFSExists(t, fs, filepath.Join(home, ".claude", "skills", "beta"), false)
}

func TestInstallBuilderValidatesAgentAndSelection(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "unsupported agent", args: []string{"alpha", "--agent", "cursor"}, want: "supported agents are claude and codex"},
		{name: "missing selection", args: []string{"--agent", "claude"}, want: "specify one skill name or use --all"},
		{name: "conflicting selection", args: []string{"alpha", "--all", "--agent", "claude"}, want: "not both"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := buildInstallCommand(newInstallTestSource(), afero.NewMemMapFs(), func() (string, error) { return t.TempDir(), nil })
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			require.ErrorContains(t, err, tt.want)
		})
	}
}

type failingWriteFs struct {
	afero.Fs
	failSuffix string
}

func (fs failingWriteFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	if strings.HasSuffix(filepath.ToSlash(name), filepath.ToSlash(fs.failSuffix)) {
		return nil, errors.New("simulated filesystem write failure")
	}
	return fs.Fs.OpenFile(name, flag, perm)
}

func TestInstallRollsBackDirectoriesOnWriteFailure(t *testing.T) {
	memFS := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	fs := failingWriteFs{Fs: memFS, failSuffix: filepath.Join("alpha", "scripts", "run.py")}
	opts := installOpts{
		source:  newInstallTestSource(),
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(context.Background(), []string{"alpha"})
	require.ErrorContains(t, err, "simulated filesystem write failure")
	assertFSExists(t, memFS, filepath.Join(home, ".claude", "skills", "alpha"), false)
}

func TestInstallAllRollsBackEarlierSkillsOnWriteFailure(t *testing.T) {
	memFS := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	fs := failingWriteFs{Fs: memFS, failSuffix: filepath.Join("beta", "SKILL.md")}
	opts := installOpts{
		source:  newInstallTestSource(),
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
		all:     true,
	}

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "simulated filesystem write failure")
	assertFSExists(t, memFS, filepath.Join(home, ".claude", "skills", "alpha"), false)
	assertFSExists(t, memFS, filepath.Join(home, ".claude", "skills", "beta"), false)
}

func TestInstallAllRollsBackCommittedSkillsOnRenameFailure(t *testing.T) {
	memFS := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	fs := failingRenameFs{Fs: memFS, failTarget: filepath.Join("skills", "beta")}
	opts := installOpts{
		source:  newInstallTestSource(),
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
		all:     true,
	}

	err := opts.Run(context.Background(), nil)
	require.ErrorContains(t, err, "simulated directory rename failure")
	assertFSExists(t, memFS, filepath.Join(home, ".claude", "skills", "alpha"), false)
	assertFSExists(t, memFS, filepath.Join(home, ".claude", "skills", "beta"), false)
}

func TestInstallRollsBackOnContextCancellationDuringWrite(t *testing.T) {
	memFS := afero.NewMemMapFs()
	home := filepath.Join(t.TempDir(), "home")
	ctx, cancel := context.WithCancel(context.Background())
	fs := cancelAfterWriteFs{Fs: memFS, cancel: cancel}
	opts := installOpts{
		source:  newInstallTestSource(),
		fs:      fs,
		homeDir: func() (string, error) { return home, nil },
		agent:   claudeAgent,
	}

	err := opts.Run(ctx, []string{"alpha"})
	require.ErrorIs(t, err, context.Canceled)
	assertFSExists(t, memFS, filepath.Join(home, ".claude", "skills", "alpha"), false)
}

type cancelAfterWriteFs struct {
	afero.Fs
	cancel context.CancelFunc
}

type failingRenameFs struct {
	afero.Fs
	failTarget string
}

func (fs failingRenameFs) Rename(oldName, newName string) error {
	if strings.HasSuffix(filepath.ToSlash(newName), filepath.ToSlash(fs.failTarget)) {
		return errors.New("simulated directory rename failure")
	}
	return fs.Fs.Rename(oldName, newName)
}

func (fs cancelAfterWriteFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	file, err := fs.Fs.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return &cancelAfterWriteFile{File: file, cancel: fs.cancel}, nil
}

type cancelAfterWriteFile struct {
	afero.File
	cancel context.CancelFunc
	once   sync.Once
}

func (file *cancelAfterWriteFile) Write(content []byte) (int, error) {
	written, err := file.File.Write(content)
	if err == nil {
		file.once.Do(file.cancel)
	}
	return written, err
}

func assertFSExists(t *testing.T, fs afero.Fs, path string, want bool) {
	t.Helper()
	exists, err := afero.Exists(fs, path)
	require.NoError(t, err)
	assert.Equal(t, want, exists)
}
