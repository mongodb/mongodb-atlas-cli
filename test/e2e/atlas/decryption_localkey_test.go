// Copyright 2022 MongoDB Inc
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

//go:build e2e || (atlas && decrypt)

package atlas_test

import (
	"embed"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e/decryption"
	"github.com/stretchr/testify/require"
)

//go:embed decryption/localKey/*
var filesLocalKey embed.FS

const localKeyTestsInputDir = "decryption/localKey"

func TestDecryptWithLocalKey(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	tmpDir := t.TempDir()
	inputFile := decryption.GenerateFileName(tmpDir, "input")
	err = decryption.DumpToTemp(filesLocalKey, decryption.GenerateFileName(localKeyTestsInputDir, "input"), inputFile)
	req.NoError(err)

	expectedContents, err := filesLocalKey.ReadFile(decryption.GenerateFileName(localKeyTestsInputDir, "output"))
	req.NoError(err)

	cmd := exec.Command(cliPath,
		"logs",
		"decrypt",
		"--file",
		inputFile,
	)
	cmd.Env = os.Environ()

	gotContents, err := e2e.RunAndGetStdOut(cmd)
	req.NoError(err, string(gotContents))

	decryption.LogsAreEqual(t, expectedContents, gotContents)
}
