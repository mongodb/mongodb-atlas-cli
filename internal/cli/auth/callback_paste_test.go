// Copyright 2026 MongoDB Inc.
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

package auth

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type staticCallbackPasteVector struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	Code   string `json:"code"`
	Issuer string `json:"issuer"`
	Paste  string `json:"paste"`
}

func loadStaticCallbackPasteVectors(t *testing.T) []staticCallbackPasteVector {
	t.Helper()

	data, err := os.ReadFile("../../../oauth-callback/testdata/paste-v1.json")
	require.NoError(t, err)
	var fixture struct {
		Format  string                      `json:"format"`
		Version string                      `json:"version"`
		Vectors []staticCallbackPasteVector `json:"vectors"`
	}
	require.NoError(t, json.Unmarshal(data, &fixture))
	require.Equal(t, staticCallbackPasteFormat, fixture.Format)
	require.Equal(t, "1", fixture.Version)
	require.NotEmpty(t, fixture.Vectors)
	return fixture.Vectors
}

func TestDecodeStaticCallbackPaste(t *testing.T) {
	for _, vector := range loadStaticCallbackPasteVectors(t) {
		t.Run(vector.Name, func(t *testing.T) {
			got, err := DecodeStaticCallbackPaste(vector.Paste, vector.State, vector.Issuer)
			require.NoError(t, err)
			require.Equal(t, vector.Code, got)
		})
	}
}

func TestDecodeStaticCallbackPasteRejectsInvalidValues(t *testing.T) {
	vector := loadStaticCallbackPasteVectors(t)[0]
	badChecksum := vector.Paste[:len(vector.Paste)-1] + "A"
	if strings.HasSuffix(vector.Paste, "A") {
		badChecksum = vector.Paste[:len(vector.Paste)-1] + "B"
	}

	tests := []struct {
		name   string
		paste  string
		state  string
		issuer string
	}{
		{name: "empty state", paste: vector.Paste, issuer: vector.Issuer},
		{name: "wrong state", paste: vector.Paste, state: vector.State + "wrong", issuer: vector.Issuer},
		{name: "wrong issuer", paste: vector.Paste, state: vector.State, issuer: vector.Issuer + "/wrong"},
		{name: "invalid base64", paste: "!abcd", state: vector.State, issuer: vector.Issuer},
		{name: "incorrect checksum", paste: badChecksum, state: vector.State, issuer: vector.Issuer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DecodeStaticCallbackPaste(tt.paste, tt.state, tt.issuer)
			require.ErrorIs(t, err, ErrInvalidStaticCallbackPaste)
		})
	}
}
