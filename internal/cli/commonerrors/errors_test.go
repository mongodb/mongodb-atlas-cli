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
	"testing"

	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestCheck(t *testing.T) {
	dummyErr := errors.New("dummy error")

	skderr := &atlasv2.GenericOpenAPIError{}
	skderr.SetModel(atlasv2.ApiError{ErrorCode: tenantClusterUpdateUnsupportedErrorCode})

	asymmetricShardErr := &atlasv2.GenericOpenAPIError{}
	asymmetricShardErr.SetModel(atlasv2.ApiError{ErrorCode: asymmetricShardUnsupportedErrorCode})

	unauthErr := &atlas.ErrorResponse{ErrorCode: unauthorizedErrorCode}

	invalidRefreshTokenErr := &atlas.ErrorResponse{ErrorCode: invalidRefreshTokenErrorCode}

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
		{
			name: "unauthorized error",
			err:  unauthErr,
			want: ErrUnauthorized,
		},
		{
			name: "invalid refresh token error",
			err:  invalidRefreshTokenErr,
			want: ErrInvalidRefreshToken,
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

func TestGetError(t *testing.T) {
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
			if got := getError(tc.err); got != tc.want {
				t.Errorf("GetError(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestGetErrorCode(t *testing.T) {
	dummyErr := errors.New("dummy error")

	atlasErr := &atlas.ErrorResponse{ErrorCode: invalidRefreshTokenErrorCode}
	atlasv2Err := &atlasv2.GenericOpenAPIError{}
	atlasv2Err.SetModel(atlasv2.ApiError{ErrorCode: tenantClusterUpdateUnsupportedErrorCode})
	atlasClustersPinnedErr := &atlasClustersPinned.GenericOpenAPIError{}
	asymmetricCode := asymmetricShardUnsupportedErrorCode
	atlasClustersPinnedErr.SetModel(atlasClustersPinned.ApiError{ErrorCode: &asymmetricCode})

	testCases := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "atlas error",
			err:  atlasErr,
			want: invalidRefreshTokenErrorCode,
		},
		{
			name: "atlasv2 error",
			err:  atlasv2Err,
			want: tenantClusterUpdateUnsupportedErrorCode,
		},
		{
			name: "atlasClusterPinned error",
			err:  atlasClustersPinnedErr,
			want: asymmetricShardUnsupportedErrorCode,
		},
		{
			name: "arbitrary error",
			err:  dummyErr,
			want: unknownErrorCode,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := getErrorCode(tc.err); got != tc.want {
				t.Errorf("GetErrorCode(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
