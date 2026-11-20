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

import "testing"

func TestApiOperation(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want ApiOperation
		ok   bool
	}{
		{
			name: "value receiver request",
			in:   "go.mongodb.org/atlas-sdk/v20240530005/admin.CreateClusterApiRequest.Execute",
			want: ApiOperation{Version: "v20240530", OperationID: "CreateCluster"},
			ok:   true,
		},
		{
			name: "same operation from a newer version",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.CreateClusterApiRequest.Execute",
			want: ApiOperation{Version: "v20250312", OperationID: "CreateCluster"},
			ok:   true,
		},
		{
			name: "pointer receiver request",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.(*DeleteAccessListEntryApiRequest).Execute",
			want: ApiOperation{Version: "v20250312", OperationID: "DeleteAccessListEntry"},
			ok:   true,
		},
		{
			name: "nested package",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/auth/code.(*SomeApiRequest).Execute",
			want: ApiOperation{Version: "v20250312", OperationID: "Some"},
			ok:   true,
		},
		{
			name: "client option",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.UseBaseURL",
			ok:   false,
		},
		{
			name: "error helper",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.AsError",
			ok:   false,
		},
		{
			name: "request parameter setter",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.ListClustersApiRequest.ItemsPerPage",
			ok:   false,
		},
		{
			name: "generated getter",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.(*ClusterDescription20240805).GetName",
			ok:   false,
		},
		{
			name: "package with no request type",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/detectaiagent.Detect",
			ok:   false,
		},
		{
			name: "not an sdk symbol",
			in:   "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store.Store).Projects",
			ok:   false,
		},
		{
			name: "no version in the module path",
			in:   "go.mongodb.org/atlas-sdk/admin.CreateClusterApiRequest.Execute",
			ok:   false,
		},
		{
			name: "request type with no operation id",
			in:   "go.mongodb.org/atlas-sdk/v20250312027/admin.ApiRequest.Execute",
			ok:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := apiOperation(tt.in)
			if ok != tt.ok {
				t.Fatalf("apiOperation(%q) ok = %v, want %v", tt.in, ok, tt.ok)
			}
			if got != tt.want {
				t.Errorf("apiOperation(%q) = %+v, want %+v", tt.in, got, tt.want)
			}
		})
	}
}
