// Copyright 2024 MongoDB Inc
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

//nolint:revive,stylecheck
package L1

import (
	"fmt"
	"strings"
)

type GroupedAndSortedCommands []Group

type Group struct {
	Name        string
	Description string
	Commands    []Command
}

type Command struct {
	OperationID       string
	Description       string
	RequestParameters RequestParameters
	Versions          []Version
}

type RequestParameters struct {
	URL             string
	QueryParameters []Parameter
	URLParameters   []Parameter
	Verb            HTTPVerb
}

type Version struct {
	Version              string
	RequestContentTypes  []string
	ResponseContentTypes []string
}

type Parameter struct {
	Name        string
	Description string
	Required    bool
}

type HTTPVerb string

const (
	DELETE HTTPVerb = "DELETE"
	GET    HTTPVerb = "GET"
	PATCH  HTTPVerb = "PATCH"
	POST   HTTPVerb = "POST"
	PUT    HTTPVerb = "PUT"
)

func ToHTTPVerb(s string) (HTTPVerb, error) {
	verb := HTTPVerb(strings.ToUpper(s))

	switch verb {
	case DELETE, GET, PATCH, POST, PUT:
		return verb, nil
	default:
		return "", fmt.Errorf("invalid HTTP verb: %s", s)
	}
}
