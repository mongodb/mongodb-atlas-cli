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
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/endpointcreds"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
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
	ErrAWSInit    = errors.New("failed to initialize AWS KMS Service")
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

	cfg.Credentials = newChainProvider(
		cfg.Credentials,
		credentials.NewStaticCredentialsProvider(ki.AccessKey, ki.SecretAccessKey, ki.SessionToken),
	)

	if _, err2 := cfg.Credentials.Retrieve(ctx); err2 != nil {
		var target *credentials.StaticCredentialsEmptyError
		if !errors.As(err, &target) {
			return err
		}
		_, _ = log.Warningf(`No credentials found for resource: AWS region="%v" endpoint="%v" key="%v"
`, ki.Region, ki.Endpoint, ki.Key)
		_, _ = log.Warningf("Note: if you have an AWS session token leave AWS access key and AWS secret access key empty")
		ki.AccessKey, err = provideInput("Provide AWS access key:", ki.AccessKey)
		if err != nil {
			return err
		}
		ki.SecretAccessKey, err = provideInput("Provide AWS secret access key:", ki.SecretAccessKey)
		if err != nil {
			return err
		}
		ki.SessionToken, err = provideInput("Provide AWS session token:", ki.SessionToken)
		if err != nil {
			return err
		}
		cfg.Credentials = credentials.NewStaticCredentialsProvider(ki.AccessKey, ki.SecretAccessKey, ki.SessionToken)
	}
	ki.cfg = cfg
	return err
}

func newChainProvider(providers ...aws.CredentialsProvider) aws.CredentialsProvider {
	return aws.NewCredentialsCache(
		aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			var errs []error
			for _, p := range providers {
				creds, err := p.Retrieve(ctx)
				if err == nil {
					return creds, nil
				}
				errs = append(errs, err)
			}

			return aws.Credentials{}, fmt.Errorf("%w: %s", ErrAWSInit, errs)
		}),
	)
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
		return nil, fmt.Errorf("%w: %w", ErrAWSDecrypt, err)
	}

	return output.Plaintext, nil
}
