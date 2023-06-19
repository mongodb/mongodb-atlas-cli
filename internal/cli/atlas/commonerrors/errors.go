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

	"go.mongodb.org/atlas-sdk/admin"
)

var (
	errClusterUnsupported = errors.New("atlas supports this command only for M10+ clusters. You can upgrade your cluster by running the 'atlas cluster upgrade' command")
)

func Check(err error) error {
	if err == nil {
		return nil
	}

	apiError, ok := admin.AsError(err)
	if ok {
		if apiError.GetErrorCode() == "TENANT_CLUSTER_UPDATE_UNSUPPORTED" {
			return errClusterUnsupported
		}
	}
	return err
}
