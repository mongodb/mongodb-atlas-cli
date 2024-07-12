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
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_clusters.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ClusterLister,ClusterDescriber,ClusterCreator,ClusterDeleter,ClusterUpdater,AtlasClusterGetterUpdater,ClusterPauser,ClusterStarter,AtlasClusterQuickStarter,SampleDataAdder,SampleDataStatusDescriber,AtlasClusterConfigurationOptionsDescriber,AtlasSharedClusterDescriber,ClusterUpgrader,AtlasSharedClusterGetterUpgrader,AtlasClusterConfigurationOptionsUpdater,ClusterTester

type ClusterLister interface {
	ProjectClusters(string, *ListOptions) (*admin.PaginatedAdvancedClusterDescription, error)
}

type ClusterDescriber interface {
	AtlasCluster(string, string) (*admin.AdvancedClusterDescription, error)
}

type AtlasClusterConfigurationOptionsDescriber interface {
	AtlasClusterConfigurationOptions(string, string) (*admin.ClusterDescriptionProcessArgs, error)
}

type AtlasClusterConfigurationOptionsUpdater interface {
	UpdateAtlasClusterConfigurationOptions(string, string, *admin.ClusterDescriptionProcessArgs) (*admin.ClusterDescriptionProcessArgs, error)
}

type AtlasSharedClusterDescriber interface {
	AtlasSharedCluster(string, string) (*atlas.Cluster, error)
}

type ClusterCreator interface {
	CreateCluster(v15 *admin.AdvancedClusterDescription) (*admin.AdvancedClusterDescription, error)
}

type ClusterDeleter interface {
	DeleteCluster(string, string) error
}

type ClusterUpdater interface {
	UpdateCluster(string, string, *admin.AdvancedClusterDescription) (*admin.AdvancedClusterDescription, error)
}

type ClusterPauser interface {
	PauseCluster(string, string) (*admin.AdvancedClusterDescription, error)
}

type ClusterStarter interface {
	StartCluster(string, string) (*admin.AdvancedClusterDescription, error)
}

type ClusterUpgrader interface {
	UpgradeCluster(string, *atlas.Cluster) (*atlas.Cluster, error)
}

type SampleDataAdder interface {
	AddSampleData(string, string) (*admin.SampleDatasetStatus, error)
}

type SampleDataStatusDescriber interface {
	SampleDataStatus(string, string) (*admin.SampleDatasetStatus, error)
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
}

// AddSampleData encapsulate the logic to manage different cloud providers.
func (s *Store) AddSampleData(groupID, clusterName string) (*admin.SampleDatasetStatus, error) {
	result, _, err := s.clientv2.ClustersApi.LoadSampleDataset(s.ctx, groupID, clusterName).Execute()
	return result, err
}

// SampleDataStatus encapsulate the logic to manage different cloud providers.
func (s *Store) SampleDataStatus(groupID, id string) (*admin.SampleDatasetStatus, error) {
	result, _, err := s.clientv2.ClustersApi.GetSampleDatasetLoadStatus(s.ctx, groupID, id).Execute()
	return result, err
}

// CreateCluster encapsulate the logic to manage different cloud providers.
func (s *Store) CreateCluster(cluster *admin.AdvancedClusterDescription) (*admin.AdvancedClusterDescription, error) {
	result, _, err := s.clientv2.ClustersApi.CreateCluster(s.ctx, cluster.GetGroupId(), cluster).Execute()
	return result, err
}

// UpdateCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateCluster(projectID, name string, cluster *admin.AdvancedClusterDescription) (*admin.AdvancedClusterDescription, error) {
	result, _, err := s.clientv2.ClustersApi.UpdateCluster(s.ctx, projectID, name, cluster).Execute()
	return result, err
}

// PauseCluster encapsulate the logic to manage different cloud providers.
func (s *Store) PauseCluster(projectID, name string) (*admin.AdvancedClusterDescription, error) {
	paused := true
	cluster := &admin.AdvancedClusterDescription{
		Paused: &paused,
	}
	return s.UpdateCluster(projectID, name, cluster)
}

// StartCluster encapsulate the logic to manage different cloud providers.
func (s *Store) StartCluster(projectID, name string) (*admin.AdvancedClusterDescription, error) {
	paused := false
	cluster := &admin.AdvancedClusterDescription{
		Paused: &paused,
	}
	return s.UpdateCluster(projectID, name, cluster)
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
func (s *Store) ProjectClusters(projectID string, opts *ListOptions) (*admin.PaginatedAdvancedClusterDescription, error) {
	res := s.clientv2.ClustersApi.ListClusters(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// AtlasCluster encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasCluster(projectID, name string) (*admin.AdvancedClusterDescription, error) {
	result, _, err := s.clientv2.ClustersApi.GetCluster(s.ctx, projectID, name).Execute()
	return result, err
}

// AtlasClusterConfigurationOptions encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasClusterConfigurationOptions(projectID, name string) (*admin.ClusterDescriptionProcessArgs, error) {
	result, _, err := s.clientv2.ClustersApi.GetClusterAdvancedConfiguration(s.ctx, projectID, name).Execute()
	return result, err
}

// UpdateAtlasClusterConfigurationOptions encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateAtlasClusterConfigurationOptions(projectID, clusterName string, args *admin.ClusterDescriptionProcessArgs) (*admin.ClusterDescriptionProcessArgs, error) {
	result, _, err := s.clientv2.ClustersApi.UpdateClusterAdvancedConfiguration(s.ctx, projectID, clusterName, args).Execute()
	return result, err
}

func (s *Store) TestClusterFailover(projectID, clusterName string) error {
	_, err := s.clientv2.ClustersApi.TestFailover(s.ctx, projectID, clusterName).Execute()
	return err
}
