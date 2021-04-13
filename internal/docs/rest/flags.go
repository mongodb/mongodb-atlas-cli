// Copyright 2021 MongoDB Inc
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

package rest

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	stringType = "string"
	boolType   = "bool"
	countType  = "count"
	falseValue = "false"
	nilValue   = "<nil>"
)

func FlagUsages(f *pflag.FlagSet) string {
	buf := new(bytes.Buffer)

	f.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		line := ""
		varname, usage := pflag.UnquoteUsage(flag)
		usage = strings.ReplaceAll(usage, "\n", "\n"+strings.Repeat(" ", 6))

		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			line = fmt.Sprintf("  * - -%s, --%s\n    - %s", flag.Shorthand, flag.Name, varname)
		} else {
			line = fmt.Sprintf("  * - --%s\n    - %s", flag.Name, varname)
		}

		required := false
		if len(flag.Annotations) != 0 {
			_, required = flag.Annotations[cobra.BashCompOneRequiredFlag]
		}
		line += fmt.Sprintf("\n    - %v", required)

		if flag.NoOptDefVal != "" {
			switch flag.Value.Type() {
			case stringType:
				line += fmt.Sprintf("[=\"%s\"]", flag.NoOptDefVal)
			case boolType:
				if flag.NoOptDefVal != "true" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			case countType:
				if flag.NoOptDefVal != "+1" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			default:
				line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
			}
		}

		line += "\n    - " + usage
		if !defaultIsZeroValue(flag) {
			if flag.Value.Type() == stringType {
				line += fmt.Sprintf(" (default %q)", flag.DefValue)
			} else {
				line += fmt.Sprintf(" (default %s)", flag.DefValue)
			}
		}
		if flag.Deprecated != "" {
			line += fmt.Sprintf(" (DEPRECATED: %s)", flag.Deprecated)
		}

		_, _ = fmt.Fprintln(buf, line)
	})

	return buf.String()
}

// defaultIsZeroValue returns true if the default value for this flag represents
// a zero value.
func defaultIsZeroValue(f *pflag.Flag) bool {
	switch f.Value.Type() {
	case boolType:
		return f.DefValue == falseValue
	case "duration":
		// Beginning in Go 1.7, duration zero values are "0s"
		return f.DefValue == "0" || f.DefValue == "0s"
	case "int", "int8", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "count", "float32", "float64":
		return f.DefValue == "0"
	case stringType:
		return f.DefValue == ""
	case "ip", "ipMask", "ipNet":
		return f.DefValue == nilValue
	case "intSlice", "stringSlice", "stringArray":
		return f.DefValue == "[]"
	default:
		switch f.Value.String() {
		case falseValue:
			return true
		case nilValue:
			return true
		case "":
			return true
		case "0":
			return true
		}
		return false
	}
}
