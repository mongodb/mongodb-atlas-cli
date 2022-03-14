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

type KMIPKeyWrapMethod string

const (
	KMIPKeyWrapMethodGet     KMIPKeyWrapMethod = "get"
	KMIPKeyWrapMethodEncrypt KMIPKeyWrapMethod = "encrypt"
)

type KMIPKeyIdentifier struct {
	KeyStoreIdentifier
	UniqueKeyID               string
	ServerName                string
	ServerPort                string
	KeyWrapMethod             KMIPKeyWrapMethod
	ServerCAFileName          string
	ClientCertificateFileName string
}

func (keyIdentifier *KMIPKeyIdentifier) DecryptKey(_, _ []byte) ([]byte, error) {
	return nil, fmt.Errorf("%s not implemented", keyIdentifier.Provider)
}
