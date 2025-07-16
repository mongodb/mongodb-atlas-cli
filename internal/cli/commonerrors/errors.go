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
	"strings"

	"go.mongodb.org/atlas-sdk/v20250312005/admin"
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

	apiError, ok := admin.AsError(err)
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

func CheckHTTPErrors(err error) error {
	if strings.Contains(err.Error(), "401") && strings.Contains(err.Error(), "Unauthorized") {
		return ErrUnauthorized
	}
	return nil
}

func IsAsymmetricShardUnsupported(err error) bool {
	apiError, ok := admin.AsError(err)
	if !ok {
		return false
	}
	return apiError.GetErrorCode() == asymmetricShardUnsupportedErrorCode
}

func IsCannotUseFlexWithClusterApis(err error) bool {
	apiError, ok := admin.AsError(err)
	if !ok {
		return false
	}
	return apiError.GetErrorCode() == "CANNOT_USE_FLEX_CLUSTER_IN_CLUSTER_API"
}
