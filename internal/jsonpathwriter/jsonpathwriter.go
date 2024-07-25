// Copyright 2021 MongoDB Inc
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

package jsonpathwriter

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/PaesslerAG/jsonpath"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/jsonwriter"
)

var ErrEmptyPath = errors.New("empty jsonpath")

func Print(w io.Writer, path string, obj any) error {
	if path == "" {
		return ErrEmptyPath
	}

	jsonString, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	var val any
	if e := json.Unmarshal(jsonString, &val); e != nil {
		return e
	}

	v, err := jsonpath.Get(path, val)
	if err != nil {
		return err
	}
	return jsonwriter.Print(w, v)
}
