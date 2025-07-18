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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var (
	errClusterUnsupported         = errors.New("atlas supports this command only for M10+ clusters. You can upgrade your cluster by running the 'atlas cluster upgrade' command")
	errOutsideVPN                 = errors.New("forbidden action outside access allow list, if you are a MongoDB employee double check your VPN connection")
	errAsymmetricShardUnsupported = errors.New("trying to run a cluster wide scaling command on an independent shard scaling cluster. Use --autoScalingMode 'independentShardScaling' instead")
	ErrUnauthorized               = errors.New(`this action requires authentication
	
To log in using your Atlas username and password, run: atlas auth login
To set credentials using API keys, run: atlas config init`)
)

const (
	asymmetricShardUnsupportedErrorCode = "ASYMMETRIC_SHARD_UNSUPPORTED"
)

func Check(err error) error {
	if err == nil {
		return nil
	}

	apiError, ok := atlasv2.AsError(err)
	if ok {
		switch apiError.GetErrorCode() {
		case "TENANT_CLUSTER_UPDATE_UNSUPPORTED":
			return errClusterUnsupported
		case "GLOBAL_USER_OUTSIDE_SUBNET":
			return errOutsideVPN
		case asymmetricShardUnsupportedErrorCode:
			return errAsymmetricShardUnsupported
		}
	}
	return err
}

// GetHumanFriendErrorMessage returns a human-friendly error message if error is status code 401.
func GetHumanFriendlyErrorMessage(err error) error {
	if err == nil {
		return nil
	}

	statusCode := GetErrorStatusCode(err)
	if statusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	}
	return nil
}

// GetErrorStatusCode returns the HTTP status code from the error.
// It checks for v2 SDK, the pinned clusters SDK and the old SDK errors.
// If the error is not an API error or is nil, it returns 0.
func GetErrorStatusCode(err error) int {
	if err == nil {
		return 0
	}
	apiError, ok := atlasv2.AsError(err)
	if ok {
		return apiError.GetError()
	}
	apiPinnedError, ok := atlasClustersPinned.AsError(err)
	if ok {
		return apiPinnedError.GetError()
	}
	var atlasErr *atlas.ErrorResponse
	if errors.As(err, &atlasErr) {
		return atlasErr.HTTPCode
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
