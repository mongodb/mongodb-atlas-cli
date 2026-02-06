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

package commonerrors

import (
	"errors"
	"net/http"

	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"golang.org/x/oauth2"
)

var (
	errClusterUnsupported         = errors.New("atlas supports this command only for M10+ clusters. You can upgrade your cluster by running the 'atlas cluster upgrade' command")
	errOutsideVPN                 = errors.New("forbidden action outside access allow list, if you are a MongoDB employee double check your VPN connection")
	errAsymmetricShardUnsupported = errors.New("trying to run a cluster wide scaling command on an independent shard scaling cluster. Use --autoScalingMode 'independentShardScaling' instead")
	ErrUnauthorized               = errors.New(`unauthorized

To log in using your Atlas username and password or to set credentials using a Service account or API keys, run: atlas auth login`)
	ErrInvalidRefreshToken = errors.New(`session expired

Please note that your session expires periodically. 
If you use Atlas CLI for automation, see https://www.mongodb.com/docs/atlas/cli/stable/atlas-cli-automate/ for best practices.
To login, run: atlas auth login`)
)

const (
	unknownErrorCode                        = "UNKNOWN_ERROR"
	asymmetricShardUnsupportedErrorCode     = "ASYMMETRIC_SHARD_UNSUPPORTED"
	tenantClusterUpdateUnsupportedErrorCode = "TENANT_CLUSTER_UPDATE_UNSUPPORTED"
	globalUserOutsideSubnetErrorCode        = "GLOBAL_USER_OUTSIDE_SUBNET"
	unauthorizedErrorCode                   = "UNAUTHORIZED"
	invalidRefreshTokenErrorCode            = "INVALID_REFRESH_TOKEN"
	invalidServiceAccountClient             = "invalid_client"
)

// Check checks the error and returns a more user-friendly error message if applicable.
func Check(err error) error {
	if err == nil {
		return nil
	}

	apiErrorCode := getErrorCode(err)

	switch apiErrorCode {
	case unauthorizedErrorCode:
		return ErrUnauthorized
	case invalidRefreshTokenErrorCode:
		return ErrInvalidRefreshToken
	case tenantClusterUpdateUnsupportedErrorCode:
		return errClusterUnsupported
	case globalUserOutsideSubnetErrorCode:
		return errOutsideVPN
	case asymmetricShardUnsupportedErrorCode:
		return errAsymmetricShardUnsupported
	case invalidServiceAccountClient: // oauth2 error
		return ErrUnauthorized
	}

	apiError := getError(err) // some `Unauthorized` errors do not have an error code, so we check the HTTP status code

	if apiError == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	return err
}

// getErrorCode extracts the error code from the error if it is an Atlas error.
// This function checks for v2 SDK, the pinned clusters SDK, the old SDK errors
// and oauth2 errors.
// If the error is not any of these errors, it returns "UNKNOWN_ERROR".
func getErrorCode(err error) string {
	if err == nil {
		return unknownErrorCode
	}

	var atlasErr *atlas.ErrorResponse
	if errors.As(err, &atlasErr) {
		return atlasErr.ErrorCode
	}
	if sdkError, ok := atlasv2.AsError(err); ok {
		return sdkError.ErrorCode
	}
	if sdkPinnedError, ok := atlasClustersPinned.AsError(err); ok {
		return sdkPinnedError.GetErrorCode()
	}
	var oauth2Err *oauth2.RetrieveError
	if errors.As(err, &oauth2Err) {
		return oauth2Err.ErrorCode
	}

	return unknownErrorCode
}

// getError extracts the HTTP error code from the error if it is an Atlas error.
// This function checks for v2 SDK, the pinned clusters SDK and the old SDK errors.
// If the error is not any of these Atlas errors, it returns 0.
func getError(err error) int {
	if err == nil {
		return 0
	}

	var atlasErr *atlas.ErrorResponse
	if errors.As(err, &atlasErr) {
		return atlasErr.HTTPCode
	}
	if apiError, ok := atlasv2.AsError(err); ok {
		return apiError.GetError()
	}
	if apiPinnedError, ok := atlasClustersPinned.AsError(err); ok {
		return apiPinnedError.GetError()
	}

	return 0
}

func IsAsymmetricShardUnsupported(err error) bool {
	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return false
	}
	return apiError.GetErrorCode() == asymmetricShardUnsupportedErrorCode
}

func IsCannotUseFlexWithClusterApis(err error) bool {
	apiError, ok := atlasv2.AsError(err)
	if !ok {
		return false
	}
	return apiError.GetErrorCode() == "CANNOT_USE_FLEX_CLUSTER_IN_CLUSTER_API"
}

func IsInvalidRefreshToken(err error) bool {
	errCode := getErrorCode(err)
	return errCode == invalidRefreshTokenErrorCode
}
