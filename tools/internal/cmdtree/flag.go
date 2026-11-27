// Copyright 2026 MongoDB Inc
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

package cmdtree

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag is one command line flag a command declares.
type Flag struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required,omitempty"`
	Default  string `json:"default,omitempty"`
}

// flags returns the flags cmd declares, in the lexicographical order that
// FlagSet.VisitAll guarantees. Flags inherited from a parent, like the root's
// --profile, are left out: they belong to every command and say nothing about
// this one.
func flags(cmd *cobra.Command) []Flag {
	var flags []Flag
	cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		flags = append(flags, Flag{
			Name:     flag.Name,
			Type:     flag.Value.Type(),
			Required: isRequired(flag),
			Default:  flag.DefValue,
		})
	})
	return flags
}

// isRequired reports whether cobra was told the command needs this flag.
func isRequired(flag *pflag.Flag) bool {
	_, required := flag.Annotations[cobra.BashCompOneRequiredFlag]
	return required
}
