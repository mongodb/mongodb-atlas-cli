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

package generate

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/evergreen-ci/shrub"
	"gopkg.in/yaml.v3"
)

const (
	e2eExecTimeoutSecs = 11400

	e2eWaterfallVariant = "e2e_waterfall"
	e2ePRVariant        = "e2e_pr"
	e2eGovVariant       = "e2e_gov"

	e2eWaterfallTaskPrefix = "e2e_waterfall_"
	e2ePRTaskPrefix        = "e2e_pr_"
	e2eGovTaskPrefix       = "e2e_gov_"

	e2eGovProfile = "__e2e_gov"

	e2eManifestName = "e2e.yaml"

	e2eTreePrefix = "test/e2e/"
)

var e2eGroups = map[string]bool{
	"pr":  true,
	"gov": true,
}

var e2eFixtureVars = map[string]map[string]string{
	"aws_role": {
		"E2E_CLOUD_ROLE_ID": "${e2e_cloud_role_id}",
	},
	"s3_bucket": {
		"E2E_TEST_BUCKET": "${e2e_test_bucket}",
	},
	"flex_instance": {
		"E2E_FLEX_INSTANCE_NAME": "${e2e_flex_instance_name}",
	},
	"kms_aws": { //nolint:gosec // Evergreen expansion references, not credentials
		"AWS_ACCESS_KEY":        "${logs_decrypt_aws_access_key}",
		"AWS_SECRET_ACCESS_KEY": "${logs_decrypt_aws_secret_access_key}",
	},
	"kms_azure": { //nolint:gosec // Evergreen expansion references, not credentials
		"AZURE_TENANT_ID":     "${logs_decrypt_azure_tenant_id}",
		"AZURE_CLIENT_ID":     "${logs_decrypt_azure_client_id}",
		"AZURE_CLIENT_SECRET": "${logs_decrypt_azure_client_secret}",
	},
	"kms_gcp": { //nolint:gosec // Evergreen expansion references, not credentials
		"GCP_CREDENTIALS": "${logs_decrypt_gcp_credentials}",
	},
}

var e2eStructuralFixtures = map[string]bool{
	"api":                    true,
	"docker":                 true,
	"flex_mtm":               true,
	"m0_mtm":                 true,
	"flex_snapshot":          true,
	"federation_settings":    true,
	"mongohouse":             true,
	"sls":                    true,
	"dedicated_search_nodes": true,
}

var e2eDockerFunctions = []string{
	"install mongodb database tools",
	"increase inotify limits",
	"authenticate to ecr",
}

type e2eManifest struct {
	Groups   []string `yaml:"groups"`
	Fixtures []string `yaml:"fixtures"`
	Needs    []string `yaml:"needs"`
	Timeout  string   `yaml:"timeout"`
}

type e2eUnit struct {
	path     string
	manifest e2eManifest
}

func (u e2eUnit) hasGroup(group string) bool {
	return slices.Contains(u.manifest.Groups, group)
}

func (u e2eUnit) hasFixture(fixture string) bool {
	return slices.Contains(u.manifest.Fixtures, fixture)
}

type E2EOptions struct {
	Variant          string
	TaskPrefix       string
	Profile          string
	IncludeAnyGroups []string
	ExcludeGroups    []string
	RequireFixtures  []string
	ExcludeFixtures  []string
}

func e2eSelectedUnits(opts E2EOptions) ([]e2eUnit, map[string]bool, error) {
	units, err := loade2eUnits()
	if err != nil {
		return nil, nil, err
	}

	for _, group := range append(opts.IncludeAnyGroups, opts.ExcludeGroups...) {
		if !e2eGroups[group] {
			return nil, nil, fmt.Errorf("unknown e2e group: %s", group)
		}
	}

	for _, fixture := range append(opts.RequireFixtures, opts.ExcludeFixtures...) {
		if !e2eStructuralFixtures[fixture] {
			if _, ok := e2eFixtureVars[fixture]; !ok {
				return nil, nil, fmt.Errorf("unknown e2e fixture: %s", fixture)
			}
		}
	}

	allPaths := make(map[string]bool, len(units))
	selected := make([]e2eUnit, 0, len(units))
	for _, unit := range units {
		allPaths[unit.path] = true
		if !e2eSelected(unit, opts) {
			continue
		}
		selected = append(selected, unit)
	}
	if len(selected) == 0 {
		return nil, nil, errors.New("no e2e packages selected")
	}

	return selected, allPaths, nil
}

type e2eUnitManifest struct {
	Path     string   `json:"path"`
	Groups   []string `json:"groups"`
	Fixtures []string `json:"fixtures"`
	Needs    []string `json:"needs"`
	Timeout  string   `json:"timeout,omitempty"`
}

func E2EManifests(opts E2EOptions) ([]byte, error) {
	selected, _, err := e2eSelectedUnits(opts)
	if err != nil {
		return nil, err
	}

	units := make([]e2eUnitManifest, 0, len(selected))
	for _, unit := range selected {
		out := e2eUnitManifest{
			Path:     unit.path,
			Groups:   unit.manifest.Groups,
			Fixtures: unit.manifest.Fixtures,
			Needs:    unit.manifest.Needs,
			Timeout:  unit.manifest.Timeout,
		}
		if out.Groups == nil {
			out.Groups = []string{}
		}
		if out.Fixtures == nil {
			out.Fixtures = []string{}
		}
		if out.Needs == nil {
			out.Needs = []string{}
		}
		units = append(units, out)
	}

	return json.MarshalIndent(units, "", "  ")
}

func E2ETasks(c *shrub.Configuration, opts E2EOptions) error {
	if opts.Variant == "" {
		return errors.New("e2e generation requires a variant")
	}
	if opts.TaskPrefix == "" {
		opts.TaskPrefix = opts.Variant + "_"
	}

	selected, allPaths, err := e2eSelectedUnits(opts)
	if err != nil {
		return err
	}

	byPath := make(map[string]e2eUnit, len(selected))
	nameByPath := make(map[string]string, len(selected))
	for _, unit := range selected {
		name, err := e2eTaskName(unit, opts.TaskPrefix)
		if err != nil {
			return err
		}
		byPath[unit.path] = unit
		nameByPath[unit.path] = name
	}

	if err := validateE2ENeeds(byPath, allPaths, opts.Variant); err != nil {
		return err
	}

	v := &shrub.Variant{
		BuildName: opts.Variant,
	}

	names := make(map[string]bool, len(selected))
	for _, unit := range selected {
		name := nameByPath[unit.path]
		if names[name] {
			return fmt.Errorf("duplicate generated task name: %s", name)
		}
		names[name] = true

		needs := make([]shrub.TaskDependency, 0, len(unit.manifest.Needs))
		for _, needed := range unit.manifest.Needs {
			dependency := shrub.TaskDependency{Name: nameByPath[needed], Variant: opts.Variant}
			dependency.SetPatchOptional(true)
			needs = append(needs, dependency)
		}

		t := e2eTask(unit, name, opts, needs)
		c.Tasks = append(c.Tasks, t)
		v.AddTasks(t.Name)
	}

	c.Variants = append(c.Variants, v)

	return nil
}

func validateE2ENeeds(selected map[string]e2eUnit, allPaths map[string]bool, variant string) error {
	for path, unit := range selected {
		for _, needed := range unit.manifest.Needs {
			if needed == path {
				return fmt.Errorf("e2e package %s cannot need itself", path)
			}
			if !allPaths[needed] {
				return fmt.Errorf("%s needs %s, which has no %s", path, needed, e2eManifestName)
			}
			if _, ok := selected[needed]; !ok {
				return fmt.Errorf("%s needs %s, which is not generated for variant %s", path, needed, variant)
			}
		}
	}

	const visiting, done = 1, 2
	state := make(map[string]int, len(selected))
	var visit func(path string) error
	visit = func(path string) error {
		switch state[path] {
		case visiting:
			return fmt.Errorf("circular e2e needs involving %s", path)
		case done:
			return nil
		}
		state[path] = visiting
		for _, needed := range selected[path].manifest.Needs {
			if err := visit(needed); err != nil {
				return err
			}
		}
		state[path] = done
		return nil
	}
	for path := range selected {
		if err := visit(path); err != nil {
			return err
		}
	}
	return nil
}

func e2eSelected(unit e2eUnit, opts E2EOptions) bool {
	for _, group := range opts.ExcludeGroups {
		if unit.hasGroup(group) {
			return false
		}
	}
	for _, fixture := range opts.ExcludeFixtures {
		if unit.hasFixture(fixture) {
			return false
		}
	}
	for _, fixture := range opts.RequireFixtures {
		if !unit.hasFixture(fixture) {
			return false
		}
	}
	if len(opts.IncludeAnyGroups) == 0 {
		return true
	}
	for _, group := range opts.IncludeAnyGroups {
		if unit.hasGroup(group) {
			return true
		}
	}
	return false
}

func e2eTaskName(unit e2eUnit, prefix string) (string, error) {
	if !strings.HasPrefix(unit.path, e2eTreePrefix) {
		return "", fmt.Errorf("e2e package outside %s: %s", e2eTreePrefix, unit.path)
	}
	return prefix + strings.ReplaceAll(strings.TrimPrefix(unit.path, e2eTreePrefix), "/", "_"), nil
}

func e2eTask(unit e2eUnit, name string, opts E2EOptions, needs []shrub.TaskDependency) *shrub.Task {
	compileDependency := shrub.TaskDependency{Name: "compile", Variant: "code_health"}
	compileDependency.SetPatchOptional(true)

	dependencies := append([]shrub.TaskDependency{compileDependency}, needs...)

	t := &shrub.Task{
		Name: name,
	}

	t = t.Tag("e2e").
		MustHaveTestResults(true).
		Dependency(dependencies...)

	if unit.hasFixture("docker") {
		t = t.Function(e2eDockerFunctions...)
	}

	vars := map[string]string{
		"E2E_TEST_PACKAGES": "./" + unit.path + "/...",
	}
	if opts.Profile != "" {
		vars["E2E_PROFILE_NAME"] = opts.Profile
	}
	if unit.manifest.Timeout != "" {
		vars["E2E_TIMEOUT"] = unit.manifest.Timeout
	}
	for _, fixture := range unit.manifest.Fixtures {
		for key, value := range e2eFixtureVars[fixture] {
			vars[key] = value
		}
	}

	t = t.Function("restore go modules").
		FunctionWithVars("e2e test", vars)

	if unit.manifest.Timeout != "" {
		t = t.ExecTimeout(e2eExecTimeoutSecs)
	}

	return t
}

func loade2eUnits() ([]e2eUnit, error) {
	module, err := e2eModulePath()
	if err != nil {
		return nil, err
	}

	repoRoot, err := e2eRepoRoot()
	if err != nil {
		return nil, err
	}

	manifestDirs, err := finde2eManifestDirs(filepath.Join(repoRoot, "test", "e2e"))
	if err != nil {
		return nil, err
	}

	for i, parent := range manifestDirs {
		for j, child := range manifestDirs {
			if i != j && e2eIsSubdir(child, parent) {
				return nil, fmt.Errorf("nested %s files are not allowed: %s and %s",
					e2eManifestName, e2eRelPath(repoRoot, parent), e2eRelPath(repoRoot, child))
			}
		}
	}

	listOut, err := exec.CommandContext(context.Background(), "go", "list", "-f", "{{.ImportPath}} {{.Dir}}", module+"/test/e2e/...").Output() //nolint:gosec // module comes from `go list -m` in this module
	if err != nil {
		return nil, fmt.Errorf("listing e2e packages: %w", err)
	}

	covered := make(map[string]bool, len(manifestDirs))
	var uncovered []string
	for _, line := range strings.Split(strings.TrimSpace(string(listOut)), "\n") {
		importPath, dir, found := strings.Cut(line, " ")
		if !found || dir == "" {
			continue
		}
		manifestDir := e2eCoveringManifestDir(dir, manifestDirs)
		if manifestDir == "" {
			uncovered = append(uncovered, strings.TrimPrefix(importPath, module+"/"))
			continue
		}
		covered[manifestDir] = true
	}

	if len(uncovered) > 0 {
		return nil, fmt.Errorf("e2e packages not covered by any %s: %s", e2eManifestName, strings.Join(uncovered, ", "))
	}

	var orphans []string
	units := make([]e2eUnit, 0, len(manifestDirs))
	for _, dir := range manifestDirs {
		if !covered[dir] {
			orphans = append(orphans, e2eRelPath(repoRoot, dir))
			continue
		}
		manifest, err := loade2eManifest(filepath.Join(dir, e2eManifestName))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e2eRelPath(repoRoot, dir), err)
		}
		if manifest == nil {
			return nil, fmt.Errorf("%s: missing %s", e2eRelPath(repoRoot, dir), e2eManifestName)
		}
		units = append(units, e2eUnit{path: e2eRelPath(repoRoot, dir), manifest: *manifest})
	}
	if len(orphans) > 0 {
		return nil, fmt.Errorf("orphan %s files covering no e2e packages: %s", e2eManifestName, strings.Join(orphans, ", "))
	}

	sort.Slice(units, func(i, j int) bool {
		return units[i].path < units[j].path
	})

	return units, nil
}

func finde2eManifestDirs(root string) ([]string, error) {
	var dirs []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			return nil
		}
		if _, err := os.Stat(filepath.Join(path, e2eManifestName)); err == nil {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

func e2eCoveringManifestDir(pkgDir string, manifestDirs []string) string {
	for _, dir := range manifestDirs {
		if pkgDir == dir || e2eIsSubdir(pkgDir, dir) {
			return dir
		}
	}
	return ""
}

func e2eIsSubdir(child, parent string) bool {
	return strings.HasPrefix(child, parent+string(filepath.Separator))
}

func e2eRelPath(repoRoot, dir string) string {
	return strings.TrimPrefix(dir, repoRoot+string(filepath.Separator))
}

func loade2eManifest(path string) (*e2eManifest, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)

	var manifest e2eManifest
	if err := decoder.Decode(&manifest); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", e2eManifestName, err)
	}

	if err := validateE2EManifest(manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

func validateE2EManifest(manifest e2eManifest) error {
	for _, group := range manifest.Groups {
		if !e2eGroups[group] {
			return fmt.Errorf("unknown e2e group: %s", group)
		}
	}
	for _, fixture := range manifest.Fixtures {
		if !e2eStructuralFixtures[fixture] {
			if _, ok := e2eFixtureVars[fixture]; !ok {
				return fmt.Errorf("unknown e2e fixture: %s", fixture)
			}
		}
	}
	for _, needed := range manifest.Needs {
		if !strings.HasPrefix(needed, e2eTreePrefix) {
			return fmt.Errorf("needs must reference an e2e package path: %s", needed)
		}
	}
	if manifest.Timeout != "" {
		if _, err := time.ParseDuration(manifest.Timeout); err != nil {
			return fmt.Errorf("invalid timeout %q: %w", manifest.Timeout, err)
		}
	}
	return nil
}

func e2eModulePath() (string, error) {
	out, err := exec.CommandContext(context.Background(), "go", "list", "-m").Output()
	if err != nil {
		return "", fmt.Errorf("getting module path: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func e2eRepoRoot() (string, error) {
	out, err := exec.CommandContext(context.Background(), "go", "env", "GOMOD").Output()
	if err != nil {
		return "", fmt.Errorf("getting go.mod path: %w", err)
	}
	gomod := strings.TrimSpace(string(out))
	if gomod == "" {
		return "", errors.New("go env GOMOD returned an empty path")
	}
	return filepath.Dir(gomod), nil
}
