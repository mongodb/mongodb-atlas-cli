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
	"testing"

	"github.com/docker/docker-credential-helpers/credentials"
)

func TestKeyring(t *testing.T) {
	var store = nativeStore()

	const url = "https://atlascli.internal/serviceAccount"
	const username = "client_id"
	const secret = "client_secret"

	c := &credentials.Credentials{
		ServerURL: url,
		Username:  username,
		Secret:    secret,
	}

	store.Add(c)

	gotSecret, gotUsername, gotErr := store.Get("https://api.github.com")
	if gotErr != nil {
		t.Errorf("Get() error = %v", gotErr)
		return
	}
	if gotSecret != secret {
		t.Errorf("Get() secret = %v, want %v", gotSecret, secret)
	}
	if gotUsername != username {
		t.Errorf("Get() username = %v, want %v", gotUsername, username)
	}
	if err := store.Delete(url); err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}
