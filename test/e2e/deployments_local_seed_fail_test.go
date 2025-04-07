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
//go:build e2e || (atlas && deployments && local && seed)

package e2e

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

func TestDeploymentsLocalSeedFail(t *testing.T) {
	const (
		deploymentName = "test-seed-fail"
		dbUsername     = "admin"
		dbUserPassword = "testpwd"
	)

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("Test failed seed setup", func(t *testing.T) {
		t.Cleanup(func() {
			// This shouldn't be required, since setup should fail
			// but in case there's a bug we might create a deployment, so try to delete it just in case
			cmd := exec.Command(cliPath,
				deploymentEntity,
				"delete",
				deploymentName,
				"--type",
				"local",
				"--force",
			)

			cmd.Env = os.Environ()

			_, _ = internal.RunAndGetStdOutAndErr(cmd)
		})

		cmd := exec.Command(cliPath,
			deploymentEntity,
			"setup",
			deploymentName,
			"--type", "local",
			"--username", dbUsername,
			"--password", dbUserPassword,
			"--initdb", "./data/db_seed_fail",
			"--bindIpAll",
			"--force",
			"--debug",
		)

		cmd.Env = os.Environ()

		_, setupErr := internal.RunAndGetStdOutAndErr(cmd)
		require.Error(t, setupErr, "expected seeding to fail")
		require.Contains(t, setupErr.Error(), "CATASTROPHIC_SEED_FAILURE_PANIC_DONT_PANIC")
	})
}
