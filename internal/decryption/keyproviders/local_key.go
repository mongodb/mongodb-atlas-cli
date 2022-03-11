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
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

type LocalKeyIdentifier struct {
	KeyStoreIdentifier
	Filename string
}

func (keyIdentifier *LocalKeyIdentifier) DecryptKey(encryptedLEK, iv []byte) ([]byte, error) {
	encodedKEK, err := os.ReadFile(keyIdentifier.Filename)
	if err != nil {
		return nil, err
	}

	kek, err := base64.StdEncoding.DecodeString(string(encodedKEK))
	if err != nil {
		return nil, err
	}

	return aesCBCDecrypt(kek, iv, encryptedLEK)
}

func aesCBCDecrypt(encryptionKey, iv, cipherText []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	decryptedText := make([]byte, len(cipherText))
	cipher.NewCBCDecrypter(cipherBlock, iv).CryptBlocks(decryptedText, cipherText)
	return decryptedText, nil
}
