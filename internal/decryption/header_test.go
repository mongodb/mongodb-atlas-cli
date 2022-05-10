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
	"encoding/base64"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/decryption/keyproviders"
)

func Test_validateMAC(t *testing.T) {
	validKey, _ := base64.StdEncoding.DecodeString("pnvb++3sbhxIJdfODOq5uIaUX8yxTuWS95VLgES30FM=")
	invalidKey, _ := base64.StdEncoding.DecodeString("pnvb++4sbhxIJdfODOq5uIaUX8yxTuWS95VLgES30FM=")
	validVersion := "0.0"
	invalidVersion := "0.1"
	validTimestamp := time.UnixMilli(1644232049921)
	invalidTimestamp := time.UnixMilli(0)
	validMAC := "qE9fUsGK0EuRrrCRAQAAAAAAAAAAAAAA"
	invalidMAC := "wrongAAAAAAAAAAAAAAAAAAAAAAAAAAA"

	testCases := []struct {
		input       HeaderRecord
		inputKey    []byte
		expectedErr bool
	}{
		{
			input: HeaderRecord{
				Timestamp: validTimestamp,
				Version:   validVersion,
				MAC:       validMAC,
			},
			inputKey:    validKey,
			expectedErr: false,
		},
		{
			input: HeaderRecord{
				Timestamp: invalidTimestamp,
				Version:   validVersion,
				MAC:       validMAC,
			},
			inputKey:    validKey,
			expectedErr: true,
		},
		{
			input: HeaderRecord{
				Timestamp: validTimestamp,
				Version:   invalidVersion,
				MAC:       validMAC,
			},
			inputKey:    validKey,
			expectedErr: true,
		},
		{
			input: HeaderRecord{
				Timestamp: validTimestamp,
				Version:   validVersion,
				MAC:       invalidMAC,
			},
			inputKey:    validKey,
			expectedErr: true,
		},
		{
			input: HeaderRecord{
				Timestamp: validTimestamp,
				Version:   validVersion,
				MAC:       validMAC,
			},
			inputKey:    invalidKey,
			expectedErr: true,
		},
	}

	for _, testCase := range testCases {
		err := testCase.input.validateMAC(testCase.inputKey)
		if testCase.expectedErr && err == nil {
			t.Errorf("expected: not nil got: %v", err)
		} else if !testCase.expectedErr && err != nil {
			t.Errorf("expected: nil got: %v", err)
		}
	}
}

func Test_validateHeaderFields(t *testing.T) {
	ts := time.Now()
	version := "0.0"
	compressionMode := "none"
	invalidCompressionMode := "foo"
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
		{
			input: AuditLogLine{
				TS:      &ts,
				Version: &version,
				KeyStoreIdentifier: AuditLogLineKeyStoreIdentifier{
					Provider: &provider,
				},
				EncryptedKey:    encryptedKey,
				MAC:             &mac,
				AuditRecordType: recordType,
			},
			expectErr: true,
		},
		{
			input: AuditLogLine{
				TS:              &ts,
				Version:         &version,
				CompressionMode: &invalidCompressionMode,
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
