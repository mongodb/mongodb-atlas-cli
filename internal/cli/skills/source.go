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
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/google/go-github/v61/github"
	"gopkg.in/yaml.v3"
)

const (
	upstreamOwner             = "mongodb"
	upstreamRepository        = "agent-skills"
	upstreamBranch            = "main"
	upstreamSkillRoot         = "skills"
	rawGitHubBaseURL          = "https://raw.githubusercontent.com"
	maxSkillFileSize          = 8 * 1024 * 1024
	minSkillPathParts         = 2
	skillDefinitionFile       = "SKILL.md"
	githubHTTPTimeout         = 30 * time.Second
	regularFilePermissions    = os.FileMode(0o644)
	executableFilePermissions = os.FileMode(0o755)
)

type skillSource interface {
	Snapshot(context.Context) (*skillSnapshot, error)
	ReadFile(context.Context, string, string) ([]byte, error)
}

type skillFile struct {
	repositoryPath string
	relativePath   string
	mode           os.FileMode
	size           int64
}

type skillSnapshot struct {
	commitSHA   string
	skills      []Skill
	filesByName map[string][]skillFile
}

func (snapshot *skillSnapshot) filesForSkill(name string) ([]skillFile, error) {
	files, ok := snapshot.filesByName[name]
	if !ok {
		return nil, fmt.Errorf("skill %q is not available in the upstream repository", name)
	}

	return files, nil
}

type githubSource struct {
	apiClient  *github.Client
	httpClient *http.Client
	rawBaseURL string
}

func newGitHubSource() *githubSource {
	httpClient := &http.Client{Timeout: githubHTTPTimeout}
	return &githubSource{
		apiClient:  github.NewClient(httpClient),
		httpClient: httpClient,
		rawBaseURL: rawGitHubBaseURL,
	}
}

func (source *githubSource) Snapshot(ctx context.Context) (*skillSnapshot, error) {
	branch, _, err := source.apiClient.Repositories.GetBranch(ctx, upstreamOwner, upstreamRepository, upstreamBranch, 0)
	if err != nil {
		return nil, fmt.Errorf("could not resolve %s/%s branch %q: %w", upstreamOwner, upstreamRepository, upstreamBranch, err)
	}

	commitSHA := branch.GetCommit().GetSHA()
	if commitSHA == "" {
		return nil, fmt.Errorf("could not resolve commit for %s/%s branch %q", upstreamOwner, upstreamRepository, upstreamBranch)
	}

	tree, _, err := source.apiClient.Git.GetTree(ctx, upstreamOwner, upstreamRepository, commitSHA, true)
	if err != nil {
		return nil, fmt.Errorf("could not list files in %s/%s at %s: %w", upstreamOwner, upstreamRepository, commitSHA, err)
	}
	if tree.GetTruncated() {
		return nil, fmt.Errorf("GitHub returned a truncated file tree for %s/%s at %s", upstreamOwner, upstreamRepository, commitSHA)
	}

	filesByName := make(map[string][]skillFile)
	for _, entry := range tree.Entries {
		if entry == nil {
			continue
		}

		file, name, ok, err := skillFileFromTreeEntry(entry)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}
		filesByName[name] = append(filesByName[name], file)
	}

	snapshot := &skillSnapshot{
		commitSHA:   commitSHA,
		skills:      make([]Skill, 0, len(filesByName)),
		filesByName: filesByName,
	}
	for name, files := range snapshot.filesByName {
		sort.Slice(files, func(i, j int) bool {
			return files[i].relativePath < files[j].relativePath
		})

		skillFileIndex := sort.Search(len(files), func(i int) bool {
			return files[i].relativePath >= skillDefinitionFile
		})
		if skillFileIndex == len(files) || files[skillFileIndex].relativePath != skillDefinitionFile {
			delete(snapshot.filesByName, name)
			continue
		}

		snapshot.skills = append(snapshot.skills, Skill{Name: name})
	}

	sort.Slice(snapshot.skills, func(i, j int) bool {
		return snapshot.skills[i].Name < snapshot.skills[j].Name
	})

	return snapshot, nil
}

func skillFileFromTreeEntry(entry *github.TreeEntry) (skillFile, string, bool, error) {
	repositoryPath := entry.GetPath()
	if !strings.HasPrefix(repositoryPath, upstreamSkillRoot+"/") {
		return skillFile{}, "", false, nil
	}

	pathUnderSkills := strings.TrimPrefix(repositoryPath, upstreamSkillRoot+"/")
	segments := strings.Split(pathUnderSkills, "/")
	if len(segments) < minSkillPathParts {
		// Ignore metadata files such as skills/OWNERS.yaml.
		return skillFile{}, "", false, nil
	}

	name := segments[0]
	if !isSafePathComponent(name) {
		return skillFile{}, "", false, fmt.Errorf("unsafe skill name in repository path %q", repositoryPath)
	}
	if entry.GetType() == "tree" {
		return skillFile{}, "", false, nil
	}
	if entry.GetType() != "blob" {
		return skillFile{}, "", false, fmt.Errorf("unsupported GitHub file type %q at %q", entry.GetType(), repositoryPath)
	}

	mode, err := regularFileMode(entry.GetMode())
	if err != nil {
		return skillFile{}, "", false, fmt.Errorf("unsupported GitHub file at %q: %w", repositoryPath, err)
	}

	relativePath := strings.Join(segments[1:], "/")
	if !isSafeRelativePath(relativePath) {
		return skillFile{}, "", false, fmt.Errorf("unsafe file path in upstream skill: %q", repositoryPath)
	}

	return skillFile{
		repositoryPath: repositoryPath,
		relativePath:   relativePath,
		mode:           mode,
		size:           int64(entry.GetSize()),
	}, name, true, nil
}

func regularFileMode(mode string) (os.FileMode, error) {
	switch mode {
	case "100644":
		return regularFilePermissions, nil
	case "100755":
		return executableFilePermissions, nil
	default:
		return 0, fmt.Errorf("mode %q is not a regular file", mode)
	}
}

func isSafePathComponent(value string) bool {
	if value == "" || value == "." || value == parentPathSegment || value != strings.TrimSpace(value) ||
		strings.TrimRight(value, " .") != value || strings.ContainsAny(value, `/\\<>:"|?*`) ||
		path.Clean(value) != value || hasUnsafeCharacters(value) {
		return false
	}

	return !isWindowsReservedDeviceName(value)
}

func isSafeRelativePath(value string) bool {
	if value == "" || value == "." || value == parentPathSegment || path.IsAbs(value) || value != path.Clean(value) ||
		strings.ContainsAny(value, `\\`) {
		return false
	}

	for _, segment := range strings.Split(value, "/") {
		if !isSafePathComponent(segment) {
			return false
		}
	}

	return true
}

func isWindowsReservedDeviceName(value string) bool {
	base, _, _ := strings.Cut(value, ".")
	base = strings.ToUpper(base)
	switch base {
	case "CON", "PRN", "AUX", "NUL", "CONIN$", "CONOUT$", "CLOCK$":
		return true
	}

	characters := []rune(base)
	if len(characters) != 4 || (string(characters[:3]) != "COM" && string(characters[:3]) != "LPT") {
		return false
	}
	last := characters[3]
	return last >= '1' && last <= '9' || last == '¹' || last == '²' || last == '³'
}

func hasUnsafeCharacters(value string) bool {
	for _, char := range value {
		if unicode.IsControl(char) || unicode.Is(unicode.Cf, char) || unicode.Is(unicode.Zl, char) || unicode.Is(unicode.Zp, char) {
			return true
		}
	}
	return false
}

func (source *githubSource) ReadFile(ctx context.Context, commitSHA, repositoryPath string) ([]byte, error) {
	if !isSafeRelativePath(repositoryPath) || !strings.HasPrefix(repositoryPath, upstreamSkillRoot+"/") {
		return nil, fmt.Errorf("unsafe upstream file path %q", repositoryPath)
	}

	segments := strings.Split(repositoryPath, "/")
	for i, segment := range segments {
		segments[i] = url.PathEscape(segment)
	}
	fileURL := strings.TrimRight(source.rawBaseURL, "/") + "/" + upstreamOwner + "/" + upstreamRepository + "/" + commitSHA + "/" + strings.Join(segments, "/")
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for %s: %w", repositoryPath, err)
	}

	response, err := source.httpClient.Do(request) //nolint:gosec // G704: rawBaseURL is fixed to GitHub in production; tests inject an httptest server.
	if err != nil {
		return nil, fmt.Errorf("could not download %s: %w", repositoryPath, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not download %s: GitHub returned %s", repositoryPath, response.Status)
	}

	content, err := io.ReadAll(io.LimitReader(response.Body, maxSkillFileSize+1))
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", repositoryPath, err)
	}
	if len(content) > maxSkillFileSize {
		return nil, fmt.Errorf("upstream file %s exceeds the %d-byte size limit", repositoryPath, maxSkillFileSize)
	}

	return content, nil
}

func skillDescription(content []byte) (string, error) {
	lines := bytes.Split(content, []byte("\n"))
	if len(lines) < 2 || !isFrontmatterDelimiter(lines[0]) {
		return "", nil
	}

	closingDelimiter := -1
	for i := 1; i < len(lines); i++ {
		if isFrontmatterDelimiter(lines[i]) {
			closingDelimiter = i
			break
		}
	}
	if closingDelimiter < 0 {
		return "", errors.New("missing closing YAML frontmatter delimiter")
	}

	var metadata struct {
		Description string `yaml:"description"`
	}
	if err := yaml.Unmarshal(bytes.Join(lines[1:closingDelimiter], []byte("\n")), &metadata); err != nil {
		return "", err
	}

	return strings.TrimSpace(metadata.Description), nil
}

func isFrontmatterDelimiter(line []byte) bool {
	return bytes.Equal(bytes.TrimSuffix(line, []byte("\r")), []byte("---"))
}
