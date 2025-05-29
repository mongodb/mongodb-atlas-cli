// Copyright 2025 MongoDB Inc
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
//go:build e2e || e2eSnap || (atlas && keyring)

package keyring

import (
	"fmt"
	"testing"

	"github.com/zalando/go-keyring"
)

func TestKeyring(t *testing.T) {
	const service = "atlascli"

	for i := range 3 {
		username := fmt.Sprintf("client_id_%d", i)
		secret := fmt.Sprintf("client_secret_%d", i)

		err := keyring.Set(service, username, secret)
		if err != nil {
			t.Fatalf("Set() error = %v", err)
		}
	}

	for i := range 3 {
		username := fmt.Sprintf("client_id_%d", i)
		secret := fmt.Sprintf("client_secret_%d", i)

		gotSecret, gotErr := keyring.Get(service, username)
		if gotErr != nil {
			t.Errorf("Get() error = %v", gotErr)
			return
		}
		if gotSecret != secret {
			t.Errorf("Get() secret = %v, want %v", gotSecret, secret)
		}
	}

	if err := keyring.DeleteAll(service); err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}
