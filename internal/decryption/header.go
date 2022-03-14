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
	"time"

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
)

type HeaderRecord struct {
	Timestamp       time.Time
	Version         string
	CompressionMode CompressionMode
	KeyProvider     keyproviders.KeyProvider
	EncryptedLEK    []byte
	IV              []byte
	AESBlock        []byte
	MAC             string
}

func (header *HeaderRecord) DecryptKey() ([]byte, error) {
	return header.KeyProvider.DecryptKey(header.EncryptedLEK, header.IV)
}

func decodeHeader(logLine *AuditLogLine, opts KeyProviderOpts) (*HeaderRecord, error) {
	keyProvider, err := logLine.KeyProvider(opts)
	if err != nil {
		return nil, err
	}
	return &HeaderRecord{
		Timestamp:       logLine.TS,
		Version:         logLine.Version,
		CompressionMode: CompressionMode(logLine.CompressionMode),
		KeyProvider:     keyProvider,
		IV:              logLine.EncryptedKey[:16],
		EncryptedLEK:    logLine.EncryptedKey[16:48],
		AESBlock:        logLine.EncryptedKey[48:],
	}, nil
}

func processHeader(logLine *AuditLogLine, opts KeyProviderOpts) (*DecryptSection, error) {
	header, err := decodeHeader(logLine, opts)
	if err != nil {
		return nil, err
	}

	lek, err := header.DecryptKey()
	if err != nil {
		return nil, err
	}

	// todo: validate header

	return &DecryptSection{
		lek:               lek,
		compressionMode:   header.CompressionMode,
		processedLogLines: 0,
	}, nil
}
