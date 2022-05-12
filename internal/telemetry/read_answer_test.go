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
			input:    map[string]interface{}{"test": "value", "test2": 2},
			name:     "test",
			expected: "value",
		},
		{
			input:    map[string]interface{}{"test": "value", "test2": 2},
			name:     "test2",
			expected: 2,
		},
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
				Test2 struct {
					Test3 string
					Test4 int
				}
			}{
				Test: "value",
				Test2: struct {
					Test3 string
					Test4 int
				}{
					Test3: "value3",
					Test4: 4,
				},
			},
			name:     "Test4",
			expected: 4,
		},
		{
			input: &struct {
				Test  string
				Test2 *struct {
					Test3 string
					Test4 int
				}
			}{
				Test: "value",
				Test2: &struct {
					Test3 string
					Test4 int
				}{
					Test3: "value3",
					Test4: 4,
				},
			},
			name:     "Test4",
			expected: 4,
		},
		{
			input: &struct {
				Test  string
				Test2 int `survey:"f"`
			}{
				Test:  "value",
				Test2: 2,
			},
			name:     "f",
			expected: 2,
		},
		{
			input: &struct {
				Test  string
				Test2 int
				Test3 interface {
					String()
				}
			}{
				Test:  "value",
				Test2: 2,
			},
			name:     "test",
			expected: "value",
		},
	}

	for _, testCase := range testCases {
		answer, err := readAnswer(testCase.input, testCase.name)
		a.NoError(err)
		a.Equal(testCase.expected, answer)
	}
}

func TestReadAnswerNotFound(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		input interface{}
		name  string
	}{
		{
			input: map[string]interface{}{"test": "value", "test2": 2},
			name:  "test3",
		},
		{
			input: &struct {
				Test  string
				Test2 int
			}{
				Test:  "value",
				Test2: 2,
			},
			name: "Test3",
		},
		{
			input: &struct {
				Test  string
				Test2 struct {
					Test3 string
					Test4 int
				}
			}{
				Test: "value",
				Test2: struct {
					Test3 string
					Test4 int
				}{
					Test3: "value3",
					Test4: 4,
				},
			},
			name: "Test5",
		},
	}

	for _, testCase := range testCases {
		_, err := readAnswer(testCase.input, testCase.name)
		a.ErrorIs(err, ErrFieldNotFound)
	}
}

func TestReadAnswerNotStructOrMap(t *testing.T) {
	a := assert.New(t)

	test := "value"

	testCases := []struct {
		input interface{}
		name  string
	}{
		{
			input: "test",
			name:  "test2",
		},
		{
			input: 1,
			name:  "test3",
		},
		{
			input: &test,
			name:  "test4",
		},
	}

	for _, testCase := range testCases {
		_, err := readAnswer(testCase.input, testCase.name)
		a.ErrorIs(err, ErrNotMapOrStruct)
	}
}
