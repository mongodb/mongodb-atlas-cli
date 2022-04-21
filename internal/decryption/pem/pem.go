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
)

type BlockType string

const (
	CertificateBlock         BlockType = "CERTIFICATE"
	RSAPrivateKeyBlock       BlockType = "RSA PRIVATE KEY"
	EncryptedPrivateKeyBlock BlockType = "ENCRYPTED PRIVATE KEY"
)

var defaultPem = &pemDecoderValidator{
	fs: afero.NewOsFs(),
}

type pemDecoderValidator struct {
	fs afero.Fs
}

func Default() DecoderValidator {
	return defaultPem
}

type DecoderValidator interface {
	Decode(filename, password string) (cert, privateKey []byte, err error)
	ValidateBlocks(filename string) (isEncrypted bool, err error)
}

var (
	errKMIPCertificateBlock       = errors.New("file does not contain a certificate block")
	errKMIPMissingPrivateKeyBlock = errors.New("file does not contain a private key block")
)

func (p *pemDecoderValidator) load(filename string) (map[BlockType]*pem.Block, error) {
	clientCertAndKey, err := afero.ReadFile(p.fs, filename)
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

func Decode(filename, password string) (cert, privateKey []byte, err error) {
	return defaultPem.Decode(filename, password)
}

func (p *pemDecoderValidator) Decode(filename, password string) (cert, privateKey []byte, err error) {
	pemBlocks, err := p.load(filename)
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
			privateKeyBytes, err := DecryptPKCS8PrivateKey(pemBlock.Bytes, []byte(password))
			if err != nil {
				return nil, nil, err
			}
			pemBlock = &pem.Block{Type: string(RSAPrivateKeyBlock), Bytes: privateKeyBytes}
			privateKey = pem.EncodeToMemory(pemBlock)
		}
	}

	return cert, privateKey, nil
}

func ValidateBlocks(filename string) (isEncrypted bool, err error) {
	return defaultPem.ValidateBlocks(filename)
}

func (p *pemDecoderValidator) ValidateBlocks(filename string) (isEncrypted bool, err error) {
	pemBlocks, err := p.load(filename)
	if err != nil {
		return false, err
	}

	_, hasPrivateKey := pemBlocks[RSAPrivateKeyBlock]
	_, hasEncryptedPrivateKey := pemBlocks[EncryptedPrivateKeyBlock]
	if !hasPrivateKey && !hasEncryptedPrivateKey {
		return false, errKMIPMissingPrivateKeyBlock
	}

	if _, hasCertBlock := pemBlocks[CertificateBlock]; !hasCertBlock {
		return hasEncryptedPrivateKey, errKMIPCertificateBlock
	}

	return hasEncryptedPrivateKey, nil
}
