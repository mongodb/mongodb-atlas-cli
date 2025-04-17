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

//go:build e2e || e2eSnap || (atlas && metrics)

package metrics

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

const (
	metricsEntity = "metrics"
)

func TestMetrics(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProjectAndCluster("metrics")

	hostname, err := g.GetHostnameAndPort()
	require.NoError(t, err)

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.Run("processes", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		process(t, cliPath, hostname, g.ProjectID)
	})

	g.Run("processes with type", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		processWithType(t, cliPath, hostname, g.ProjectID)
	})

	g.Run("databases", func(_ *testing.T) {
		databases(g, cliPath, hostname, g.ProjectID)
	})

	g.Run("disks", func(_ *testing.T) {
		disks(g, cliPath, hostname, g.ProjectID)
	})
}

func process(t *testing.T, cliPath, hostname, projectID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		metricsEntity,
		"processes",
		hostname,
		"--granularity=PT30M",
		"--period=P1DT12H",
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	var metrics *atlasv2.ApiMeasurementsGeneralViewAtlas
	require.NoError(t, json.Unmarshal(resp, &metrics), string(resp))
	assert.NotEmpty(t, metrics.Measurements)
}

func processWithType(t *testing.T, cliPath, hostname, projectID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		metricsEntity,
		"processes",
		hostname,
		"--granularity=PT30M",
		"--period=P1DT12H",
		"--type=MAX_PROCESS_CPU_USER",
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	var metrics *atlasv2.ApiMeasurementsGeneralViewAtlas
	require.NoError(t, json.Unmarshal(resp, &metrics), string(resp))
	assert.NotEmpty(t, metrics.Measurements)
}

func databases(g *internal.AtlasE2ETestGenerator, cliPath, hostname, projectID string) {
	g.Run("databases list", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			metricsEntity,
			"databases",
			"list",
			hostname,
			"--projectId", projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var db atlasv2.PaginatedDatabase
		require.NoError(t, json.Unmarshal(resp, &db), string(resp))
		assert.NotEmpty(t, db.GetTotalCount())
	})

	g.Run("databases describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			metricsEntity,
			"databases",
			"describe",
			hostname,
			"config",
			"--granularity=PT30M",
			"--period=P1DT12H",
			"--projectId", projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var metrics atlasv2.ApiMeasurementsGeneralViewAtlas
		require.NoError(t, json.Unmarshal(resp, &metrics), string(resp))
		assert.NotEmpty(t, metrics.Measurements)
	})
}

func disks(g *internal.AtlasE2ETestGenerator, cliPath, hostname, projectID string) {
	g.Run("disks list", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			metricsEntity,
			"disks",
			"list",
			hostname,
			"--projectId", projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var d atlasv2.PaginatedDiskPartition
		require.NoError(t, json.Unmarshal(resp, &d), string(resp))
		assert.Positive(t, d.GetTotalCount())
	})

	g.Run("disks describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			metricsEntity,
			"disks",
			"describe",
			hostname,
			"data",
			"--granularity=PT30M",
			"--period=P1DT12H",
			"--projectId", projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var metrics atlasv2.ApiMeasurementsGeneralViewAtlas
		require.NoError(t, json.Unmarshal(resp, &metrics), string(resp))
		assert.NotEmpty(t, metrics.Measurements)
	})
}
