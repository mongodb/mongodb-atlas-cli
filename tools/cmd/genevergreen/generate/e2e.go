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

	e2eSidecarName = "e2e.yaml"

	e2eTreePrefix = "test/e2e/"
)

var e2eGroups = map[string]bool{
	"pr":    true,
	"gov":   true,
	"noapi": true,
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
	"cluster":             true,
	"docker":              true,
	"flex_mtm":            true,
	"m0_mtm":              true,
	"flex_snapshot":       true,
	"federation_settings": true,
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

type e2ePackage struct {
	path     string
	manifest e2eManifest
}

func (p e2ePackage) hasGroup(group string) bool {
	return slices.Contains(p.manifest.Groups, group)
}

func (p e2ePackage) hasFixture(fixture string) bool {
	return slices.Contains(p.manifest.Fixtures, fixture)
}

type E2EOptions struct {
	Variant          string
	TaskPrefix       string
	Profile          string
	IncludeAnyGroups []string
	ExcludeGroups    []string
	ExcludeFixtures  []string
}

func E2ETasks(c *shrub.Configuration, opts E2EOptions) error {
	if opts.Variant == "" {
		return errors.New("e2e generation requires a variant")
	}
	if opts.TaskPrefix == "" {
		opts.TaskPrefix = opts.Variant + "_"
	}

	packages, err := loade2ePackages()
	if err != nil {
		return err
	}

	for _, group := range append(opts.IncludeAnyGroups, opts.ExcludeGroups...) {
		if !e2eGroups[group] {
			return fmt.Errorf("unknown e2e group: %s", group)
		}
	}

	for _, fixture := range opts.ExcludeFixtures {
		if !e2eStructuralFixtures[fixture] {
			if _, ok := e2eFixtureVars[fixture]; !ok {
				return fmt.Errorf("unknown e2e fixture to exclude: %s", fixture)
			}
		}
	}

	selected := make([]e2ePackage, 0, len(packages))
	for _, pkg := range packages {
		if !e2eSelected(pkg, opts) {
			continue
		}
		selected = append(selected, pkg)
	}
	if len(selected) == 0 {
		return fmt.Errorf("no e2e packages selected for variant %s", opts.Variant)
	}

	byPath := make(map[string]e2ePackage, len(selected))
	nameByPath := make(map[string]string, len(selected))
	for _, pkg := range selected {
		name, err := e2eTaskName(pkg, opts.TaskPrefix)
		if err != nil {
			return err
		}
		byPath[pkg.path] = pkg
		nameByPath[pkg.path] = name
	}

	if err := validateE2ENeeds(byPath, opts.Variant); err != nil {
		return err
	}

	v := &shrub.Variant{
		BuildName: opts.Variant,
	}

	names := make(map[string]bool, len(selected))
	for _, pkg := range selected {
		name := nameByPath[pkg.path]
		if names[name] {
			return fmt.Errorf("duplicate generated task name: %s", name)
		}
		names[name] = true

		needs := make([]shrub.TaskDependency, 0, len(pkg.manifest.Needs))
		for _, needed := range pkg.manifest.Needs {
			dependency := shrub.TaskDependency{Name: nameByPath[needed], Variant: opts.Variant}
			dependency.SetPatchOptional(true)
			needs = append(needs, dependency)
		}

		t := e2eTask(pkg, name, opts, needs)
		c.Tasks = append(c.Tasks, t)
		v.AddTasks(t.Name)
	}

	c.Variants = append(c.Variants, v)

	return nil
}

func validateE2ENeeds(byPath map[string]e2ePackage, variant string) error {
	for path, pkg := range byPath {
		for _, needed := range pkg.manifest.Needs {
			if needed == path {
				return fmt.Errorf("e2e package %s cannot need itself", path)
			}
			if _, ok := byPath[needed]; !ok {
				return fmt.Errorf("%s needs %s, which is not generated for variant %s", path, needed, variant)
			}
		}
	}

	const visiting, done = 1, 2
	state := make(map[string]int, len(byPath))
	var visit func(path string) error
	visit = func(path string) error {
		switch state[path] {
		case visiting:
			return fmt.Errorf("circular e2e needs involving %s", path)
		case done:
			return nil
		}
		state[path] = visiting
		for _, needed := range byPath[path].manifest.Needs {
			if err := visit(needed); err != nil {
				return err
			}
		}
		state[path] = done
		return nil
	}
	for path := range byPath {
		if err := visit(path); err != nil {
			return err
		}
	}
	return nil
}

func e2eSelected(pkg e2ePackage, opts E2EOptions) bool {
	for _, group := range opts.ExcludeGroups {
		if pkg.hasGroup(group) {
			return false
		}
	}
	for _, fixture := range opts.ExcludeFixtures {
		if pkg.hasFixture(fixture) {
			return false
		}
	}
	if len(opts.IncludeAnyGroups) == 0 {
		return true
	}
	for _, group := range opts.IncludeAnyGroups {
		if pkg.hasGroup(group) {
			return true
		}
	}
	return false
}

func e2eTaskName(pkg e2ePackage, prefix string) (string, error) {
	if !strings.HasPrefix(pkg.path, e2eTreePrefix) {
		return "", fmt.Errorf("e2e package outside %s: %s", e2eTreePrefix, pkg.path)
	}
	return prefix + strings.ReplaceAll(strings.TrimPrefix(pkg.path, e2eTreePrefix), "/", "_"), nil
}

func e2eTask(pkg e2ePackage, name string, opts E2EOptions, needs []shrub.TaskDependency) *shrub.Task {
	compileDependency := shrub.TaskDependency{Name: "compile", Variant: "code_health"}
	compileDependency.SetPatchOptional(true)

	dependencies := append([]shrub.TaskDependency{compileDependency}, needs...)

	t := &shrub.Task{
		Name: name,
	}

	t = t.Tag("e2e").
		MustHaveTestResults(true).
		Dependency(dependencies...)

	if pkg.hasFixture("docker") {
		t = t.Function(e2eDockerFunctions...)
	}

	vars := map[string]string{
		"E2E_TEST_PACKAGES": "./" + pkg.path + "/...",
	}
	if opts.Profile != "" {
		vars["E2E_PROFILE_NAME"] = opts.Profile
	}
	if pkg.manifest.Timeout != "" {
		vars["E2E_TIMEOUT"] = pkg.manifest.Timeout
	}
	for _, fixture := range pkg.manifest.Fixtures {
		for key, value := range e2eFixtureVars[fixture] {
			vars[key] = value
		}
	}

	t = t.Function("restore go modules").
		FunctionWithVars("e2e test", vars)

	if pkg.manifest.Timeout != "" {
		t = t.ExecTimeout(e2eExecTimeoutSecs)
	}

	return t
}

func loade2ePackages() ([]e2ePackage, error) {
	module, err := e2eModulePath()
	if err != nil {
		return nil, err
	}

	listOut, err := exec.CommandContext(context.Background(), "go", "list", "-f", "{{.ImportPath}} {{.Dir}}", module+"/test/e2e/...").Output() //nolint:gosec // module comes from `go list -m` in this module
	if err != nil {
		return nil, fmt.Errorf("listing e2e packages: %w", err)
	}

	repoRoot, err := e2eRepoRoot()
	if err != nil {
		return nil, err
	}

	var packages []e2ePackage
	packageDirs := make(map[string]bool)
	var missing []string
	for _, line := range strings.Split(strings.TrimSpace(string(listOut)), "\n") {
		importPath, dir, found := strings.Cut(line, " ")
		if !found || dir == "" {
			continue
		}
		pkg := e2ePackage{path: strings.TrimPrefix(importPath, module+"/")}
		packageDirs[dir] = true

		manifest, err := loade2eManifest(filepath.Join(dir, e2eSidecarName))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", pkg.path, err)
		}
		if manifest == nil {
			missing = append(missing, pkg.path)
			continue
		}
		pkg.manifest = *manifest
		packages = append(packages, pkg)
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("e2e packages without %s: %s", e2eSidecarName, strings.Join(missing, ", "))
	}

	if err := checkOrphanSidecars(filepath.Join(repoRoot, "test", "e2e"), packageDirs); err != nil {
		return nil, err
	}

	sort.Slice(packages, func(i, j int) bool {
		return packages[i].path < packages[j].path
	})

	return packages, nil
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
		return nil, fmt.Errorf("parsing %s: %w", e2eSidecarName, err)
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

func checkOrphanSidecars(root string, packageDirs map[string]bool) error {
	var orphans []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			return nil
		}
		if _, err := os.Stat(filepath.Join(path, e2eSidecarName)); err == nil && !packageDirs[path] {
			orphans = append(orphans, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(orphans) > 0 {
		return fmt.Errorf("orphan %s files in non-package directories: %s", e2eSidecarName, strings.Join(orphans, ", "))
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
