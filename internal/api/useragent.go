// Copyright 2026 MongoDB Inc
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

import "net/http"

// userAgentTransport sets the User-Agent on every request.
//
// The clients built by atlas-cli-core only set a User-Agent on the service
// account and OAuth token paths, so without this the API key, user account and
// unauthenticated paths would send Go's default User-Agent and the request would
// not be attributable to the Atlas CLI.
type userAgentTransport struct {
	base      http.RoundTripper
	userAgent string
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// A RoundTripper must not modify the request it is given.
	clone := req.Clone(req.Context())
	clone.Header.Set("User-Agent", t.userAgent)

	return t.transport().RoundTrip(clone)
}

func (t *userAgentTransport) transport() http.RoundTripper {
	if t.base != nil {
		return t.base
	}

	return http.DefaultTransport
}
