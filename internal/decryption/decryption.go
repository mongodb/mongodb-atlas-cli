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
	"io"
	"strings"

	"github.com/mongodb/mongocli/internal/decryption/keyproviders"
)

type DecryptConfig struct {
	lek             []byte
	compressionMode CompressionMode
}

// Decrypt decrypts the content of an audit log file using the metadata found in the file,
// the credentials provided by the user and the AES-GCM algorithm.
// The decrypted audit log records are saved in the out stream.
func Decrypt(logReader io.Reader, out io.Writer, credentialsProvider keyproviders.CredentialsProvider) error {
	auditLogFormat, logLines, err := readAuditLogFile(logReader)
	if err != nil {
		return err
	}

	auditLogEncoding := newAuditLogEncoding(auditLogFormat)
	output := buildOutput(out)
	var decryptConfig *DecryptConfig
	var logRecordIdx uint64

	for idx, line := range logLines {
		lineNb := idx + 1
		if len(line) > 0 {
			logLine, err := auditLogEncoding.Parse(line)
			if err != nil {
				panicIfError(output.Errorf(lineNb, "error parsing line %d, %s", lineNb, err))
				continue
			}

			switch logLine.AuditRecordType {
			case AuditHeaderRecord:
				decryptConfig, err = processHeader(logLine, credentialsProvider)
				if err != nil {
					panicIfError(output.Errorf(lineNb, `error processing header line %d: %s`, lineNb, err))
				}
				logRecordIdx = 0
			case AuditLogRecord:
				logRecordIdx++
				if decryptConfig == nil {
					panicIfError(output.Warningf(lineNb, `line %d skipped, the header record for current section is missing or corrupted`, lineNb))
				} else {
					decryptedLogRecord, err := processLogRecord(decryptConfig, logLine, lineNb, logRecordIdx)
					if err != nil {
						panicIfError(output.Error(lineNb, err))
					} else {
						panicIfError(output.LogRecord(lineNb, decryptedLogRecord))
					}
				}
			default:
				panicIfError(output.Errorf(lineNb, `line %d skipped, unknown auditRecordType="%s"`, lineNb, logLine.AuditRecordType))
			}
		}
	}

	return nil
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func readAuditLogFile(logReader io.Reader) (AuditLogFormat, []string, error) {
	const LineBreak = "\n"
	auditLogFormat := BSON

	data, err := io.ReadAll(logReader)
	if err != nil {
		return auditLogFormat, nil, err
	}

	const jsonStartChar = '{'
	if data[0] == jsonStartChar {
		auditLogFormat = JSON
	}

	return auditLogFormat, strings.Split(string(data), LineBreak), nil
}
