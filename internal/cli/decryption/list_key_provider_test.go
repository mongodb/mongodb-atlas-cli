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

//go:build unit
// +build unit

package decryption

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/spf13/afero"
)

func TestListKeyProviderBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		KeyProvidersListBuilder(),
		0,
		[]string{
			flag.File,
			flag.Output,
		},
	)
}

func TestKeyProviderListOpts_Run(t *testing.T) {
	fileJSON := []byte(`{"ts":{"$date":{"$numberLong":"1644232049921"}},"version":"0.0","compressionMode":"zstd","keyStoreIdentifier":{"provider":"local","filename":"localKey"},"encryptedKey":{"$binary":{"base64":"+yjPCaKKE1M8fZmPGzGHkyfHYxaw34okpavsHzpd8iPVx2+JjOhXwXw5E2FdI5Rcb5JgmcPUFRPISh/7Si1R/g==","subType":"0"}},"MAC":"qE9fUsGK0EuRrrCRAQAAAAAAAAAAAAAA","auditRecordType":"header"}
{"ts":{"$date":{"$numberLong":"1644232049922"}},"log":"1Lu4o8XVMM/Rg7GKAQAAAAEAAAAAAAAA/8tXQ36mEd90OaAOzCOSti7N5a2jr0B9ek48/uvyteG/zUJHyM16Hs3wMEhDqTQGBwGhWSHEqXh0/5Jbz6tXsYHhDTMr1BOsn1zaavZScx/CkO5+Hd8Vx+zeFPREtQTe1y+JngXSIroezeyV0/zF4YC4vpug+OZtrEQLNEgwT2bjaqUyaKDbmzCNetd2Ff/eFfMFzinbzKVgXAC7T4YmDuowqXommEXLIBiYh2u4VagwJKZRw5OGZjnvqwyVpSPgGqLxGKUoFigh3NgC6EuGi17VIs5BLRZOIw7+OfbPgQQiKzjCxCk="}
{"ts":{"$date":{"$numberLong":"1644232049921"}},"version":"0.0","compressionMode":"zstd","keyStoreIdentifier":{"provider":"kmip","uid":"uniqueKeyID","kmipServerName":["kmipServerName"],"kmipPort":{"$numberInt":"8081"},"keyWrapMethod":"get"},"encryptedKey":{"$binary":{"base64":"+yjPCaKKE1M8fZmPGzGHkyfHYxaw34okpavsHzpd8iPVx2+JjOhXwXw5E2FdI5Rcb5JgmcPUFRPISh/7Si1R/g==","subType":"0"}},"MAC":"qE9fUsGK0EuRrrCRAQAAAAAAAAAAAAAA","auditRecordType":"header"}
{"ts":{"$date":{"$numberLong":"1644232049922"}},"log":"1Lu4o8XVMM/Rg7GKAQAAAAEAAAAAAAAA/8tXQ36mEd90OaAOzCOSti7N5a2jr0B9ek48/uvyteG/zUJHyM16Hs3wMEhDqTQGBwGhWSHEqXh0/5Jbz6tXsYHhDTMr1BOsn1zaavZScx/CkO5+Hd8Vx+zeFPREtQTe1y+JngXSIroezeyV0/zF4YC4vpug+OZtrEQLNEgwT2bjaqUyaKDbmzCNetd2Ff/eFfMFzinbzKVgXAC7T4YmDuowqXommEXLIBiYh2u4VagwJKZRw5OGZjnvqwyVpSPgGqLxGKUoFigh3NgC6EuGi17VIs5BLRZOIw7+OfbPgQQiKzjCxCk="}`)

	listOpts := &KeyProviderListOpts{
		file: "test",
		fs:   afero.NewMemMapFs(),
	}
	bufOut := new(bytes.Buffer)
	_ = listOpts.InitOutput(bufOut, listTmpl)()
	_ = afero.WriteFile(listOpts.fs, "test", fileJSON, 0600)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	expected := `local: Filename = localKey
kmip: Unique Key ID = "uniqueKeyID" KMIP Server Name = "[kmipServerName]" KMIP Port = "8081" Key Wrap Method = "get"
`
	if bufOut.String() != expected {
		t.Fatalf("Run() expected: %s got: %v", expected, bufOut.String())
	}
}
