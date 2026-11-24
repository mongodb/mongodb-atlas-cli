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

package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed test.go.tmpl
var testTemplateContent string

const (
	defaultSpecsDir = "test/e2e/auto/specs"
	defaultOutDir   = "test/e2e/auto"

	specSuffix = ".arazzo.yaml"

	e2eManifest = "groups:\n  - gov\nfixtures:\n  - api\n"

	dirPermission  = 0o750
	filePermission = 0o600
)

// fileModel renders one generated test file for one Arazzo spec.
type fileModel struct {
	Package      string
	NeedsRequire bool
	Workflows    []testModel
}

func main() {
	var specsDir, outDir string

	rootCmd := &cobra.Command{
		Use:   "arazzo-generator",
		Short: "CLI which generates `atlas api` based tests from Arazzo specs",
		RunE: func(command *cobra.Command, _ []string) error {
			return run(command.Context(), specsDir, outDir, command.OutOrStdout())
		},
	}

	rootCmd.Flags().StringVar(&specsDir, "specs-dir", defaultSpecsDir, "Directory containing the Arazzo specs")
	rootCmd.Flags().StringVar(&outDir, "out-dir", defaultOutDir, "Directory where the generated tests are written")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(_ context.Context, specsDir, outDir string, out io.Writer) error {
	specPaths, err := findSpecs(specsDir)
	if err != nil {
		return err
	}
	if len(specPaths) == 0 {
		return fmt.Errorf("no Arazzo specs found in %s", specsDir)
	}

	index := buildCommandIndex()
	testTemplate := template.Must(template.New("test").Parse(testTemplateContent))

	for _, specPath := range specPaths {
		if err := generateSpec(specPath, outDir, index, testTemplate, out); err != nil {
			return err
		}
	}

	return nil
}

func findSpecs(specsDir string) ([]string, error) {
	entries, err := os.ReadDir(specsDir)
	if err != nil {
		return nil, fmt.Errorf("reading specs directory %s: %w", specsDir, err)
	}

	specPaths := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), specSuffix) {
			continue
		}
		specPaths = append(specPaths, filepath.Join(specsDir, entry.Name()))
	}
	sort.Strings(specPaths)

	return specPaths, nil
}

func generateSpec(
	specPath, outDir string,
	index map[string]commandInfo,
	testTemplate *template.Template,
	out io.Writer,
) error {
	doc, err := loadDocument(specPath)
	if err != nil {
		return err
	}

	pkg := packageName(specPath)
	model := fileModel{Package: pkg}
	for _, workflow := range doc.Workflows {
		test, err := buildTest(workflow, index)
		if err != nil {
			return fmt.Errorf("%s: workflow %q: %w", specPath, workflow.WorkflowID, err)
		}
		if test.NeedsRequire {
			model.NeedsRequire = true
		}
		model.Workflows = append(model.Workflows, *test)
	}

	var rendered bytes.Buffer
	if err := testTemplate.Execute(&rendered, model); err != nil {
		return fmt.Errorf("%s: rendering template: %w", specPath, err)
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		return fmt.Errorf("%s: formatting generated code: %w\n%s", specPath, err, rendered.String())
	}

	targetDir := filepath.Join(outDir, pkg)
	if err := os.MkdirAll(targetDir, dirPermission); err != nil {
		return fmt.Errorf("creating %s: %w", targetDir, err)
	}

	testPath := filepath.Join(targetDir, pkg+"_test.go")
	if err := os.WriteFile(testPath, formatted, filePermission); err != nil {
		return fmt.Errorf("writing %s: %w", testPath, err)
	}

	manifestPath := filepath.Join(targetDir, "e2e.yaml")
	if err := os.WriteFile(manifestPath, []byte(e2eManifest), filePermission); err != nil {
		return fmt.Errorf("writing %s: %w", manifestPath, err)
	}

	fmt.Fprintf(out, "Generated %s from %s\n", testPath, specPath)

	return nil
}

var nonAlphanumeric = regexp.MustCompile(`[^a-zA-Z0-9]`)

// packageName derives the Go package name for a spec, for example
// clusters_flags.arazzo.yaml becomes clustersflags.
func packageName(specPath string) string {
	base := strings.TrimSuffix(filepath.Base(specPath), specSuffix)
	return strings.ToLower(nonAlphanumeric.ReplaceAllString(base, ""))
}
