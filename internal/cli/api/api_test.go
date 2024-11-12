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
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestSplitShortAndLongDescription(t *testing.T) {
	tests := []struct {
		name        string
		description string
		wantShort   string
		wantLong    string
	}{
		{
			name:        "empty string",
			description: "",
			wantShort:   "",
			wantLong:    "",
		},
		{
			name:        "single sentence",
			description: "This is a single sentence.",
			wantShort:   "This is a single sentence.",
			wantLong:    "",
		},
		{
			name:        "two sentences",
			description: "First sentence. Second sentence.",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence.",
		},
		{
			name:        "multiple sentences with spaces",
			description: "First sentence.   Second sentence.    Third sentence.",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence. Third sentence.",
		},
		{
			name:        "sentence without period",
			description: "This is a sentence without period",
			wantShort:   "This is a sentence without period.",
			wantLong:    "",
		},
		{
			name:        "multiple sentences with extra periods",
			description: "This is version 1.2.3. Second sentence. Third sentence.",
			wantShort:   "This is version 1.2.3.",
			wantLong:    "Second sentence. Third sentence.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShort, gotLong := splitShortAndLongDescription(tt.description)
			if gotShort != tt.wantShort {
				t.Errorf("splitShortAndLongDescription() gotShort = %v, want %v", gotShort, tt.wantShort)
			}
			if gotLong != tt.wantLong {
				t.Errorf("splitShortAndLongDescription() gotLong = %v, want %v", gotLong, tt.wantLong)
			}
		})
	}
}

func TestAddParameters(t *testing.T) {
	testCmd := &cobra.Command{}
	err := addParameters(testCmd, []api.Parameter{
		{Name: "test1", Description: "description1", Required: false, Type: api.ParameterType{IsArray: false, Type: api.TypeString}},
		{Name: "test2", Description: "description2", Required: true, Type: api.ParameterType{IsArray: true, Type: api.TypeBool}},
	})

	require.NoError(t, err)

	test1 := testCmd.Flag("test1")
	require.NotNil(t, test1)
	require.Equal(t, "test1", test1.Name)
	require.Equal(t, "description1", test1.Usage)
	require.Equal(t, "string", test1.Value.Type())
	_, required1 := test1.Annotations[cobra.BashCompOneRequiredFlag]
	require.False(t, required1)

	test2 := testCmd.Flag("test2")
	require.NotNil(t, test2)
	require.Equal(t, "test2", test2.Name)
	require.Equal(t, "description2", test2.Usage)
	require.Equal(t, "boolSlice", test2.Value.Type())
	_, required2 := test2.Annotations[cobra.BashCompOneRequiredFlag]
	require.True(t, required2)
}

func TestAddParametersDuplicates(t *testing.T) {
	testCmd := &cobra.Command{}
	err := addParameters(testCmd, []api.Parameter{
		{Name: "test1", Description: "description1", Required: false, Type: api.ParameterType{IsArray: false, Type: api.TypeString}},
		{Name: "test1", Description: "description2", Required: true, Type: api.ParameterType{IsArray: true, Type: api.TypeBool}},
	})
	require.Error(t, err)
}
