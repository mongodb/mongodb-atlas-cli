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
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestGenerateCmds(t *testing.T) {
	cliCmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"testa"},
	}
	cliCmd.Flags().StringP("flag1", "f", "default1", "flag1")
	generatedData := generateCmds(cliCmd)

	expectedData := map[string]cmdData{
		"test": {
			Aliases: []string{"testa"},
			Flags: map[string]flagData{
				"flag1": {
					Type:    "string",
					Default: "default1",
					Short:   "f",
				},
			},
		},
	}

	if !reflect.DeepEqual(generatedData, expectedData) {
		t.Fatalf("got: %v, expected: %v", generatedData, expectedData)
	}
}
