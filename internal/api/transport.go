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
	"net/http"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
)

func authenticatedTransport(authenticatedConfig store.AuthenticatedConfig, httpTransport http.RoundTripper) http.RoundTripper {
	username := authenticatedConfig.PublicAPIKey()
	password := authenticatedConfig.PrivateAPIKey()

	if username != "" && password != "" {
		return &digest.Transport{
			Username:  username,
			Password:  password,
			Transport: httpTransport,
		}
	}

	return &transport{
		authenticatedConfig: authenticatedConfig,
		base:                httpTransport,
	}
}

type transport struct {
	authenticatedConfig store.AuthenticatedConfig
	base                http.RoundTripper
}

func (tr *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := tr.authenticatedConfig.Token()
	if err == nil {
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}

	return tr.base.RoundTrip(req)
}
