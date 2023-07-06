// Copyright 2022 MongoDB Inc
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

package oauth

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"go.mongodb.org/atlas/auth"
)

const (
	timeout                   = 5 * time.Second
	keepAlive                 = 30 * time.Second
	maxIdleConns              = 5
	maxIdleConnsPerHost       = 4
	idleConnTimeout           = 30 * time.Second
	expectContinueTimeout     = 1 * time.Second
	cloudGovServiceURL        = "https://cloud.mongodbgov.com/"
	userAgentContainerPostfix = "container"
	userAgentNativePostfix    = "native"
)

var userAgentPostfix = userAgentNativePostfix

var defaultTransport = &http.Transport{
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

type ServiceGetter interface {
	Service() string
	OpsManagerURL() string
	ClientID() string
}

const (
	ClientID    = "0oabtxactgS3gHIR0297" // ClientID for production
	GovClientID = "0oabtyfelbTBdoucy297" // GovClientID for production
)

func addUserAgentHostParameter() {
	if !strings.Contains(config.UserAgent, userAgentPostfix) {
		config.UserAgent = fmt.Sprintf("%s;%s)",
			strings.TrimSuffix(config.UserAgent, ")"), userAgentPostfix)
	}
}

func FlowWithConfig(c ServiceGetter) (*auth.Config, error) {
	client := http.DefaultClient
	client.Transport = defaultTransport
	id := ClientID
	if c.Service() == config.CloudGovService {
		id = GovClientID
	}
	if c.ClientID() != "" {
		id = c.ClientID()
	}
	if _, runningFromContainer := os.LookupEnv("CONTAINER"); runningFromContainer {
		userAgentPostfix = userAgentContainerPostfix
	}
	addUserAgentHostParameter()

	authOpts := []auth.ConfigOpt{
		auth.SetUserAgent(config.UserAgent),
		auth.SetClientID(id),
		auth.SetScopes([]string{"openid", "profile", "offline_access"}),
	}
	if configURL := c.OpsManagerURL(); configURL != "" {
		authOpts = append(authOpts, auth.SetAuthURL(c.OpsManagerURL()))
	} else if c.Service() == config.CloudGovService {
		authOpts = append(authOpts, auth.SetAuthURL(cloudGovServiceURL))
	}
	return auth.NewConfigWithOptions(client, authOpts...)
}
