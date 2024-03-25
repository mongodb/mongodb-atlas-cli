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

package logs

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
)

func TestDecryptBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DecryptBuilder(),
		0,
		[]string{
			flag.File,
			flag.Out,
			flag.AWSAccessKey,
			flag.AWSSecretKey,
			flag.AWSSessionToken,
			flag.AzureClientID,
			flag.AzureTenantID,
			flag.AzureSecret,
			flag.GCPServiceAccountKey,
		},
	)
}

func TestDecrypt_Run(t *testing.T) {
	fileJSON := []byte(`{"ts":{"$date":{"$numberLong":"1644232049921"}},"version":"0.0","compressionMode":"zstd","keyStoreIdentifier":{"provider":"local","filename":"localKey"},"encryptedKey":{"$binary":{"base64":"+yjPCaKKE1M8fZmPGzGHkyfHYxaw34okpavsHzpd8iPVx2+JjOhXwXw5E2FdI5Rcb5JgmcPUFRPISh/7Si1R/g==","subType":"0"}},"MAC":"qE9fUsGK0EuRrrCRAQAAAAAAAAAAAAAA","auditRecordType":"header"}
{"ts":{"$date":{"$numberLong":"1644232049922"}},"log":"1Lu4o8XVMM/Rg7GKAQAAAAEAAAAAAAAA/8tXQ36mEd90OaAOzCOSti7N5a2jr0B9ek48/uvyteG/zUJHyM16Hs3wMEhDqTQGBwGhWSHEqXh0/5Jbz6tXsYHhDTMr1BOsn1zaavZScx/CkO5+Hd8Vx+zeFPREtQTe1y+JngXSIroezeyV0/zF4YC4vpug+OZtrEQLNEgwT2bjaqUyaKDbmzCNetd2Ff/eFfMFzinbzKVgXAC7T4YmDuowqXommEXLIBiYh2u4VagwJKZRw5OGZjnvqwyVpSPgGqLxGKUoFigh3NgC6EuGi17VIs5BLRZOIw7+OfbPgQQiKzjCxCk="}`)

	testCases := []struct {
		input []byte
	}{
		{input: fileJSON},
	}

	for _, testCase := range testCases {
		listOpts := &DecryptOpts{
			inFileName: "log",
			awsOpts: DecryptAWSOpts{
				awsAccessKey:       "test",
				awsSecretAccessKey: "test",
				awsSessionToken:    "test",
			},
		}
		listOpts.Out = "decryptedAuditLog"
		listOpts.Fs = afero.NewMemMapFs()
		_ = afero.WriteFile(listOpts.Fs, "log", testCase.input, 0600)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	}
}
