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

type Decryption struct {
	opts KeyProviderOpts
}

type Option func(d *Decryption)

func NewDecryption(options ...Option) *Decryption {
	d := &Decryption{}
	for _, opt := range options {
		opt(d)
	}
	return d
}

func WithAWSOpts(accessKey, secretAccessKey, sessionToken string) func(d *Decryption) {
	return func(d *Decryption) {
		d.opts.AWS = &KeyProviderAWSOpts{
			AccessKey:       accessKey,
			SecretAccessKey: secretAccessKey,
			SessionToken:    sessionToken,
		}
	}
}

func WithGCPOpts(serviceAccountKey string) func(d *Decryption) {
	return func(d *Decryption) {
		d.opts.GCP = &KeyProviderGCPOpts{
			ServiceAccountKey: serviceAccountKey,
		}
	}
}

func WithAzureOpts(tenantID, clientID, secret string) func(d *Decryption) {
	return func(d *Decryption) {
		d.opts.Azure = &KeyProviderAzureOpts{
			TenantID: tenantID,
			ClientID: clientID,
			Secret:   secret,
		}
	}
}

// Decrypt decrypts the content of an audit log file using the metadata found in the file,
// the credentials provided by the user and the AES-GCM algorithm.
// The decrypted audit log records are saved in the out stream.
func (d *Decryption) Decrypt(logReader io.ReadSeeker, out io.Writer) error {
	logLineScanner, err := readAuditLogFile(logReader)
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
			if outputErr := output.Errorf(lineNb, logLine, "error parsing line %d, %v", lineNb, err); outputErr != nil {
				return outputErr
			}
			if decryptSection != nil {
				// even if log record is corrupted, consider it was encrypted,
				// so the lastKeyInvocationCount should be incremented
				decryptSection.lastKeyInvocationCount++
			}
			continue
		}

		switch logLine.AuditRecordType {
		case AuditHeaderRecord:
			decryptSection.zeroLEK()
			if decryptSection, err = processHeader(logLine, d.opts); err != nil {
				if outputErr := output.Errorf(lineNb, logLine, `error processing header line %d: %s`, lineNb, err); outputErr != nil {
					return outputErr
				}
			}
		case AuditLogRecord:
			if err := decryptAuditLogRecord(decryptSection, logLine, output, lineNb); err != nil {
				return err
			}
		default:
			if outputErr := output.Errorf(lineNb, logLine, `line %d skipped, unknown auditRecordType="%s"`, lineNb, logLine.AuditRecordType); outputErr != nil {
				return outputErr
			}
		}
	}
	decryptSection.zeroLEK()
	if err := logLineScanner.Err(); err != nil {
		lineNb := idx + 1
		if outputErr := output.Errorf(lineNb, nil, "error parsing line %d, %v", lineNb, err); outputErr != nil {
			return outputErr
		}
	}

	return nil
}

func decryptAuditLogRecord(decryptSection *DecryptSection, logLine *AuditLogLine, output AuditLogOutput, lineNb int) error {
	if decryptSection == nil {
		return output.Warningf(lineNb, logLine, `line %d skipped, the header record for current section is missing or corrupted`, lineNb)
	}

	decryptedLogRecord, keyInvocationCount, err := processLogRecord(decryptSection, logLine, lineNb)
	if err != nil {
		// even if log record is corrupted, consider it was encrypted,
		// so the lastKeyInvocationCount should be incremented
		decryptSection.lastKeyInvocationCount++
		return output.Error(lineNb, logLine, err)
	}

	err = validateLogRecord(decryptSection, keyInvocationCount)
	decryptSection.lastKeyInvocationCount = keyInvocationCount
	if err != nil {
		if outputErr := output.Error(lineNb, logLine, err); outputErr != nil {
			return outputErr
		}
	}

	return output.LogRecord(lineNb, decryptedLogRecord)
}
