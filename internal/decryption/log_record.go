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
	"encoding/binary"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (logLine *AuditLogLine) decodeLogRecord() (*EncryptedLogRecord, error) {
	aad, err := logLine.logAdditionalAuthData()
	if err != nil {
		return nil, err
	}

	log, err := base64.StdEncoding.DecodeString(logLine.Log)
	if err != nil {
		return nil, err
	}

	TagData := log[0:12]
	KeyInitCountData := log[12:16]
	KeyInvocationCountData := log[16:24]
	LogData := log[24:]

	return &EncryptedLogRecord{
		CipherText:         append(LogData, TagData...),
		AAD:                aad,
		IV:                 append(KeyInitCountData, KeyInvocationCountData...),
		KeyInitCount:       binary.LittleEndian.Uint32(KeyInitCountData),
		KeyInvocationCount: binary.LittleEndian.Uint64(KeyInvocationCountData),
	}, nil
}

func processLogRecord(decryptConfig *DecryptConfig, logLine *AuditLogLine, lineNb int, expectedLogRecordIdx uint64) (interface{}, error) {
	encryptedLogRecord, err := logLine.decodeLogRecord()
	if err != nil {
		return nil, fmt.Errorf("line %v is corrupted, %v", lineNb, err)
	}

	err = validateLogLine(encryptedLogRecord, expectedLogRecordIdx)
	if err != nil {
		return nil, err
	}

	decryptedLog, err := aesGCMDecrypt(encryptedLogRecord, decryptConfig.lek)
	if err != nil {
		return nil, fmt.Errorf("error decrypting line %v, %v, %v", lineNb, err, decryptConfig.lek)
	}

	decompressedLogRecord, err := decompress(decryptConfig.compressionMode, decryptedLog)
	if err != nil {
		return nil, fmt.Errorf("error decompressing line %v, %v", lineNb, err)
	}

	var bsonParsedLogRecord interface{}
	err = bson.Unmarshal(decompressedLogRecord, &bsonParsedLogRecord)
	if err != nil {
		return nil, fmt.Errorf("error parsing decrypted line %v, %v", lineNb, err)
	}

	return bsonParsedLogRecord, nil
}

func (logLine *AuditLogLine) logAdditionalAuthData() ([]byte, error) {
	const AADByteSize = 8

	timestampMs, err := logLine.UTCTimestampValue()
	if err != nil {
		return nil, err
	}

	additionalAuthData := make([]byte, AADByteSize)
	binary.LittleEndian.PutUint64(additionalAuthData, timestampMs)
	return additionalAuthData, nil
}

func validateLogLine(encryptedLogRecord *EncryptedLogRecord, expectedLogRecordIdx uint64) error {
	if expectedLogRecordIdx != encryptedLogRecord.KeyInvocationCount {
		return fmt.Errorf("logRecordIdx missmatch, expected: %v, actual: %v", encryptedLogRecord.KeyInvocationCount, expectedLogRecordIdx)
	}

	return nil
}
