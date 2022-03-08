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
	"encoding/json"
	"strconv"
	"time"

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
)

type AuditRecordType string
type AuditLogFormat string
type AuditLogEncoding struct {
	format AuditLogFormat
}

type AuditLogLine struct {
	UTCTimestamp struct {
		Date struct {
			Value string `json:"$numberLong"`
		} `json:"$date"`
	} `json:"ts"`
	AuditRecordType    AuditRecordType `json:"auditRecordType,omitempty"`
	Log                string          `json:"log,omitempty"`
	Version            string          `json:"version,omitempty"`
	CompressionMode    string          `json:"compressionMode,omitempty"`
	KeyStoreIdentifier struct {
		Provider string `json:"provider"`
		// todo: add the rest of the fields for kmip & cloud kms
		Filename string `json:"filename,omitempty"`
	} `json:"keyStoreIdentifier,omitempty"`
	EncryptedKey struct {
		Binary struct {
			Base64  string `json:"base64"`
			SubType string `json:"subType"`
		} `json:"$binary"`
	} `json:"encryptedKey,omitempty"`
	MAC string `json:"mac,omitempty"`
}

type HeaderRecord struct {
	UTCTimestamp       uint64
	Version            string
	CompressionMode    CompressionMode
	KeyStoreIdentifier keyproviders.KeyStoreIdentifier
	EncryptedLEK       []byte
	IV                 []byte
	AESBlock           []byte
	MAC                string
}

type HeaderAAD struct {
	TS      time.Time `json:"ts"`
	Version string    `json:"version"`
}

type EncryptedLogRecord struct {
	CipherText         []byte
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
		if err := json.Unmarshal([]byte(value), &line); err != nil {
			return nil, err
		}
	} else {
		panic("not implemented for format: " + encoding.format)
	}

	return &line, nil
}

func (logLine *AuditLogLine) UTCTimestampValue() (uint64, error) {
	const Base10 = 10
	const TimestampBitSize = 64
	return strconv.ParseUint(logLine.UTCTimestamp.Date.Value, Base10, TimestampBitSize)
}
