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
	"path"
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

	keyFile := path.Join(tmp, "localKey")
	err = dumpToTempFile(files, path.Join(LocalKeyTestsInputDir, "localKey"), keyFile)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := os.RemoveAll(tmp); err != nil {
			t.Fatal(err)
		}
	})

	for i := 1; i <= 4; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile, err := dumpToTemp(files, LocalKeyTestsInputDir, i, "input", tmp)
			if err != nil {
				t.Fatal(err)
			}

			expectedContents, err := files.ReadFile(generateFileName(LocalKeyTestsInputDir, i, "output"))
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

			if equal, err := logsAreEqual(expectedContents, gotContents); !equal {
				t.Fatalf("decryption unexpected: expected %v, got %v, %v", string(expectedContents), string(gotContents), err)
			}
		})
	}
}
