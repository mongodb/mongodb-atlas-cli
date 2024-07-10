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

package store

import (
	"errors"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

var errTScheduledPolicyItemNotFound = errors.New("scheduled policy item not found")

//go:generate mockgen -destination=../mocks/mock_backup_compliance.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store CompliancePolicyDescriber,CompliancePolicyUpdater,CompliancePolicyEncryptionAtRestEnabler,CompliancePolicyEncryptionAtRestDisabler,CompliancePolicyEnabler,CompliancePolicyCopyProtectionEnabler,CompliancePolicyCopyProtectionDisabler,CompliancePolicyPointInTimeRestoresEnabler,CompliancePolicyOnDemandPolicyCreator,CompliancePolicyScheduledPolicyCreator,CompliancePolicyScheduledPolicyDeleter,CompliancePolicyScheduledPolicyUpdater

type CompliancePolicyDescriber interface {
	DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
}
type CompliancePolicyPointInTimeRestoresEnabler interface {
	EnablePointInTimeRestore(projectID string, restoreWindowDays int) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyEnabler interface {
	EnableCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyCopyProtectionEnabler interface {
	EnableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}
type CompliancePolicyCopyProtectionDisabler interface {
	DisableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}
type CompliancePolicyEncryptionAtRestUpdater interface {
	UpdateEncryptionAtRest(projectID string, enable bool) (*atlasv2.DataProtectionSettings20231001, error)
}
type CompliancePolicyUpdater interface {
	CompliancePolicyDescriber
	UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error)
}

type CompliancePolicyEncryptionAtRestEnabler interface {
	EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyEncryptionAtRestDisabler interface {
	DisableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyOnDemandPolicyCreator interface {
	CreateOnDemandPolicy(projectID string, policy *atlasv2.BackupComplianceOnDemandPolicyItem) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyScheduledPolicyCreator interface {
	CreateScheduledPolicy(projectID string, policy *atlasv2.BackupComplianceScheduledPolicyItem) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyScheduledPolicyUpdater interface {
	UpdateScheduledPolicy(projectID string, policy *atlasv2.BackupComplianceScheduledPolicyItem) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

type CompliancePolicyScheduledPolicyDeleter interface {
	DeleteScheduledPolicy(projectID, scheduledPolicyID string) (*atlasv2.DataProtectionSettings20231001, error)
	CompliancePolicyDescriber
}

func (s *Store) DescribeCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	return result, err
}

func (s *Store) UpdateCompliancePolicy(projectID string, opts *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, opts).Execute()
	return result, err
}

func (s *Store) EnablePointInTimeRestore(projectID string, restoreWindowDays int) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetRestoreWindowDays(restoreWindowDays)
	compliancePolicy.SetPitEnabled(true)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) EnableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetEncryptionAtRestEnabled(true)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) EnableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetCopyProtectionEnabled(true)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}
func (s *Store) DisableEncryptionAtRest(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetEncryptionAtRestEnabled(false)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) DisableCopyProtection(projectID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}
	compliancePolicy.SetCopyProtectionEnabled(false)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}
func (s *Store) EnableCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy := newEmptyCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) CreateOnDemandPolicy(projectID string, policy *atlasv2.BackupComplianceOnDemandPolicyItem) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}

	compliancePolicy.SetOnDemandPolicyItem(*policy)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) CreateScheduledPolicy(projectID string, policy *atlasv2.BackupComplianceScheduledPolicyItem) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}

	compliancePolicy.SetScheduledPolicyItems(append(compliancePolicy.GetScheduledPolicyItems(), *policy))
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) UpdateScheduledPolicy(projectID string, policy *atlasv2.BackupComplianceScheduledPolicyItem) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}

	for idx, i := range compliancePolicy.GetScheduledPolicyItems() {
		if i.GetId() == policy.GetId() {
			compliancePolicy.GetScheduledPolicyItems()[idx] = *policy
			compliancePolicy.SetScheduledPolicyItems(append(compliancePolicy.GetScheduledPolicyItems(), *policy))
			return s.updateDataProtectionSettings(projectID, compliancePolicy)
		}
	}

	return nil, errTScheduledPolicyItemNotFound
}

func (s *Store) DeleteScheduledPolicy(projectID, scheduledPolicyID string) (*atlasv2.DataProtectionSettings20231001, error) {
	compliancePolicy, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettings(s.ctx, projectID).Execute()
	if err != nil {
		return nil, err
	}

	items := make([]atlasv2.BackupComplianceScheduledPolicyItem, 0, len(compliancePolicy.GetScheduledPolicyItems()))
	found := false
	for _, i := range compliancePolicy.GetScheduledPolicyItems() {
		if i.GetId() == scheduledPolicyID {
			found = true
			continue
		}
		items = append(items, i)
	}

	if !found {
		return nil, errTScheduledPolicyItemNotFound
	}

	compliancePolicy.SetScheduledPolicyItems(items)
	return s.updateDataProtectionSettings(projectID, compliancePolicy)
}

func (s *Store) updateDataProtectionSettings(projectID string, compliancePolicy *atlasv2.DataProtectionSettings20231001) (*atlasv2.DataProtectionSettings20231001, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettings(s.ctx, projectID, compliancePolicy).Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func newEmptyCompliancePolicy(projectID, authorizedEmail, authorizedFirstName, authorizedLastName string) *atlasv2.DataProtectionSettings20231001 {
	policy := atlasv2.NewDataProtectionSettings20231001(authorizedEmail, authorizedFirstName, authorizedLastName)
	policy.SetProjectId(projectID)
	return policy
}
