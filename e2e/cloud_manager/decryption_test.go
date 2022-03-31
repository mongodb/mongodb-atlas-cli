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
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
)

//go:embed decryption/localKey/*
var files embed.FS

const LocalKeyTestsInputDir = "decryption/localKey"

func TestDecrypt(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tmp := t.TempDir()

	t.Cleanup(func() {
		if err := os.RemoveAll(tmp); err != nil {
			t.Fatal(err)
		}
	})

	for i := 1; i <= 4; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile, err := e2e.DumpToTemp(files, LocalKeyTestsInputDir, i, "input", tmp)
			if err != nil {
				t.Fatal(err)
			}
			keyFile, err := e2e.DumpToTemp(files, LocalKeyTestsInputDir, i, "localKey", tmp)
			if err != nil {
				t.Fatal(err)
			}

			expectedContents, err := files.ReadFile(e2e.GenerateFileName(LocalKeyTestsInputDir, i, "output"))
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(cliPath,
				entity,
				"logs",
				"decrypt",
				"--file",
				inputFile,
				"--localKeyFile",
				keyFile,
			)
			cmd.Env = os.Environ()
			gotContents, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("unexpected error: %v, resp: %v", err, string(gotContents))
			}

			if equal, err := e2e.LogsAreEqual(expectedContents, gotContents); !equal {
				t.Fatalf("decryption unexpected: expected %v, got %v, %v", string(expectedContents), string(gotContents), err)
			}
		})
	}
}
