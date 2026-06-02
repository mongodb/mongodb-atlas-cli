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

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedactPEMPrivateKeys(t *testing.T) {
	const certBlock = "-----BEGIN CERTIFICATE-----\nMIIFBzCC\n-----END CERTIFICATE-----\n"
	const pkcs8Key = "-----BEGIN PRIVATE KEY-----\nMIIJQwIBADA\n-----END PRIVATE KEY-----"
	const rsaKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAK\n-----END RSA PRIVATE KEY-----"
	const ecKey = "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIP\n-----END EC PRIVATE KEY-----"
	const pubKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkq\n-----END PUBLIC KEY-----"

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no key material unchanged",
			input: "plain text response",
			want:  "plain text response",
		},
		{
			name:  "PKCS8 private key redacted",
			input: pkcs8Key,
			want:  "REDACTED",
		},
		{
			name:  "RSA private key redacted",
			input: rsaKey,
			want:  "REDACTED",
		},
		{
			name:  "EC private key redacted",
			input: ecKey,
			want:  "REDACTED",
		},
		{
			name:  "cert preserved key redacted in mixed payload",
			input: certBlock + pkcs8Key,
			want:  certBlock + "REDACTED",
		},
		{
			name:  "public key unchanged",
			input: pubKey,
			want:  pubKey,
		},
		{
			name:  "certificate alone unchanged",
			input: certBlock,
			want:  certBlock,
		},
		{
			name:  "multiple private keys all redacted",
			input: pkcs8Key + "\n" + rsaKey,
			want:  "REDACTED\nREDACTED",
		},
		{
			name:  "json-encoded PEM (escaped newlines) redacted",
			input: `{"key":"-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAK\n-----END RSA PRIVATE KEY-----"}`,
			want:  `{"key":"REDACTED"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := redactPEMPrivateKeys([]byte(tt.input))
			assert.Equal(t, tt.want, string(got))
		})
	}
}
