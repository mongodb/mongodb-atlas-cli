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
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/e2e"
	"github.com/mongodb/mongodb-atlas-cli/e2e/decryption"
	"github.com/stretchr/testify/require"
	exec "golang.org/x/sys/execabs"
)

//go:embed decryption/localKey/*
var files embed.FS

const localKeyTestsInputDir = "decryption/localKey"

func TestDecrypt(t *testing.T) {
	cliPath, err := e2e.Bin()
	req := require.New(t)
	req.NoError(err)

	tmp := t.TempDir()

	keyFileContent, ok := os.LookupEnv("LOCAL_KEY")
	req.True(ok, "Local key not found")
	keyFile := path.Join(tmp, "localKey")
	err = os.WriteFile(keyFile, []byte(keyFileContent), fs.ModePerm)
	req.NoError(err)

	t.Cleanup(func() {
		err = os.RemoveAll(tmp)
		req.NoError(err)
	})

	for i := 1; i <= 5; i++ {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			inputFile := decryption.GenerateFileNameCase(tmp, i, "input")
			err := decryption.DumpToTemp(files, decryption.GenerateFileNameCase(localKeyTestsInputDir, i, "input"), inputFile)
			req.NoError(err)

			expectedContents, err := files.ReadFile(decryption.GenerateFileNameCase(localKeyTestsInputDir, i, "output"))
			req.NoError(err)

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
			req.NoError(err, string(gotContents))

			equal, err := decryption.LogsAreEqual(expectedContents, gotContents)
			req.NoError(err)
			req.True(equal, "expected %v, got %v", string(expectedContents), string(gotContents))
		})
	}
}
