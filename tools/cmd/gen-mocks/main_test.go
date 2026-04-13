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
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteMockFiles(t *testing.T) {
	dir := t.TempDir()

	mockFile := filepath.Join(dir, "create_mock_test.go")
	normalFile := filepath.Join(dir, "create.go")
	nestedDir := filepath.Join(dir, "sub")
	if err := os.MkdirAll(nestedDir, 0o755); err != nil {
		t.Fatal(err)
	}
	nestedMock := filepath.Join(nestedDir, "list_mock_test.go")

	for _, f := range []string{mockFile, normalFile, nestedMock} {
		if err := os.WriteFile(f, []byte("package foo"), 0o600); err != nil {
			t.Fatal(err)
		}
	}

	if err := deleteMockFiles(dir); err != nil {
		t.Fatalf("deleteMockFiles: %v", err)
	}

	for _, f := range []string{mockFile, nestedMock} {
		if _, err := os.Stat(f); !os.IsNotExist(err) {
			t.Errorf("expected %s to be deleted", f)
		}
	}
	if _, err := os.Stat(normalFile); err != nil {
		t.Errorf("expected %s to still exist: %v", normalFile, err)
	}
}

func TestFindPackagesWithGenerate(t *testing.T) {
	dir := t.TempDir()

	// Package with mockgen directive — should be discovered
	pkgWithMock := filepath.Join(dir, "internal", "cli", "foo")
	if err := os.MkdirAll(pkgWithMock, 0o755); err != nil {
		t.Fatal(err)
	}
	mockgenContent := "package foo\n\n//go:generate go tool go.uber.org/mock/mockgen -typed -destination=foo_mock_test.go -package=foo -source=foo.go\n"
	if err := os.WriteFile(filepath.Join(pkgWithMock, "foo.go"), []byte(mockgenContent), 0o600); err != nil {
		t.Fatal(err)
	}
	// Second file in same package — should not duplicate the package entry
	if err := os.WriteFile(filepath.Join(pkgWithMock, "bar.go"), []byte("package foo\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	// Package without mockgen directive — should NOT be discovered
	pkgNone := filepath.Join(dir, "internal", "cli", "bar")
	if err := os.MkdirAll(pkgNone, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pkgNone, "bar.go"), []byte("package bar\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	pkgs, err := findPackagesWithGenerate(dir)
	if err != nil {
		t.Fatalf("findPackagesWithGenerate: %v", err)
	}
	if len(pkgs) != 1 {
		t.Fatalf("expected 1 package, got %d: %v", len(pkgs), pkgs)
	}
	if pkgs[0] != pkgWithMock {
		t.Errorf("expected %s, got %s", pkgWithMock, pkgs[0])
	}
}

func TestIsMockFile(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"create_mock_test.go", true},
		{"mock_store.go", true},
		{"create.go", false},
		{"create_test.go", false},
		{"mock_store.txt", false},
	}
	for _, c := range cases {
		if got := isMockFile(c.name); got != c.want {
			t.Errorf("isMockFile(%q) = %v, want %v", c.name, got, c.want)
		}
	}
}
