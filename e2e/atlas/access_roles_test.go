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
// +build e2e atlas,generic

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

const aws = "AWS"

func TestAccessRoles(t *testing.T) {
	n, err := e2e.RandInt(255)
	assert.NoError(t, err)

	cliPath, err := e2e.Bin()
	assert.NoError(t, err)

	projectName := fmt.Sprintf("e2e-integration-access-roles-%v", n)
	projectID, err := createProject(projectName)
	assert.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			cloudProvidersEntity,
			accessRolesEntity,
			awsEntity,
			"create",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var iamRole mongodbatlas.AWSIAMRole
		if err := json.Unmarshal(resp, &iamRole); a.NoError(err) {
			a.Equal(aws, iamRole.ProviderName)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			cloudProvidersEntity,
			accessRolesEntity,
			"list",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var roles mongodbatlas.CloudProviderAccessRoles
		if err := json.Unmarshal(resp, &roles); a.NoError(err) {
			a.Len(roles.AWSIAMRoles, 1)
		}
	})
}
