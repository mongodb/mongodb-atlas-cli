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
	"strings"

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

	// Header
	UniqueKeyID   string
	ServerNames   []string
	ServerPort    int
	KeyWrapMethod KMIPKeyWrapMethod

	// CLI
	ServerCAFileName          string
	ClientCertificateFileName string
}

// KMIPEncryptedKey encrypted LEK and tag, BSON marshaled.
type KMIPEncryptedKey struct {
	IV  []byte
	Key []byte
}

var ErrKMIPServerCAMissing = errors.New("server CA missing")
var ErrKMIPClientCertificateMissing = errors.New("client certificate missing")

func (ki *KMIPKeyIdentifier) ValidateCredentials() error {
	if ki.ServerCAFileName == "" || ki.ClientCertificateFileName == "" {
		fmt.Fprintf(os.Stderr, `No credentials found for resource: KMIP uniqueKeyID="%v" serverNames="%v" serverPort="%v" keyWrapMethod="%v"
`, ki.UniqueKeyID, strings.Join(ki.ServerNames, "; "), ki.ServerPort, ki.KeyWrapMethod)
	}

	if ki.ServerCAFileName == "" {
		f, err := provideInput("Provide server CA filename:", "")
		if err != nil {
			return err
		}
		ki.ServerCAFileName = f
		if ki.ServerCAFileName == "" {
			return ErrKMIPServerCAMissing
		}
	}

	if ki.ClientCertificateFileName == "" {
		f, err := provideInput("Provide client certificate filename:", "")
		if err != nil {
			return err
		}
		ki.ClientCertificateFileName = f
		if ki.ClientCertificateFileName == "" {
			return ErrKMIPClientCertificateMissing
		}
	}

	return nil
}

// DecryptKey decrypts LEK using KMIP get or decrypt methods.
func (ki *KMIPKeyIdentifier) DecryptKey(encryptedKey []byte) ([]byte, error) {
	if len(ki.ServerNames) == 0 {
		return nil, errors.New("server name is not provided")
	}

	kmipEncryptedKey, err := ki.decodeEncryptedKey(encryptedKey)
	if err != nil {
		return nil, err
	}

	var clientError error
	collectErr := func(err error, serverName string) {
		if clientError == nil {
			clientError = err
		} else {
			clientError = fmt.Errorf("'%s': %s", serverName, err.Error())
		}
	}

	for _, serverName := range ki.ServerNames {
		kmipClient, err := ki.kmipClient(serverName)
		if err != nil {
			// init KMIP client error (invalid config), skip other KMIP servers
			return nil, err
		}

		var key []byte
		if ki.KeyWrapMethod == KMIPKeyWrapMethodEncrypt {
			key, err = ki.decryptWithKeyWrapMethodEncrypt(kmipClient, kmipEncryptedKey)
		} else {
			key, err = ki.decryptWithKeyWrapMethodGet(kmipClient, kmipEncryptedKey)
		}

		if err == nil {
			// key successfully decrypted, skip other KMIP servers and return decrypted key
			return key, nil
		}
		collectErr(err, serverName)
	}

	return nil, clientError
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

func (ki *KMIPKeyIdentifier) kmipClient(serverName string) (*kmip.Client, error) {
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
		Hostname:          serverName,
		Port:              ki.ServerPort,
		RootCertificate:   rootCA,
		ClientPrivateKey:  clientCertAndKey,
		ClientCertificate: clientCertAndKey,
	})
}
