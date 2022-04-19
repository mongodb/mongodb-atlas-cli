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

package pem

import (
	"encoding/pem"
	"errors"

	"github.com/spf13/afero"
	"go.step.sm/crypto/pemutil"
)

type BlockType string

const (
	CertificateBlock         BlockType = "CERTIFICATE"
	RSAPrivateKeyBlock       BlockType = "RSA PRIVATE KEY"
	EncryptedPrivateKeyBlock BlockType = "ENCRYPTED PRIVATE KEY"
)

var (
	ErrKMIPCertificateBlock       = errors.New("file does not contain a certificate block")
	ErrKMIPMissingPrivateKeyBlock = errors.New("file does not contain a private key block")
)

func load(fs afero.Fs, filename string) (map[BlockType]*pem.Block, error) {
	clientCertAndKey, err := afero.ReadFile(fs, filename)
	if err != nil {
		return nil, err
	}

	pemBlocks := map[BlockType]*pem.Block{}
	for {
		var pemBlock *pem.Block
		pemBlock, clientCertAndKey = pem.Decode(clientCertAndKey)
		if pemBlock == nil {
			break
		}
		pemBlocks[BlockType(pemBlock.Type)] = pemBlock
	}

	return pemBlocks, nil
}

func Decode(fs afero.Fs, filename, password string) (cert, privateKey []byte, err error) {
	pemBlocks, err := load(fs, filename)
	if err != nil {
		return nil, nil, err
	}

	for blockType, pemBlock := range pemBlocks {
		switch blockType {
		case CertificateBlock:
			cert = pem.EncodeToMemory(pemBlock)
		case RSAPrivateKeyBlock:
			privateKey = pem.EncodeToMemory(pemBlock)
		case EncryptedPrivateKeyBlock:
			privateKeyBytes, err := pemutil.DecryptPKCS8PrivateKey(pemBlock.Bytes, []byte(password))
			if err != nil {
				return nil, nil, err
			}
			pemBlock = &pem.Block{Type: string(RSAPrivateKeyBlock), Bytes: privateKeyBytes}
			privateKey = pem.EncodeToMemory(pemBlock)
		}
	}

	return cert, privateKey, nil
}

func ValidateBlocks(fs afero.Fs, filename string) (isEncrypted bool, err error) {
	pemBlocks, err := load(fs, filename)
	if err != nil {
		return false, err
	}

	_, hasPrivateKey := pemBlocks[RSAPrivateKeyBlock]
	_, hasEncryptedPrivateKey := pemBlocks[EncryptedPrivateKeyBlock]
	if !hasPrivateKey && !hasEncryptedPrivateKey {
		return false, ErrKMIPMissingPrivateKeyBlock
	}

	if _, hasCertBlock := pemBlocks[CertificateBlock]; !hasCertBlock {
		return hasEncryptedPrivateKey, ErrKMIPCertificateBlock
	}

	return hasEncryptedPrivateKey, nil
}
