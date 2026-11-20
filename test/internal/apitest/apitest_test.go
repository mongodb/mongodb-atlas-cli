// Copyright 2026 MongoDB Inc
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

package apitest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONValue(t *testing.T) {
	body := map[string]any{
		"results": []any{
			map[string]any{
				"instanceSizes": []any{
					map[string]any{"name": "M10"},
				},
			},
		},
		"totalCount": 1.0,
	}

	assert.Equal(t, "M10", JSONValue(body, "results/0/instanceSizes/0/name"))
	assert.Nil(t, JSONValue(body, "results/1/instanceSizes/0/name"))
	assert.Nil(t, JSONValue(body, "results/notanindex"))
	assert.Equal(t, body, JSONValue(body, ""))
}

func TestJSONString(t *testing.T) {
	body := map[string]any{
		"version":    "8.0",
		"diskSizeGB": 30.0,
		"paused":     false,
	}

	assert.Equal(t, "8.0", JSONString(body, "version"))
	assert.Equal(t, "30", JSONString(body, "diskSizeGB"))
	assert.Equal(t, "false", JSONString(body, "paused"))
	assert.Empty(t, JSONString(body, "missing"))
}

func TestPreviousMajorVersion(t *testing.T) {
	previous, err := PreviousMajorVersion("8.0")
	require.NoError(t, err)
	assert.Equal(t, "7.0", previous)

	_, err = PreviousMajorVersion("8")
	require.ErrorContains(t, err, "invalid MongoDB major version")

	_, err = PreviousMajorVersion("eight.zero")
	require.ErrorContains(t, err, "invalid MongoDB major version")
}
