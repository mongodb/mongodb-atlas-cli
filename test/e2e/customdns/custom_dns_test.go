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
//go:build e2e || e2eSnap || (atlas && generic)

package customdns

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

const (
	awsEntity       = "aws"
	customDNSEntity = "customDns"
)

func TestCustomDNS(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("customDNS")

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Enable", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDNSEntity,
			awsEntity,
			"enable",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var dns atlasv2.AWSCustomDNSEnabled
		require.NoError(t, json.Unmarshal(resp, &dns))
		a.True(dns.Enabled)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDNSEntity,
			awsEntity,
			"describe",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var dns atlasv2.AWSCustomDNSEnabled
		require.NoError(t, json.Unmarshal(resp, &dns))
		a.True(dns.Enabled)
	})

	g.Run("Disable", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDNSEntity,
			awsEntity,
			"disable",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var dns atlasv2.AWSCustomDNSEnabled
		require.NoError(t, json.Unmarshal(resp, &dns))
		a.False(dns.Enabled)
	})
}
