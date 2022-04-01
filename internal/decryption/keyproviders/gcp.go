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
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/option"
)

type GCPKeyIdentifier struct {
	KeyStoreIdentifier

	//Header
	ProjectID string
	Location  string
	KeyRing   string
	KeyName   string

	//CLI
	ServiceAccountKey string

	client *secretmanager.Client
}

func (ki *GCPKeyIdentifier) ValidateCredentials() error {
	var err error

	if ki.ServiceAccountKey != "" {
		ki.client, err = secretmanager.NewClient(context.Background(), option.WithCredentialsFile(ki.ServiceAccountKey))
		if err != nil {
			return err
		}
	}

	ki.client, err = secretmanager.NewClient(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, `No credentials found for resource: GCP location="%v" projectID="%v" keyRing="%v" keyName="%v"
`, ki.Location, ki.ProjectID, ki.KeyRing, ki.KeyName)

		json, err := provideInput("Provide service account key JSON filename:", "")
		if err != nil {
			return err
		}
		ki.client, err = secretmanager.NewClient(context.Background(), option.WithCredentialsFile(json))
		if err != nil {
			return err
		}
	}
	return nil
}

func (ki *GCPKeyIdentifier) DecryptKey(_ []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}
