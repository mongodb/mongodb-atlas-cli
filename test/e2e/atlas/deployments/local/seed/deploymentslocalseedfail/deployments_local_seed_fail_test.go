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

package deploymentslocalseedfail

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

const (
	deploymentEntity = "deployments"
)

func TestDeploymentsLocalSeedFail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	req := require.New(t)
	mode, err := internal.TestRunMode()
	req.NoError(err)

	if mode != internal.TestModeLive {
		t.Skip("skipping test in snapshot mode")
	}

	const (
		deploymentName = "test-seed-fail"
		dbUsername     = "admin"
		dbUserPassword = "testpwd"
	)

	cliPath, err := internal.AtlasCLIBin()
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
			"--initdb", "./testdata/db_seed_fail",
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
