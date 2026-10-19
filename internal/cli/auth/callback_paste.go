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
	"bytes"
	"crypto/hkdf"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"unicode/utf8"
)

const (
	staticCallbackPasteFormat        = "atlas-cli-paste"
	staticCallbackPasteInfo          = staticCallbackPasteFormat
	staticCallbackPasteChecksumBytes = 4
	maxStaticCallbackPayloadBytes    = 255 * sha256.Size
)

// ErrInvalidStaticCallbackPaste indicates that a compact callback response
// could not be decoded or its copy-check checksum did not match.
var ErrInvalidStaticCallbackPaste = errors.New("invalid static callback paste value")

// StaticCallbackResponse contains the authorization response fields packaged
// by the hosted callback page. State and redirect validation remain the
// responsibility of the active OAuth transaction.
type StaticCallbackResponse struct {
	Code   string
	Issuer string
}

type staticCallbackPayload struct {
	Code   string `json:"code"`
	Issuer string `json:"iss"`
}

// DecodeStaticCallbackPaste decodes the hosted callback page's v1 compact
// response using the state from the active OAuth transaction. The four-byte
// checksum detects copy errors only; it is not an integrity check.
func DecodeStaticCallbackPaste(value, state string) (StaticCallbackResponse, error) {
	if state == "" {
		return StaticCallbackResponse{}, ErrInvalidStaticCallbackPaste
	}

	payloadBytes, checksum, err := decryptStaticCallbackPaste(value, state)
	if err != nil {
		return StaticCallbackResponse{}, err
	}
	payload, err := decodeStaticCallbackPayload(payloadBytes)
	if err != nil {
		return StaticCallbackResponse{}, err
	}
	if !staticCallbackPasteChecksumMatches(payload.Code, checksum) {
		return StaticCallbackResponse{}, ErrInvalidStaticCallbackPaste
	}

	return StaticCallbackResponse(payload), nil
}

func decryptStaticCallbackPaste(value, state string) ([]byte, string, error) {
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

	stream, err := hkdf.Key(sha256.New, []byte(state), nil, staticCallbackPasteInfo, len(ciphertext))
	if err != nil {
		return nil, "", ErrInvalidStaticCallbackPaste
	}
	payloadBytes := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		payloadBytes[i] = b ^ stream[i]
	}
	return payloadBytes, checksum, nil
}

func decodeStaticCallbackPayload(payloadBytes []byte) (staticCallbackPayload, error) {
	if !utf8.Valid(payloadBytes) {
		return staticCallbackPayload{}, ErrInvalidStaticCallbackPaste
	}

	decoder := json.NewDecoder(bytes.NewReader(payloadBytes))
	decoder.DisallowUnknownFields()
	var payload staticCallbackPayload
	if err := decoder.Decode(&payload); err != nil {
		return staticCallbackPayload{}, ErrInvalidStaticCallbackPaste
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return staticCallbackPayload{}, ErrInvalidStaticCallbackPaste
	}
	if payload.Code == "" || payload.Issuer == "" {
		return staticCallbackPayload{}, ErrInvalidStaticCallbackPaste
	}

	return payload, nil
}

func staticCallbackPasteChecksumMatches(code, checksum string) bool {
	digest := sha256.Sum256([]byte(code))
	expectedChecksum := base64.RawURLEncoding.EncodeToString(digest[:])[:staticCallbackPasteChecksumBytes]
	return subtle.ConstantTimeCompare([]byte(checksum), []byte(expectedChecksum)) == 1
}
