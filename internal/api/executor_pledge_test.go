// Copyright 2026 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// withTempStateDir points pledge state at a fresh temp dir for the duration of the test.
func withTempStateDir(t *testing.T) {
	t.Helper()
	t.Setenv("ATLAS_PLEDGE_STATE_DIR", t.TempDir())
}

// currentSessionKey resolves the pledge session key for the current process.
func currentSessionKey(t *testing.T) pledge.SessionKey {
	t.Helper()
	k, err := pledge.ResolveSessionKey()
	require.NoError(t, err)
	return k
}

func makeExecutor(t *testing.T, httpCode int) *Executor {
	t.Helper()
	ctrl := gomock.NewController(t)

	configProvider := NewMockConfigProvider(ctrl)
	configProvider.EXPECT().GetBaseURL().Return("https://cloud.mongodb.com", nil).AnyTimes()

	conv, err := NewDefaultCommandConverter(configProvider)
	require.NoError(t, err)

	httpClient := NewMockDoer(ctrl)
	httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: httpCode,
		Body:       io.NopCloser(strings.NewReader(`{}`)),
	}, nil).AnyTimes()

	logger := NewMockLogger(ctrl)
	logger.EXPECT().IsDebugLevel().Return(false).AnyTimes()

	exec, err := NewExecutor(conv, httpClient, NewFormatter(), logger)
	require.NoError(t, err)
	return exec
}

func readOnlyRequest(opID string, tier shared_api.PermissionTier) CommandRequest {
	v := shared_api.NewStableVersion(2023, 11, 15)
	return CommandRequest{
		Command: shared_api.Command{
			OperationID: opID,
			Permission:  tier,
			RequestParameters: shared_api.RequestParameters{
				URL:  "/api/atlas/v2/groups",
				Verb: "http.MethodGet",
			},
			Versions: []shared_api.CommandVersion{{
				Version:              v,
				RequestContentType:   "json",
				ResponseContentTypes: []string{"json"},
			}},
		},
		ContentType: "json",
		Format:      "json",
		Version:     v,
	}
}

// TestPledgeBlocksWriteUnderReadonly verifies that a write operation is rejected
// when the active pledge is readonly.
func TestPledgeBlocksWriteUnderReadonly(t *testing.T) {
	withTempStateDir(t)

	k := currentSessionKey(t)
	pf, err := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	require.NoError(t, err)
	require.NoError(t, pledge.Save(k, pf))

	exec := makeExecutor(t, http.StatusOK)
	req := readOnlyRequest("deleteCluster", shared_api.PermissionWrite)

	_, err = exec.ExecuteCommand(context.Background(), req)
	require.Error(t, err)
	// Non-TTY path: returns a token-bearing error (not BlockedError directly).
	require.Contains(t, err.Error(), "atlas pledge [readonly]")
	require.Contains(t, err.Error(), "atlas pledge allow")
}

// TestPledgeAllowsReadUnderReadonly verifies that a read operation is allowed
// when the active pledge is readonly.
func TestPledgeAllowsReadUnderReadonly(t *testing.T) {
	withTempStateDir(t)

	k := currentSessionKey(t)
	pf, err := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	require.NoError(t, err)
	require.NoError(t, pledge.Save(k, pf))

	exec := makeExecutor(t, http.StatusOK)
	req := readOnlyRequest("listClusters", shared_api.PermissionRead)

	resp, err := exec.ExecuteCommand(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

// TestPledgeAllowsAllUnderAdmin verifies that a write operation is permitted
// when the active pledge is admin (highest tier).
func TestPledgeAllowsAllUnderAdmin(t *testing.T) {
	withTempStateDir(t)

	k := currentSessionKey(t)
	pf, err := pledge.NewPledgeFile(pledge.ProfileAdmin, nil)
	require.NoError(t, err)
	require.NoError(t, pledge.Save(k, pf))

	exec := makeExecutor(t, http.StatusOK)
	req := readOnlyRequest("deleteCluster", shared_api.PermissionWrite)

	resp, err := exec.ExecuteCommand(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

// TestPledgeNoPledgeAllowsEverything verifies that without an active pledge
// all operations are permitted.
func TestPledgeNoPledgeAllowsEverything(t *testing.T) {
	withTempStateDir(t)
	// No pledge saved — Load will return ErrNoPledge.

	exec := makeExecutor(t, http.StatusOK)
	req := readOnlyRequest("deleteOrg", shared_api.PermissionAdmin)

	resp, err := exec.ExecuteCommand(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
