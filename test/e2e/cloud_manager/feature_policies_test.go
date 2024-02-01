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

//go:build e2e || (generic && (cloudmanager || om60))

package cloud_manager_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/andreangiolillo/mongocli-test/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	policyExternallyManagedLock = "EXTERNALLY_MANAGED_LOCK"
	policyDisableUserManagement = "DISABLE_USER_MANAGEMENT"
)

func TestFeaturePolicies(t *testing.T) {
	n, err := e2e.RandInt(255)
	require.NoError(t, err)

	projectName := fmt.Sprintf("e2e-maintenance-proj-%v", n)
	projectID, err := e2e.CreateProject(projectName)
	require.NoError(t, err)

	cliPath, err := e2e.Bin()
	require.NoError(t, err)
	t.Cleanup(func() {
		e2e.DeleteProjectWithRetry(t, projectID)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			featurePolicies,
			"update",
			"--name",
			"test",
			"--policy",
			policyExternallyManagedLock,
			"--policy",
			policyDisableUserManagement,
			"-o=json",
			"--projectId",
			projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var policy *opsmngr.FeaturePolicy
		err2 := json.Unmarshal(resp, &policy)
		require.NoError(t, err2, string(resp))
		assert.ElementsMatch(t,
			policy.Policies, []*opsmngr.Policy{
				{Policy: policyExternallyManagedLock},
				{Policy: policyDisableUserManagement},
			},
		)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			featurePolicies,
			"list",
			"-o=json",
			"--projectId",
			projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var policy *opsmngr.FeaturePolicy
		err2 := json.Unmarshal(resp, &policy)
		require.NoError(t, err2, string(resp))
		assert.NotEmpty(t, policy.Policies)
	})
}
