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

//go:build e2e || (decrypt && (cloudmanager || om60))

package cloud_manager_test

import (
	"embed"
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/test/e2e"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/test/e2e/decryption"
	"github.com/stretchr/testify/require"
)

//go:embed decryption/kmip/*
var filesKmip embed.FS

const kmipTestsInputDir = "decryption/kmip"

func decodeAndWriteToPath(t *testing.T, encodedText, filepath string) error {
	t.Helper()

	decoded, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, decoded, fs.ModePerm)
}

func dumpCertsToTemp(t *testing.T, tmpDir string) (caFile, certFile string) {
	t.Helper()
	kmipCA := os.Getenv("KMIP_CA")
	require.NotEmpty(t, kmipCA)
	caFile = path.Join(tmpDir, "tls-rootCA.pem")
	require.NoError(t, decodeAndWriteToPath(t, kmipCA, caFile))

	kmipCert := os.Getenv("KMIP_CERT")
	require.NotEmpty(t, kmipCert)
	certFile = path.Join(tmpDir, "tls-localhost.pem")
	require.NoError(t, decodeAndWriteToPath(t, kmipCert, certFile))

	return caFile, certFile
}

func TestDecryptWithKMIP(t *testing.T) {
	cliPath, err := e2e.Bin()
	req := require.New(t)
	req.NoError(err)

	tmpDir := t.TempDir()
	caFile, certFile := dumpCertsToTemp(t, tmpDir)
	for i := 1; i <= 2; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile := decryption.GenerateFileNameCase(tmpDir, i, "input")
			err = decryption.DumpToTemp(filesKmip, decryption.GenerateFileNameCase(kmipTestsInputDir, i, "input"), inputFile)
			req.NoError(err)

			expectedContents, err2 := filesKmip.ReadFile(decryption.GenerateFileNameCase(kmipTestsInputDir, i, "output"))
			req.NoError(err2, string(expectedContents))

			cmd := exec.Command(cliPath,
				entity,
				"logs",
				"decrypt",
				"--file",
				inputFile,
				"--kmipServerCAFile",
				caFile,
				"--kmipClientCertificateFile",
				certFile,
			)
			cmd.Env = os.Environ()

			gotContents, err := cmd.CombinedOutput()
			req.NoError(err, string(gotContents))
			decryption.LogsAreEqual(t, expectedContents, gotContents)
		})
	}
}
