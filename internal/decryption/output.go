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
	"encoding/json"
	"fmt"
)

type AuditLogOutput struct {
	filepath  string
	Warningf  func(lineNb int, format string, a ...interface{})
	Error     func(lineNb int, err error)
	Errorf    func(lineNb int, format string, a ...interface{})
	LogRecord func(lineNb int, logRecord interface{})
}

func buildOutput(filepath string) AuditLogOutput {
	return AuditLogOutput{
		filepath: filepath,
		Warningf: func(lineNb int, format string, a ...interface{}) {
			fmt.Printf(format+"\n", a)
		},
		Error: func(lineNb int, err error) {
			fmt.Printf("%s\n", err)
		},
		Errorf: func(lineNb int, format string, a ...interface{}) {
			fmt.Printf(format+"\n", a)
		},
		LogRecord: func(lineNb int, logRecord interface{}) {
			jsonVal, _ := json.Marshal(logRecord)
			fmt.Printf("%s\n", jsonVal)
		},
	}
}
