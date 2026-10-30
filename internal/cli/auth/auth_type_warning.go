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
	missingAuthTypeWarning = "Warning: tokens are ignored without an auth type. " +
		"Set MONGODB_ATLAS_AUTH_TYPE to user_account, or user_account_legacy for device-code tokens.\n"
	deviceCodeTokenWarning = "Warning: these tokens are not from the User Account auth server and will fail to refresh. " +
		"For device-code tokens, set MONGODB_ATLAS_AUTH_TYPE=user_account_legacy.\n"
	authCodeTokenWarning = "Warning: these tokens are from the User Account auth server and will fail to refresh with user_account_legacy. " +
		"Set MONGODB_ATLAS_AUTH_TYPE=user_account.\n"
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
		if fromAuthServer, ok := issuedByAuthServer(profile); ok && !fromAuthServer {
			_, _ = log.Warning(deviceCodeTokenWarning)
		}
	case config.UserAccountLegacy:
		if fromAuthServer, ok := issuedByAuthServer(profile); ok && fromAuthServer {
			_, _ = log.Warning(authCodeTokenWarning)
		}
	case config.ServiceAccount, config.APIKeys, config.NoAuth:
	}
}

// issuedByAuthServer reports whether the access token's iss is the Authorization Server; ok is false when unknown.
func issuedByAuthServer(profile *config.Profile) (fromAuthServer, ok bool) {
	issuer, ok := accessTokenIssuer(profile.AccessToken())
	if !ok {
		return false, false
	}
	expected, ok := expectedAuthServerIssuer(profile)
	if !ok {
		return false, false
	}
	return issuer == expected, true
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
