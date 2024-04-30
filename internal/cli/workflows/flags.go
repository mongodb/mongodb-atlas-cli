// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workflows

import (
	"strconv"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

var GlobalFlagsToArgs = map[string]string{
	flag.Profile:      "1",
	flag.ProfileShort: "1",
	flag.Debug:        "0",
	flag.DebugShort:   "0",
}

func shouldRemoveFlagAndArgs(flags map[string]string, arg string) (int, error) {
	trimmedArg := strings.TrimLeft(arg, "-")

	flagArgsCount := flags[trimmedArg]
	if flagArgsCount == "" {
		flagArgsCount = GlobalFlagsToArgs[trimmedArg]
	}

	if flagArgsCount != "" {
		flagArgsCountValue, err := strconv.Atoi(flagArgsCount)
		if err != nil {
			return -1, err
		}
		return flagArgsCountValue, nil
	}

	return -1, nil
}

func RemoveFlagsAndArgs(flagsToBeRemoved map[string]string, argsToBeRemoved map[string]bool, args []string) ([]string, error) {
	var newArgs []string
	for i := 0; i < len(args); i++ {
		arg := args[i]

		// remove global flags
		flagArgsCount, err := shouldRemoveFlagAndArgs(flagsToBeRemoved, arg)
		if err != nil {
			_, _ = log.Debugf("Error while removing flags and arg(s): %s\n", err)
			return nil, err
		}

		if flagArgsCount >= 0 {
			_, _ = log.Debugf("Skipped flag %s and %d arg(s)\n", args[i], flagArgsCount)
			i += flagArgsCount // skip the args after flag
			continue
		}

		// remove args
		if argsToBeRemoved[arg] {
			_, _ = log.Debugf("Skipped arg %s\n", args[i])
			continue
		}

		newArgs = append(newArgs, arg)
	}

	_, _ = log.Debugf("Removing 3 first terms in %s \n", newArgs)
	return newArgs[3:], nil
}
