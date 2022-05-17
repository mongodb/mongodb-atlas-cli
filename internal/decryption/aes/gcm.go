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

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

type GCMInput struct {
	Key []byte
	IV  []byte
	AAD []byte
	Tag []byte
}

var ErrAESGCMecrypt = errors.New("aes-gcm decrypt error")

func (input *GCMInput) Decrypt(cipherText []byte) (decrypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%w: %v", ErrAESGCMecrypt, r)
			decrypted = nil
		}
	}()

	cipherBlock, err := aes.NewCipher(input.Key)
	if err != nil {
		return nil, err
	}

	gcmBlockChiper, err := cipher.NewGCMWithTagSize(cipherBlock, len(input.Tag))
	if err != nil {
		return nil, err
	}

	cipherTextWithTag := append([]byte{}, cipherText...)
	cipherTextWithTag = append(cipherTextWithTag, input.Tag...)

	return gcmBlockChiper.Open(nil, input.IV, cipherTextWithTag, input.AAD)
}
