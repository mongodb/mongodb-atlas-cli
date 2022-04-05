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
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
)

//go:embed decryption/kmip/*
var filesKmip embed.FS

const KmipTestsInputDir = "decryption/kmip"

func decodeAndWriteToPath(encodedText, filepath string) error {
	if encodedText == "" {
		return errors.New("unexpected empty value")
	}

	decoded, err := base64.StdEncoding.DecodeString(encodedText)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, decoded, fs.ModePerm)
}

func dumpCertsToTemp(tmpDir string) (caFile, certFile string, err error) {
	caFile = tmpDir + "/tls-rootCA.pem"
	certFile = tmpDir + "/tls-localhost.pem"

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
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tmpDir := t.TempDir()

	t.Cleanup(func() {
		if errCleanup := os.RemoveAll(tmpDir); errCleanup != nil {
			t.Fatal(errCleanup)
		}
	})

	caFile, certFile, err := dumpCertsToTemp(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= 2; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile, err := e2e.DumpToTemp(filesKmip, KmipTestsInputDir, i, "input", tmpDir)
			if err != nil {
				t.Fatal(err)
			}

			expectedContents, err := filesKmip.ReadFile(e2e.GenerateFileName(KmipTestsInputDir, i, "output"))
			if err != nil {
				t.Fatal(err)
			}

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
			if err != nil {
				t.Fatalf("unexpected error: %v, resp: %s", err, string(gotContents))
			}

			if equal, err := e2e.LogsAreEqual(expectedContents, gotContents); !equal {
				t.Fatalf("decryption unexpected: expected %v, got %v, %v", string(expectedContents), string(gotContents), err)
			}
		})
	}
}
