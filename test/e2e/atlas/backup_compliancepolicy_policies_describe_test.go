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

//go:build e2e || (atlas && backup && compliancepolicy)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

func TestPoliciesDescribe(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("describe-compliance-policy-policies")
	err = enableCompliancePolicy(g.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("unable to enable compliance policy: %w", err))
	}

	cmd := exec.Command(cliPath,
		backupsEntity,
		compliancepolicyEntity,
		policiesEntity,
		"describe",
		"--projectId",
		g.projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, outputErr := cmd.CombinedOutput()

	r.NoError(outputErr, string(resp))

	a := assert.New(t)
	var result atlasv2.DataProtectionSettings
	err = json.Unmarshal(resp, &result)
	a.NoError(err, string(resp))
	a.NotEmpty(result)

}
