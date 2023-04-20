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

package atlas

import (
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_performance_advisor.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas PerformanceAdvisorNamespacesLister,PerformanceAdvisorSlowQueriesLister,PerformanceAdvisorIndexesLister,PerformanceAdvisorSlowOperationThresholdEnabler,PerformanceAdvisorSlowOperationThresholdDisabler
type PerformanceAdvisorNamespacesLister interface {
	PerformanceAdvisorNamespaces(string, string, *atlas.NamespaceOptions) (*atlas.Namespaces, error)
}

type PerformanceAdvisorSlowQueriesLister interface {
	PerformanceAdvisorSlowQueries(string, string, *atlas.SlowQueryOptions) (*atlas.SlowQueries, error)
}

type PerformanceAdvisorIndexesLister interface {
	PerformanceAdvisorIndexes(string, string, *atlas.SuggestedIndexOptions) (*atlas.SuggestedIndexes, error)
}

type PerformanceAdvisorSlowOperationThresholdEnabler interface {
	EnablePerformanceAdvisorSlowOperationThreshold(string) error
}

type PerformanceAdvisorSlowOperationThresholdDisabler interface {
	DisablePerformanceAdvisorSlowOperationThreshold(string) error
}

// PerformanceAdvisorNamespaces encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorNamespaces(projectID, processName string, opts *atlas.NamespaceOptions) (*atlas.Namespaces, error) {
	result, _, err := s.client.PerformanceAdvisor.GetNamespaces(s.ctx, projectID, processName, opts)
	return result, err
}

// PerformanceAdvisorSlowQueries encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorSlowQueries(projectID, processName string, opts *atlas.SlowQueryOptions) (*atlas.SlowQueries, error) {
	result, _, err := s.client.PerformanceAdvisor.GetSlowQueries(s.ctx, projectID, processName, opts)
	return result, err
}

// PerformanceAdvisorIndexes encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorIndexes(projectID, processName string, opts *atlas.SuggestedIndexOptions) (*atlas.SuggestedIndexes, error) {
	result, _, err := s.client.PerformanceAdvisor.GetSuggestedIndexes(s.ctx, projectID, processName, opts)
	return result, err
}

// EnablePerformanceAdvisorSlowOperationThreshold encapsulates the logic to manage different cloud providers.
func (s *Store) EnablePerformanceAdvisorSlowOperationThreshold(projectID string) error {
	_, err := s.client.PerformanceAdvisor.EnableManagedSlowOperationThreshold(s.ctx, projectID)
	return err
}

// DisablePerformanceAdvisorSlowOperationThreshold encapsulates the logic to manage different cloud providers.
func (s *Store) DisablePerformanceAdvisorSlowOperationThreshold(projectID string) error {
	_, err := s.client.PerformanceAdvisor.DisableManagedSlowOperationThreshold(s.ctx, projectID)
	return err
}
