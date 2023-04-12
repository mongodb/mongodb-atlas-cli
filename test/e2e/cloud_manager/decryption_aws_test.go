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

//go:build e2e || (decrypt && (cloudmanager || om50 || om60))

package cloud_manager_test

import (
	"embed"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e/decryption"
	"github.com/stretchr/testify/require"
)

//go:embed decryption/aws/*
var filesAWS embed.FS

const awsTestsInputDir = "decryption/aws"

func TestDecryptWithAWS(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.Bin()
	req.NoError(err)

	tmpDir := t.TempDir()
	inputFile := decryption.GenerateFileName(tmpDir, "input")
	err = decryption.DumpToTemp(filesAWS, decryption.GenerateFileName(awsTestsInputDir, "input"), inputFile)
	req.NoError(err)

	expectedContents, err := filesAWS.ReadFile(decryption.GenerateFileName(awsTestsInputDir, "output"))
	req.NoError(err)

	cmd := exec.Command(cliPath,
		entity,
		"logs",
		"decrypt",
		"--file",
		inputFile,
	)
	cmd.Env = os.Environ()

	gotContents, err := cmd.CombinedOutput()
	req.NoError(err, string(gotContents))
	decryption.LogsAreEqual(t, expectedContents, gotContents)
}
