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

func URL(val interface{}) error {
	var u string
	var ok bool
	if u, ok = val.(string); !ok {
		return fmt.Errorf("'%v' is not a valid URL", val)
	}
	if !strings.HasSuffix(u, "/") {
		return fmt.Errorf("'%s' must have a trailing slash", u)
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid URL", u)
	}

	return nil
}

func ObjectID(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid Id", s)
	}
	return nil
}

func Credentials() error {
	if config.PrivateAPIKey() == "" || config.PublicAPIKey() == "" {
		return errors.New("missing credentials")
	}
	return nil
}
