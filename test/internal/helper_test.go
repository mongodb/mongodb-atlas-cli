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

package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTransientError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			name: "connection reset by peer as seen against cloud-dev",
			err:  errors.New(`Delete "https://cloud-dev.mongodb.com/api/atlas/v2/groups/123/clusters/c": read tcp 10.128.23.36:60008->3.228.247.77:443: read: connection reset by peer`),
			want: true,
		},
		{
			name: "truncated response",
			err:  errors.New(`Get "https://cloud-dev.mongodb.com/api/atlas/v2/groups/123/clusters/c": EOF`),
			want: true,
		},
		{
			name: "replication lag right after cluster creation",
			err:  errors.New(`HTTP 400 Bad Request (Error code: "OPERATION_INVALID_MEMBER_REPLICATION_LAG")`),
			want: true,
		},
		{
			name: "wrapped transient error",
			err:  fmt.Errorf("error watching cluster: %w", errors.New("read: connection reset by peer")),
			want: true,
		},
		{
			name: "cluster not found is a real answer, not a blip",
			err:  errors.New(`HTTP 404 (Error code: "CLUSTER_NOT_FOUND")`),
			want: false,
		},
		{
			name: "bad request is not retried",
			err:  errors.New("HTTP 400 Bad Request (Error code: \"INVALID_ATTRIBUTE\")"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsTransientError(tt.err))
		})
	}
}

func TestIsClusterGone(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			name: "cluster deleted while we watched it",
			err:  errors.New(`HTTP 404 (Error code: "CLUSTER_NOT_FOUND") Detail: No cluster named c exists`),
			want: true,
		},
		{
			name: "project already closed",
			err:  errors.New(`HTTP 404 (Error code: "GROUP_NOT_FOUND")`),
			want: true,
		},
		{
			name: "wrapped not found",
			err:  fmt.Errorf("failed to delete cluster: %w", errors.New(`(Error code: "CLUSTER_NOT_FOUND")`)),
			want: true,
		},
		{
			name: "cluster still exists but is busy",
			err:  errors.New(`HTTP 409 (Error code: "CANNOT_CLOSE_GROUP_ACTIVE_ATLAS_CLUSTERS")`),
			want: false,
		},
		{
			name: "network failure is not a missing cluster",
			err:  errors.New("read: connection reset by peer"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isClusterGone(tt.err))
		})
	}
}
