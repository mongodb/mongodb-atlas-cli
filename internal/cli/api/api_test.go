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
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

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

type fixedValueProvider struct{}

func (*fixedValueProvider) ValueForFlag(flagName string) (*string, error) {
	if flagName == "b" {
		return nil, nil
	}

	return pointer.Get("default-for-" + flagName), nil
}

func TestSetUnTouchedFlags(t *testing.T) {
	// Create cobra command with flags
	testCmd := &cobra.Command{}
	testCmd.Flags().String("a", "a", "will remain unchanged by test")
	testCmd.Flags().String("b", "b", "will remain unchanged by test, it also won't be changed by the fixedValueProvider")
	testCmd.Flags().String("c", "c", "will be changed by test")
	testCmd.Flags().String("d", "d", "will remain unchanged by test")
	testCmd.Flags().String("e", "e", "will be changed by test")

	// Parse flags, this will also set `Changed`
	// Adding test here in case cobra behavior changes
	require.NoError(t, testCmd.ParseFlags([]string{"--c", "user-changed-c", "--e", "user-changed-e"}))

	// This is the state before PreRunE
	require.Equal(t, "a", testCmd.Flag("a").Value.String())
	require.Equal(t, "b", testCmd.Flag("b").Value.String())
	require.Equal(t, "user-changed-c", testCmd.Flag("c").Value.String())
	require.Equal(t, "d", testCmd.Flag("d").Value.String())
	require.Equal(t, "user-changed-e", testCmd.Flag("e").Value.String())

	require.False(t, testCmd.Flag("a").Changed)
	require.False(t, testCmd.Flag("b").Changed)
	require.True(t, testCmd.Flag("c").Changed)
	require.False(t, testCmd.Flag("d").Changed)
	require.True(t, testCmd.Flag("e").Changed)

	// The actual method we're testing
	require.NoError(t, setUnTouchedFlags(&fixedValueProvider{}, testCmd))

	// Verify that untouched values have been updated with their respective new values
	require.Equal(t, "default-for-a", testCmd.Flag("a").Value.String())
	require.Equal(t, "b", testCmd.Flag("b").Value.String())
	require.Equal(t, "user-changed-c", testCmd.Flag("c").Value.String())
	require.Equal(t, "default-for-d", testCmd.Flag("d").Value.String())
	require.Equal(t, "user-changed-e", testCmd.Flag("e").Value.String())
	require.True(t, testCmd.Flag("a").Changed)
	require.False(t, testCmd.Flag("b").Changed)
	require.True(t, testCmd.Flag("c").Changed)
	require.True(t, testCmd.Flag("d").Changed)
	require.True(t, testCmd.Flag("e").Changed)
}
