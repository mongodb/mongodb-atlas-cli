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
	"os"
	"strings"
)

type DecryptConfig struct {
	lek             []byte
	compressionMode CompressionMode
}

func Decrypt(filepath, outputFilepath string) error {
	auditLogFormat, logLines, err := readAuditLogFile(filepath)
	if err != nil {
		return err
	}

	auditLogEncoding := newAuditLogEncoding(auditLogFormat)
	output := buildOutput(outputFilepath)
	var decryptConfig *DecryptConfig
	var logRecordIdx uint64

	for idx, line := range logLines {
		lineNb := idx + 1
		if len(line) > 0 {
			logLine, err := auditLogEncoding.Parse(line)
			if err != nil {
				output.Errorf(lineNb, "error parsing line %d, %s", lineNb, err)
				continue
			}

			switch logLine.AuditRecordType {
			case AuditHeaderRecord:
				decryptConfig, err = processHeader(logLine)
				if err != nil {
					output.Errorf(lineNb, `error processing header line %d: %s`, lineNb, err)
				}
				logRecordIdx = 0
			case AuditLogRecord:
				logRecordIdx++
				if decryptConfig == nil {
					output.Warningf(lineNb, `line %d skipped, the header record for current section is missing or corrupted`, lineNb)
				} else {
					decryptedLogRecord, err := processLogRecord(decryptConfig, logLine, lineNb, logRecordIdx)
					if err != nil {
						output.Error(lineNb, err)
					} else {
						output.LogRecord(lineNb, decryptedLogRecord)
					}
				}
			default:
				output.Errorf(lineNb, `line %d skipped, unknown auditRecordType="%s"`, lineNb, logLine.AuditRecordType)
			}
		}
	}

	return nil
}

func readAuditLogFile(filepath string) (AuditLogFormat, []string, error) {
	const LineBreak = "\n"
	auditLogFormat := BSON

	data, err := os.ReadFile(filepath)
	if err != nil {
		return auditLogFormat, nil, err
	}

	const jsonStartChar = '{'
	if data[0] == jsonStartChar {
		auditLogFormat = JSON
	}

	return auditLogFormat, strings.Split(string(data), LineBreak), nil
}
