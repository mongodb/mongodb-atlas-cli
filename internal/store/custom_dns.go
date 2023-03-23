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

//go:generate mockgen -destination=../mocks/mock_custom_dns.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store CustomDNSEnabler,CustomDNSDisabler,CustomDNSDescriber

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlasv2"
)

type CustomDNSEnabler interface {
	EnableCustomDNS(string) (*atlas.AWSCustomDNSEnabled, error)
}

type CustomDNSDisabler interface {
	DisableCustomDNS(string) (*atlas.AWSCustomDNSEnabled, error)
}

type CustomDNSDescriber interface {
	DescribeCustomDNS(string) (*atlas.AWSCustomDNSEnabled, error)
}

// EnableCustomDNS encapsulates the logic to manage different cloud providers.
func (s *Store) EnableCustomDNS(projectID string) (*atlas.AWSCustomDNSEnabled, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		customDNSSetting := &atlas.AWSCustomDNSEnabled{
			Enabled: true,
		}
		result, _, err := s.clientv2.AWSClustersDNSApi.ToggleAWSCustomDNS(s.ctx, projectID).AWSCustomDNSEnabled(*customDNSSetting).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DisableCustomDNS encapsulates the logic to manage different cloud providers.
func (s *Store) DisableCustomDNS(projectID string) (*atlas.AWSCustomDNSEnabled, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		customDNSSetting := &atlas.AWSCustomDNSEnabled{
			Enabled: false,
		}
		result, _, err := s.clientv2.AWSClustersDNSApi.ToggleAWSCustomDNS(s.ctx, projectID).AWSCustomDNSEnabled(*customDNSSetting).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DescribeCustomDNS encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeCustomDNS(projectID string) (*atlas.AWSCustomDNSEnabled, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.AWSClustersDNSApi.GetAWSCustomDNS(s.ctx, projectID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
