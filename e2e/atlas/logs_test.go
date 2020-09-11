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
	"testing"

	"github.com/mongodb/mongocli/e2e"
)

func TestLogs(t *testing.T) {
	const logsEntity = "logs"

	clusterName, err := deployCluster("")
	if err != nil {
		t.Fatalf("failed to deploy a cluster: %v", err)
	}
	defer func() {
		if e := deleteCluster(clusterName, ""); e != nil {
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
	t.Run("Download mongodb.gz", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		logFile := "mongodb.gz"
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
	})

	t.Run("Download mongos.gz", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		logFile := "mongos.gz"
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
	})

	t.Run("Download mongodb-audit-log.gz", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		logFile := "mongodb-audit-log.gz"
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
	})

	t.Run("Download mongos-audit-log.gz", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		logFile := "mongos-audit-log.gz"
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
	})
}
