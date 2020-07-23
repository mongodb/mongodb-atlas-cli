// Copyright 2020 MongoDB Inc
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

package output

import (
	"os"
	"text/template"

	"github.com/mongodb/mongocli/internal/json"
)

type Config interface {
	Output() string
}

const jsonFormat = "json"

// Print outputs v to os.Stdout while handling configured formats,
// if the optional t is given then it's processed as a go-template
func Print(c Config, t string, v interface{}) error {
	if c.Output() == jsonFormat {
		return json.PrettyPrint(v)
	}
	if t != "" {
		tmpl, err := template.New("output").Parse(t)
		if err != nil {
			return err
		}
		return tmpl.Execute(os.Stdout, v)
	}
	return nil
}
