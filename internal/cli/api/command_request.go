// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"io"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewCommandRequestFromCobraCommand(cobraCommand *cobra.Command, apiCommand api.Command, content io.Reader, version string) (*api.CommandRequest, error) {
	return &api.CommandRequest{
		Command:    apiCommand,
		Content:    content, // content has to be set by caller
		Format:     "json",  // part of CLOUDP-280747
		Parameters: cobraFlagsToRequestParameters(cobraCommand),
		Version:    version,
	}, nil
}

func cobraFlagsToRequestParameters(cobraCommand *cobra.Command) map[string][]string {
	parameters := make(map[string][]string)
	var flagsToIgnore = map[string]struct{}{
		"file":    {},
		"version": {},
	}

	cobraCommand.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		if _, ignoreFlag := flagsToIgnore[flag.Name]; ignoreFlag {
			return
		}

		if values, ok := flag.Value.(pflag.SliceValue); ok {
			parameters[flag.Name] = values.GetSlice()
		} else {
			parameters[flag.Name] = []string{flag.Value.String()}
		}
	})

	return parameters
}
