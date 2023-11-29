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

//go:build e2e || (atlas && cleanup)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestCleanup(t *testing.T) {
	req := require.New(t)
	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	t.Run("trying to delete org invitations", func(t *testing.T) {
		t.Parallel()
		deleteOrgInvitations(t, cliPath)
	})
	t.Run("trying to delete org teams", func(t *testing.T) {
		t.Parallel()
		deleteOrgTeams(t, cliPath)
	})

	cmd := exec.Command(cliPath,
		projectEntity,
		"list",
		"--limit=500",
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	req.NoError(err, string(resp))

	var projects admin.PaginatedAtlasGroup
	err = json.Unmarshal(resp, &projects)
	req.NoError(err, string(resp))
	t.Logf("%s\n", resp)
	for _, project := range projects.Results {
		projectID := project.GetId()
		if projectID == os.Getenv("MCLI_PROJECT_ID") {
			t.Log("skipping project", projectID)
			continue
		}
		func(projectId string) {
			t.Run(fmt.Sprintf("trying to delete project %s", projectId), func(t *testing.T) {
				t.Parallel()
				t.Run(fmt.Sprintf("trying to delete project's %s resources", projectId), func(t *testing.T) {
					t.Parallel()
					for _, provider := range []string{"aws", "gcp", "azure"} {
						func(provider string) {
							t.Run(fmt.Sprintf("delete network peers for %s", provider), func(t *testing.T) {
								t.Parallel()
								deleteAllNetworkPeers(t, cliPath, projectID, provider)
							})
							t.Run(fmt.Sprintf("delete private endpoints for %s", provider), func(t *testing.T) {
								t.Parallel()
								deleteAllPrivateEndpoints(t, cliPath, projectID, provider)
							})
						}(provider)
					}
					t.Run("delete all clusters", func(t *testing.T) {
						t.Parallel()
						deleteAllClustersForProject(t, cliPath, projectID)
					})
					t.Run("delete datapipelines", func(t *testing.T) {
						t.Parallel()
						deleteDatapipelinesForProject(t, cliPath, projectID)
					})
					t.Run("delete data deferations", func(t *testing.T) {
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
				t.Cleanup(func() {
					deleteProjectWithRetry(t, projectID)
				})
			})
		}(project.GetId())
	}
}
