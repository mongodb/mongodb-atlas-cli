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
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/cmd/genevergreen/generate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixtureFeed(t *testing.T) *feed {
	t.Helper()
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	t.Cleanup(srv.Close)
	f, err := fetchFeed(t.Context(), srv.URL+"/full.json")
	require.NoError(t, err)
	return f
}

func baseMatrix() generate.ServerMatrix {
	return generate.ServerMatrix{
		"8.0": {"rhel80"},
		"9.0": {"rhel80"},
	}
}

func TestUpdate(t *testing.T) {
	res, err := update(baseMatrix(), fixtureFeed(t), nil)
	require.NoError(t, err)

	assert.Equal(t, generate.ServerMatrix{
		// rhel10 added to an existing version; rhel82 maps to rhel80; ubuntu2004 skipped.
		"8.0": {"rhel10", "rhel80"},
		// Union of 9.0.1 and 9.0.2; rhel8 and rhel93 mapped; macos skipped; RC ubuntu2604 ignored.
		"9.0": {"debian12", "debian13", "rhel80", "rhel90"},
		// New version newer than 9.0; rhel9 mapped; windows skipped.
		"10.0": {"amazon2023", "rhel90"},
		// 7.0 and 8.1 are older than 9.0 and not in the matrix, so not added.
	}, res.Matrix)
	assert.Equal(t, []string{"10.0"}, res.AddedVersions)
	assert.Equal(t, []pair{
		{"8.0", "rhel10"},
		{"9.0", "debian12"},
		{"9.0", "debian13"},
		{"9.0", "rhel90"},
		{"10.0", "amazon2023"},
		{"10.0", "rhel90"},
	}, res.Added)
	assert.Empty(t, res.Removed)
	// newos2 is s390x only, so it's not considered.
	assert.Equal(t, []pair{{"9.0", "newos1"}}, res.Unknown)
}

func TestUpdateNeverRemoves(t *testing.T) {
	current := generate.ServerMatrix{"9.0": {"suse16"}}
	f := &feed{Versions: []release{{Version: "9.0.1", ProductionRelease: true, Downloads: []download{{Arch: "x86_64", Target: "rhel80"}}}}}
	res, err := update(current, f, nil)
	require.NoError(t, err)
	assert.Equal(t, generate.ServerMatrix{"9.0": {"rhel80", "suse16"}}, res.Matrix)
}

func TestUpdateNewVersionWithoutSupportedDistro(t *testing.T) {
	f := &feed{Versions: []release{{Version: "10.0.0", ProductionRelease: true, Downloads: []download{
		{Arch: "x86_64", Target: "newos1"},
		{Arch: "x86_64", Target: "macos"},
	}}}}
	_, err := update(baseMatrix(), f, nil)
	require.ErrorContains(t, err, "new version 10.0")
}

func TestUpdateRejectsFeed(t *testing.T) {
	tests := map[string]*feed{
		"no GA releases":     {Versions: []release{{Version: "10.0.0-rc0"}}},
		"empty feed":         {},
		"invalid GA version": {Versions: []release{{Version: "9.0", ProductionRelease: true}}},
	}
	for name, f := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := update(baseMatrix(), f, nil)
			assert.Error(t, err)
		})
	}
}

func TestUpdateAppliesExceptions(t *testing.T) {
	exceptions := []exception{
		{Version: "9.0", Distro: "rhel90", Action: actionRemove, Jira: "CLOUDP-1"},
		{Version: "9.0", Distro: "suse16", Action: actionAdd, Jira: "CLOUDP-2"},
		{Version: "8.0", Distro: "rhel80", Action: actionRemove, Jira: "CLOUDP-3"},
	}
	res, err := update(baseMatrix(), fixtureFeed(t), exceptions)
	require.NoError(t, err)
	assert.Equal(t, []string{"rhel10"}, res.Matrix["8.0"])
	assert.Equal(t, []string{"debian12", "debian13", "rhel80", "suse16"}, res.Matrix["9.0"])
	assert.Equal(t, []pair{{"8.0", "rhel80"}}, res.Removed)
}

func TestUpdateRejectsExceptions(t *testing.T) {
	tests := map[string]exception{
		"version not in matrix":   {Version: "7.0", Distro: "rhel80", Action: actionAdd, Jira: "CLOUDP-1"},
		"removes the last distro": {Version: "10.0", Distro: "amazon2023", Action: actionRemove, Jira: "CLOUDP-1"},
	}
	for name, e := range tests {
		t.Run(name, func(t *testing.T) {
			f := &feed{Versions: []release{{Version: "10.0.0", ProductionRelease: true, Downloads: []download{{Arch: "x86_64", Target: "amazon2023"}}}}}
			_, err := update(baseMatrix(), f, []exception{e})
			assert.Error(t, err)
		})
	}
}

func TestParseExceptions(t *testing.T) {
	e, err := parseExceptions([]byte(`
- {version: "8.2", distro: debian13, action: add, jira: CLOUDP-415313}
- {version: "7.0", distro: rhel10, action: remove, jira: CLOUDP-452014}
`))
	require.NoError(t, err)
	assert.Equal(t, []exception{
		{Version: "8.2", Distro: "debian13", Action: actionAdd, Jira: "CLOUDP-415313"},
		{Version: "7.0", Distro: "rhel10", Action: actionRemove, Jira: "CLOUDP-452014"},
	}, e)

	e, err = parseExceptions(nil)
	require.NoError(t, err)
	assert.Empty(t, e)
}

func TestParseExceptionsRejects(t *testing.T) {
	tests := map[string]string{
		"patch version":  `- {version: "8.2.1", distro: debian13, action: add, jira: CLOUDP-1}`,
		"unknown distro": `- {version: "8.2", distro: rhel70, action: add, jira: CLOUDP-1}`,
		"invalid action": `- {version: "8.2", distro: debian13, action: skip, jira: CLOUDP-1}`,
		"missing jira":   `- {version: "8.2", distro: debian13, action: add}`,
		"invalid jira":   `- {version: "8.2", distro: debian13, action: add, jira: cloudp-1}`,
		"unknown field":  `- {version: "8.2", distro: debian13, action: add, jira: CLOUDP-1, reason: x}`,
		"second document": `- {version: "8.2", distro: debian13, action: add, jira: CLOUDP-1}
---
- {version: "7.0", distro: rhel10, action: remove, jira: CLOUDP-2}`,
		"duplicate": `- {version: "8.2", distro: debian13, action: add, jira: CLOUDP-1}
- {version: "8.2", distro: debian13, action: remove, jira: CLOUDP-2}`,
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := parseExceptions([]byte(input))
			assert.Error(t, err)
		})
	}
}

func TestFetchFeedNotFound(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	defer srv.Close()
	_, err := fetchFeed(t.Context(), srv.URL)
	require.ErrorContains(t, err, "404")
}

func TestMarshalMatrixOrdersVersionsNumerically(t *testing.T) {
	m := generate.ServerMatrix{"10.0": {"rhel80"}, "9.0": {"debian12", "rhel80"}}
	out, err := marshalMatrix(m)
	require.NoError(t, err)
	s := string(out)
	assert.Less(t, strings.Index(s, `"9.0"`), strings.Index(s, `"10.0"`))

	parsed, err := generate.ParseServerMatrix(out)
	require.NoError(t, err)
	assert.Equal(t, m, parsed)
}

func TestCommittedExceptionsAreValid(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "build", "ci", "server-matrix-exceptions.yaml"))
	require.NoError(t, err)
	_, err = parseExceptions(b)
	require.NoError(t, err)
}

func TestMarshalMatrixRoundTrips(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "build", "ci", "server-matrix.json"))
	require.NoError(t, err)
	m, err := generate.ParseServerMatrix(b)
	require.NoError(t, err)
	out, err := marshalMatrix(m)
	require.NoError(t, err)
	assert.True(t, bytes.Equal(b, out), "server-matrix.json is not in servermatrix output format")
}
