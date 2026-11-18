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

// Package apitest provides helpers for the API-driven tests generated from
// the Arazzo specs under test/auto/specs.
package apitest

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

const (
	cloudService    = "cloud"
	govService      = "cloudgov"
	orgIDProfileKey = "org_id"

	randomSuffixBytes = 4
	versionParts      = 2
	majorPart         = 0
	minorPart         = 1

	groupNotFoundCode   = "GROUP_NOT_FOUND"
	maxCleanupAttempts  = 10
	initialCleanupDelay = 1
	cleanupBackoffStep  = 2
)

// Run executes an `atlas api` command with the given arguments, optionally
// piping stdin as the request body, and asserts it succeeds. It returns the
// command stdout.
func Run(t *testing.T, stdin string, args []string) []byte {
	t.Helper()

	out, err := execute(t, stdin, args)
	require.NoError(t, err, string(out))

	return out
}

// RunExpectError executes an `atlas api` command and asserts it fails, for
// example when the API is expected to reject the request.
func RunExpectError(t *testing.T, stdin string, args []string) []byte {
	t.Helper()

	out, err := execute(t, stdin, args)
	require.Error(t, err, string(out))

	return out
}

func execute(t *testing.T, stdin string, args []string) ([]byte, error) {
	t.Helper()

	cmd, err := buildCommand(stdin, args)
	if err != nil {
		require.NoError(t, err)
	}

	return cmd.Output()
}

func buildCommand(stdin string, args []string) (*exec.Cmd, error) {
	cliPath, err := internal.AtlasCLIBin()
	if err != nil {
		return nil, err
	}

	cmdArgs := append([]string{"api"}, args...)
	cmdArgs = append(cmdArgs, "-P", internal.ProfileName())

	cmd := exec.Command(cliPath, cmdArgs...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	return cmd, nil
}

// Cleanup registers a best-effort `atlas api` command that runs when the test
// finishes, so resources created by a failing test are still released.
func Cleanup(t *testing.T, args []string) {
	t.Helper()

	t.Cleanup(func() {
		if _, err := runCommand(args); err != nil {
			fmt.Fprintf(os.Stderr, "warning: cleanup command failed: %v\n", err)
		}
	})
}

// CleanupWithRetry registers a best-effort `atlas api` command retried with
// exponential backoff. Group-not-found responses count as success.
func CleanupWithRetry(t *testing.T, args []string) {
	t.Helper()

	t.Cleanup(func() {
		delay := initialCleanupDelay
		for attempts := 1; attempts <= maxCleanupAttempts; attempts++ {
			out, err := runCommand(args)
			if err == nil || strings.Contains(string(out), groupNotFoundCode) {
				return
			}
			fmt.Fprintf(os.Stderr, "warning: cleanup attempt %d/%d failed: %v\n", attempts, maxCleanupAttempts, err)
			time.Sleep(time.Duration(delay) * time.Second)
			delay *= cleanupBackoffStep
		}
	})
}

func runCommand(args []string) ([]byte, error) {
	cmd, err := buildCommand("", args)
	if err != nil {
		return nil, err
	}

	return cmd.Output()
}

// OrgID returns the organization ID configured in the test profile.
func OrgID(t *testing.T) string {
	t.Helper()

	profile, err := internal.ProfileData()
	require.NoError(t, err)

	orgID := profile[orgIDProfileKey]
	require.NotEmpty(t, orgID, "org_id is not set in the test profile")

	return orgID
}

// Service returns the service the test profile targets, mirroring the value
// the CLI profile stores and internal.IsGov reads.
func Service() string {
	if internal.IsGov() {
		return govService
	}

	return cloudService
}

// RandomSuffix returns a random lowercase hexadecimal suffix.
func RandomSuffix(t *testing.T) string {
	t.Helper()

	buf := make([]byte, randomSuffixBytes)
	_, err := rand.Read(buf)
	require.NoError(t, err)

	return hex.EncodeToString(buf)
}

// Marshal serializes a request body built by a generated test.
func Marshal(t *testing.T, body map[string]any) string {
	t.Helper()

	out, err := json.Marshal(body)
	require.NoError(t, err)

	return string(out)
}

// Unmarshal decodes an `atlas api` command response.
func Unmarshal(t *testing.T, out []byte) map[string]any {
	t.Helper()

	var body map[string]any
	require.NoError(t, json.Unmarshal(out, &body))

	return body
}

// JSONValue resolves a JSON pointer such as results/0/version against a
// decoded response body. It returns nil when the pointer does not resolve.
func JSONValue(body any, pointer string) any {
	current := body
	if pointer == "" {
		return current
	}

	for _, rawToken := range strings.Split(pointer, "/") {
		token := strings.ReplaceAll(strings.ReplaceAll(rawToken, "~1", "/"), "~0", "~")
		switch node := current.(type) {
		case map[string]any:
			current = node[token]
		case []any:
			index, err := strconv.Atoi(token)
			if err != nil || index < 0 || index >= len(node) {
				return nil
			}
			current = node[index]
		default:
			return nil
		}
	}

	return current
}

// JSONString resolves a JSON pointer against a decoded response body and
// renders the value as a string.
func JSONString(body any, pointer string) string {
	switch value := JSONValue(body, pointer).(type) {
	case string:
		return value
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case nil:
		return ""
	default:
		return jsonScalar(value)
	}
}

func jsonScalar(value any) string {
	out, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(out)
}

// PreviousMajorVersion decrements the major version and keeps the minor
// version, mirroring getPreviousMajorVersion in the clusters flags test.
func PreviousMajorVersion(version string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != versionParts {
		return "", errInvalidVersion{version: version}
	}

	major, err := strconv.Atoi(parts[majorPart])
	if err != nil {
		return "", errInvalidVersion{version: version}
	}

	return strconv.Itoa(major-1) + "." + parts[minorPart], nil
}

type errInvalidVersion struct {
	version string
}

func (e errInvalidVersion) Error() string {
	return "invalid MongoDB major version: " + e.version
}
