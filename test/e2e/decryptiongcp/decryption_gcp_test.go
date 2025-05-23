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

package decryptiongcp

import (
	"embed"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal/decryption"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/*
var filesGCP embed.FS

const gcpTestsInputDir = "testdata"

func TestDecryptWithGCP(t *testing.T) {
	_ = internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	req := require.New(t)

	cliPath, err := internal.AtlasCLIBin()
	req.NoError(err)

	tmpDir := t.TempDir()

	GCPCredentialsContent, ok := os.LookupEnv("GCP_CREDENTIALS")
	req.True(ok, "GCP Credentials not found")
	GCPCredentialsFile := path.Join(tmpDir, "gcp_credentials.json")
	err = os.WriteFile(GCPCredentialsFile, []byte(GCPCredentialsContent), fs.ModePerm)
	req.NoError(err)
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GCPCredentialsFile)

	inputFile := decryption.GenerateFileName(tmpDir, "input")
	err = decryption.DumpToTemp(filesGCP, decryption.GenerateFileName(gcpTestsInputDir, "input"), inputFile)
	req.NoError(err)

	expectedContents, err := filesGCP.ReadFile(decryption.GenerateFileName(gcpTestsInputDir, "output"))
	req.NoError(err)

	cmd := exec.Command(cliPath,
		"logs",
		"decrypt",
		"--file",
		inputFile,
	)
	cmd.Env = os.Environ()

	gotContents, err := internal.RunAndGetStdOut(cmd)
	req.NoError(err, string(gotContents))

	decryption.LogsAreEqual(t, expectedContents, gotContents)
}
