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

	"github.com/mongodb/mongocli/internal/decryption/aes"
	"go.mongodb.org/mongo-driver/bson"
)

func (logLine *AuditLogLine) decodeLogRecord() (*DecodedLogRecord, error) {
	if logLine.Log == nil {
		return nil, fmt.Errorf("missing log")
	}

	log, err := base64.StdEncoding.DecodeString(*logLine.Log)
	if err != nil {
		return nil, err
	}

	TagData := log[0:12]
	KeyInitCountData := log[12:16]
	KeyInvocationCountData := log[16:24]
	LogData := log[24:]

	return &DecodedLogRecord{
		CipherText:         LogData,
		Tag:                TagData,
		AAD:                logLine.logAdditionalAuthData(),
		IV:                 append(KeyInitCountData, KeyInvocationCountData...),
		KeyInitCount:       binary.LittleEndian.Uint32(KeyInitCountData),
		KeyInvocationCount: binary.LittleEndian.Uint64(KeyInvocationCountData),
	}, nil
}

func processLogRecord(decryptConfig *DecryptSection, logLine *AuditLogLine, lineNb int) (bsonData interface{}, keyInvocationCount uint64, err error) {
	encryptedLogRecord, decodeErr := logLine.decodeLogRecord()
	if decodeErr != nil {
		return nil, 0, fmt.Errorf("line %v is corrupted, %v", lineNb, decodeErr)
	}

	gcm := &aes.GCMInput{
		Key: decryptConfig.lek,
		AAD: encryptedLogRecord.AAD,
		IV:  encryptedLogRecord.IV,
		Tag: encryptedLogRecord.Tag,
	}
	decryptedLog, decryptErr := gcm.Decrypt(encryptedLogRecord.CipherText)

	if decryptErr != nil {
		return nil, 0, fmt.Errorf("error decrypting line %v, %v, %v", lineNb, decryptErr, decryptConfig.lek)
	}

	decompressedLogRecord, decompressErr := decompress(decryptConfig.compressionMode, decryptedLog)
	if decompressErr != nil {
		return nil, 0, fmt.Errorf("error decompressing line %v, %v", lineNb, decompressErr)
	}

	var bsonParsedLogRecord map[string]interface{}
	if bsonErr := bson.Unmarshal(decompressedLogRecord, &bsonParsedLogRecord); bsonErr != nil {
		return nil, 0, fmt.Errorf("error parsing decrypted line %v, %v", lineNb, bsonErr)
	}

	if _, ok := bsonParsedLogRecord["ts"]; !ok {
		bsonParsedLogRecord["ts"] = logLine.TS
	}

	return bsonParsedLogRecord, encryptedLogRecord.KeyInvocationCount, nil
}

func (logLine *AuditLogLine) logAdditionalAuthData() []byte {
	const AADByteSize = 8

	additionalAuthData := make([]byte, AADByteSize)
	binary.LittleEndian.PutUint64(additionalAuthData, uint64(logLine.TS.UnixMilli()))
	return additionalAuthData
}

func validateLogRecord(decryptSection *DecryptSection, keyInvocationCount uint64) error {
	expected := decryptSection.lastKeyInvocationCount + 1
	if expected != keyInvocationCount {
		return fmt.Errorf("logRecordIdx missmatch, expected: %v, actual: %v", expected, keyInvocationCount)
	}
	return nil
}
