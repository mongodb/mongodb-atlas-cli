// Copyright 2021 MongoDB Inc
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
//go:build e2e || (atlas && generic)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const aws = "AWS"

func TestAccessRoles(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("accessRoles")

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			cloudProvidersEntity,
			accessRolesEntity,
			awsEntity,
			"create",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var iamRole atlasv2.CloudProviderAccessRole
		require.NoError(t, json.Unmarshal(resp, &iamRole))
		assert.Equal(t, aws, iamRole.ProviderName)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			cloudProvidersEntity,
			accessRolesEntity,
			"list",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var roles atlasv2.CloudProviderAccessRoles
		require.NoError(t, json.Unmarshal(resp, &roles))
		assert.Len(t, roles.GetAwsIamRoles(), 1)
	})
}
