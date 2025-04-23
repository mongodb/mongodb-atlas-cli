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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_clusters.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ClusterLister,ClusterDescriber,ClusterCreator,ClusterDeleter,ClusterUpdater,AtlasClusterGetterUpdater,ClusterPauser,ClusterStarter,AtlasClusterQuickStarter,SampleDataAdder,SampleDataStatusDescriber,AtlasClusterConfigurationOptionsDescriber,AtlasSharedClusterDescriber,ClusterUpgrader,AtlasSharedClusterGetterUpgrader,AtlasClusterConfigurationOptionsUpdater,ClusterTester,ClusterDescriberStarter

type ClusterLister interface {
	ProjectClusters(string, *ListOptions) (*atlasClustersPinned.PaginatedAdvancedClusterDescription, error)
	ListFlexClusters(*atlasv2.ListFlexClustersApiParams) (*atlasv2.PaginatedFlexClusters20241113, error)
}

type ClusterDescriber interface {
	AtlasCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
	FlexCluster(string, string) (*atlasv2.FlexClusterDescription20241113, error)
}

type ClusterDescriberStarter interface {
	ClusterDescriber
	ClusterStarter
}

type AtlasClusterConfigurationOptionsDescriber interface {
	AtlasClusterConfigurationOptions(string, string) (*atlasClustersPinned.ClusterDescriptionProcessArgs, error)
}

type AtlasClusterConfigurationOptionsUpdater interface {
	UpdateAtlasClusterConfigurationOptions(string, string, *atlasClustersPinned.ClusterDescriptionProcessArgs) (*atlasClustersPinned.ClusterDescriptionProcessArgs, error)
}

type AtlasSharedClusterDescriber interface {
	AtlasSharedCluster(string, string) (*atlas.Cluster, error)
}

type ClusterCreator interface {
	CreateCluster(v15 *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error)
	CreateFlexCluster(string, *atlasv2.FlexClusterDescriptionCreate20241113) (*atlasv2.FlexClusterDescription20241113, error)
}

type ClusterDeleter interface {
	DeleteCluster(string, string) error
	DeleteFlexCluster(string, string) error
}

type ClusterUpdater interface {
	UpdateCluster(string, string, *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error)
	UpdateFlexCluster(string, string, *atlasv2.FlexClusterDescriptionUpdate20241113) (*atlasv2.FlexClusterDescription20241113, error)
}

type ClusterPauser interface {
	PauseCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
}

type ClusterStarter interface {
	StartCluster(string, string) (*atlasClustersPinned.AdvancedClusterDescription, error)
}

type ClusterUpgrader interface {
	UpgradeCluster(string, *atlas.Cluster) (*atlas.Cluster, error)
	UpgradeFlexCluster(string, *atlasv2.AtlasTenantClusterUpgradeRequest20240805) (*atlasv2.FlexClusterDescription20241113, error)
}

type SampleDataAdder interface {
	AddSampleData(string, string) (*atlasv2.SampleDatasetStatus, error)
}

type SampleDataStatusDescriber interface {
	SampleDataStatus(string, string) (*atlasv2.SampleDatasetStatus, error)
}

type ClusterTester interface {
	TestClusterFailover(string, string) error
}

type AtlasClusterGetterUpdater interface {
	ClusterDescriber
	ClusterUpdater
}

type AtlasSharedClusterGetterUpgrader interface {
	AtlasSharedClusterDescriber
	ClusterDescriber
	ClusterUpgrader
}

type AtlasClusterQuickStarter interface {
	SampleDataAdder
	SampleDataStatusDescriber
	CloudProviderRegionsLister
	ClusterLister
	DatabaseUserCreator
	DatabaseUserDescriber
	ProjectIPAccessListCreator
	ClusterDescriber
	ClusterCreator
	ProjectMDBVersionLister
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

// UpdateCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateCluster(projectID, name string, cluster *atlasClustersPinned.AdvancedClusterDescription) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	result, _, err := s.clientClusters.ClustersApi.UpdateCluster(s.ctx, projectID, name, cluster).Execute()
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

// StartCluster encapsulate the logic to manage different cloud providers.
func (s *Store) StartCluster(projectID, name string) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	paused := false
	cluster := &atlasClustersPinned.AdvancedClusterDescription{
		Paused: &paused,
	}
	return s.UpdateCluster(projectID, name, cluster)
}

// DeleteCluster encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteCluster(projectID, name string) error {
	_, err := s.clientClusters.ClustersApi.DeleteCluster(s.ctx, projectID, name).Execute()
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

// AtlasCluster encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasCluster(projectID, name string) (*atlasClustersPinned.AdvancedClusterDescription, error) {
	result, _, err := s.clientClusters.ClustersApi.GetCluster(s.ctx, projectID, name).Execute()
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
