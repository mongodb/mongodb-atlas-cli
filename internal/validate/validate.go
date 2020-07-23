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

package validate

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/mongodb/mongocli/internal/config"
)

// toString tries to cast an interface to string
func toString(val interface{}) (string, error) {
	var u string
	var ok bool
	if u, ok = val.(string); !ok {
		return "", fmt.Errorf("'%v' is not valid", val)
	}
	return u, nil
}

// URL validates a value is a valid URL for the cli store
func URL(val interface{}) error {
	s, err := toString(val)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(s, "/") {
		return fmt.Errorf("'%s' must have a trailing slash", s)
	}
	_, err = url.ParseRequestURI(s)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid URL", s)
	}

	return nil
}

// OptionalURL validates a value is a valid URL for the cli store
func OptionalURL(val interface{}) error {
	if val == nil {
		return nil
	}
	s, err := toString(val)
	if err != nil {
		return err
	}
	if s == "" {
		return nil
	}

	return URL(val)
}

// OptionalObjectID validates a value is a valid ObjectID
func OptionalObjectID(val interface{}) error {
	if val == nil {
		return nil
	}
	s, err := toString(val)
	if err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	return ObjectID(s)
}

// ObjectID validates a value is a valid ObjectID
func ObjectID(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid Id", s)
	}
	return nil
}

// Credentials validates public and private API keys have been set
func Credentials() error {
	if config.PrivateAPIKey() == "" || config.PublicAPIKey() == "" {
		return errors.New("missing credentials")
	}
	return nil
}

func EnsureFlagHasValidValue(value, flag string, validValues []string) error {
	for _, v := range(validValues) {
		if value == v {
			return nil
		}
	}

	return fmt.Errorf("%v is an invalid value for %v. It must be one of %v", value, flag, strings.Join(validValues, ","))
}
