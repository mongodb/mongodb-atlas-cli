// Copyright 2023 MongoDB Inc
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
	"fmt"

	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_backup.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CompliancePolicyDescriber,CompliancePolicy,CompliancePolicyEncryptionAtRestEnabler,CompliancePolicyEncryptionAtRestDisabler,CompliancePolicyEnabler
type CompliancePolicyDescriber interface {
	DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings, error)
}
type CompliancePolicyUpdater interface {
	UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, error)
}

type CompliancePolicyEnabler interface {
	EnableCompliancePolicy(projectID string, authorizedEmail string) (*atlasv2.DataProtectionSettings, error)
	CompliancePolicyDescriber
}
type CompliancePolicyEncryptionAtRestUpdater interface {
	UpdateEncryptionAtRest(projectID string, enable bool) (*atlasv2.DataProtectionSettings, error)
}

type CompliancePolicy interface {
	CompliancePolicyDescriber
	CompliancePolicyUpdater
}

type CompliancePolicyEncryptionAtRestEnabler interface {
	EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings, error)
	CompliancePolicyDescriber
}

type CompliancePolicyEncryptionAtRestDisabler interface {
	DisableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings, error)
	CompliancePolicyDescriber
}

func (s *Store) DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	return result, err
}

func (s *Store) UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, opts).Execute()
	return result, err
}

func (s *Store) EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetEncryptionAtRestEnabled(true)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Store) DisableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetEncryptionAtRestEnabled(false)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Store) EnableCompliancePolicy(projectID string, authorizedEmail string) (*atlasv2.DataProtectionSettings, error) {
	compliancePolicy := newEmptyCompliancePolicy(projectID, authorizedEmail)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, fmt.Errorf("couldn't enable compliance policy: %w", err)
	}
	return result, nil
}

func newEmptyCompliancePolicy(projectID string, authorizedEmail string) *atlasv2.DataProtectionSettings {
	policy := atlasv2.NewDataProtectionSettings()
	policy.SetAuthorizedEmail(authorizedEmail)
	policy.SetProjectId(projectID)
	return policy
}
