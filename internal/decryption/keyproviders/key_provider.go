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

import "fmt"

type KeyStoreProvider string

const (
	LocalKey KeyStoreProvider = "local"
	AWS      KeyStoreProvider = "aws"
	Azure    KeyStoreProvider = "azure"
	GCP      KeyStoreProvider = "gcp"
	KMIP     KeyStoreProvider = "kmip"
)

type AWSKeyIdentifier struct {
	Key      string
	Region   string
	Endpoint string
}

type AzureKeyIdentifier struct {
	Environment      string
	KeyVaultEndpoint string
	KeyName          string
	KeyVersion       string
}

type GCPKeyIdentifier struct {
	ProjectID string
	Location  string
	KeyRing   string
	KeyName   string
}

type KeyStoreIdentifier struct {
	Provider KeyStoreProvider
	Filename string
	KMIP     *KMIPKeyIdentifier
	AWS      *AWSKeyIdentifier
	Azure    *AzureKeyIdentifier
	GCP      *GCPKeyIdentifier
}

func DecryptLEK(keyStore KeyStoreIdentifier, encryptedLEK, iv []byte) ([]byte, error) {
	switch keyStore.Provider {
	case LocalKey:
		return decryptWithLocalKey(keyStore.Filename, encryptedLEK, iv)
	case KMIP, AWS, Azure, GCP:
		return nil, fmt.Errorf(`KeyStoreProvider "%s" is not implemented`, keyStore.Provider)
	default:
		return nil, fmt.Errorf(`KeyStoreProvider "%s" is not supported`, keyStore.Provider)
	}
}
