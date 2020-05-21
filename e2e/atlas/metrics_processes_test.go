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
// +build e2e

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestAtlasMetrics(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	atlasEntity := "atlas"
	metricsEntity := "metrics"
	clusterName := "myReplicaSet"
	deployCluster(cliPath, atlasEntity, clusterName)
	hostname, err := getHostnameAndPort(cliPath, atlasEntity)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("processes", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			metricsEntity,
			"processes",
			hostname,
			"--granularity=PT30M",
			"--period=P1DT12H")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		metrics := &mongodbatlas.ProcessMeasurements{}
		err = json.Unmarshal(resp, &metrics)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if metrics.Measurements == nil {
			t.Errorf("there are no measurements")
		}

		if len(metrics.Measurements) == 0 {
			t.Errorf("got=%#v\nwant=%#v\n", 0, "len(metrics.Measurements) > 0")
		}

	})

	deleteCluster(cliPath, atlasEntity, clusterName)

}

func getHostnameAndPort(cliPath, atlasEntity string) (string, error) {
	cmd := exec.Command(cliPath,
		atlasEntity,
		"processes",
		"list")

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	var processes []*mongodbatlas.Process
	err = json.Unmarshal(resp, &processes)

	if err != nil {
		return "", fmt.Errorf("unexpected error: %v", err)
	}

	if len(processes) == 0 {
		return "", fmt.Errorf("got=%#v\nwant=%#v\n", 0, "len(processes) > 0")
	}

	return processes[0].Hostname + ":" + strconv.Itoa(processes[0].Port), nil

}

func deployCluster(cliPath, atlasEntity, clusterName string) error {
	cmd := exec.Command(cliPath,
		atlasEntity,
		"clusters",
		"create",
		clusterName,
		"--region=US_EAST_1",
		"--members=3",
		"--tier=M10",
		"--provider=AWS",
		"--mdbVersion=4.0",
		"--diskSizeGB=10")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("unexpected error: %v, resp: %v", err, string(resp))
	}

	cmd = exec.Command(cliPath,
		"atlas",
		"clusters",
		"watch",
		clusterName)
	cmd.Env = os.Environ()
	resp, err = cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("unexpected error: %v, resp: %v", err, string(resp))
	}

	if !strings.Contains(string(resp), "Cluster available at:") {
		return fmt.Errorf("got=%#v\nwant=%#v\n", string(resp), "Cluster available at:")
	}

	return nil
}

func deleteCluster(cliPath, atlasEntity, clusterName string) error {
	cmd := exec.Command(cliPath, atlasEntity, "clusters", "delete", clusterName, "--force")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("unexpected error: %v, resp: %v", err, string(resp))
	}

	expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
	if string(resp) != expected {
		return fmt.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
	}
	return nil
}
