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

//go:build unit

package commonerrors

import (
	"errors"
	"testing"

	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestCheck(t *testing.T) {
	dummyErr := errors.New("dummy error")

	skderr := &atlasv2.GenericOpenAPIError{}
	skderr.SetModel(atlasv2.ApiError{ErrorCode: "TENANT_CLUSTER_UPDATE_UNSUPPORTED"})

	asymmetricShardErr := &atlasv2.GenericOpenAPIError{}
	asymmetricShardErr.SetModel(atlasv2.ApiError{ErrorCode: asymmetricShardUnsupportedErrorCode})

	testCases := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "nil",
			err:  nil,
			want: nil,
		},
		{
			name: "unsupported cluster update",
			err:  skderr,
			want: errClusterUnsupported,
		},
		{
			name: "arbitrary error",
			err:  dummyErr,
			want: dummyErr,
		},
		{
			name: "asymmetric shard unsupported",
			err:  asymmetricShardErr,
			want: errAsymmetricShardUnsupported,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Check(tc.err); !errors.Is(got, tc.want) {
				t.Errorf("Check(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestGetHumanFriendlyErrorMessage(t *testing.T) {
	dummyErr := errors.New("dummy error")
	testErr := &atlasv2.GenericOpenAPIError{}
	testErr.SetModel(atlasv2.ApiError{Error: 401})

	testCases := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "unauthorized error",
			err:  testErr,
			want: ErrUnauthorized,
		},
		{
			name: "arbitrary error",
			err:  dummyErr,
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := GetHumanFriendlyErrorMessage(tc.err); !errors.Is(got, tc.want) {
				t.Errorf("GetHumanFriendlyErrorMessage(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestGetErrorStatusCode(t *testing.T) {
	dummyErr := errors.New("dummy error")

	unauthorizedCode := 401
	forbiddenCode := 403
	notFoundCode := 404

	atlasErr := &atlas.ErrorResponse{HTTPCode: unauthorizedCode}
	atlasv2Err := &atlasv2.GenericOpenAPIError{}
	atlasv2Err.SetModel(atlasv2.ApiError{Error: forbiddenCode})
	atlasClustersPinnedErr := &atlasClustersPinned.GenericOpenAPIError{}
	atlasClustersPinnedErr.SetModel(atlasClustersPinned.ApiError{Error: &notFoundCode})

	testCases := []struct {
		name string
		err  error
		want int
	}{
		{
			name: "atlas unauthorized error",
			err:  atlasErr,
			want: unauthorizedCode,
		},
		{
			name: "atlasv2 forbidden error",
			err:  atlasv2Err,
			want: forbiddenCode,
		},
		{
			name: "atlasClusterPinned not found error",
			err:  atlasClustersPinnedErr,
			want: notFoundCode,
		},
		{
			name: "arbitrary error",
			err:  dummyErr,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := GetErrorStatusCode(tc.err); got != tc.want {
				t.Errorf("GetErrorStatusCode(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
