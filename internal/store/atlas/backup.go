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
	"errors"
	"fmt"

	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_backup.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CompliancePolicyDescriber,CompliancePolicy,EnableEncryptionAtRestStore,CompliancePolicyPoliciesCreator
type CompliancePolicyDescriber interface {
	DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings, error)
}
type CompliancePolicyUpdater interface {
	UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, error)
}

type EnableEncryptionAtRestStore interface {
	EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings, error)
	CompliancePolicyDescriber
}
type CompliancePolicy interface {
	CompliancePolicyDescriber
	CompliancePolicyUpdater
}
type CompliancePolicyPoliciesCreator interface {
	CreatePolicyItem(projectID string, policyItem *atlasv2.DiskBackupApiPolicyItem) (*atlasv2.DataProtectionSettings, error)
	ProjectLister
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
		return nil, fmt.Errorf("couldn't update encryption at rest: %w", err)
	}
	compliancePolicy.SetEncryptionAtRestEnabled(true)

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, fmt.Errorf("couldn't update encryption at rest: %w", err)
	}
	return result, nil
}

func (s *Store) CreatePolicyItem(projectID string, policyItem *atlasv2.DiskBackupApiPolicyItem) (*atlasv2.DataProtectionSettings, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}

	err = addItem(compliancePolicy, policyItem)
	if err != nil {
		return nil, err
	}

	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func addItem(compliancePolicy *atlasv2.DataProtectionSettings, item *atlasv2.DiskBackupApiPolicyItem) error {
	if compliancePolicy == nil || item == nil {
		return errors.New("either compliancePolicy or item is nil")
	}
	if item.FrequencyType == "ondemand" {
		compliancePolicy.OnDemandPolicyItem = item
	} else {
		compliancePolicy.ScheduledPolicyItems = append(compliancePolicy.ScheduledPolicyItems, *item)
	}

	return nil
}
