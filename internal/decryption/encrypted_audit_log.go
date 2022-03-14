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
	"fmt"
	"time"

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
	"go.mongodb.org/mongo-driver/bson"
)

type AuditRecordType string
type AuditLogFormat string
type AuditLogEncoding struct {
	format AuditLogFormat
}

type AuditLogLine struct {
	TS                 *time.Time
	AuditRecordType    AuditRecordType
	Version            *string
	CompressionMode    *string
	KeyStoreIdentifier struct {
		Provider *keyproviders.KeyStoreProvider
		// localKey
		Filename *string
		// kmip
		UniqueKeyID    *string
		KmipServerName *string
		KmipPort       *string
		KeyWrapMethod  *keyproviders.KMIPKeyWrapMethod
		// aws
		Key      *string
		Region   *string
		Endpoint *string
		// azure & gcp
		KeyName *string
		// azure
		Environment      *string
		KeyVaultEndpoint *string
		KeyVersion       *string
		// gcp
		ProjectID *string
		Location  *string
		KeyRing   *string
	}
	EncryptedKey []byte
	MAC          *string
	Log          *string
}

func (logLine *AuditLogLine) KeyProvider(opts KeyProviderOpts) (keyproviders.KeyProvider, error) {
	if logLine.AuditRecordType != AuditHeaderRecord {
		return nil, fmt.Errorf("not a valid header line")
	}

	if logLine.KeyStoreIdentifier.Provider == nil {
		return nil, fmt.Errorf("keyProvider not set")
	}

	switch *logLine.KeyStoreIdentifier.Provider {
	case keyproviders.LocalKey:
		return &keyproviders.LocalKeyIdentifier{
			Filename: opts.LocalKeyFileName,
		}, nil
	case keyproviders.KMIP:
		return &keyproviders.KMIPKeyIdentifier{
			UniqueKeyID:               *logLine.KeyStoreIdentifier.UniqueKeyID,
			ServerName:                *logLine.KeyStoreIdentifier.KmipServerName,
			ServerPort:                *logLine.KeyStoreIdentifier.KmipPort,
			KeyWrapMethod:             *logLine.KeyStoreIdentifier.KeyWrapMethod,
			ServerCAFileName:          opts.KMIPServerCAFileName,
			ClientCertificateFileName: opts.KMIPClientCertificateFileName,
		}, nil
	default:
		return nil, fmt.Errorf("keyProvider %s not implemented", *logLine.KeyStoreIdentifier.Provider)
	}
}

type HeaderAAD struct {
	TS      time.Time `json:"ts"`
	Version string    `json:"version"`
}

type DecodedLogRecord struct {
	CipherText         []byte
	Tag                []byte
	IV                 []byte
	AAD                []byte
	KeyInitCount       uint32
	KeyInvocationCount uint64
}

const (
	AuditHeaderRecord AuditRecordType = "header"
	AuditLogRecord    AuditRecordType = ""
)

const (
	JSON AuditLogFormat = "JSON"
	BSON AuditLogFormat = "BSON"
)

func newAuditLogEncoding(format AuditLogFormat) AuditLogEncoding {
	return AuditLogEncoding{
		format: format,
	}
}

func (encoding *AuditLogEncoding) Parse(value string) (*AuditLogLine, error) {
	var line AuditLogLine

	if encoding.format == JSON {
		if err := bson.UnmarshalExtJSON([]byte(value), true, &line); err != nil {
			return nil, err
		}
	} else {
		// todo: implement BSON
		return nil, fmt.Errorf("not implemented for format: %s", encoding.format)
	}

	return &line, nil
}
