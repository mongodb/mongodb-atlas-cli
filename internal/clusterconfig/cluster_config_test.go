// Copyright 2020 MongoDB Inc
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

package clusterconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312003/admin"
)

func TestSetTags(t *testing.T) {
	cases := []struct {
		name         string
		cluster      *atlasClustersPinned.AdvancedClusterDescription
		providedTags map[string]string
		expectedTags *[]atlasClustersPinned.ResourceTag
	}{
		{
			name:         "pinned cluster with no tags",
			cluster:      &atlasClustersPinned.AdvancedClusterDescription{},
			providedTags: map[string]string{"key2": "value2"},
			expectedTags: &[]atlasClustersPinned.ResourceTag{
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "pinned cluster with provided tags",
			cluster: &atlasClustersPinned.AdvancedClusterDescription{
				Tags: &[]atlasClustersPinned.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"key2": "value2"},
			expectedTags: &[]atlasClustersPinned.ResourceTag{
				{Key: "key", Value: "value"},
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "pinned cluster with provided tags and empty key",
			cluster: &atlasClustersPinned.AdvancedClusterDescription{
				Tags: &[]atlasClustersPinned.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"": "value2"},
			expectedTags: &[]atlasClustersPinned.ResourceTag{
				{Key: "key", Value: "value"},
			},
		},
		{
			name: "pinned cluster with provided tags and empty value",
			cluster: &atlasClustersPinned.AdvancedClusterDescription{
				Tags: &[]atlasClustersPinned.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"key": ""},
			expectedTags: &[]atlasClustersPinned.ResourceTag{
				{Key: "key", Value: "value"},
			},
		},
		{
			name: "pinned cluster with provided tags and empty key and value",
			cluster: &atlasClustersPinned.AdvancedClusterDescription{
				Tags: &[]atlasClustersPinned.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"": ""},
			expectedTags: &[]atlasClustersPinned.ResourceTag{
				{Key: "key", Value: "value"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			SetTags(tc.cluster, tc.providedTags)
			assert.Equal(t, tc.expectedTags, tc.cluster.Tags)
		})
	}
}

func TestSetTagsLatest(t *testing.T) {
	cases := []struct {
		name         string
		cluster      *atlasv2.ClusterDescription20240805
		providedTags map[string]string
		expectedTags *[]atlasv2.ResourceTag
	}{
		{
			name:         "latest cluster with no tags",
			cluster:      &atlasv2.ClusterDescription20240805{},
			providedTags: map[string]string{"key2": "value2"},
			expectedTags: &[]atlasv2.ResourceTag{
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "latest cluster with provided tags",
			cluster: &atlasv2.ClusterDescription20240805{
				Tags: &[]atlasv2.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"key2": "value2"},
			expectedTags: &[]atlasv2.ResourceTag{
				{Key: "key", Value: "value"},
				{Key: "key2", Value: "value2"},
			},
		},
		{
			name: "latest cluster with provided tags and empty key",
			cluster: &atlasv2.ClusterDescription20240805{
				Tags: &[]atlasv2.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"": "value2"},
			expectedTags: &[]atlasv2.ResourceTag{
				{Key: "key", Value: "value"},
			},
		},
		{
			name: "latest cluster with provided tags and empty value",
			cluster: &atlasv2.ClusterDescription20240805{
				Tags: &[]atlasv2.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"key": ""},
			expectedTags: &[]atlasv2.ResourceTag{
				{Key: "key", Value: "value"},
			},
		},
		{
			name: "latest cluster with provided tags and empty key and value",
			cluster: &atlasv2.ClusterDescription20240805{
				Tags: &[]atlasv2.ResourceTag{
					{Key: "key", Value: "value"},
				},
			},
			providedTags: map[string]string{"": ""},
			expectedTags: &[]atlasv2.ResourceTag{
				{Key: "key", Value: "value"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			SetTagsLatest(tc.cluster, tc.providedTags)
			assert.Equal(t, tc.expectedTags, tc.cluster.Tags)
		})
	}
}
