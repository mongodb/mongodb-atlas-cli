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
//go:build e2e || e2eSnap || config

package e2e_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const completionEntity = "completion"

func TestAtlasCLIAutocomplete(t *testing.T) {
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)
	options := []string{"zsh", "bash", "fish", "powershell"}
	for _, option := range options {
		o := option
		t.Run(o, func(t *testing.T) {
			t.Parallel()
			cmd := exec.Command(cliPath, completionEntity, o)
			cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
			resp, err := RunAndGetStdOut(cmd)
			require.NoError(t, err, string(resp))
		})
	}
}
