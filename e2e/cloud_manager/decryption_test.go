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
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/mongodb/mongocli/e2e"
)

//go:embed decryption/*
var files embed.FS

func generateFileName(dir string, i int, suffix string) string {
	return path.Join(dir, fmt.Sprintf("test%v-%v", i, suffix))
}

func dumpToTemp(dir string, i int, suffix string) (string, error) {
	inputFile := generateFileName("decryption", i, suffix)
	outputFile := generateFileName(dir, i, suffix)

	content, err := files.ReadFile(inputFile)
	if err != nil {
		return "", err
	}

	return outputFile, ioutil.WriteFile(outputFile, content, fs.ModePerm)
}

func TestDecrypt(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tmp := t.TempDir()

	t.Cleanup(func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Fatal(err)
		}
	})

	for i := 1; i <= 4; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile, err := dumpToTemp(tmp, i, "input")
			if err != nil {
				t.Fatal(err)
			}
			keyFile, err := dumpToTemp(tmp, i, "localKey")
			if err != nil {
				t.Fatal(err)
			}

			outputFileContents, err := files.ReadFile(generateFileName("decryption", i, "output"))
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
			resp, err := cmd.CombinedOutput()

			if err != nil {
				t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
			}

			if !bytes.Equal(resp, outputFileContents) {
				t.Fatalf("decryption unexpected: expected %v, got %v", string(outputFileContents), string(resp))
			}
		})
	}
}
