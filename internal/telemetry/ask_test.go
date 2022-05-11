// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telemetry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAnswer(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		input    interface{}
		name     string
		expected interface{}
	}{
		{
			input: &struct {
				Test  string
				Test2 int
			}{
				Test:  "value",
				Test2: 2,
			},
			name:     "Test",
			expected: "value",
		},
		{
			input: &struct {
				Test  string
				Test2 int
			}{
				Test:  "value",
				Test2: 2,
			},
			name:     "Test2",
			expected: 2,
		},
		{
			input:    map[string]interface{}{"test": "value", "test2": 2},
			name:     "test",
			expected: "value",
		},
		{
			input:    map[string]interface{}{"test": "value", "test2": 2},
			name:     "test2",
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		answer, err := readAnswer(testCase.input, testCase.name)
		a.NoError(err)
		a.Equal(testCase.expected, answer)
	}
}
