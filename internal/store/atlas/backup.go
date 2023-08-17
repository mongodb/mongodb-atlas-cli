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
	"net/http"

	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_backup.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CompliancePolicyDescriber,CompliancePolicy,CompliancePolicyItemUpdater

type CompliancePolicyDescriber interface {
	DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings, error)
}
type CompliancePolicyUpdater interface {
	UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, error)
}
type CompliancePolicyItemUpdater interface {
	UpdatePolicyItem(projectID string, policyItem *atlasv2.DiskBackupApiPolicyItem) (*atlasv2.DataProtectionSettings, *http.Response, error)
}
type CompliancePolicy interface {
	CompliancePolicyDescriber
	CompliancePolicyUpdater
}

func (s *Store) DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	return result, err
}

func (s *Store) UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, opts).Execute()
	return result, err
}

func (s *Store) UpdateCompliancePolicyAndGetResponse(projectID string, opts *atlasv2.DataProtectionSettings) (*atlasv2.DataProtectionSettings, *http.Response, error) {
	result, httpResp, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, opts).Execute()
	return result, httpResp, err
}

func (s *Store) UpdatePolicyItem(projectID string, policyItem *atlasv2.DiskBackupApiPolicyItem) (*atlasv2.DataProtectionSettings, *http.Response, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't update compliance policy: %w", err)
	}

	err = replaceItem(compliancePolicy, policyItem)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't update compliance policy: %w", err)
	}

	result, httpResp, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't update compliance policy: %w", err)
	}
	return result, httpResp, err
}

// replaceItem searches for a DiskBackupApiPolicyItem within the provided DataProtectionSettings by its ID.
// If a matching item is found, it replaces the existing item with the provided item **in place**.
func replaceItem(compliancePolicy *atlasv2.DataProtectionSettings, item *atlasv2.DiskBackupApiPolicyItem) error {
	items := compliancePolicy.GetScheduledPolicyItems()
	for i, existingItem := range items {
		if existingItem.GetId() == item.GetId() {
			items[i] = *item
			return nil
		}
	}
	onDemandItem := compliancePolicy.GetOnDemandPolicyItem()
	if onDemandItem.GetId() == item.GetId() {
		compliancePolicy.SetOnDemandPolicyItem(*item)
		return nil
	}
	return errors.New("did not find a policy item with a matching ID")
}
