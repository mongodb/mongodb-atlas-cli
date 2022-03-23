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
	"errors"
	"time"

	"github.com/mongodb/mongocli/internal/decryption/aes"
	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
	"go.mongodb.org/mongo-driver/bson"
)

type HeaderEncryptedBlock struct {
	LEK   []byte
	IV    []byte
	Block []byte
}

type HeaderRecord struct {
	Timestamp       time.Time
	Version         string
	CompressionMode CompressionMode
	KeyProvider     keyproviders.KeyProvider
	EncryptedBlock  HeaderEncryptedBlock
	MAC             string
}

func (h *HeaderRecord) DecryptKey() ([]byte, error) {
	return h.KeyProvider.DecryptKey(h.EncryptedBlock.LEK, h.EncryptedBlock.IV)
}

func decodeHeader(logLine *AuditLogLine, opts KeyProviderOpts) (*HeaderRecord, error) {
	keyProvider, err := logLine.KeyProvider(opts)
	if err != nil {
		return nil, err
	}
	return &HeaderRecord{
		Timestamp:       *logLine.TS,
		Version:         *logLine.Version,
		CompressionMode: CompressionMode(*logLine.CompressionMode),
		KeyProvider:     keyProvider,
		EncryptedBlock: HeaderEncryptedBlock{
			IV:    logLine.EncryptedKey[:16],
			LEK:   logLine.EncryptedKey[16:48],
			Block: logLine.EncryptedKey[48:],
		},
		MAC: *logLine.MAC,
	}, nil
}

func validateHeaderFields(logLine *AuditLogLine) error {
	if logLine.TS == nil {
		return errors.New("missing timestamp")
	}

	if logLine.Version == nil {
		return errors.New("missing version")
	}

	if logLine.CompressionMode == nil {
		return errors.New("missing compression mode")
	}

	c := CompressionMode(*logLine.CompressionMode)

	if c != CompressionModeNone && c != CompressionModeZstd {
		return errors.New("invalid compression mode")
	}
	if logLine.KeyStoreIdentifier.Provider == nil {
		return errors.New("missing provider")
	}

	if logLine.EncryptedKey == nil {
		return errors.New("missing encrypted key")
	}

	if logLine.MAC == nil {
		return errors.New("missing mac")
	}

	if logLine.AuditRecordType != AuditHeaderRecord {
		return errors.New("incorrect header record")
	}

	return nil
}

func processHeader(logLine *AuditLogLine, opts KeyProviderOpts) (*DecryptSection, error) {
	err := validateHeaderFields(logLine)
	if err != nil {
		return nil, err
	}

	header, err := decodeHeader(logLine, opts)
	if err != nil {
		return nil, err
	}

	lek, err := header.DecryptKey()
	if err != nil {
		return nil, err
	}

	err = header.validateMAC(lek)
	if err != nil {
		return nil, err
	}

	return &DecryptSection{
		lek:                    lek,
		compressionMode:        header.CompressionMode,
		lastKeyInvocationCount: 0,
	}, nil
}

func (h *HeaderRecord) generateAAD() ([]byte, error) {
	aadData := struct {
		TS      time.Time
		Version string
	}{
		TS:      h.Timestamp,
		Version: h.Version,
	}

	aad, err := bson.Marshal(aadData)
	if err != nil {
		return nil, err
	}

	return aad, nil
}

func (h *HeaderRecord) validateMAC(decryptedKey []byte) error {
	aad, err := h.generateAAD()
	if err != nil {
		return err
	}

	mac, err := base64.StdEncoding.DecodeString(h.MAC)
	if err != nil {
		return err
	}

	tag := mac[:12]
	iv := mac[12:24]
	text := mac[24:]

	input := aes.GCMInput{
		Key: decryptedKey,
		IV:  iv,
		AAD: aad,
		Tag: tag,
	}
	_, err = input.Decrypt(text)
	return err
}
