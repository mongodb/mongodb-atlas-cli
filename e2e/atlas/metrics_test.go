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
// +build e2e atlas,metrics

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestMetrics(t *testing.T) {
	const metricsEntity = "metrics"

	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hostname, err := getHostnameAndPort()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := cli()
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

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

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

	t.Run("databases list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			metricsEntity,
			"databases",
			"list",
			hostname)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		databases := &mongodbatlas.ProcessDatabasesResponse{}
		err = json.Unmarshal(resp, &databases)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if databases.TotalCount != 2 {
			t.Errorf("got=%#v\nwant=%#v\n", databases.TotalCount, 2)
		}
	})

	t.Run("disks list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			metricsEntity,
			"disks",
			"list",
			hostname)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		disks := &mongodbatlas.ProcessDisksResponse{}
		err = json.Unmarshal(resp, &disks)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if disks.TotalCount != 1 {
			t.Errorf("got=%#v\nwant=%#v\n", disks.TotalCount, 1)
		}
	})

	t.Run("disks describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			metricsEntity,
			"disks",
			"describe",
			hostname,
			"data",
			"--granularity=PT30M",
			"--period=P1DT12H")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		metrics := &mongodbatlas.ProcessDiskMeasurements{}
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
	if err := deleteCluster(clusterName); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
