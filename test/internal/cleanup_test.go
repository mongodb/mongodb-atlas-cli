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

package internal

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

// list of keys to delete as clean up.
func getKeysToDelete() map[string]struct{} {
	return map[string]struct{}{
		"mongodb-atlas-operator-api-key": {},
		"e2e-test-helper":                {},
		"e2e-test-atlas-org":             {},
	}
}

func TestCleanup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	mode, err := TestRunMode()
	if err != nil {
		t.Fatal(err)
	}

	if mode != TestModeLive {
		t.Skip("skipping test in snapshot mode")
	}

	req := require.New(t)
	cliPath, err := AtlasCLIBin()
	req.NoError(err)

	t.Run("trying to delete org invitations", func(t *testing.T) {
		t.Parallel()
		deleteOrgInvitations(t, cliPath)
	})
	t.Run("trying to delete org teams", func(t *testing.T) {
		t.Parallel()
		deleteOrgTeams(t, cliPath)
	})
	t.Run("trying to delete all org idps", func(t *testing.T) {
		if IsGov() {
			t.Skip("idps are not available on gov")
		}
		t.Parallel()
		deleteAllIDPs(t, cliPath)
	})
	args := []string{projectEntity,
		"list",
		"--limit=500",
		"-o=json",
		"-P",
		ProfileName(),
	}
	if orgID, set := os.LookupEnv("MONGODB_ATLAS_ORG_ID"); set {
		args = append(args, "--orgId", orgID)
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	req.NoError(err, string(resp))
	var projects admin.PaginatedAtlasGroup
	req.NoError(json.Unmarshal(resp, &projects), string(resp))
	t.Logf("projects:\n%s\n", resp)
	for _, project := range projects.GetResults() {
		projectID := project.GetId()
		if projectID == ProfileProjectID() || projectID == os.Getenv("MONGODB_ATLAS_PROJECT_ID") {
			// we have to clean up data federations from default project
			// as this is the only project configured for data federation
			// (has a configured awsRoleId)
			t.Run("delete data federations", func(t *testing.T) {
				t.Parallel()
				deleteAllDataFederations(t, cliPath, projectID)
			})
			t.Cleanup(func() {
				deleteKeys(t, cliPath, getKeysToDelete())
			})

			t.Log("skip deleting default project", projectID)
			continue
		}

		t.Run("trying to delete project "+projectID, func(t *testing.T) {
			t.Parallel()
			t.Cleanup(func() {
				DeleteProjectWithRetry(t, projectID)
				deleteKeys(t, cliPath, getKeysToDelete())
			})
			for _, provider := range []string{"aws", "gcp", "azure"} {
				p := provider
				t.Run("delete network peers for "+p, func(t *testing.T) {
					t.Parallel()
					deleteAllNetworkPeers(t, cliPath, projectID, p)
				})
				t.Run("delete private endpoints for "+p, func(t *testing.T) {
					t.Parallel()
					deleteAllPrivateEndpoints(t, cliPath, projectID, p)
				})
			}
			t.Run("delete all streams", func(t *testing.T) {
				t.Parallel()
				deleteAllStreams(t, cliPath, projectID)
			})

			t.Run("delete all clusters", func(t *testing.T) {
				t.Parallel()
				deleteAllClustersForProject(t, cliPath, projectID)
			})
			t.Run("delete data federations", func(t *testing.T) {
				t.Parallel()
				deleteAllDataFederations(t, cliPath, projectID)
			})
			t.Run("delete all serverless instances", func(t *testing.T) {
				if IsGov() {
					t.Skip("serverless is not available on gov")
				}
				t.Parallel()
				deleteAllServerlessInstances(t, cliPath, projectID)
			})
		})
	}
}
