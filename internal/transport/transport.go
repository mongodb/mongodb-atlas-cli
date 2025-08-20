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

package transport

import (
	"net/http"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	atlasauth "go.mongodb.org/atlas/auth"
)

func Default() *http.Transport {
	return transport.Default()
}

func Telemetry() *http.Transport {
	return transport.Telemetry()
}

func NewDigestTransport(username, password string, base http.RoundTripper) *digest.Transport {
	return transport.NewDigestTransport(username, password, base)
}

func NewAccessTokenTransport(token *atlasauth.Token, base http.RoundTripper, saveToken func(*atlasauth.Token) error) (http.RoundTripper, error) {
	return transport.NewAccessTokenTransport(token, base, version.Version, saveToken)
}
