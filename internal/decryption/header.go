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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/decryption/aes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/decryption/keyproviders"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrTimestampMissing       = errors.New("missing timestamp")
	ErrVersionMissing         = errors.New("missing version")
	ErrCompressionModeMissing = errors.New("missing compression mode")
	ErrCompressionModeInvalid = errors.New("invalid compression mode")
	ErrProviderMissing        = errors.New("missing provider")
	ErrEncryptedKeyMissing    = errors.New("missing encrypted key")
	ErrMACMissing             = errors.New("missing mac")
	ErrHeaderRecordInvalid    = errors.New("incorrect header record")
)

type HeaderRecord struct {
	Timestamp       time.Time
	Version         string
	CompressionMode CompressionMode
	KeyProvider     keyproviders.KeyProvider
	EncryptedKey    []byte
	MAC             string
}

func (h *HeaderRecord) DecryptKey() ([]byte, error) {
	if err := h.KeyProvider.ValidateCredentials(); err != nil {
		return nil, err
	}
	return h.KeyProvider.DecryptKey(h.EncryptedKey)
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
		EncryptedKey:    logLine.EncryptedKey,
		MAC:             *logLine.MAC,
	}, nil
}

func validateHeaderFields(logLine *AuditLogLine) error {
	if logLine.TS == nil {
		return ErrTimestampMissing
	}

	if logLine.Version == nil {
		return ErrVersionMissing
	}

	if logLine.CompressionMode == nil {
		return ErrCompressionModeMissing
	}

	c := CompressionMode(*logLine.CompressionMode)

	if c != CompressionModeNone && c != CompressionModeZstd {
		return ErrCompressionModeInvalid
	}
	if logLine.KeyStoreIdentifier.Provider == nil {
		return ErrProviderMissing
	}

	if logLine.EncryptedKey == nil {
		return ErrEncryptedKeyMissing
	}

	if logLine.MAC == nil {
		return ErrMACMissing
	}

	if logLine.AuditRecordType != AuditHeaderRecord {
		return ErrHeaderRecordInvalid
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

	return bson.Marshal(aadData)
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
