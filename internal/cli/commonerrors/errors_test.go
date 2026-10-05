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
	"slices"
	"strings"
	"testing"

	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312025/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"golang.org/x/oauth2"
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

func TestServiceAccountEnvVars(t *testing.T) {
	testCases := []struct {
		name    string
		environ []string
		want    []string
	}{
		{
			name:    "no env vars",
			environ: []string{"HOME=/home/user"},
			want:    nil,
		},
		{
			name:    "atlas prefix",
			environ: []string{"MONGODB_ATLAS_CLIENT_ID=id", "MONGODB_ATLAS_CLIENT_SECRET=secret"},
			want:    []string{"MONGODB_ATLAS_CLIENT_ID", "MONGODB_ATLAS_CLIENT_SECRET"},
		},
		{
			name:    "mcli prefix takes over",
			environ: []string{"MONGODB_ATLAS_CLIENT_ID=id", "MCLI_CLIENT_ID=id"},
			want:    []string{"MCLI_CLIENT_ID"},
		},
		{
			name:    "any mcli var switches the prefix",
			environ: []string{"MCLI_OPS_MANAGER_URL=https://cloud-dev.mongodb.com/", "MONGODB_ATLAS_CLIENT_ID=id"},
			want:    nil,
		},
		{
			name:    "empty values are ignored",
			environ: []string{"MONGODB_ATLAS_CLIENT_ID=", "MONGODB_ATLAS_CLIENT_SECRET=secret"},
			want:    []string{"MONGODB_ATLAS_CLIENT_SECRET"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := serviceAccountEnvVars(tc.environ); !slices.Equal(got, tc.want) {
				t.Errorf("serviceAccountEnvVars(%v) = %v, want %v", tc.environ, got, tc.want)
			}
		})
	}
}

func TestCheck_invalidClient(t *testing.T) {
	invalidClientErr := &oauth2.RetrieveError{ErrorCode: invalidServiceAccountClient}

	t.Run("with service account env vars set", func(t *testing.T) {
		t.Setenv("MCLI_CLIENT_ID", "id")
		t.Setenv("MCLI_CLIENT_SECRET", "")

		got := Check(invalidClientErr)
		if !errors.Is(got, ErrUnauthorized) {
			t.Fatalf("Check() = %v, want ErrUnauthorized", got)
		}
		if !strings.Contains(got.Error(), "MCLI_CLIENT_ID") || !strings.Contains(got.Error(), EnvVarsDocsURL) {
			t.Errorf("Check() = %v, want the env var name and docs URL in the message", got)
		}
	})

	t.Run("without service account env vars", func(t *testing.T) {
		// setting MCLI vars to empty also pins the prefix detection to MCLI_,
		// keeping the result deterministic on hosts with MONGODB_ATLAS_* vars set
		t.Setenv("MCLI_CLIENT_ID", "")
		t.Setenv("MCLI_CLIENT_SECRET", "")

		if got := Check(invalidClientErr); !errors.Is(got, ErrUnauthorized) || got.Error() != ErrUnauthorized.Error() {
			t.Errorf("Check() = %v, want exactly ErrUnauthorized", got)
		}
	})
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
