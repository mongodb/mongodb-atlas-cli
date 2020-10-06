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

package backupstore

import "go.mongodb.org/ops-manager/opsmngr"

type AdminOpts struct {
	Assignment               bool
	EncryptedCredentials     bool
	SSL                      bool
	WriteConcern             string
	URI                      string
	LoadFactor               int64
	MaxCapacityGB            int64
	MMAPV1CompressionSetting string
	StorePath                string
	WTCompressionSetting     string
	ID                       string
	Label                    []string
}

func (opts *AdminOpts) NewBackupStore() *opsmngr.BackupStore {
	backupStore := &opsmngr.BackupStore{
		AdminBackupConfig: opsmngr.AdminBackupConfig{
			ID:           opts.ID,
			URI:          opts.URI,
			WriteConcern: opts.WriteConcern,
			Labels:       opts.Label,
		},
	}

	if opts.SSL {
		backupStore.SSL = &opts.SSL
	}

	if opts.EncryptedCredentials {
		backupStore.EncryptedCredentials = &opts.EncryptedCredentials
	}

	if opts.Assignment {
		backupStore.AssignmentEnabled = &opts.Assignment
	}

	if opts.MaxCapacityGB != 0 {
		backupStore.MaxCapacityGB = &opts.MaxCapacityGB
	}

	if opts.LoadFactor != 0 {
		backupStore.LoadFactor = &opts.LoadFactor
	}

	return backupStore
}

func (opts *AdminOpts) NewFileSystemConfiguration() *opsmngr.FileSystemStoreConfiguration {
	return &opsmngr.FileSystemStoreConfiguration{
		BackupStore:              *opts.NewBackupStore(),
		MMAPV1CompressionSetting: opts.MMAPV1CompressionSetting,
		StorePath:                opts.StorePath,
		WTCompressionSetting:     opts.WTCompressionSetting,
	}
}
