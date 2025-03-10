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

package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GroupedAndSortedCommands []Group

type Group struct {
	Name        string
	Description string
	Commands    []Command
}

type Command struct {
	OperationID       string
	Aliases           []string
	Description       string
	RequestParameters RequestParameters
	Versions          []Version
	Watcher           *WatcherProperties
}

type RequestParameters struct {
	URL             string
	QueryParameters []Parameter
	URLParameters   []Parameter
	Verb            string
}

type Version struct {
	Version              string
	Sunset               *time.Time
	RequestContentType   string
	ResponseContentTypes []string
}

type Parameter struct {
	Name        string
	Aliases     []string
	Short       string
	Description string
	Required    bool
	Type        ParameterType
	Example     string
}

type ParameterType struct {
	IsArray bool
	Type    ParameterConcreteType
}

type ParameterConcreteType string

const (
	TypeString ParameterConcreteType = "string"
	TypeInt    ParameterConcreteType = "int"
	TypeBool   ParameterConcreteType = "bool"
)

type HTTPVerb string

func ToHTTPVerb(s string) (string, error) {
	verb := strings.ToUpper(s)

	switch verb {
	case http.MethodDelete:
		return "http.MethodDelete", nil
	case http.MethodGet:
		return "http.MethodGet", nil
	case http.MethodPatch:
		return "http.MethodPatch", nil
	case http.MethodPost:
		return "http.MethodPost", nil
	case http.MethodPut:
		return "http.MethodPut", nil
	default:
		return "", fmt.Errorf("invalid HTTP verb: %s", s)
	}
}

type WatcherProperties struct {
	Get    WatcherGetProperties
	Expect *WatcherExpectProperties
}

type WatcherGetProperties struct {
	OperationID string
	Version     string
	Params      map[string]string
}

type WatcherExpectProperties struct {
	HTTPCode int
	Match    *WatcherMatchProperties
}

type WatcherMatchProperties struct {
	Path   string
	Values []string
}
