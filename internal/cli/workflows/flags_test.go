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

//go:build unit

package workflows

import (
	"testing"

	"github.com/go-test/deep"
)

func TestRemoveGlobalFlagsAndArgs(t *testing.T) {
	flagsToBeRemoved := map[string]string{
		"debug": "1",
	}
	argsToBeRemoved := map[string]bool{}
	args := []string{"binary", "command", "subcommand", "--debug", "value1", "arg1"}

	expectedNewArgs := []string{"arg1"} // "--debug" and "value1" should be removed

	newArgs, err := RemoveFlagsAndArgs(flagsToBeRemoved, argsToBeRemoved, args)
	if err != nil {
		t.Fatalf("unexpected error:  %v\n", err)
	}
	if diff := deep.Equal(newArgs, expectedNewArgs); diff != nil {
		t.Error(diff)
	}
}

func TestRemoveArgs(t *testing.T) {
	flagsToBeRemoved := map[string]string{}
	argsToBeRemoved := map[string]bool{
		"arg1": true,
	}
	args := []string{"binary", "command", "subcommand", "arg1", "arg2", "arg3"}

	expectedNewArgs := []string{"arg2", "arg3"} // "arg1" should be removed

	newArgs, err := RemoveFlagsAndArgs(flagsToBeRemoved, argsToBeRemoved, args)
	if err != nil {
		t.Fatalf("unexpected error:  %v\n", err)
	}
	if diff := deep.Equal(newArgs, expectedNewArgs); diff != nil {
		t.Error(diff)
	}
}

func TestNoRemoval(t *testing.T) {
	flagsToBeRemoved := map[string]string{
		"debug": "1",
	}
	argsToBeRemoved := map[string]bool{
		"arg1": true,
	}
	args := []string{"binary", "command", "subcommand", "--verbose", "arg2"}

	expectedNewArgs := []string{"--verbose", "arg2"} // No removal

	newArgs, err := RemoveFlagsAndArgs(flagsToBeRemoved, argsToBeRemoved, args)
	if err != nil {
		t.Fatalf("unexpected error:  %v\n", err)
	}
	if diff := deep.Equal(newArgs, expectedNewArgs); diff != nil {
		t.Error(diff)
	}
}
