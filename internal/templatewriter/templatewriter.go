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

package templatewriter

import (
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
	"text/template"
)

const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
)

var funcMap = template.FuncMap{
	"valueOrEmptySlice": valueOrEmptySlice,
	"formatAliases":     formatAliases,
}

func valueOrEmptySlice(slice any) (result any) {
	if slice == nil {
		return result
	}

	k := reflect.TypeOf(slice).Kind()
	if (k == reflect.Slice || k == reflect.Ptr) && reflect.ValueOf(slice).IsNil() {
		return result
	}

	return slice
}

// formatAliases formats command aliases for display.
// Returns a formatted string like " [aliases: cmd1, c2]" if aliases exist,
// or an empty string if no aliases are present.
func formatAliases(aliases []string) string {
	if len(aliases) == 0 {
		return ""
	}
	return " [aliases: " + strings.Join(aliases, ", ") + "]"
}

// newTabWriter returns a tabwriter that handles tabs(`\t) to space columns evenly.
func newTabWriter(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, 0)
}

// Print outputs v to os.Stdout while handling configured formats,
// if the optional t is given then it's processed as a go-template,
// this template will be handled with a tabwriter so you can use tabs (\t)
// and new lines (\n) to space your content evenly.
func Print(writer io.Writer, t string, v any) error {
	tmpl, err := template.New("output").Funcs(funcMap).Parse(t)
	if err != nil {
		return err
	}
	w := newTabWriter(writer)

	if err := tmpl.Execute(w, v); err != nil {
		return err
	}
	return w.Flush()
}
