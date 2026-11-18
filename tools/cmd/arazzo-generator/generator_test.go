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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveExpr(t *testing.T) {
	cases := []struct {
		expr     string
		expected string
		ok       bool
	}{
		{expr: "$inputs.projectName", expected: "projectName", ok: true},
		{expr: "$steps.createProject.outputs.projectId", expected: "createProject_projectId", ok: true},
		{expr: "$response.body#/id", ok: false},
		{expr: "AWS", ok: false},
	}

	for _, tc := range cases {
		resolved, ok := resolveExpr(tc.expr)
		assert.Equal(t, tc.ok, ok, tc.expr)
		assert.Equal(t, tc.expected, resolved, tc.expr)
	}
}

func TestParseTemplate(t *testing.T) {
	parts, err := parseTemplate("clustersFlags-{$inputs.rand}")
	require.NoError(t, err)
	require.Len(t, parts, 2)
	assert.Equal(t, "clustersFlags-", parts[0].literal)
	assert.Equal(t, "rand", parts[0+1].expr)

	parts, err = parseTemplate("{$steps.discoverMongoDBMajorVersion.outputs.mongoDBMajorVersion}")
	require.NoError(t, err)
	require.Len(t, parts, 1)
	assert.Equal(t, "discoverMongoDBMajorVersion_mongoDBMajorVersion", parts[0].expr)

	_, err = parseTemplate("prefix-{$inputs.rand")
	require.ErrorContains(t, err, "unbalanced")
}

func TestParseCriterion(t *testing.T) {
	parsed, err := parseCriterion("$statusCode == 200")
	require.NoError(t, err)
	assert.Equal(t, criterion{kind: criterionStatus, status: 200}, parsed)

	parsed, err = parseCriterion(`$response.body#/id != null`)
	require.NoError(t, err)
	assert.Equal(t, criterion{kind: criterionNotNull, pointer: "id"}, parsed)

	parsed, err = parseCriterion(`$response.body#/stateName == 'IDLE'`)
	require.NoError(t, err)
	assert.Equal(t, criterion{kind: criterionEqual, pointer: "stateName", expected: `"IDLE"`, expectedRaw: `'IDLE'`}, parsed)

	parsed, err = parseCriterion(`$response.body#/diskSizeGB == 30`)
	require.NoError(t, err)
	assert.Equal(t, criterion{kind: criterionEqual, pointer: "diskSizeGB", expected: `"30"`, expectedRaw: "30"}, parsed)

	parsed, err = parseCriterion(`$response.body#/name == $inputs.clusterName`)
	require.NoError(t, err)
	assert.Equal(t, criterion{kind: criterionEqual, pointer: "name", expected: "clusterName", expectedRaw: "$inputs.clusterName"}, parsed)

	_, err = parseCriterion(`$response.body#/stateName ~= 'IDLE'`)
	require.ErrorContains(t, err, "unsupported criterion")

	_, err = parseCriterion(`$response.body#/stateName != 'IDLE'`)
	require.ErrorContains(t, err, "comparisons against null")
}

func TestRenderGoValue(t *testing.T) {
	rendered, err := renderGoValue(map[string]any{
		"name":        "$inputs.clusterName",
		"clusterType": "REPLICASET",
		"diskSizeGB":  30,
		"paused":      false,
		"tags":        []any{map[string]any{"key": "env"}},
	})
	require.NoError(t, err)
	assert.Equal(
		t,
		`map[string]any{"clusterType": "REPLICASET", "diskSizeGB": 30, "name": clusterName, "paused": false, "tags": []any{map[string]any{"key": "env"}}}`,
		rendered,
	)

	_, err = renderGoValue("$response.body#/id")
	require.ErrorContains(t, err, "unsupported expression")
}

func TestBodyVersion(t *testing.T) {
	version, err := bodyVersion("application/vnd.atlas.2023-02-01+json")
	require.NoError(t, err)
	assert.Equal(t, "2023-02-01", version)

	_, err = bodyVersion("application/json")
	require.ErrorContains(t, err, "unsupported content type")
}

func TestMergeWatchSteps(t *testing.T) {
	index := map[string]commandInfo{
		"createGroupCluster": {
			hasWatcher:            true,
			watcherGetOperationID: "getGroupCluster",
		},
		"getGroupCluster": {},
		"deleteGroup":     {},
	}

	steps := []Step{
		{StepID: "create", OperationID: "createGroupCluster"},
		{StepID: "waitForReady", OperationID: "getGroupCluster", OnFailure: []FailureAction{{Type: "retry", RetryAfter: 30, RetryLimit: 60}}},
		{StepID: "delete", OperationID: "deleteGroup"},
	}

	merged, watches, err := mergeWatchSteps(steps, index)
	require.NoError(t, err)
	require.Len(t, merged, 2)
	assert.Equal(t, "create", merged[0].StepID)
	assert.Equal(t, "delete", merged[1].StepID)
	assert.Equal(t, map[string]int{"create": 1800}, watches)

	_, _, err = mergeWatchSteps([]Step{
		{StepID: "delete", OperationID: "deleteGroup"},
		{StepID: "waitFor", OperationID: "getGroupCluster", OnFailure: []FailureAction{{Type: "retry"}}},
	}, index)
	require.ErrorContains(t, err, "does not watch it")
}

func TestPackageName(t *testing.T) {
	assert.Equal(t, "clustersflags", packageName("test/e2e/auto/specs/clusters_flags.arazzo.yaml"))
}
