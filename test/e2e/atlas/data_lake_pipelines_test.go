// Copyright 2023 MongoDB Inc
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
//go:build e2e || (atlas && datalakepipeline)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDataLakePipelines(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	g := newAtlasE2ETestGenerator(t)
	g.enableBackup = true
	g.generateProjectAndCluster("dataLakePipeline")

	t.Run("Load Sample Data", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"load",
			g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var sampleDatasetJob *atlasv2.SampleDatasetStatus
		err = json.Unmarshal(resp, &sampleDatasetJob)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(g.clusterName, sampleDatasetJob.GetClusterName())

		cmd = exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"watch",
			sampleDatasetJob.GetId(),
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err = e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		t.Log(string(resp))
	})

	var snapshotID string
	t.Run("Generate Snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create", g.clusterName,
			"--desc", "snapshot 1",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var snapshot *atlasv2.BackupSnapshot
		err = json.Unmarshal(resp, &snapshot)
		req.NoError(err)

		snapshotID = snapshot.GetId()

		cmd = exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch", snapshotID,
			"--clusterName", g.clusterName,
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err = e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		t.Log(string(resp))
	})

	const pipelineName = "sample_mflix.movies"
	var pipelineID, pipelineRunID string

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"create", pipelineName,
			"--sourceType", "ON_DEMAND_CPS",
			"--sourceClusterName", g.clusterName,
			"--sourceDatabaseName", "sample_mflix",
			"--sourceCollectionName", "movies",
			"--sinkType", "DLS",
			"--sinkMetadataProvider", "AWS",
			"--sinkMetadataRegion", "US_EAST_1",
			"--sinkPartitionField", "year,title",
			"--transform", "EXCLUDE:fullplot",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err, string(resp))

		a := assert.New(t)
		var pipeline *atlasv2.DataLakeIngestionPipeline
		req.NoError(json.Unmarshal(resp, &pipeline))
		pipelineID = *pipeline.Id
		a.Equal(pipelineName, *pipeline.Name)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"watch", pipelineName,
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		t.Log(string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"describe", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		a := assert.New(t)
		var pipeline atlasv2.DataLakeIngestionPipeline
		req.NoError(json.Unmarshal(resp, &pipeline))
		a.Equal(pipelineName, *pipeline.Name)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"ls",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var r []atlasv2.DataLakeIngestionPipeline
		a := assert.New(t)
		req.NoError(json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"update", pipelineName,
			"--sinkMetadataProvider", "AWS",
			"--sinkMetadataRegion", "US_EAST_2",
			"--sinkPartitionField", "year,title",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err, string(resp))

		a := assert.New(t)
		var pipeline *atlasv2.DataLakeIngestionPipeline
		req.NoError(json.Unmarshal(resp, &pipeline))
		a.Equal(pipelineName, *pipeline.Name)
	})

	t.Run("Trigger", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"trigger", pipelineName,
			"--snapshotId", snapshotID,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err, string(resp))

		a := assert.New(t)
		var run *atlasv2.IngestionPipelineRun
		req.NoError(json.Unmarshal(resp, &run))
		pipelineRunID = *run.Id
		a.Equal(pipelineID, *run.PipelineId)
	})

	t.Run("Pause", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"pause", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err, string(resp))

		a := assert.New(t)
		var pipeline *atlasv2.DataLakeIngestionPipeline
		req.NoError(json.Unmarshal(resp, &pipeline))
		a.Equal(pipelineName, *pipeline.Name)
	})

	t.Run("Start", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"start", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err, string(resp))

		a := assert.New(t)
		var pipeline *atlasv2.DataLakeIngestionPipeline
		req.NoError(json.Unmarshal(resp, &pipeline))
		a.Equal(pipelineName, *pipeline.Name)
	})

	t.Run("Runs List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"runs",
			"ls",
			"--pipeline", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var r *atlasv2.PaginatedPipelineRun
		a := assert.New(t)
		req.NoError(json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	t.Run("Runs Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"runs",
			"describe", pipelineRunID,
			"--pipeline", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		a := assert.New(t)
		var run *atlasv2.IngestionPipelineRun
		req.NoError(json.Unmarshal(resp, &run))
		pipelineRunID = *run.Id
		a.Equal(pipelineID, *run.PipelineId)
	})

	t.Run("Runs Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"runs",
			"watch", pipelineRunID,
			"--pipeline", pipelineName,
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		t.Log(string(resp))
	})

	t.Run("Datasets Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"datasets",
			"delete", pipelineRunID,
			"--pipeline", pipelineName,
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		t.Log(string(resp))
	})

	t.Run("AvailableSchedules List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"availableschedules",
			"ls",
			"--pipeline", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var r []atlasv2.DiskBackupApiPolicyItem
		a := assert.New(t)
		req.NoError(json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	t.Run("AvailableSnapshots List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"availablesnapshots",
			"ls",
			"--pipeline", pipelineName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var r *atlasv2.PaginatedBackupSnapshot
		a := assert.New(t)
		req.NoError(json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datalakePipelineEntity,
			"delete", pipelineName,
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("'%s' deleted\n", pipelineName)
		assert.Equal(t, expected, string(resp))
	})
}
