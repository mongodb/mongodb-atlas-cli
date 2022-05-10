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
//go:build unit
// +build unit

package featurepolicies

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/spf13/afero"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockFeatureControlPoliciesUpdater(ctrl)
	defer ctrl.Finish()

	expected := &opsmngr.FeaturePolicy{}

	t.Run("flags", func(t *testing.T) {
		opts := &UpdateOpts{
			store: mockStore,
		}
		p, _ := opts.newFeatureControl()
		mockStore.
			EXPECT().UpdateFeatureControlPolicy(opts.ConfigProjectID(), p).
			Return(expected, nil).
			Times(1)

		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		fileYML := `{
    "created":"2022-03-15T15:03:32Z",
    "externalManagementSystem": {
        "name": "mongodb-enterprise-operator",
        "version": ""
    },
    "policies": [
        {
            "policy": "DISABLE_SET_MONGOD_VERSION"
        },
        {
            "disabledParams": [],
            "policy": "EXTERNALLY_MANAGED_LOCK"
        },
        {
            "policy": "DISABLE_AUTHENTICATION_MECHANISMS",
            "disabledParams": []
        },
        {
            "policy": "DISABLE_SET_MONGOD_CONFIG",
            "disabledParams": [
                "syslog.verbosity",
                "net.tls.mode",
                "syslog.timeStampFormat",
                "net.tls.disabledProtocols",
                "setParameter.suppressNoTLSPeerCertificateWarning"
            ]
        }
    ]
}`
		fileName := "atlas_cluster_create_test.json"
		_ = afero.WriteFile(appFS, fileName, []byte(fileYML), 0600)

		opts := &UpdateOpts{
			filename: fileName,
			fs:       appFS,
			store:    mockStore,
		}

		p, _ := opts.newFeatureControl()
		mockStore.
			EXPECT().UpdateFeatureControlPolicy(opts.ConfigProjectID(), p).
			Return(expected, nil).
			Times(1)
		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.Output, flag.Name, flag.SystemID, flag.Policy},
	)
}
