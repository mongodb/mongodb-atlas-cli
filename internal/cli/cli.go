// Copyright 2024 MongoDB Inc
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

package cli

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tangzero/inflector"
)

type AutoFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

// GenerateAliases return aliases for use such that they are:
// a version all lower case, a version with dashes, a singular versions with the same rules.
func GenerateAliases(use string, extra ...string) []string {
	aliases := make([]string, 0)

	if lower := strings.ToLower(use); lower != use {
		aliases = append(aliases, lower)
	}
	if dash := inflector.Dasherize(use); dash != use {
		aliases = append(aliases, dash)
	}
	if singular := inflector.Singularize(use); singular != use {
		aliases = append(aliases, singular)
		if lower := strings.ToLower(singular); lower != singular {
			aliases = append(aliases, lower)
		}
		if dash := inflector.Dasherize(singular); dash != singular {
			aliases = append(aliases, dash)
		}
	}
	aliases = append(aliases, extra...)
	return aliases
}

// ReturnValueForSetting returns a boolean value that is useful when working with boolean flags to inform
// whether the given option should be active or inactive.
func ReturnValueForSetting(enableFlag, disableFlag bool) *bool {
	var valueToSet bool
	if enableFlag && disableFlag {
		return nil
	}
	if enableFlag {
		valueToSet = true
		return &valueToSet
	}
	if disableFlag {
		valueToSet = false
		return &valueToSet
	}
	return nil
}
