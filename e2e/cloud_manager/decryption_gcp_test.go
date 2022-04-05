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
)

//go:embed decryption/gcp/*
var filesGCP embed.FS

const GCPTestsInputDir = "decryption/kmip"

func TestDecryptWithGCP(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tmpDir := t.TempDir()

	GCPCredentialsContent, ok := os.LookupEnv("GCP_CREDENTIALS")
	if !ok {
		t.Fatal("GCP Credentials not found")
	}
	GCPCredentialsFile := path.Join(tmpDir, "gcp_credentials.json")
	err = os.WriteFile(GCPCredentialsFile, []byte(GCPCredentialsContent), fs.ModePerm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GCPCredentialsFile)

	t.Cleanup(func() {
		if errCleanup := os.RemoveAll(tmpDir); errCleanup != nil {
			t.Fatal(errCleanup)
		}
	})

	i := 1
	t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
		inputFile, err := e2e.DumpToTemp(filesGCP, GCPTestsInputDir, i, "input", tmpDir)
		if err != nil {
			t.Fatal(err)
		}

		expectedContents, err := filesKmip.ReadFile(e2e.GenerateFileName(GCPTestsInputDir, i, "output"))
		if err != nil {
			t.Fatal(err)
		}

		cmd := exec.Command(cliPath,
			entity,
			"logs",
			"decrypt",
			"--file",
			inputFile,
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
