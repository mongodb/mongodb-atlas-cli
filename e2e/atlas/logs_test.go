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
// +build e2e atlas,logs

package atlas_test

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongocli/e2e"
)

func TestLogs(t *testing.T) {
	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("failed to deploy a cluster: %v", err)
	}
	defer func() {
		if e := deleteCluster(clusterName); e != nil {
			t.Errorf("error deleting test cluster: %v", e)
		}
	}()

	hostname, err := getHostname()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	logTypes := []string{
		"mongodb.gz",
		"mongos.gz",
		"mongodb-audit-log.gz",
		"mongos-audit-log.gz",
	}
	for _, logType := range logTypes {
		lt := logType
		t.Run("Download "+lt, func(t *testing.T) {
			downloadLogTmpPath(t, cliPath, hostname, lt)
		})
	}

	t.Run("Download mongodb.gz no output path", func(t *testing.T) {
		downloadLog(t, cliPath, hostname, "mongodb.gz")
	})
}

func downloadLogTmpPath(t *testing.T, cliPath, hostname, logFile string) {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filepath := dir + logFile

	cmd := exec.Command(cliPath,
		atlasEntity,
		logsEntity,
		"download",
		hostname,
		logFile,
		"--out",
		filepath,
	)

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}

	if _, err := os.Stat(filepath); err != nil {
		t.Fatalf("%v has not been downloaded", filepath)
	}
}

func downloadLog(t *testing.T, cliPath, hostname, logFile string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		atlasEntity,
		logsEntity,
		"download",
		hostname,
		logFile,
	)

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}

	outputFile := strings.ReplaceAll(logFile, ".gz", ".log.gz")
	if _, err := os.Stat(outputFile); err != nil {
		t.Fatalf("%v has not been downloaded", logFile)
	}
}
