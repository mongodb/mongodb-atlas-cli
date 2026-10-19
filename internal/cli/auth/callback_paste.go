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
	"crypto/hkdf"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"unicode/utf8"
)

const (
	staticCallbackPasteFormat        = "atlas-cli-paste"
	staticCallbackPasteChecksumBytes = 4
	maxStaticCallbackPayloadBytes    = 255 * sha256.Size
)

// ErrInvalidStaticCallbackPaste indicates that a compact callback response
// could not be decoded or its copy-check checksum did not match.
var ErrInvalidStaticCallbackPaste = errors.New("invalid static callback paste value")

// DecodeStaticCallbackPaste decodes the hosted callback page's v1 compact
// response using the state from the active OAuth transaction. expectedIssuer
// is the issuer from discovery when authorization_response_iss_parameter_supported
// is true; pass an empty string when the server does not support the parameter.
// The four-byte checksum detects copy errors only; it is not an integrity check.
func DecodeStaticCallbackPaste(value, state, expectedIssuer string) (string, error) {
	if state == "" {
		return "", ErrInvalidStaticCallbackPaste
	}

	codeBytes, checksum, err := decryptStaticCallbackPaste(value, state, expectedIssuer)
	if err != nil {
		return "", err
	}
	if !utf8.Valid(codeBytes) || len(codeBytes) == 0 {
		return "", ErrInvalidStaticCallbackPaste
	}
	code := string(codeBytes)
	if !staticCallbackPasteChecksumMatches(code, checksum) {
		return "", ErrInvalidStaticCallbackPaste
	}

	return code, nil
}

func decryptStaticCallbackPaste(value, state, issuer string) ([]byte, string, error) {
	encoded := value
	maxEncodedBytes := base64.RawURLEncoding.EncodedLen(maxStaticCallbackPayloadBytes) + staticCallbackPasteChecksumBytes
	if len(encoded) <= staticCallbackPasteChecksumBytes || len(encoded) > maxEncodedBytes {
		return nil, "", ErrInvalidStaticCallbackPaste
	}

	checksumStart := len(encoded) - staticCallbackPasteChecksumBytes
	checksum := encoded[checksumStart:]
	ciphertext, err := base64.RawURLEncoding.Strict().DecodeString(encoded[:checksumStart])
	if err != nil || len(ciphertext) == 0 || len(ciphertext) > maxStaticCallbackPayloadBytes {
		return nil, "", ErrInvalidStaticCallbackPaste
	}

	stream, err := hkdf.Key(sha256.New, []byte(state), nil, staticCallbackPasteHKDFInfo(issuer), len(ciphertext))
	if err != nil {
		return nil, "", ErrInvalidStaticCallbackPaste
	}
	codeBytes := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		codeBytes[i] = b ^ stream[i]
	}
	return codeBytes, checksum, nil
}

func staticCallbackPasteHKDFInfo(issuer string) string {
	if issuer != "" {
		return issuer
	}
	return staticCallbackPasteFormat
}

func staticCallbackPasteChecksumMatches(code, checksum string) bool {
	digest := sha256.Sum256([]byte(code))
	expectedChecksum := base64.RawURLEncoding.EncodeToString(digest[:])[:staticCallbackPasteChecksumBytes]
	return subtle.ConstantTimeCompare([]byte(checksum), []byte(expectedChecksum)) == 1
}
