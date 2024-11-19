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
	"net/url"
	"time"

	"github.com/rapid7/go-get-proxied/proxy"
)

const (
	timeout               = 5 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleConns          = 5
	maxIdleConnsPerHost   = 4
	idleConnTimeout       = 30 * time.Second
	expectContinueTimeout = 1 * time.Second
	CloudGovServiceURL    = "https://cloud.mongodbgov.com/"
	telemetryTimeout      = 1 * time.Second
)

func proxyFromSettingsAndEnv(req *http.Request) (*url.URL, error) {
	switch req.URL.Scheme {
	case "http":
		p := proxy.NewProvider("").GetHTTPProxy(req.URL.String())
		if p != nil {
			return p.URL(), nil
		}
	case "https":
		p := proxy.NewProvider("").GetHTTPSProxy(req.URL.String())
		if p != nil {
			return p.URL(), nil
		}
	}
	return nil, nil
}

var DefaultTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   timeout,
		KeepAlive: keepAlive,
	}).DialContext,
	MaxIdleConns:          maxIdleConns,
	MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	Proxy:                 proxyFromSettingsAndEnv,
	IdleConnTimeout:       idleConnTimeout,
	ExpectContinueTimeout: expectContinueTimeout,
}

var TelemetryTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   telemetryTimeout,
		KeepAlive: keepAlive,
	}).DialContext,
	MaxIdleConns:          maxIdleConns,
	MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	Proxy:                 proxyFromSettingsAndEnv,
	IdleConnTimeout:       idleConnTimeout,
	ExpectContinueTimeout: expectContinueTimeout,
}
