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
	"errors"
	"fmt"
	"hash/crc32"

	kmsv1 "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var ErrDataCorruptedInTransit = errors.New("decrypt: response corrupted in-transit")

type GCPKeyIdentifier struct {
	KeyStoreIdentifier

	// Header
	ProjectID string
	Location  string
	KeyRing   string
	KeyName   string

	// CLI
	ServiceAccountKey string

	client *kmsv1.KeyManagementClient
}

func (ki *GCPKeyIdentifier) ValidateCredentials() error {
	var err error

	if ki.ServiceAccountKey != "" {
		ki.client, err = kmsv1.NewKeyManagementClient(context.Background(), option.WithCredentialsFile(ki.ServiceAccountKey))
		if err == nil {
			return nil
		}
	}

	ki.client, err = kmsv1.NewKeyManagementClient(context.Background())
	if err != nil {
		_, _ = log.Warningf(`No credentials found for resource: GCP location="%v" projectID="%v" keyRing="%v" keyName="%v"
`, ki.Location, ki.ProjectID, ki.KeyRing, ki.KeyName)

		var json string
		json, err = provideInput("Provide service account key JSON filename:", "")
		if err != nil {
			return err
		}
		ki.client, err = kmsv1.NewKeyManagementClient(context.Background(), option.WithCredentialsFile(json))
		if err == nil {
			return nil
		}
	}
	return err
}

func crc32c(data []byte) int64 {
	t := crc32.MakeTable(crc32.Castagnoli)
	return int64(crc32.Checksum(data, t))
}

func (ki *GCPKeyIdentifier) DecryptKey(key []byte) ([]byte, error) {
	defer func() {
		ki.client.Close()
		ki.client = nil
	}()

	resourceName := fmt.Sprintf("projects/%v/locations/%v/keyRings/%v/cryptoKeys/%v", ki.ProjectID, ki.Location, ki.KeyRing, ki.KeyName)
	req := &kmspb.DecryptRequest{
		Name:             resourceName,
		Ciphertext:       key,
		CiphertextCrc32C: wrapperspb.Int64(crc32c(key)),
	}
	result, err := ki.client.Decrypt(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if crc32c(result.GetPlaintext()) != result.PlaintextCrc32C.Value {
		return nil, ErrDataCorruptedInTransit
	}
	return result.GetPlaintext(), nil
}
