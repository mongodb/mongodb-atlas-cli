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

package container

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestPortMappingFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    PortMapping
		expected string
	}{
		{
			name: "All fields populated",
			input: PortMapping{
				HostAddress:       "127.0.0.1",
				HostPort:          8080,
				ContainerPort:     80,
				ContainerProtocol: "tcp",
			},
			expected: "127.0.0.1:8080:80/tcp",
		},
		{
			name: "No host address",
			input: PortMapping{
				HostPort:          8080,
				ContainerPort:     80,
				ContainerProtocol: "udp",
			},
			expected: "8080:80/udp",
		},
		{
			name: "No host port",
			input: PortMapping{
				HostAddress:       "192.168.1.100",
				ContainerPort:     443,
				ContainerProtocol: "tcp",
			},
			expected: "192.168.1.100::443/tcp",
		},
		{
			name: "No container protocol",
			input: PortMapping{
				HostAddress:   "10.0.0.1",
				HostPort:      5000,
				ContainerPort: 5000,
			},
			expected: "10.0.0.1:5000:5000",
		},
		{
			name: "Only container port",
			input: PortMapping{
				ContainerPort: 8080,
			},
			expected: ":8080",
		},
		{
			name: "Host port and container port only",
			input: PortMapping{
				HostPort:      3000,
				ContainerPort: 3000,
			},
			expected: "3000:3000",
		},
		{
			name: "Host address and container port only",
			input: PortMapping{
				HostAddress:   "localhost",
				ContainerPort: 8080,
			},
			expected: "localhost::8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := portMappingFlag(tt.input)
			if result != tt.expected {
				t.Errorf("portMappingFlag(%+v) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParsePortMapping(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []PortMapping
	}{
		{
			name:  "All fields populated",
			input: "127.0.0.1:8080->80/tcp",
			expected: []PortMapping{
				{
					HostAddress:       "127.0.0.1",
					HostPort:          8080,
					ContainerPort:     80,
					ContainerProtocol: "tcp",
				},
			},
		},
		{
			name:  "No host address",
			input: "8080->80/udp",
			expected: []PortMapping{{
				HostPort:          8080,
				ContainerPort:     80,
				ContainerProtocol: "udp",
			}},
		},
		{
			name:  "No host port",
			input: "192.168.1.100->443/tcp",
			expected: []PortMapping{{
				HostAddress:       "192.168.1.100",
				ContainerPort:     443,
				ContainerProtocol: "tcp",
			}},
		},
		{
			name:  "No container protocol",
			input: "10.0.0.1:5000->5000",
			expected: []PortMapping{{
				HostAddress:   "10.0.0.1",
				HostPort:      5000,
				ContainerPort: 5000,
			}},
		},
		{
			name:  "Only container port",
			input: "8080",
			expected: []PortMapping{{
				ContainerPort: 8080,
			}},
		},
		{
			name:  "Host port and container port only",
			input: "3000->3000",
			expected: []PortMapping{{
				HostPort:      3000,
				ContainerPort: 3000,
			}},
		},
		{
			name:  "Host address and container port only",
			input: "localhost->8080",
			expected: []PortMapping{{
				HostAddress:   "localhost",
				ContainerPort: 8080,
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := parsePortMapping(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if diff := deep.Equal(result, tt.expected); diff != nil {
				t.Errorf("parsePortMapping(%s) diff: %+v", tt.input, diff)
			}
		})
	}
}

func TestJsonL(t *testing.T) {
	testCases := []struct {
		input   string
		expect  []map[string]any
		wantErr bool
	}{
		{
			input: `{"a":1,"b":2}`,
			expect: []map[string]any{
				{"a": 1.0, "b": 2.0},
			},
			wantErr: false,
		},
		{
			input: `{"a":1,"b":2}
{"c":3,"d":4}`,
			expect: []map[string]any{
				{"a": 1.0, "b": 2.0},
				{"c": 3.0, "d": 4.0},
			},
			wantErr: false,
		},
		{
			input: `{"a":1,"b":2}
{"c":3,"d":`,
			expect:  nil,
			wantErr: true,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			got, err := readJsonl[map[string]any](strings.NewReader(testCase.input))
			if err != nil {
				if testCase.wantErr {
					return
				}
				t.Fatal(err)
			}
			if testCase.wantErr {
				t.Fatal("error expected")
			}
			if diff := deep.Equal(testCase.expect, got); diff != nil {
				t.Fatal(diff)
			}
		})
	}
}
