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
	"errors"
	"io"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrNoBytesToRead  = errors.New("no bytes to read")
	ErrSeekNotAllowed = errors.New("impossible to seek bytes")
)

type AuditLogFormat string

const (
	JSON AuditLogFormat = "JSON"
	BSON AuditLogFormat = "BSON"
)

func peekFirstByte(reader io.ReadSeeker) (byte, error) {
	b := make([]byte, 1)

	n, err := reader.Read(b)
	if err != nil {
		return 0, err
	}

	if n != 1 {
		return 0, ErrNoBytesToRead
	}

	c, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	if c != 0 {
		return 0, ErrSeekNotAllowed
	}
	return b[0], nil
}

func readAuditLogFile(reader io.ReadSeeker) (AuditLogFormat, auditLogScanner, error) {
	auditLogFormat := BSON

	b, err := peekFirstByte(reader)
	if err != nil {
		return auditLogFormat, nil, err
	}

	if b == '{' {
		auditLogFormat = JSON
	}

	var scanner auditLogScanner
	switch auditLogFormat {
	case BSON:
		scanner = newBSONScanner(reader)
	case JSON:
		scanner = newJSONScanner(reader)
	}
	return auditLogFormat, scanner, err
}

type auditLogScanner interface {
	Scan() bool
	Err() error
	Bytes() []byte
	AuditLogLine() (*AuditLogLine, error)
}

func newBSONScanner(r io.Reader) *bsonScanner {
	return &bsonScanner{r: r}
}

type bsonScanner struct {
	r   io.Reader
	buf []byte
	err error
}

func (s *bsonScanner) Scan() bool {
	raw, err := bson.NewFromIOReader(s.r)
	if err != nil {
		if err != io.EOF {
			s.err = err
		}
		return false
	}
	s.buf = raw
	return true
}

func (s *bsonScanner) Err() error {
	return s.err
}

func (s *bsonScanner) Bytes() []byte {
	return s.buf
}

func (s *bsonScanner) AuditLogLine() (*AuditLogLine, error) {
	var logLine AuditLogLine
	err := bson.Unmarshal(s.Bytes(), &logLine)
	return &logLine, err
}

func newJSONScanner(r io.Reader) *jsonScanner {
	return &jsonScanner{r: bufio.NewScanner(r)}
}

type jsonScanner struct {
	r *bufio.Scanner
}

func (s *jsonScanner) Scan() bool {
	return s.r.Scan()
}

func (s *jsonScanner) Err() error {
	return s.r.Err()
}

func (s *jsonScanner) Bytes() []byte {
	return s.r.Bytes()
}

func (s *jsonScanner) AuditLogLine() (*AuditLogLine, error) {
	var logLine AuditLogLine
	err := bson.UnmarshalExtJSON(s.Bytes(), true, &logLine)
	return &logLine, err
}
