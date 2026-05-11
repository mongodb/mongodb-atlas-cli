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

package pledge

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const hmacKeyLen = 32

// loadOrCreateHMACKey returns the per-host HMAC key, creating one if absent.
// The key file is mode 0600.
func loadOrCreateHMACKey() ([]byte, error) {
	dir, err := StateDir()
	if err != nil {
		return nil, err
	}
	if err := ensureDir(dir); err != nil {
		return nil, err
	}
	keyPath := filepath.Join(dir, ".hmac")

	data, err := os.ReadFile(keyPath)
	if err == nil && len(data) == hex.EncodedLen(hmacKeyLen) {
		key, decErr := hex.DecodeString(string(data))
		if decErr == nil {
			return key, nil
		}
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("reading HMAC key: %w", err)
	}

	// Generate a fresh key.
	key := make([]byte, hmacKeyLen)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("generating HMAC key: %w", err)
	}
	encoded := []byte(hex.EncodeToString(key))
	if err := os.WriteFile(keyPath, encoded, 0o600); err != nil {
		return nil, fmt.Errorf("writing HMAC key: %w", err)
	}
	return key, nil
}

// SignHMAC produces a hex HMAC-SHA256 over data using the per-host key.
func SignHMAC(data []byte) (string, error) {
	key, err := loadOrCreateHMACKey()
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// VerifyHMAC checks that sig is the correct HMAC-SHA256 of data.
func VerifyHMAC(data []byte, sig string) (bool, error) {
	expected, err := SignHMAC(data)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(expected), []byte(sig)), nil
}
