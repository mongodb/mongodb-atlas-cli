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

func TestCustomDNS(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("customDNS")

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("Enable", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDNSEntity,
			awsEntity,
			"enable",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var dns atlasv2.AWSCustomDNSEnabled
		require.NoError(t, json.Unmarshal(resp, &dns))
		a.True(dns.Enabled)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDNSEntity,
			awsEntity,
			"describe",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var dns atlasv2.AWSCustomDNSEnabled
		require.NoError(t, json.Unmarshal(resp, &dns))
		a.True(dns.Enabled)
	})

	t.Run("Disable", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDNSEntity,
			awsEntity,
			"disable",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var dns atlasv2.AWSCustomDNSEnabled
		require.NoError(t, json.Unmarshal(resp, &dns))
		a.False(dns.Enabled)
	})
}
