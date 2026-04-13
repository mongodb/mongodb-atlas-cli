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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func main() {
	const root = "internal"

	fmt.Println("==> Cleaning up mock files")
	if err := deleteMockFiles(root); err != nil {
		fmt.Fprintf(os.Stderr, "cleanup failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("==> Discovering packages")
	pkgs, err := findPackagesWithGenerate(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "discovery failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("    Found %d packages\n", len(pkgs))

	fmt.Printf("==> Generating mocks (%d workers)\n", runtime.NumCPU())
	failures := runParallel(pkgs)
	if len(failures) > 0 {
		fmt.Fprintln(os.Stderr, "==> Generation failed:")
		for _, f := range failures {
			fmt.Fprintln(os.Stderr, f)
		}
		os.Exit(1)
	}
	fmt.Println("==> Done")
}

func deleteMockFiles(root string) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && isMockFile(d.Name()) {
			return os.Remove(path)
		}
		return nil
	})
}

func findPackagesWithGenerate(root string) ([]string, error) {
	seen := make(map[string]bool)
	var pkgs []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		s := string(content)
		if strings.Contains(s, "//go:generate") && strings.Contains(s, "mockgen") {
			dir := filepath.Dir(path)
			if !seen[dir] {
				seen[dir] = true
				pkgs = append(pkgs, dir)
			}
		}
		return nil
	})
	return pkgs, err
}

type generateResult struct {
	pkg    string
	output []byte
	err    error
}

func runParallel(pkgs []string) []string {
	jobs := make(chan string, len(pkgs))
	for _, p := range pkgs {
		jobs <- p
	}
	close(jobs)

	results := make(chan generateResult, len(pkgs))

	var wg sync.WaitGroup
	for range runtime.NumCPU() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pkg := range jobs {
				cmd := exec.Command("go", "generate", "./"+filepath.ToSlash(pkg))
				out, err := cmd.CombinedOutput()
				results <- generateResult{pkg: pkg, output: out, err: err}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var failures []string
	for r := range results {
		if r.err != nil {
			failures = append(failures, fmt.Sprintf("FAIL %s:\n%s", r.pkg, r.output))
		}
	}
	return failures
}

// isMockFile reports whether a filename belongs to a generated mock.
func isMockFile(name string) bool {
	return strings.Contains(name, "mock") && filepath.Ext(name) == ".go"
}
