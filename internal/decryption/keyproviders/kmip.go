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

	"github.com/mongodb/mongocli/internal/decryption/aes"
	"github.com/mongodb/mongocli/internal/decryption/kmip"
	"go.mongodb.org/mongo-driver/bson"
)

type KMIPKeyWrapMethod string

const (
	KMIPKeyWrapMethodGet     KMIPKeyWrapMethod = "get"
	KMIPKeyWrapMethodEncrypt KMIPKeyWrapMethod = "encrypt"
)

// LocalKeyIdentifier config for the KMIP speaking server used to encrypt the Log Encryption Key (LEK).
type KMIPKeyIdentifier struct {
	KeyStoreIdentifier
	UniqueKeyID               string
	ServerName                []string
	ServerPort                int
	KeyWrapMethod             KMIPKeyWrapMethod
	ServerCAFileName          string
	ClientCertificateFileName string
}

// KMIPEncryptedKey encrypted LEK and tag, BSON marshaled.
type KMIPEncryptedKey struct {
	IV  []byte
	Key []byte
}

// DecryptKey decrypts LEK using KMIP get or decrypt methods.
func (ki *KMIPKeyIdentifier) DecryptKey(encryptedKey []byte) ([]byte, error) {
	kmipEncryptedKey, err := ki.decodeEncryptedKey(encryptedKey)
	if err != nil {
		return nil, err
	}

	kmipClient, err := ki.kmipClient()
	if err != nil {
		return nil, err
	}

	if ki.KeyWrapMethod == KMIPKeyWrapMethodEncrypt {
		return ki.decryptWithKeyWrapMethodEncrypt(kmipClient, kmipEncryptedKey)
	}
	return ki.decryptWithKeyWrapMethodGet(kmipClient, kmipEncryptedKey)
}

func (ki *KMIPKeyIdentifier) decryptWithKeyWrapMethodEncrypt(kmipClient *kmip.Client, kmipEncryptedKey *KMIPEncryptedKey) ([]byte, error) {
	decryptedLEK, err := kmipClient.Decrypt(ki.UniqueKeyID, kmipEncryptedKey.Key, kmipEncryptedKey.IV)
	if err != nil {
		return nil, err
	}
	return decryptedLEK.Data, nil
}

func (ki *KMIPKeyIdentifier) decryptWithKeyWrapMethodGet(kmipClient *kmip.Client, kmipEncryptedKey *KMIPEncryptedKey) ([]byte, error) {
	kek, err := kmipClient.GetSymmetricKey(ki.UniqueKeyID)
	if err != nil {
		return nil, err
	}

	tag := kmipEncryptedKey.Key[0:12]
	cipherText := kmipEncryptedKey.Key[12:]
	aad := []byte(ki.UniqueKeyID)
	iv := kmipEncryptedKey.IV

	gcm := &aes.GCMInput{Key: kek, AAD: aad, IV: iv, Tag: tag}
	decryptedLEK, err := gcm.Decrypt(cipherText)
	if err != nil {
		return nil, err
	}
	return decryptedLEK, nil
}

func (ki *KMIPKeyIdentifier) decodeEncryptedKey(encryptedKey []byte) (*KMIPEncryptedKey, error) {
	var kmipEncryptedKey KMIPEncryptedKey
	if err := bson.Unmarshal(encryptedKey, &kmipEncryptedKey); err != nil {
		return nil, err
	}
	return &kmipEncryptedKey, nil
}

func (ki *KMIPKeyIdentifier) kmipClient() (*kmip.Client, error) {
	clientCertAndKey, err := os.ReadFile(ki.ClientCertificateFileName)
	if err != nil {
		return nil, err
	}

	rootCA, err := os.ReadFile(ki.ServerCAFileName)
	if err != nil {
		return nil, err
	}

	version := kmip.V12
	if ki.KeyWrapMethod == KMIPKeyWrapMethodGet {
		version = kmip.V10
	}

	return kmip.NewClient(&kmip.Config{
		Version:           version,
		IP:                ki.ServerName[0],
		Port:              ki.ServerPort,
		Hostname:          ki.ServerName[0],
		RootCertificate:   rootCA,
		ClientPrivateKey:  clientCertAndKey,
		ClientCertificate: clientCertAndKey,
	})
}
