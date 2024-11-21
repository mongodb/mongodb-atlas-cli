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
	"net"
	"net/http"
	"time"

	"github.com/mongodb-forks/digest"
	atlasauth "go.mongodb.org/atlas/auth"
)

const (
	telemetryTimeout      = 1 * time.Second
	timeout               = 5 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleConns          = 5
	maxIdleConnsPerHost   = 4
	idleConnTimeout       = 30 * time.Second
	expectContinueTimeout = 1 * time.Second
)

var defaultTransport = newTransport(timeout)

func Default() *http.Transport {
	return defaultTransport
}

var telemetryTransport = newTransport(telemetryTimeout)

func Telemetry() *http.Transport {
	return telemetryTransport
}

func newTransport(timeout time.Duration) *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		Proxy:                 http.ProxyFromEnvironment,
		IdleConnTimeout:       idleConnTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
	}
}

func NewDigestTransport(username, password string, base http.RoundTripper) *digest.Transport {
	return &digest.Transport{
		Username:  username,
		Password:  password,
		Transport: base,
	}
}

func NewAccessTokenTransport(token *atlasauth.Token, base http.RoundTripper) http.RoundTripper {
	return &tokenTransport{
		token: token,
		base:  base,
	}
}

type tokenTransport struct {
	token *atlasauth.Token
	base  http.RoundTripper
}

func (tr *tokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr.token.SetAuthHeader(req)
	return tr.base.RoundTrip(req)
}
