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
	"text/tabwriter"
	"text/template"

	"github.com/mongodb/mongocli/internal/json"
)

type Config interface {
	Output() string
}

const jsonFormat = "json"

// Print outputs v to os.Stdout while handling configured formats,
// if the optional t is given then it's processed as a go-template,
// this template will be handled with a tabwriter so you can use tabs (\t)
// and new lines (\n) to space your content evenly.
func Print(c Config, t string, v interface{}) error {
	if c.Output() == jsonFormat {
		return json.PrettyPrint(v)
	}
	if t != "" {
		tmpl, err := template.New("output").Parse(t)
		if err != nil {
			return err
		}
		// tabwriter will handle tabs(`\t) to space columns evenly, each column will use a tab(\t) of 8 spaces
		// with a minimum padding of 2 characters per column so columns don't touch each other if they are too wide
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)

		if err := tmpl.Execute(w, v); err != nil {
			return err
		}
		return w.Flush()
	}
	return nil
}
