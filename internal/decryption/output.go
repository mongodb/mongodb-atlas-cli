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

type AuditLogOutput struct {
	Warningf  func(lineNb int, format string, a ...interface{}) error
	Error     func(lineNb int, err error) error
	Errorf    func(lineNb int, format string, a ...interface{}) error
	LogRecord func(lineNb int, logRecord interface{}) error
}

func buildOutput(out io.Writer) AuditLogOutput {
	newLine := []byte{'\n'}
	writeLine := func(value []byte) error {
		if _, err := out.Write(value); err != nil {
			return err
		}
		_, err := out.Write(newLine)
		return err
	}

	return AuditLogOutput{
		Warningf: func(lineNb int, format string, a ...interface{}) error {
			return writeLine([]byte(fmt.Sprintf(format, a)))
		},
		Error: func(lineNb int, err error) error {
			return writeLine([]byte(err.Error()))
		},
		Errorf: func(lineNb int, format string, a ...interface{}) error {
			return writeLine([]byte(fmt.Sprintf(format, a...)))
		},
		LogRecord: func(lineNb int, logRecord interface{}) error {
			jsonVal, err := bson.MarshalExtJSON(logRecord, false, false)
			if err != nil {
				return err
			}
			return writeLine(jsonVal)
		},
	}
}
