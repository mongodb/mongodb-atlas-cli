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
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type AuditLogOutput interface {
	Warningf(lineNb int, logLine *AuditLogLine, format string, a ...any) error
	Error(lineNb int, logLine *AuditLogLine, err error) error
	Errorf(lineNb int, logLine *AuditLogLine, format string, a ...any) error
	LogRecord(lineNb int, logRecord any) error
}

type auditLogOutputWriter struct {
	out io.Writer
}

type AuditLogErrorLevel string

const (
	AuditLogErrorLevelError   AuditLogErrorLevel = "decryptionLogError"
	AuditLogErrorLevelWarning AuditLogErrorLevel = "decryptionLogWarning"
)

type AuditLogError struct {
	Level AuditLogErrorLevel
	Line  int
	TS    *time.Time
	Err   error
}

func (e AuditLogError) MarshalBSON() ([]byte, error) {
	doc := bson.D{
		bson.E{Key: "atype", Value: string(e.Level)},
		bson.E{Key: "line", Value: e.Line},
		bson.E{Key: "error", Value: e.Err.Error()},
	}

	if e.TS != nil {
		doc = append(doc, bson.E{Key: "ts", Value: e.TS.UnixMilli()})
	}

	return bson.Marshal(doc)
}

func (e AuditLogError) Error() string {
	return e.Err.Error()
}

func NewAuditLogOutput(out io.Writer) AuditLogOutput {
	return &auditLogOutputWriter{
		out: out,
	}
}

func (l *auditLogOutputWriter) writeRecord(value any) error {
	jsonVal, err := bson.MarshalExtJSON(value, false, false)
	if err != nil {
		return err
	}

	_, err = l.out.Write(jsonVal)
	if err != nil {
		return err
	}
	_, err = l.out.Write([]byte{'\n'})
	return err
}

func (l *auditLogOutputWriter) Warningf(lineNb int, logLine *AuditLogLine, format string, a ...any) error {
	e := AuditLogError{
		Level: AuditLogErrorLevelWarning,
		Line:  lineNb,
		Err:   fmt.Errorf(format, a...),
	}
	if logLine != nil {
		e.TS = logLine.TS
	}

	return l.writeRecord(e)
}

func (l *auditLogOutputWriter) Error(lineNb int, logLine *AuditLogLine, err error) error {
	e := AuditLogError{
		Level: AuditLogErrorLevelError,
		Line:  lineNb,
		Err:   err,
	}
	if logLine != nil {
		e.TS = logLine.TS
	}

	return l.writeRecord(e)
}

func (l *auditLogOutputWriter) Errorf(lineNb int, logLine *AuditLogLine, format string, a ...any) error {
	return l.Error(lineNb, logLine, fmt.Errorf(format, a...))
}

func (l *auditLogOutputWriter) LogRecord(_ int, logRecord any) error {
	return l.writeRecord(logRecord)
}
