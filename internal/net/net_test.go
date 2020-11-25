// Copyright 2020 MongoDB Inc
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
package net

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const publicIP = "212.129.77.106"

func TestNewIPAddress(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	ip, err := NewIPAddress()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if ip != publicIP {
		t.Fatalf("expected %s, got %s", publicIP, ip)
	}
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("https://api.ipify.org", responseMock)
	srv := httptest.NewServer(handler)
	return srv
}

func responseMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(publicIP))
}
