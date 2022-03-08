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

package decryption

import (
	"encoding/base64"

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
)

func decodeHeader(logLine *AuditLogLine) (*HeaderRecord, error) {
	timestamp, err := logLine.UTCTimestampValue()
	if err != nil {
		return nil, err
	}
	encryptedKey, _ := base64.StdEncoding.DecodeString(logLine.EncryptedKey.Binary.Base64)
	return &HeaderRecord{
		UTCTimestamp:    timestamp,
		Version:         logLine.Version,
		CompressionMode: CompressionMode(logLine.CompressionMode),
		KeyStoreIdentifier: keyproviders.KeyStoreIdentifier{
			Provider: keyproviders.LocalKey,
			Filename: logLine.KeyStoreIdentifier.Filename,
		},
		IV:           encryptedKey[:16],
		EncryptedLEK: encryptedKey[16:48],
		AESBlock:     encryptedKey[48:],
	}, nil
}
func validateHeader(_ *HeaderRecord) error {
	// todo
	return nil
}

func processHeader(logLine *AuditLogLine, credentialsProvider keyproviders.CredentialsProvider) (*DecryptConfig, error) {
	header, err := decodeHeader(logLine)
	if err != nil {
		return nil, err
	}

	lek, err := keyproviders.DecryptLEK(header.KeyStoreIdentifier, header.EncryptedLEK, header.IV, credentialsProvider)
	if err != nil {
		return nil, err
	}

	if err := validateHeader(header); err != nil {
		return nil, err
	}

	return &DecryptConfig{
		lek:             lek,
		compressionMode: header.CompressionMode,
	}, nil
}
