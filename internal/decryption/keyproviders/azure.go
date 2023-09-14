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
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
)

type AzureKeyIdentifier struct {
	KeyStoreIdentifier

	// Header
	KeyName          string
	Environment      string // not used
	KeyVaultEndpoint string
	KeyVersion       string

	// CLI
	ClientID string
	TenantID string
	Secret   string

	credentials azcore.TokenCredential
}

func (ki *AzureKeyIdentifier) ValidateCredentials() error {
	var err error

	ki.credentials, err = azidentity.NewClientSecretCredential(ki.TenantID, ki.ClientID, ki.Secret, nil)
	if err == nil {
		return nil
	}

	ki.credentials, err = azidentity.NewDefaultAzureCredential(nil)
	if err == nil {
		return nil
	}

	return err
}

func (ki *AzureKeyIdentifier) DecryptKey(key []byte) ([]byte, error) {
	client, err := azkeys.NewClient(ki.KeyVaultEndpoint, ki.credentials, nil)
	if err != nil {
		return nil, err
	}

	algo := azkeys.EncryptionAlgorithmRSAOAEP
	r, err := client.Decrypt(context.Background(), ki.KeyName, ki.KeyVersion, azkeys.KeyOperationParameters{
		Value:     key,
		Algorithm: &algo,
	}, nil)
	if err != nil {
		return nil, err
	}
	return r.Result, nil
}
