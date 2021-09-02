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
//go:build e2e || (atlas && livemigrations)
// +build e2e atlas,livemigrations

package atlas_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/require"
)

func TestLinkToken(t *testing.T) {
	cliPath, err := e2e.Bin()
	r := require.New(t)
	r.NoError(err)

	// Cleanup, do a delete in case a token already exists
	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			liveMigrationsEntity,
			"link",
			"delete",
			"-o=json")
		cmd.Env = os.Environ()
	})

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			liveMigrationsEntity,
			"link",
			"create",
			"--accessListIp",
			"1.2.3.4,5.6.7.8")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			liveMigrationsEntity,
			"link",
			"delete")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))
	})
}
