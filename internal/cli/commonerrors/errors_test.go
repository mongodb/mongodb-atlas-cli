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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCheck(t *testing.T) {
	dummyErr := errors.New("dummy error")

	skderr := &admin.GenericOpenAPIError{}
	skderr.SetModel(admin.ApiError{ErrorCode: pointer.Get("TENANT_CLUSTER_UPDATE_UNSUPPORTED")})
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Check(tc.err); !errors.Is(got, tc.want) {
				t.Errorf("Check(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
