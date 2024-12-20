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
	"net/http"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"go.mongodb.org/atlas/auth"
)

const (
	cloudGovServiceURL = "https://cloud.mongodbgov.com/"
)

type ServiceGetter interface {
	Service() string
	OpsManagerURL() string
	ClientID() string
}

const (
	ClientID    = "0oabtxactgS3gHIR0297" // ClientID for production
	GovClientID = "0oabtyfelbTBdoucy297" // GovClientID for production
)

func FlowWithConfig(c ServiceGetter, client *http.Client) (*auth.Config, error) {
	id := ClientID
	if c.Service() == config.CloudGovService {
		id = GovClientID
	}
	if c.ClientID() != "" {
		id = c.ClientID()
	}

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
