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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/endpointcreds"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type AWSKeyIdentifier struct {
	KeyStoreIdentifier
	// Header
	Key      string
	Region   string
	Endpoint string

	// CLI
	AccessKey       string
	SecretAccessKey string
	SessionToken    string

	cfg aws.Config
}

var (
	ErrAWSDecrypt = errors.New("unable to decrypt data key with AWS KMS Service")
)

func (ki *AWSKeyIdentifier) ValidateCredentials() error {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(ki.Region),
		config.WithEndpointCredentialOptions(func(options *endpointcreds.Options) {
			options.Endpoint = ki.Endpoint
		}),
	)
	ki.cfg = cfg
	return err
}

// DecryptKey attempts to decrypt the encrypted key using AWS KMS.
func (ki *AWSKeyIdentifier) DecryptKey(encryptedKey []byte) ([]byte, error) {
	ctx := context.Background()
	service := kms.NewFromConfig(ki.cfg)

	input := &kms.DecryptInput{
		CiphertextBlob: encryptedKey,
	}
	output, err := service.Decrypt(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrAWSDecrypt, err)
	}

	return output.Plaintext, nil
}
