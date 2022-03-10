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

package logs

import (
	"testing"

	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/test"
)

func TestDecryptBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DecryptBuilder(),
		0,
		[]string{
			flag.File,
			flag.Out,
			flag.AzureSecret,
			flag.AzureClientID,
			flag.AzureTenantID,
			flag.AzureClientID,
			flag.AWSSecretKey,
			flag.AWSSecretKey,
			flag.AWSSessionToken,
			flag.GCPServiceAccountKey,
		},
	)
}
func TestDecrypt_Run(t *testing.T) {
	listOpts := &DecryptOpts{
		inFileName: "test",
	}

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
