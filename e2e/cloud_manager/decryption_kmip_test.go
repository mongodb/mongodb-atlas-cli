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
	"encoding/base64"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/e2e"
	"github.com/mongodb/mongodb-atlas-cli/e2e/decryption"
	"github.com/stretchr/testify/require"
	exec "golang.org/x/sys/execabs"
)

//go:embed decryption/kmip/*
var filesKmip embed.FS

const kmipTestsInputDir = "decryption/kmip"

var errEmptyText = errors.New("unexpected empty value")

func decodeAndWriteToPath(encodedText, filepath string) error {
	if encodedText == "" {
		return errEmptyText
	}

	decoded, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, decoded, fs.ModePerm)
}

func dumpCertsToTemp(tmpDir string) (caFile, certFile string, err error) {
	caFile = path.Join(tmpDir, "tls-rootCA.pem")
	certFile = path.Join(tmpDir, "tls-localhost.pem")

	if err := decodeAndWriteToPath(os.Getenv("KMIP_CA"), caFile); err != nil {
		return "", "", err
	}

	if err := decodeAndWriteToPath(os.Getenv("KMIP_CERT"), certFile); err != nil {
		return "", "", err
	}

	return caFile, certFile, nil
}

func TestDecryptWithKMIP(t *testing.T) {
	cliPath, err := e2e.Bin()
	req := require.New(t)
	req.NoError(err)

	tmpDir := t.TempDir()

	t.Cleanup(func() {
		err = os.RemoveAll(tmpDir)
		req.NoError(err)
	})

	caFile, certFile, err := dumpCertsToTemp(tmpDir)
	req.NoError(err)

	for i := 1; i <= 2; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile := decryption.GenerateFileNameCase(tmpDir, i, "input")
			err := decryption.DumpToTemp(filesKmip, decryption.GenerateFileNameCase(kmipTestsInputDir, i, "input"), inputFile)
			req.NoError(err)

			expectedContents, err := filesKmip.ReadFile(decryption.GenerateFileNameCase(kmipTestsInputDir, i, "output"))
			req.NoError(err)

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

			equal, err := decryption.LogsAreEqual(expectedContents, gotContents)
			req.NoError(err)
			req.True(equal, "expected %v, got %v", string(expectedContents), string(gotContents))
		})
	}
}
