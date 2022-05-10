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
	"os"
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/e2e"
	exec "golang.org/x/sys/execabs"
)

func TestKeyProviders(t *testing.T) {
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

	t.Run("List", func(t *testing.T) {
		filepath := path.Join(tmp, "input.json")

		err := os.WriteFile(filepath, []byte(`{"ts":{"$date":{"$numberLong":"1644232049921"}},"version":"0.0","compressionMode":"zstd","keyStoreIdentifier":{"provider":"local","filename":"localKey"},"encryptedKey":{"$binary":{"base64":"+yjPCaKKE1M8fZmPGzGHkyfHYxaw34okpavsHzpd8iPVx2+JjOhXwXw5E2FdI5Rcb5JgmcPUFRPISh/7Si1R/g==","subType":"0"}},"MAC":"qE9fUsGK0EuRrrCRAQAAAAAAAAAAAAAA","auditRecordType":"header"}
{"ts":{"$date":{"$numberLong":"1644232049922"}},"log":"1Lu4o8XVMM/Rg7GKAQAAAAEAAAAAAAAA/8tXQ36mEd90OaAOzCOSti7N5a2jr0B9ek48/uvyteG/zUJHyM16Hs3wMEhDqTQGBwGhWSHEqXh0/5Jbz6tXsYHhDTMr1BOsn1zaavZScx/CkO5+Hd8Vx+zeFPREtQTe1y+JngXSIroezeyV0/zF4YC4vpug+OZtrEQLNEgwT2bjaqUyaKDbmzCNetd2Ff/eFfMFzinbzKVgXAC7T4YmDuowqXommEXLIBiYh2u4VagwJKZRw5OGZjnvqwyVpSPgGqLxGKUoFigh3NgC6EuGi17VIs5BLRZOIw7+OfbPgQQiKzjCxCk="}`), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		cmd := exec.Command(cliPath,
			entity,
			"logs",
			"keyProviders",
			"list",
			"--file",
			filepath,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if len(resp) == 0 {
			t.Fatalf(`expected len(resp) > 0, got 0`)
		}
	})
}
