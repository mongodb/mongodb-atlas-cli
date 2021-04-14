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

// +build unit

package store

import (
	"testing"

	"go.mongodb.org/atlas/mongodbatlas"
)

type auth struct {
	username string
	password string
}

func (a auth) PublicAPIKey() string {
	return a.username
}

func (a auth) PrivateAPIKey() string {
	return a.password
}

var _ CredentialsGetter = &auth{}

func TestService(t *testing.T) {
	c, err := New(Service("cloud"))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.service != "cloud" {
		t.Errorf("New() service = %s; expected %s", c.service, "cloud")
	}
}

func TestWithBaseURL(t *testing.T) {
	c, err := New(Service("cloud"), WithBaseURL("http://test"))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.baseURL != "http://test" {
		t.Errorf("New() baseURL = %s; expected %s", c.baseURL, "http://test")
	}
}

func TestSkipVerify(t *testing.T) {
	c, err := New(Service("cloud"), SkipVerify())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if !c.skipVerify {
		t.Error("New() skipVerify not set")
	}
}

func TestWithPublicPathBaseURL(t *testing.T) {
	c, err := New(Service("cloud"), WithBaseURL("http://test"), WithPublicPathBaseURL())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	expected := "http://test" + mongodbatlas.APIPublicV1Path
	if c.baseURL != expected {
		t.Errorf("New() baseURL = %s; expected %s", c.baseURL, expected)
	}
}

func TestWithAuthentication(t *testing.T) {
	a := auth{
		username: "username",
		password: "password",
	}
	c, err := New(Service("cloud"), WithAuthentication(a))

	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.username != a.username {
		t.Errorf("New() username = %s; expected %s", c.username, a.username)
	}
	if c.password != a.password {
		t.Errorf("New() password = %s; expected %s", c.password, a.password)
	}
}
