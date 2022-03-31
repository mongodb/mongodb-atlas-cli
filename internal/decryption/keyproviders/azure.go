// Copyright 2022 MongoDB Inc
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

package keyproviders

import (
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	_ "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type AzureKeyIdentifier struct {
	KeyStoreIdentifier
	KeyName          string
	Environment      string
	KeyVaultEndpoint string
	KeyVersion       string
	ClientID         string
	TenantID         string
	Secret           string

	credentials azcore.TokenCredential
}

func (ki *AzureKeyIdentifier) ValidateCredentials() error {
	var err error

	ki.credentials, err = azidentity.NewClientSecretCredential(ki.TenantID, ki.ClientID, ki.Secret, nil)
	if err == nil {
		return nil
	}

	ki.credentials, err = azidentity.NewEnvironmentCredential(nil)
	if err == nil {
		return nil
	}

	tenantID, err := provideInput("Provide Tenant ID:", ki.TenantID)
	if err != nil {
		return err
	}

	clientID, err := provideInput("Provide Client ID:", ki.ClientID)
	if err != nil {
		return err
	}

	secret, err := provideInput("Provide Secret:", ki.Secret)
	if err != nil {
		return err
	}

	ki.credentials, err = azidentity.NewClientSecretCredential(tenantID, clientID, secret, nil)
	if err == nil {
		return nil
	}

	ki.credentials, err = azidentity.NewDefaultAzureCredential(nil)
	if err == nil {
		return nil
	}

	return err
}

func (ki *AzureKeyIdentifier) DecryptKey(_ []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}
