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

package store

import (
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate go tool go.uber.org/mock/mockgen -destination=../mocks/mock_clusters.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ClusterDescriber

type ClusterDescriber interface { //nolint:iface // right now requires some refactor to deployment commands
	AtlasCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	FlexCluster(string, string) (*atlasv2.FlexClusterDescription20241113, error)
	LatestAtlasCluster(string, string) (*atlasv2.ClusterDescription20240805, error)
}

// AddSampleData encapsulate the logic to manage different cloud providers.
func (s *Store) AddSampleData(groupID, clusterName string) (*atlasv2.SampleDatasetStatus, error) {
	result, _, err := s.clientv2.ClustersApi.LoadSampleDataset(s.ctx, groupID, clusterName).Execute()
	return result, err
}

// SampleDataStatus encapsulate the logic to manage different cloud providers.
func (s *Store) SampleDataStatus(groupID, id string) (*atlasv2.SampleDatasetStatus, error) {
	result, _, err := s.clientv2.ClustersApi.GetSampleDatasetLoadStatus(s.ctx, groupID, id).Execute()
	return result, err
}

// CreateCluster encapsulate the logic to manage different cloud providers.
func (s *Store) CreateCluster(cluster *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	result, _, err := s.clientClusters.ClustersApi.CreateCluster(s.ctx, cluster.GetGroupId(), cluster).Execute()
	return result, err
}

// CreateClusterLatest uses the latest API version to create a cluster.
func (s *Store) CreateClusterLatest(cluster *atlasv2.ClusterDescription20240805) (*atlasv2.ClusterDescription20240805, error) {
	result, _, err := s.clientv2.ClustersApi.CreateCluster(s.ctx, cluster.GetGroupId(), cluster).Execute()
	return result, err
}

// UpdateCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateCluster(projectID, name string, cluster *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	result, _, err := s.clientClusters.ClustersApi.UpdateCluster(s.ctx, projectID, name, cluster).Execute()
	return result, err
}

// UpdateClusterLatest uses the latest API version to update a cluster.
func (s *Store) UpdateClusterLatest(projectID, name string, cluster *atlasv2.ClusterDescription20240805) (*atlasv2.ClusterDescription20240805, error) {
	result, _, err := s.clientv2.ClustersApi.UpdateCluster(s.ctx, projectID, name, cluster).Execute()
	return result, err
}

// PauseCluster encapsulate the logic to manage different cloud providers.
func (s *Store) PauseCluster(projectID, name string) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	paused := true
	cluster := &atlasClustersPinned.AdvancedClusterDescription{
		Paused: &paused,
	}
	return s.UpdateCluster(projectID, name, cluster)
}

// PauseClusterLatest uses the latest API version to pause a cluster.
func (s *Store) PauseClusterLatest(projectID, name string) (*atlasv2.ClusterDescription20240805, error) {
	paused := true
	cluster := &atlasv2.ClusterDescription20240805{
		Paused: &paused,
	}
	return s.UpdateClusterLatest(projectID, name, cluster)
}

// StartCluster encapsulate the logic to manage different cloud providers.
func (s *Store) StartCluster(projectID, name string) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	paused := false
	cluster := &atlasClustersPinned.AdvancedClusterDescription{
		Paused: &paused,
	}
	return s.UpdateCluster(projectID, name, cluster)
}

// StartClusterLatest uses the latest API version to start a cluster.
func (s *Store) StartClusterLatest(projectID, name string) (*atlasv2.ClusterDescription20240805, error) {
	paused := false
	cluster := &atlasv2.ClusterDescription20240805{
		Paused: &paused,
	}
	return s.UpdateClusterLatest(projectID, name, cluster)
}

// GetClusterAutoScalingConfig uses the latest API version to get the auto scaling configuration of a cluster.
func (s *Store) GetClusterAutoScalingConfig(projectID, name string) (*atlasv2.ClusterDescriptionAutoScalingModeConfiguration, error) {
	result, _, err := s.clientv2.ClustersApi.AutoScalingConfiguration(s.ctx, projectID, name).Execute()
	return result, err
}

// DeleteCluster encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteCluster(projectID, name string) error {
	_, err := s.clientv2.ClustersApi.DeleteCluster(s.ctx, projectID, name).Execute()
	return err
}

// AtlasSharedCluster encapsulates the logic to fetch details of one shared cluster.
func (s *Store) AtlasSharedCluster(projectID, name string) (*atlas.Cluster, error) {
	result, _, err := s.client.Clusters.Get(s.ctx, projectID, name)
	return result, err
}

// UpgradeCluster encapsulate the logic to upgrade shared clusters in a project.
func (s *Store) UpgradeCluster(projectID string, cluster *atlas.Cluster) (*atlas.Cluster, error) {
	result, _, err := s.client.Clusters.Upgrade(s.ctx, projectID, cluster)
	return result, err
}

// ProjectClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectClusters(projectID string, opts *ListOptions) (*atlasClustersPinned.PaginatedAdvancedClusterDescription, error) {
	res := s.clientClusters.ClustersApi.ListClusters(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// LatestProjectClusters lists the clusters using the latest API version.
func (s *Store) LatestProjectClusters(projectID string, opts *ListOptions) (*atlasv2.PaginatedClusterDescription20240805, error) {
	res := s.clientv2.ClustersApi.ListClusters(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// AtlasCluster encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasCluster(projectID, name string) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	result, _, err := s.clientClusters.ClustersApi.GetCluster(s.ctx, projectID, name).Execute()
	return result, err
}

// LatestAtlasCluster uses the latest API version to get a cluster.
func (s *Store) LatestAtlasCluster(projectID, name string) (*atlasv2.ClusterDescription20240805, error) {
	result, _, err := s.clientv2.ClustersApi.GetCluster(s.ctx, projectID, name).Execute()
	return result, err
}

// AtlasClusterConfigurationOptions encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasClusterConfigurationOptions(projectID, name string) (*atlasClustersPinned.ClusterDescriptionProcessArgs, error) {
	result, _, err := s.clientClusters.ClustersApi.GetClusterAdvancedConfiguration(s.ctx, projectID, name).Execute()
	return result, err
}

// UpdateAtlasClusterConfigurationOptions encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateAtlasClusterConfigurationOptions(projectID, clusterName string, args *atlasClustersPinned.ClusterDescriptionProcessArgs) (*atlasClustersPinned.ClusterDescriptionProcessArgs, error) {
	result, _, err := s.clientClusters.ClustersApi.UpdateClusterAdvancedConfiguration(s.ctx, projectID, clusterName, args).Execute()
	return result, err
}

func (s *Store) TestClusterFailover(projectID, clusterName string) error {
	_, err := s.clientv2.ClustersApi.TestFailover(s.ctx, projectID, clusterName).Execute()
	return err
}
