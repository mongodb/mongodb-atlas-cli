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

// adapted from: https://github.com/smallstep/crypto/blob/master/pemutil/pkcs8.go
// Apache License 2.0

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"  //nolint:gosec // used as part of the des cbc standard
	"crypto/sha1" //nolint:gosec // used as part of the sha1 standard
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

type prfParam struct {
	Algo      asn1.ObjectIdentifier
	NullParam asn1.RawValue
}

type pbkdf2Params struct {
	Salt           []byte
	IterationCount int
	PrfParam       prfParam `asn1:"optional"`
}

type pbkdf2Algorithms struct {
	Algo         asn1.ObjectIdentifier
	PBKDF2Params pbkdf2Params
}

type pbkdf2Encs struct {
	EncryAlgo asn1.ObjectIdentifier
	IV        []byte
}

type pbes2Params struct {
	KeyDerivationFunc pbkdf2Algorithms
	EncryptionScheme  pbkdf2Encs
}

type encryptedlAlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters pbes2Params
}

type encryptedPrivateKeyInfo struct {
	Algo       encryptedlAlgorithmIdentifier
	PrivateKey []byte
}

var (
	oidPKCS5PBKDF2    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 12}
	oidPBES2          = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 13}
	oidHMACWithSHA256 = asn1.ObjectIdentifier{1, 2, 840, 113549, 2, 9}
	oidAES128CBC      = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 2}
	oidAES196CBC      = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 22}
	oidAES256CBC      = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 1, 42}
	oidDESCBC         = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 7}
	oidD3DESCBC       = asn1.ObjectIdentifier{1, 2, 840, 113549, 3, 7}
)

//nolint
func DecryptPKCS8PrivateKey(data, password []byte) ([]byte, error) {
	var pki encryptedPrivateKeyInfo
	if _, err := asn1.Unmarshal(data, &pki); err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key %w", err)
	}

	if !pki.Algo.Algorithm.Equal(oidPBES2) {
		return nil, errors.New("unsupported encrypted PEM: only PBES2 is supported")
	}

	if !pki.Algo.Parameters.KeyDerivationFunc.Algo.Equal(oidPKCS5PBKDF2) {
		return nil, errors.New("unsupported encrypted PEM: only PBKDF2 is supported")
	}

	encParam := pki.Algo.Parameters.EncryptionScheme
	kdfParam := pki.Algo.Parameters.KeyDerivationFunc.PBKDF2Params

	iv := encParam.IV
	salt := kdfParam.Salt
	iter := kdfParam.IterationCount

	// pbkdf2 hash function
	keyHash := sha1.New
	if kdfParam.PrfParam.Algo.Equal(oidHMACWithSHA256) {
		keyHash = sha256.New
	}

	var symkey []byte
	var block cipher.Block
	var err error
	switch {
	// AES-128-CBC, AES-192-CBC, AES-256-CBC
	case encParam.EncryAlgo.Equal(oidAES128CBC):
		symkey = pbkdf2.Key(password, salt, iter, 16, keyHash)
		block, err = aes.NewCipher(symkey)
	case encParam.EncryAlgo.Equal(oidAES196CBC):
		symkey = pbkdf2.Key(password, salt, iter, 24, keyHash)
		block, err = aes.NewCipher(symkey)
	case encParam.EncryAlgo.Equal(oidAES256CBC):
		symkey = pbkdf2.Key(password, salt, iter, 32, keyHash)
		block, err = aes.NewCipher(symkey)
	// DES, TripleDES
	case encParam.EncryAlgo.Equal(oidDESCBC):
		symkey = pbkdf2.Key(password, salt, iter, 8, keyHash)
		block, err = des.NewCipher(symkey)
	case encParam.EncryAlgo.Equal(oidD3DESCBC):
		symkey = pbkdf2.Key(password, salt, iter, 24, keyHash)
		block, err = des.NewTripleDESCipher(symkey)
	default:
		return nil, fmt.Errorf("unsupported encrypted PEM: unknown algorithm %v", encParam.EncryAlgo)
	}
	if err != nil {
		return nil, err
	}

	data = pki.PrivateKey
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	// Blocks are padded using a scheme where the last n bytes of padding are all
	// equal to n. It can pad from 1 to blocksize bytes inclusive. See RFC 1423.
	// For example:
	//	[x y z 2 2]
	//	[x y 7 7 7 7 7 7 7]
	// If we detect a bad padding, we assume it is an invalid password.
	blockSize := block.BlockSize()
	dlen := len(data)
	if dlen == 0 || dlen%blockSize != 0 {
		return nil, errors.New("error decrypting PEM: invalid padding")
	}

	last := int(data[dlen-1])
	if dlen < last {
		return nil, x509.IncorrectPasswordError
	}
	if last == 0 || last > blockSize {
		return nil, x509.IncorrectPasswordError
	}
	for _, val := range data[dlen-last:] {
		if int(val) != last {
			return nil, x509.IncorrectPasswordError
		}
	}

	return data[:dlen-last], nil
}
