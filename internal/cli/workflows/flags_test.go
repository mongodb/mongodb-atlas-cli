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
	"reflect"
	"testing"
)

func TestRemoveGlobalFlagsAndArgs(t *testing.T) {
	flagsToBeRemoved := map[string]string{
		"debug": "1",
	}
	argsToBeRemoved := map[string]bool{}
	args := []string{"--debug", "value1", "arg1"}

	expectedNewArgs := []string{"arg1"} // "--debug" and "value1" should be removed

	newArgs, err := RemoveFlagsAndArgs(flagsToBeRemoved, argsToBeRemoved, args)

	if !reflect.DeepEqual(newArgs, expectedNewArgs) {
		t.Errorf("Expected newArgs %v, but got %v", expectedNewArgs, newArgs)
	}

	if err != nil {
		t.Errorf("Expected error %v, but got error %v", nil, err)
	}
}

func TestRemoveArgs(t *testing.T) {
	flagsToBeRemoved := map[string]string{}
	argsToBeRemoved := map[string]bool{
		"arg1": true,
	}
	args := []string{"arg1", "arg2", "arg3"}

	expectedNewArgs := []string{"arg2", "arg3"} // "arg1" should be removed

	newArgs, err := RemoveFlagsAndArgs(flagsToBeRemoved, argsToBeRemoved, args)

	if !reflect.DeepEqual(newArgs, expectedNewArgs) {
		t.Errorf("Expected newArgs %v, but got %v", expectedNewArgs, newArgs)
	}

	if err != nil {
		t.Errorf("Expected error %v, but got error %v", nil, err)
	}
}

func TestNoRemoval(t *testing.T) {
	flagsToBeRemoved := map[string]string{
		"debug": "1",
	}
	argsToBeRemoved := map[string]bool{
		"arg1": true,
	}
	args := []string{"--verbose", "arg2"}

	expectedNewArgs := []string{"--verbose", "arg2"} // No removal

	newArgs, err := RemoveFlagsAndArgs(flagsToBeRemoved, argsToBeRemoved, args)

	if !reflect.DeepEqual(newArgs, expectedNewArgs) {
		t.Errorf("Expected newArgs %v, but got %v", expectedNewArgs, newArgs)
	}

	if err != nil {
		t.Errorf("Expected error %v, but got error %v", nil, err)
	}
}
