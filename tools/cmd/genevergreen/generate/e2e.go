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
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/evergreen-ci/shrub"
)

const (
	e2eExecTimeoutSecs = 11400

	e2eWaterfallVariant = "e2e_waterfall"
	e2ePRVariant        = "e2e_pr"

	e2eWaterfallDisplay = "E2E Waterfall"
	e2ePRDisplay        = "E2E PR"

	e2eWaterfallPkgPrefix = "test/e2e/atlas/"
	e2ePRPkgPrefix        = "test/e2e/pr/"

	e2eWaterfallTaskPrefix = "e2e_waterfall_"
	e2ePRTaskPrefix        = "e2e_pr_"
)

var goExpansions = map[string]any{
	"go_root":      "/opt/golang/go1.26",
	"go_bin":       "/opt/golang/go1.26/bin",
	"go_base_path": "",
}

func E2EWaterfallTasks(c *shrub.Configuration) error {
	return e2eTasks(c, e2eWaterfallVariant, e2eWaterfallDisplay, e2eWaterfallPkgPrefix, e2eWaterfallTaskPrefix)
}

func E2EPRTasks(c *shrub.Configuration) error {
	return e2eTasks(c, e2ePRVariant, e2ePRDisplay, e2ePRPkgPrefix, e2ePRTaskPrefix)
}

func e2eTasks(c *shrub.Configuration, variantName, displayName, pkgPrefix, taskPrefix string) error {
	packages, err := e2ePackages(pkgPrefix)
	if err != nil {
		return err
	}

	v := &shrub.Variant{
		BuildName:        variantName,
		BuildDisplayName: displayName,
		DistroRunOn:      []string{runOn},
		Expansions:       goExpansions,
	}

	for _, pkg := range packages {
		name := taskPrefix + strings.ReplaceAll(strings.TrimPrefix(pkg, pkgPrefix), "/", "_")

		dependency := shrub.TaskDependency{Name: "compile", Variant: "code_health"}
		dependency.SetPatchOptional(true)

		t := &shrub.Task{
			Name: name,
		}

		t = t.Tag("e2e").
			MustHaveTestResults(true).
			ExecTimeout(e2eExecTimeoutSecs).
			Dependency(dependency).
			Function("install mongodb database tools").
			Function("increase inotify limits").
			Function("authenticate to ecr").
			Function("restore go modules").
			FunctionWithVars("e2e test", e2eTestVars(pkg))

		c.Tasks = append(c.Tasks, t)
		v.AddTasks(t.Name)
	}

	c.Variants = append(c.Variants, v)

	return nil
}

func e2eTestVars(pkg string) map[string]string {
	return map[string]string{ //nolint:gosec // Evergreen expansion references, not credentials
		"E2E_TEST_PACKAGES":      "./" + pkg + "/...",
		"E2E_CLOUD_ROLE_ID":      "${e2e_cloud_role_id}",
		"E2E_TEST_BUCKET":        "${e2e_test_bucket}",
		"E2E_FLEX_INSTANCE_NAME": "${e2e_flex_instance_name}",
		"E2E_TIMEOUT":            "3h",
		"GCP_CREDENTIALS":        "${logs_decrypt_gcp_credentials}",
		"AWS_ACCESS_KEY":         "${logs_decrypt_aws_access_key}",
		"AWS_SECRET_ACCESS_KEY":  "${logs_decrypt_aws_secret_access_key}",
		"AZURE_TENANT_ID":        "${logs_decrypt_azure_tenant_id}",
		"AZURE_CLIENT_ID":        "${logs_decrypt_azure_client_id}",
		"AZURE_CLIENT_SECRET":    "${logs_decrypt_azure_client_secret}",
	}
}

func e2ePackages(pkgPrefix string) ([]string, error) {
	moduleOut, err := exec.CommandContext(context.Background(), "go", "list", "-m").Output()
	if err != nil {
		return nil, fmt.Errorf("getting module path: %w", err)
	}
	module := strings.TrimSpace(string(moduleOut))

	listOut, err := exec.CommandContext(context.Background(), "go", "list", module+"/test/...").Output() //nolint:gosec // module comes from `go list -m` in this module
	if err != nil {
		return nil, fmt.Errorf("listing test packages: %w", err)
	}

	var packages []string
	for _, line := range strings.Split(strings.TrimSpace(string(listOut)), "\n") {
		pkg := strings.TrimPrefix(line, module+"/")
		if pkg == line {
			continue
		}
		if strings.HasPrefix(pkg, "test/internal") {
			continue
		}
		if !strings.HasPrefix(pkg, pkgPrefix) {
			continue
		}
		packages = append(packages, pkg)
	}

	if len(packages) == 0 {
		return nil, fmt.Errorf("no e2e packages found under %s", pkgPrefix)
	}

	return packages, nil
}
