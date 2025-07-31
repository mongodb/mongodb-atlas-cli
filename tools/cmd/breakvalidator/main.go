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
	"os"

	"github.com/spf13/cobra"
)

type flagData struct {
	Type    string `json:"type,omitempty"`
	Default string `json:"default,omitempty"`
	Short   string `json:"short,omitempty"`
}

type cmdData struct {
	Aliases []string            `json:"aliases,omitempty"`
	Flags   map[string]flagData `json:"flags,omitempty"`
}

func buildRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "breakvalidator",
		Short: "CLI tool to validate breaking changes in the CLI.",
	}
	rootCmd.AddCommand(buildGenerateCmd())
	rootCmd.AddCommand(buildValidateCmd())

	return rootCmd
}

func main() {
	rootCmd := buildRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
