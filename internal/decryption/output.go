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

	"go.mongodb.org/mongo-driver/bson"
)

type AuditLogOutput interface {
	Warningf(lineNb int, format string, a ...interface{}) error
	Error(lineNb int, err error) error
	Errorf(lineNb int, format string, a ...interface{}) error
	LogRecord(lineNb int, logRecord interface{}) error
}

type auditLogOutputImpl struct {
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
	Err   error
}

func (e AuditLogError) MarshalBSON() ([]byte, error) {
	doc := bson.D{
		bson.E{Key: "atype", Value: string(e.Level)},
		bson.E{Key: "line", Value: e.Line},
		bson.E{Key: "error", Value: e.Err.Error()},
	}
	return bson.Marshal(doc)
}

func (e AuditLogError) Error() string {
	return e.Err.Error()
}

func NewAuditLogOutput(out io.Writer) AuditLogOutput {
	return &auditLogOutputImpl{
		out: out,
	}
}

func (l *auditLogOutputImpl) writeRecord(value interface{}) error {
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

func (l *auditLogOutputImpl) Warningf(lineNb int, format string, a ...interface{}) error {
	e := AuditLogError{
		Level: AuditLogErrorLevelWarning,
		Line:  lineNb,
		Err:   fmt.Errorf(format, a...),
	}

	return l.writeRecord(e)
}

func (l *auditLogOutputImpl) Error(lineNb int, err error) error {
	e := AuditLogError{
		Level: AuditLogErrorLevelError,
		Line:  lineNb,
		Err:   err,
	}

	return l.writeRecord(e)
}

func (l *auditLogOutputImpl) Errorf(lineNb int, format string, a ...interface{}) error {
	return l.Error(lineNb, fmt.Errorf(format, a...))
}

func (l *auditLogOutputImpl) LogRecord(_ int, logRecord interface{}) error {
	return l.writeRecord(logRecord)
}
