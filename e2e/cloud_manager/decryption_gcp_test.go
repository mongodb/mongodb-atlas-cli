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
//go:build e2e || (decrypt && (cloudmanager || om44 || om50))

package cloud_manager_test

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/require"
)

//go:embed decryption/gcp/*
var filesGCP embed.FS

const GCPTestsInputDir = "decryption/gcp"

func TestDecryptWithGCP(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.Bin()
	req.NoError(err)

	tmpDir := t.TempDir()

	GCPCredentialsContent, ok := os.LookupEnv("GCP_CREDENTIALS")
	req.True(ok, "GCP Credentials not found")
	GCPCredentialsFile := path.Join(tmpDir, "gcp_credentials.json")
	err = os.WriteFile(GCPCredentialsFile, []byte(GCPCredentialsContent), fs.ModePerm)
	req.NoError(err)
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GCPCredentialsFile)

	t.Cleanup(func() {
		err = os.RemoveAll(tmpDir)
		req.NoError(err)
	})

	i := 1
	t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
		inputFile, err := dumpToTemp(filesGCP, GCPTestsInputDir, i, "input", tmpDir)
		req.NoError(err)

		expectedContents, err := filesGCP.ReadFile(generateFileName(GCPTestsInputDir, i, "output"))
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

		equal, err := logsAreEqual(expectedContents, gotContents)
		req.NoError(err)
		req.True(equal, "expected %v, got %v", string(expectedContents), string(gotContents))
	})
}
