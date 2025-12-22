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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

// PerformanceAdvisorNamespaces encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorNamespaces(opts *atlasv2.ListPerformanceAdvisorNamespacesApiParams) (*atlasv2.Namespaces, error) {
	request := s.clientv2.PerformanceAdvisorApi.
		ListPerformanceAdvisorNamespaces(s.ctx, opts.GroupId, opts.ProcessId)
	if opts.Duration != nil {
		request = request.Duration(*opts.Duration)
	}
	if opts.Since != nil {
		request = request.Since(*opts.Since)
	}
	result, _, err := request.Execute()
	return result, err
}

// PerformanceAdvisorSlowQueries encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorSlowQueries(opts *atlasv2.ListSlowQueryLogsApiParams) (*atlasv2.PerformanceAdvisorSlowQueryList, error) {
	request := s.clientv2.PerformanceAdvisorApi.ListSlowQueryLogs(s.ctx, opts.GroupId, opts.ProcessId)
	if opts.Duration != nil {
		request = request.Duration(*opts.Duration)
	}
	if opts.Since != nil {
		request = request.Since(*opts.Since)
	}
	if opts.Namespaces != nil {
		request = request.Namespaces(*opts.Namespaces)
	}
	if opts.NLogs != nil {
		request = request.NLogs(*opts.NLogs)
	}
	result, _, err := request.Execute()
	return result, err
}

// PerformanceAdvisorIndexes encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorIndexes(opts *atlasv2.ListSuggestedIndexesApiParams) (*atlasv2.PerformanceAdvisorResponse, error) {
	request := s.clientv2.PerformanceAdvisorApi.
		ListSuggestedIndexes(s.ctx, opts.GroupId, opts.ProcessId)
	if opts.Namespaces != nil {
		request = request.Namespaces(*opts.Namespaces)
	}
	if opts.Duration != nil {
		request = request.Duration(*opts.Duration)
	}
	if opts.Since != nil {
		request = request.Since(*opts.Since)
	}
	if opts.NExamples != nil {
		request = request.NExamples(*opts.NExamples)
	}
	if opts.NIndexes != nil {
		request = request.NIndexes(*opts.NIndexes)
	}
	result, _, err := request.Execute()
	return result, err
}

// EnablePerformanceAdvisorSlowOperationThreshold encapsulates the logic to manage different cloud providers.
func (s *Store) EnablePerformanceAdvisorSlowOperationThreshold(projectID string) error {
	_, err := s.clientv2.PerformanceAdvisorApi.EnableManagedSlowMs(s.ctx, projectID).Execute()
	return err
}

// DisablePerformanceAdvisorSlowOperationThreshold encapsulates the logic to manage different cloud providers.
func (s *Store) DisablePerformanceAdvisorSlowOperationThreshold(projectID string) error {
	_, err := s.clientv2.PerformanceAdvisorApi.DisableManagedSlowMs(s.ctx, projectID).Execute()
	return err
}
