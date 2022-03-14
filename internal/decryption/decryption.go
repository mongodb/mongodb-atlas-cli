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
)

type DecryptSection struct {
	lek               []byte
	compressionMode   CompressionMode
	processedLogLines uint64
}

type KeyProviderOpts struct {
	LocalKeyFileName              string
	KMIPServerCAFileName          string
	KMIPClientCertificateFileName string
}

// Decrypt decrypts the content of an audit log file using the metadata found in the file,
// the credentials provided by the user and the AES-GCM algorithm.
// The decrypted audit log records are saved in the out stream.
func Decrypt(logReader io.Reader, out io.Writer, opts KeyProviderOpts) error {
	auditLogFormat, logLines, err := readAuditLogFile(logReader)
	if err != nil {
		return err
	}

	auditLogEncoding := newAuditLogEncoding(auditLogFormat)
	output := buildOutput(out)
	var decryptSection *DecryptSection

	for idx, line := range logLines {
		lineNb := idx + 1
		logLine, err := auditLogEncoding.Parse(line)
		if err != nil {
			if outputErr := output.Errorf(lineNb, "error parsing line %d, %v", lineNb, err); outputErr != nil {
				return outputErr
			}
			continue
		}

		switch logLine.AuditRecordType {
		case AuditHeaderRecord:
			if decryptSection, err = processHeader(logLine, opts); err != nil {
				if outputErr := output.Errorf(lineNb, `error processing header line %d: %s`, lineNb, err); outputErr != nil {
					return outputErr
				}
			}
		case AuditLogRecord:
			if decryptSection == nil {
				if outputErr := output.Warningf(lineNb, `line %d skipped, the header record for current section is missing or corrupted`, lineNb); outputErr != nil {
					return outputErr
				}
				continue
			}

			decryptedLogRecord, err := processLogRecord(decryptSection, logLine, lineNb)
			decryptSection.processedLogLines++
			if err != nil {
				if outputErr := output.Error(lineNb, err); outputErr != nil {
					return outputErr
				}
			} else {
				if outputErr := output.LogRecord(lineNb, decryptedLogRecord); outputErr != nil {
					return outputErr
				}
			}
		default:
			if outputErr := output.Errorf(lineNb, `line %d skipped, unknown auditRecordType="%s"`, lineNb, logLine.AuditRecordType); outputErr != nil {
				return outputErr
			}
		}
	}

	return nil
}

func readAuditLogFile(logReader io.Reader) (AuditLogFormat, []string, error) {
	const LineBreak = "\n"
	auditLogFormat := BSON

	data, err := io.ReadAll(logReader)
	if err != nil {
		return auditLogFormat, nil, err
	}

	const jsonStartChar = '{'
	if len(data) > 0 && data[0] == jsonStartChar {
		auditLogFormat = JSON
	}

	return auditLogFormat, strings.Split(string(data), LineBreak), nil
}
