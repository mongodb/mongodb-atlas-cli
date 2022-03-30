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

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
)

type AuditRecordType string

type AuditLogLineKeyStoreIdentifier struct {
	Provider *keyproviders.KeyStoreProvider `json:"provider,omitempty"`
	// localKey
	Filename *string `json:"filename,omitempty"`
	// kmip
	UID            *string                         `json:"uniqueKeyID,omitempty"`
	KMIPServerName []string                        `json:"kmipServerName,omitempty"`
	KMIPPort       *int                            `json:"kmipPort,omitempty"`
	KeyWrapMethod  *keyproviders.KMIPKeyWrapMethod `json:"keyWrapMethod,omitempty"`
	// aws
	Key      *string `json:"key,omitempty"`
	Region   *string `json:"region,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	// azure & gcp
	KeyName *string `json:"keyName,omitempty"`
	// azure
	Environment      *string `json:"environment,omitempty"`
	KeyVaultEndpoint *string `json:"keyVaultEndpoint,omitempty"`
	KeyVersion       *string `json:"keyVersion,omitempty"`
	// gcp
	ProjectID *string `json:"projectId,omitempty"`
	Location  *string `json:"location,omitempty"`
	KeyRing   *string `json:"keyRing,omitempty"`
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

func (logLine *AuditLogLine) KeyProvider(opts *KeyProviderOpts) (keyproviders.KeyProvider, error) {
	if logLine.AuditRecordType != AuditHeaderRecord {
		return nil, errors.New("not a valid header line")
	}

	if logLine.KeyStoreIdentifier.Provider == nil {
		return nil, errors.New("keyProvider not set")
	}

	switch *logLine.KeyStoreIdentifier.Provider {
	case keyproviders.LocalKey:
		return &keyproviders.LocalKeyIdentifier{
			Filename: opts.Local.KeyFileName,
		}, nil
	case keyproviders.KMIP:
		return &keyproviders.KMIPKeyIdentifier{
			UniqueKeyID:               *logLine.KeyStoreIdentifier.UID,
			ServerNames:               logLine.KeyStoreIdentifier.KMIPServerName,
			ServerPort:                *logLine.KeyStoreIdentifier.KMIPPort,
			KeyWrapMethod:             *logLine.KeyStoreIdentifier.KeyWrapMethod,
			ServerCAFileName:          opts.KMIP.ServerCAFileName,
			ClientCertificateFileName: opts.KMIP.ClientCertificateFileName,
		}, nil
	default:
		return nil, fmt.Errorf("keyProvider %s not implemented", *logLine.KeyStoreIdentifier.Provider)
	}
}

const (
	AuditHeaderRecord AuditRecordType = "header"
	AuditLogRecord    AuditRecordType = ""
)
