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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_performance_advisor.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store PerformanceAdvisorNamespacesLister,PerformanceAdvisorSlowQueriesLister,PerformanceAdvisorIndexesLister
type PerformanceAdvisorNamespacesLister interface {
	PerformanceAdvisorNamespaces(string, string, *opsmngr.NamespaceOptions) (*opsmngr.Namespaces, error)
}

type PerformanceAdvisorSlowQueriesLister interface {
	PerformanceAdvisorSlowQueries(string, string, *opsmngr.SlowQueryOptions) (*opsmngr.SlowQueries, error)
}

type PerformanceAdvisorIndexesLister interface {
	PerformanceAdvisorIndexes(string, string, *opsmngr.SuggestedIndexOptions) (*opsmngr.SuggestedIndexes, error)
}

// PerformanceAdvisorNamespaces encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorNamespaces(projectID, processName string, opts *opsmngr.NamespaceOptions) (*opsmngr.Namespaces, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.PerformanceAdvisor.GetNamespaces(s.ctx, projectID, processName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// PerformanceAdvisorSlowQueries encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorSlowQueries(projectID, processName string, opts *opsmngr.SlowQueryOptions) (*opsmngr.SlowQueries, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.PerformanceAdvisor.GetSlowQueries(s.ctx, projectID, processName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// PerformanceAdvisorIndexes encapsulates the logic to manage different cloud providers.
func (s *Store) PerformanceAdvisorIndexes(projectID, processName string, opts *opsmngr.SuggestedIndexOptions) (*opsmngr.SuggestedIndexes, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.PerformanceAdvisor.GetSuggestedIndexes(s.ctx, projectID, processName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
