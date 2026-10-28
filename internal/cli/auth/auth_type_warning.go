// Copyright 2026 MongoDB Inc
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

package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
)

const (
	missingAuthTypeWarning = "Warning: an access and refresh token are set but no auth type is, so they are not used. " +
		"Set MONGODB_ATLAS_AUTH_TYPE to user_account, or user_account_legacy for device-code tokens.\n"
	deviceCodeTokenWarning = "Warning: the access token was not issued by the User Account authorization server, " +
		"so refreshing it will fail. If it came from the device-code flow, set MONGODB_ATLAS_AUTH_TYPE=user_account_legacy.\n"
)

// WarnOnTokenAuthTypeMismatch warns when the profile's tokens cannot work with its auth type.
func WarnOnTokenAuthTypeMismatch(profile *config.Profile) {
	if profile.AccessToken() == "" || profile.RefreshToken() == "" {
		return
	}
	switch profile.AuthType() {
	case "":
		_, _ = log.Warning(missingAuthTypeWarning)
	case config.UserAccount:
		issuer, ok := accessTokenIssuer(profile.AccessToken())
		if !ok {
			return
		}
		if expected, ok := expectedAuthServerIssuer(profile); ok && issuer != expected {
			_, _ = log.Warning(deviceCodeTokenWarning)
		}
	case config.UserAccountLegacy, config.ServiceAccount, config.APIKeys, config.NoAuth:
	}
}

// accessTokenIssuer reads iss without verifying the signature; it only selects a warning.
func accessTokenIssuer(accessToken string) (string, bool) {
	claims := jwt.RegisteredClaims{}
	if _, _, err := jwt.NewParser().ParseUnverified(accessToken, &claims); err != nil || claims.Issuer == "" {
		return "", false
	}
	return claims.Issuer, true
}

// expectedAuthServerIssuer resolves the issuer refresh would use, so gov and
// MONGODB_ATLAS_AUTH_SERVER_URL are handled.
func expectedAuthServerIssuer(profile *config.Profile) (string, bool) {
	flow, err := transport.FlowForAuthServer(profile, http.DefaultClient, version.Version)
	if err != nil {
		return "", false
	}
	return flow.AuthServerURL.String(), true
}
