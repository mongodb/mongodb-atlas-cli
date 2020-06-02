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

package atlas

import "testing"

func TestGetHostNameAndPort(t *testing.T) {
	t.Run("valid parameter", func(t *testing.T) {
		host, port, err := getHostnameAndPort("test:2000")
		if err != nil {
			t.Fatalf("getHostnameAndPort unexpecteted err: %#v\n", err)
		}
		if host != "test" {
			t.Errorf("Expected '%s', got '%s'\n", "test", host)
		}
		if port != 2000 {
			t.Errorf("Expected '%d', got '%d'\n", 2000, port)
		}
	})
	t.Run("incomplete format", func(t *testing.T) {
		_, _, err := getHostnameAndPort("test")
		if err == nil {
			t.Fatal("getHostnameAndPort should return an error\n")
		}
	})
	t.Run("incomplete format", func(t *testing.T) {
		_, _, err := getHostnameAndPort(":test")
		if err == nil {
			t.Fatal("getHostnameAndPort should return an error\n")
		}
	})
}
