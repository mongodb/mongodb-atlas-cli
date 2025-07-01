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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

type ListOptions struct {
	PageNum      int
	ItemsPerPage int
	IncludeCount bool
}

type ContainersListOptions struct {
	ListOptions
	ProviderName string
}

// PeeringConnections encapsulates the logic to manage different cloud providers.
func (s *Store) PeeringConnections(projectID string, opts *ContainersListOptions) ([]atlasv2.BaseNetworkPeeringConnectionSettings, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringConnections(s.ctx, projectID).
		ItemsPerPage(opts.ItemsPerPage).
		PageNum(opts.PageNum).
		ProviderName(opts.ProviderName).Execute()
	if err != nil {
		return nil, err
	}
	return result.GetResults(), nil
}

// PeeringConnection encapsulates the logic to manage different cloud providers.
func (s *Store) PeeringConnection(projectID, peerID string) (*atlasv2.BaseNetworkPeeringConnectionSettings, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.GetPeeringConnection(s.ctx, projectID, peerID).Execute()
	return result, err
}

// DeletePeeringConnection encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePeeringConnection(projectID, peerID string) error {
	_, _, err := s.clientv2.NetworkPeeringApi.DeletePeeringConnection(s.ctx, projectID, peerID).Execute()
	return err
}

// CreatePeeringConnection encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePeeringConnection(projectID string, peer *atlasv2.BaseNetworkPeeringConnectionSettings) (*atlasv2.BaseNetworkPeeringConnectionSettings, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.CreatePeeringConnection(s.ctx, projectID, peer).Execute()
	return result, err
}

// ContainersByProvider encapsulates the logic to manage different cloud providers.
func (s *Store) ContainersByProvider(projectID string, opts *ContainersListOptions) ([]atlasv2.CloudProviderContainer, error) {
	res := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).IncludeCount(opts.IncludeCount).ProviderName(opts.ProviderName)
	}
	result, _, err := res.Execute()
	if err != nil {
		return nil, err
	}
	return result.GetResults(), nil
}

const maxPerPage = 100

// AzureContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AzureContainers(projectID string) ([]atlasv2.CloudProviderContainer, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
		PageNum(0).
		ItemsPerPage(maxPerPage).
		ProviderName("Azure").
		Execute()
	if err != nil {
		return nil, err
	}
	return result.GetResults(), nil
}

// AWSContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AWSContainers(projectID string) ([]atlasv2.CloudProviderContainer, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
		PageNum(0).
		ItemsPerPage(maxPerPage).
		ProviderName("AWS").
		Execute()

	if err != nil {
		return nil, err
	}
	return result.GetResults(), nil
}

// GCPContainers encapsulates the logic to manage different cloud providers.
func (s *Store) GCPContainers(projectID string) ([]atlasv2.CloudProviderContainer, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
		PageNum(0).
		ItemsPerPage(maxPerPage).
		ProviderName("GCP").
		Execute()
	if err != nil {
		return nil, err
	}
	return result.GetResults(), nil
}

// AllContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AllContainers(projectID string, opts *ListOptions) ([]atlasv2.CloudProviderContainer, error) {
	res := s.clientv2.NetworkPeeringApi.ListPeeringContainers(s.ctx, projectID)
	if opts != nil {
		res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
	}
	result, _, err := res.Execute()
	if err != nil {
		return nil, err
	}
	return result.GetResults(), nil
}

// DeleteContainer encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteContainer(projectID, containerID string) error {
	_, err := s.clientv2.NetworkPeeringApi.DeletePeeringContainer(s.ctx, projectID, containerID).Execute()
	return err
}

// CreateContainer encapsulates the logic to manage different cloud providers.
func (s *Store) CreateContainer(projectID string, container *atlasv2.CloudProviderContainer) (*atlasv2.CloudProviderContainer, error) {
	result, _, err := s.clientv2.NetworkPeeringApi.CreatePeeringContainer(s.ctx, projectID, container).Execute()
	return result, err
}
