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
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
)

type KeyStoreProvider string

const (
	LocalKey KeyStoreProvider = "local"
	KMIP     KeyStoreProvider = "kmip"
	AWS      KeyStoreProvider = "aws"
	GCP      KeyStoreProvider = "gcp"
	Azure    KeyStoreProvider = "azure"
)

type KeyProvider interface {
	ValidateCredentials() error
	DecryptKey(encryptedLEK []byte) ([]byte, error)
}

type KeyStoreIdentifier struct {
	Provider KeyStoreProvider
}

func provideInput(prompt, defaultInput string) (string, error) {
	if strings.TrimSpace(strings.ToLower(os.Getenv("CI"))) == "true" {
		return defaultInput, nil
	}

	var input string
	err := telemetry.TrackAskOne(&survey.Input{
		Message: prompt,
		Default: defaultInput,
	}, &input)
	if err != nil {
		return "", err
	}

	return input, nil
}
