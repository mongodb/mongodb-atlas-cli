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
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
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

	credentials credentials.Value
}

func (ki *AWSKeyIdentifier) ValidateCredentials() error {
	p := &credentials.ChainProvider{
		VerboseErrors: false,
		Providers: []credentials.Provider{
			&credentials.StaticProvider{Value: credentials.Value{
				AccessKeyID:     ki.AccessKey,
				SecretAccessKey: ki.SecretAccessKey,
				SessionToken:    ki.SessionToken,
			}},
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{},
		},
	}
	cred := credentials.NewCredentials(p)
	v, err := cred.Get()
	if err != nil {
		if err != credentials.ErrNoValidProvidersFoundInChain {
			return err
		}
		fmt.Fprintf(os.Stderr, `No credentials found for resource: AWS region="%v" endpoint="%v" key="%v"
`, ki.Region, ki.Endpoint, ki.Key)
		fmt.Fprintln(os.Stderr, "Note: if you have an AWS session token leave AWS access key and AWS secret access key empty")
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
		cred = credentials.NewStaticCredentials(ki.AccessKey, ki.SecretAccessKey, ki.SessionToken)
		v, err = cred.Get()
		if err != nil {
			return err
		}
	}
	ki.credentials = v
	return nil
}

func (ki *AWSKeyIdentifier) DecryptKey(_ []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}
