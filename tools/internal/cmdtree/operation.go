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

package cmdtree

import (
	"strconv"
	"strings"
)

const (
	// apiRequestSuffix is what the Atlas SDK appends to an operation id to name
	// the type carrying the request for it.
	apiRequestSuffix = "ApiRequest"
	// executeMethod is the method that performs the request.
	executeMethod = "Execute"
	// versionLength is the "v" plus the date of an SDK module version.
	versionLength = len("v20060102")
)

// ApiOperation is one Atlas SDK operation a command calls.
type ApiOperation struct {
	Version     string `json:"version"`
	OperationID string `json:"operationId"`
}

// apiOperation parses a canonical SDK symbol into the API version and the
// operation it performs. It reports false for any SDK symbol that is not an
// ApiRequest.Execute call, which is every client option, error helper and
// generated getter the walk also passes through.
//
//	go.mongodb.org/atlas-sdk/v20240530005/admin.CreateClusterApiRequest.Execute
//	  -> {v20240530, CreateCluster}
func apiOperation(name string) (ApiOperation, bool) {
	rest, ok := strings.CutPrefix(name, sdkPrefix)
	if !ok {
		return ApiOperation{}, false
	}

	module, rest, ok := strings.Cut(rest, qualifiedSep)
	if !ok || !isVersion(module) {
		return ApiOperation{}, false
	}

	dot := strings.LastIndex(rest, typeSep)
	if dot < 0 || rest[dot+1:] != executeMethod {
		return ApiOperation{}, false
	}

	receiver := rest[:dot]
	if slash := strings.LastIndex(receiver, qualifiedSep); slash >= 0 {
		receiver = receiver[slash+1:]
	}
	if dot := strings.LastIndex(receiver, typeSep); dot >= 0 {
		receiver = receiver[dot+1:]
	}
	receiver = strings.TrimSuffix(strings.TrimPrefix(receiver, pointerReceiver), ")")

	operationID, ok := strings.CutSuffix(receiver, apiRequestSuffix)
	if !ok || operationID == "" {
		return ApiOperation{}, false
	}
	return ApiOperation{Version: module[:versionLength], OperationID: operationID}, true
}

// isVersion reports whether module names an SDK version, like v20240530005.
// Only the date part of it identifies the API version.
func isVersion(module string) bool {
	if len(module) < versionLength || !strings.HasPrefix(module, "v") {
		return false
	}
	_, err := strconv.Atoi(module[1:])
	return err == nil
}
