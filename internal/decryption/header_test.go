// Copyright 2020 MongoDB Inc
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
	"testing"
	"time"

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
)

func Test_validateHeaderFields(t *testing.T) {
	ts := time.Now()
	version := "0.0"
	compressionMode := "none"
	provider := keyproviders.LocalKey
	encryptedKey := []byte{0, 1, 2, 3}
	mac := "mac"
	recordType := AuditHeaderRecord

	testCases := []struct {
		input     AuditLogLine
		expectErr bool
	}{
		{
			input:     AuditLogLine{},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				Version:         &version,
				CompressionMode: &compressionMode,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				EncryptedKey:    encryptedKey,
				MAC:             &mac,
				AuditRecordType: recordType,
			},
			expectErr: false,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				Version:         &version,
				CompressionMode: &compressionMode,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				EncryptedKey: encryptedKey,
				MAC:          &mac,
			},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				Version:         &version,
				CompressionMode: &compressionMode,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				EncryptedKey:    encryptedKey,
				AuditRecordType: recordType,
			},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				Version:         &version,
				CompressionMode: &compressionMode,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				MAC:             &mac,
				AuditRecordType: recordType,
			},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				Version:         &version,
				CompressionMode: &compressionMode,
				MAC:             &mac,
				EncryptedKey:    encryptedKey,
				AuditRecordType: recordType,
			},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				CompressionMode: &compressionMode,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				MAC:             &mac,
				EncryptedKey:    encryptedKey,
				AuditRecordType: recordType,
			},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				Version:         &version,
				CompressionMode: &compressionMode,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				MAC:             &mac,
				EncryptedKey:    encryptedKey,
				AuditRecordType: recordType,
			},
			expectErr: true,
		},
	}
	for _, testCase := range testCases {
		err := validateHeaderFields(&testCase.input)
		if testCase.expectErr && err == nil {
			t.Errorf("expected: not nil got: %v", err)
		} else if !testCase.expectErr && err != nil {
			t.Errorf("expected: nil got: %v", err)
		}
	}
}
