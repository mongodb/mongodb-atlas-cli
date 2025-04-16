// Copyright 2020 MongoDB Inc
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
//go:build e2e || e2eSnap || (atlas && logs)

package logs

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

func TestLogs(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())

	g.GenerateProjectAndCluster("logs")

	hostname, err := g.GetHostname()
	require.NoError(t, err)

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)
	logTypes := []string{
		"mongodb.gz",
		"mongos.gz",
		"mongodb-audit-log.gz",
		"mongos-audit-log.gz",
	}
	for _, logType := range logTypes {
		lt := logType
		g.Run("Download "+lt, func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
			downloadLogTmpPath(t, cliPath, hostname, lt, g.ProjectID)
		})
	}

	g.Run("Download mongodb.gz no output path", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		downloadLog(t, cliPath, hostname, "mongodb.gz", g.ProjectID)
	})
}

func downloadLogTmpPath(t *testing.T, cliPath, hostname, logFile, projectID string) {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	filepath := dir + logFile

	cmd := exec.Command(cliPath,
		logsEntity,
		"download",
		hostname,
		logFile,
		"--out",
		filepath,
		"--projectId",
		projectID,
	)

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		require.Contains(t, string(resp), "returned file is empty")
	} else {
		if _, err := os.Stat(filepath); err != nil {
			t.Fatalf("%v has not been downloaded", filepath)
		}
	}
	_ = os.Remove(filepath)
}

func downloadLog(t *testing.T, cliPath, hostname, logFile, projectID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		logsEntity,
		"download",
		hostname,
		logFile,
		"--projectId",
		projectID,
	)

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		require.Contains(t, string(resp), "returned file is empty")
	} else {
		outputFile := strings.ReplaceAll(logFile, ".gz", ".log.gz")
		if _, err := os.Stat(outputFile); err != nil {
			t.Fatalf("%v has not been downloaded", logFile)
		}
		_ = os.Remove(outputFile)
	}
	_ = os.Remove(logFile)
}
