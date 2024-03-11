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
	"errors"
	"fmt"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/decryption/keyproviders"
)

var (
	ErrInvalidHeaderLine       = errors.New("not a valid header line")
	ErrKeyProviderMissing      = errors.New("key provider not set")
	ErrKeyProviderNotSupported = errors.New("key provider not supported")
)

type AuditRecordType string

type AuditLogLineKeyStoreIdentifier struct {
	Provider *keyproviders.KeyStoreProvider `json:"provider,omitempty"`
	// localKey
	Filename string `json:"filename,omitempty"`
	// kmip
	UID            string                         `json:"uniqueKeyID,omitempty"`
	KMIPServerName []string                       `json:"kmipServerName,omitempty"`
	KMIPPort       int                            `json:"kmipPort,omitempty"`
	KeyWrapMethod  keyproviders.KMIPKeyWrapMethod `json:"keyWrapMethod,omitempty"`
}

type AuditLogLine struct {
	TS                 *time.Time
	AuditRecordType    AuditRecordType
	Version            *string
	CompressionMode    *string
	KeyStoreIdentifier AuditLogLineKeyStoreIdentifier
	EncryptedKey       []byte
	MAC                *string
	Log                *string
}

func (logLine *AuditLogLine) KeyProvider(opts KeyProviderOpts) (keyproviders.KeyProvider, error) {
	if logLine.AuditRecordType != AuditHeaderRecord {
		return nil, ErrInvalidHeaderLine
	}

	if logLine.KeyStoreIdentifier.Provider == nil {
		return nil, ErrKeyProviderMissing
	}

	switch *logLine.KeyStoreIdentifier.Provider {
	case keyproviders.LocalKey:
		if opts.Local == nil {
			return nil, fmt.Errorf("%w: %s", ErrKeyProviderNotSupported, *logLine.KeyStoreIdentifier.Provider)
		}
		return &keyproviders.LocalKeyIdentifier{
			HeaderFilename: logLine.KeyStoreIdentifier.Filename,
			Filename:       opts.Local.KeyFileName,
		}, nil
	case keyproviders.KMIP:
		if opts.KMIP == nil {
			return nil, fmt.Errorf("%w: %s", ErrKeyProviderNotSupported, *logLine.KeyStoreIdentifier.Provider)
		}
		return &keyproviders.KMIPKeyIdentifier{
			UniqueKeyID:               logLine.KeyStoreIdentifier.UID,
			ServerNames:               logLine.KeyStoreIdentifier.KMIPServerName,
			ServerPort:                logLine.KeyStoreIdentifier.KMIPPort,
			KeyWrapMethod:             logLine.KeyStoreIdentifier.KeyWrapMethod,
			ServerCAFileName:          opts.KMIP.ServerCAFileName,
			ClientCertificateFileName: opts.KMIP.ClientCertificateFileName,
			ClientCertificatePassword: opts.KMIP.ClientCertificatePassword,
			Username:                  opts.KMIP.Username,
			Password:                  opts.KMIP.Password,
		}, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrKeyProviderNotSupported, *logLine.KeyStoreIdentifier.Provider)
	}
}

const (
	AuditHeaderRecord AuditRecordType = "header"
	AuditLogRecord    AuditRecordType = ""
)
