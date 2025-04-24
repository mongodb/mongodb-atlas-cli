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

//go:build unit

package api

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestCobraFlagsToRequestParameters(t *testing.T) {
	cmd := func(args string) *cobra.Command {
		t.Helper()

		cmd := new(cobra.Command)
		cmd.Flags().String("foo", "foo-default-value", "foo-usage")
		cmd.Flags().Int("bar", 42, "bar-usage")
		cmd.Flags().Bool("baz", true, "baz-usage")
		cmd.Flags().StringSlice("qux", []string{"qux", "quz"}, "qux-usage")
		require.NoError(t, cmd.ParseFlags(strings.Split(args, " ")))
		return cmd
	}

	tests := []struct {
		name               string
		command            *cobra.Command
		expectedParameters map[string][]string
	}{
		{
			name:               "empty string",
			command:            cmd(""),
			expectedParameters: map[string][]string{},
		},
		{
			name:    "--foo fooo",
			command: cmd("--foo fooo"),
			expectedParameters: map[string][]string{
				"foo": {"fooo"},
			},
		},
		{
			name:    "basic type",
			command: cmd("--foo fooo --bar 2 --baz=false"),
			expectedParameters: map[string][]string{
				"foo": {"fooo"},
				"bar": {"2"},
				"baz": {"false"},
			},
		},
		{
			name:    "slice 1",
			command: cmd("--qux foo,bar"),
			expectedParameters: map[string][]string{
				"qux": {"foo", "bar"},
			},
		},
		{
			name:    "slice 2",
			command: cmd("--qux foo --qux bar"),
			expectedParameters: map[string][]string{
				"qux": {"foo", "bar"},
			},
		},
		{
			name:    "all",
			command: cmd("--foo f --qux foo --bar 99 --baz --qux bar"),
			expectedParameters: map[string][]string{
				"qux": {"foo", "bar"},
				"foo": {"f"},
				"bar": {"99"},
				"baz": {"true"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameters := cobraFlagsToRequestParameters(tt.command)
			require.Equal(t, tt.expectedParameters, parameters)
		})
	}
}
