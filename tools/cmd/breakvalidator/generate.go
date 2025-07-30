// Copyright 2025 MongoDB Inc
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

package main

import (
	"encoding/json"
	"io"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func generateCmd(cmd *cobra.Command) cmdData {
	data := cmdData{}
	if len(cmd.Aliases) > 0 {
		data.Aliases = cmd.Aliases
	}
	flags := false
	data.Flags = map[string]flagData{}
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		data.Flags[f.Name] = flagData{
			Type:    f.Value.Type(),
			Default: f.DefValue,
			Short:   f.Shorthand,
		}
		flags = true
	})
	if !flags {
		data.Flags = nil
	}
	return data
}

func generateCmds(cmd *cobra.Command) map[string]cmdData {
	data := map[string]cmdData{}
	data[cmd.CommandPath()] = generateCmd(cmd)
	for _, c := range cmd.Commands() {
		for k, v := range generateCmds(c) {
			data[k] = v
		}
	}
	return data
}

func generateCmdRun(output io.Writer) error {
	cliCmd := root.Builder()
	data := generateCmds(cliCmd)
	return json.NewEncoder(output).Encode(data)
}

func buildGenerateCmd() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate the CLI command structure.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return generateCmdRun(cmd.OutOrStdout())
		},
	}
	return generateCmd
}
