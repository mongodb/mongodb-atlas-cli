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
)

type DecryptSection struct {
	lek                    []byte
	compressionMode        CompressionMode
	lastKeyInvocationCount uint64
}

func (s *DecryptSection) zeroLEK() {
	if s == nil {
		return
	}
	for i := range s.lek {
		s.lek[i] = 0
	}
}

// Decrypt decrypts the content of an audit log file using the metadata found in the file,
// the credentials provided by the user and the AES-GCM algorithm.
// The decrypted audit log records are saved in the out stream.
func Decrypt(logReader io.ReadSeeker, out io.Writer, opts *KeyProviderOpts) error {
	_, logLineScanner, err := readAuditLogFile(logReader)
	if err != nil {
		return err
	}

	output := NewAuditLogOutput(out)
	var decryptSection *DecryptSection

	idx := 0
	for ; logLineScanner.Scan(); idx++ {
		lineNb := idx + 1
		logLine, err := logLineScanner.AuditLogLine()
		if err != nil {
			if outputErr := output.Errorf(lineNb, "error parsing line %d, %v", lineNb, err); outputErr != nil {
				return outputErr
			}
			continue
		}

		switch logLine.AuditRecordType {
		case AuditHeaderRecord:
			decryptSection.zeroLEK()
			if decryptSection, err = processHeader(logLine, opts); err != nil {
				if outputErr := output.Errorf(lineNb, `error processing header line %d: %s`, lineNb, err); outputErr != nil {
					return outputErr
				}
			}
		case AuditLogRecord:
			if err := decryptAuditLogRecord(decryptSection, logLine, output, lineNb); err != nil {
				return err
			}
		default:
			if outputErr := output.Errorf(lineNb, `line %d skipped, unknown auditRecordType="%s"`, lineNb, logLine.AuditRecordType); outputErr != nil {
				return outputErr
			}
		}
	}
	decryptSection.zeroLEK()
	if err := logLineScanner.Err(); err != nil {
		lineNb := idx + 1
		if outputErr := output.Errorf(lineNb, "error parsing line %d, %v", lineNb, err); outputErr != nil {
			return outputErr
		}
	}

	return nil
}

func decryptAuditLogRecord(decryptSection *DecryptSection, logLine *AuditLogLine, output AuditLogOutput, lineNb int) error {
	if decryptSection == nil {
		return output.Warningf(lineNb, `line %d skipped, the header record for current section is missing or corrupted`, lineNb)
	}

	decryptedLogRecord, keyInvocationCount, err := processLogRecord(decryptSection, logLine, lineNb)
	if err != nil {
		return output.Error(lineNb, err)
	}

	err = validateLogRecord(decryptSection, keyInvocationCount)
	decryptSection.lastKeyInvocationCount = keyInvocationCount
	if err != nil {
		if outputErr := output.Error(lineNb, err); outputErr != nil {
			return outputErr
		}
	}

	return output.LogRecord(lineNb, decryptedLogRecord)
}
