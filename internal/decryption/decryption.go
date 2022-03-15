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
	"bufio"
	"fmt"
	"io"

	"go.mongodb.org/mongo-driver/bson"
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
func Decrypt(logReader io.ReadSeeker, out io.Writer, opts KeyProviderOpts) error {
	_, logLines, err := readAuditLogFile(logReader)
	if err != nil {
		return err
	}

	output := buildOutput(out)
	var decryptSection *DecryptSection

	for idx, logLine := range logLines {
		lineNb := idx + 1
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

func peekFirstByte(reader io.ReadSeeker) (byte, error) {
	b := make([]byte, 1)

	n, err := reader.Read(b)
	if err != nil {
		return 0, err
	}

	if n != 1 {
		return 0, fmt.Errorf("no bytes to read")
	}

	c, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	if c != 0 {
		return 0, fmt.Errorf("impossible to seek bytes")
	}
	return b[0], nil
}

func readAuditLogFile(reader io.ReadSeeker) (AuditLogFormat, []*AuditLogLine, error) {
	auditLogFormat := BSON

	b, err := peekFirstByte(reader)
	if err != nil {
		return auditLogFormat, nil, err
	}

	if b == '{' {
		auditLogFormat = JSON
	}

	var logLines []*AuditLogLine

	switch auditLogFormat {
	case BSON:
		logLines, err = readAuditLogFileBSON(reader)
	case JSON:
		logLines, err = readAuditLogFileJSON(reader)
	}
	return auditLogFormat, logLines, err
}

func readAuditLogFileBSON(reader io.ReadSeeker) ([]*AuditLogLine, error) {
	var logLines []*AuditLogLine
	for {
		raw, err := bson.NewFromIOReader(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var logLine AuditLogLine
		err = bson.Unmarshal(raw, &logLine)
		if err != nil {
			return nil, err
		}
		logLines = append(logLines, &logLine)
	}
	return logLines, nil
}

func readAuditLogFileJSON(reader io.ReadSeeker) ([]*AuditLogLine, error) {
	var logLines []*AuditLogLine
	s := bufio.NewScanner(reader)
	for s.Scan() {
		var logLine AuditLogLine
		err := bson.UnmarshalExtJSON(s.Bytes(), true, &logLine)
		if err != nil {
			return nil, err
		}
		logLines = append(logLines, &logLine)
	}
	err := s.Err()
	if err != nil {
		return nil, err
	}
	return logLines, nil
}
