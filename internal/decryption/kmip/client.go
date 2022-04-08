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

package kmip

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"

	kmip "github.com/gemalto/kmip-go"
	"github.com/gemalto/kmip-go/kmip14"
	"github.com/gemalto/kmip-go/kmip20"
	"github.com/gemalto/kmip-go/ttlv"
)

// Attributes key attributes required by Create request operation.
type Attributes struct {
	CryptographicAlgorithm kmip14.CryptographicAlgorithm
	CryptographicLength    int32
	CryptographicUsageMask kmip14.CryptographicUsageMask
}

// CreateRequestV20 used to Create symmetric key operation for KMIP 2.0+ servers.
type CreateRequestV20 struct {
	ObjectType kmip20.ObjectType
	Attributes Attributes
}

// CreateResponse response message for create operation.
type CreateResponse struct {
	UniqueIdentifier string
}

// GetRequest used for Get request operation.
type GetRequest struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
}

// GetResponse response of Get operation.
type GetResponse struct {
	ObjectType       kmip14.ObjectType
	UniqueIdentifier string
	SymmetricKey     kmip.SymmetricKey
	PrivateKey       kmip.PrivateKey
}

// EncryptRequest used for Encrypt request operation.
type EncryptRequest struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
	Data             []byte
}

// EncryptResponse response of Encrypt operation.
type EncryptResponse struct {
	UniqueIdentifier string
	Data             []byte
	IVCounterNonce   []byte
}

// DecryptRequest used for Decrypt request operation.
type DecryptRequest struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
	Data             []byte
	IVCounterNonce   []byte
}

// DecryptResponse response of Decrypt operation.
type DecryptResponse struct {
	UniqueIdentifier string
	Data             []byte
}

// Symmetric key value.
type KeyValue struct {
	KeyMaterial []byte
}

// KMIP protocol version.
type Version struct {
	Major int
	Minor int
}

var V10 = Version{Major: 1, Minor: 0} // first KMIP version
var V12 = Version{Major: 1, Minor: 2} //nolint:gomnd // KMIP version that implemented encrypt / decrypt
var V20 = Version{Major: 2, Minor: 0} //nolint:gomnd // KMIP major version change (create operation signature changed)

var versions = map[Version]bool{V10: true, V12: true, V20: true}

// cipherSuites is a list of enabled TLS 1.0–1.2 cipher suites.
var cipherSuites = []uint16{
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
}

// Client client used to communicate with a KMIP speaking server.
type Client struct {
	version       Version
	tlsConfig     tls.Config
	requestHeader kmip.RequestHeader
	ip            string
	port          int
}

// Config structure used to configure a KMIP client.
type Config struct {
	Version           Version
	IP                string
	Port              int
	Hostname          string
	Username          string
	Password          string
	ClientPrivateKey  []byte
	ClientCertificate []byte
	RootCertificate   []byte
}

var (
	ErrCertificateLoad               = errors.New("failed to load certificate")
	ErrCertificateLoadRoot           = fmt.Errorf("%w: %s", ErrCertificateLoad, "root certificate")
	ErrCertificateLoadClient         = fmt.Errorf("%w: %s", ErrCertificateLoad, "client certificate")
	ErrKMIPVersionInvalid            = errors.New("invalid KMIP version")
	ErrServerHostnameIPMissing       = errors.New("both server hostname and IP are not provided")
	ErrServerPortMissing             = errors.New("server port is not provided")
	ErrRootCertMissing               = errors.New("root certificate is not provided")
	ErrClientCertMissing             = errors.New("client certificate is not provided")
	ErrClientKeyMissing              = errors.New("client private key is not provided")
	ErrKMIPReqFailure                = errors.New("kmip request failure")
	ErrKMIPGetOpFailure              = errors.New("failed to perform get operation")
	ErrKMIPDecodeFailure             = errors.New("failed to decode")
	ErrKMIPDecodeKeyBlockFailure     = errors.New("failed to decode key block")
	ErrKMIPPerformCreateSymmetricKey = errors.New("failed to perform KMIP create symmetric key operation")
	ErrKMIPDecodeCreateSymmetricKey  = errors.New("failed to decode KMIP create symmetric key response")
	ErrKMIPPerformEncrypt            = errors.New("failed to perform KMIP encrypt operation")
	ErrKMIPDecodeEncrypt             = errors.New("failed to decode KMIP encrypt response")
	ErrKMIPPerformDecrypt            = errors.New("failed to perform KMIP decrypt operation")
	ErrKMIPDecodeDecrypt             = errors.New("failed to decode KMIP decrypt response")
)

// NewClient creates a new KMIP client and initializes all the values required for establishing connection.
func NewClient(config *Config) (*Client, error) {
	if err := validate(config); err != nil {
		return nil, err
	}

	rootCAs := x509.NewCertPool()
	if !rootCAs.AppendCertsFromPEM(config.RootCertificate) {
		return nil, ErrCertificateLoadRoot
	}

	certificate, err := tls.X509KeyPair(config.ClientCertificate, config.ClientPrivateKey)
	if err != nil {
		return nil, ErrCertificateLoadClient
	}

	hostname := config.Hostname
	if hostname == "" {
		hostname = config.IP
	}

	kc := &Client{
		version: config.Version,
		ip:      config.IP,
		port:    config.Port,
		requestHeader: kmip.RequestHeader{
			ProtocolVersion: kmip.ProtocolVersion{
				ProtocolVersionMajor: config.Version.Major,
				ProtocolVersionMinor: config.Version.Minor,
			},
			BatchCount: 1,
		},
		tlsConfig: tls.Config{
			ServerName:         hostname,
			CipherSuites:       cipherSuites,
			RootCAs:            rootCAs,
			Certificates:       []tls.Certificate{certificate},
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,
		},
	}

	if config.Username != "" && config.Password != "" {
		kc.requestHeader.Authentication = &kmip.Authentication{
			Credential: []kmip.Credential{
				{
					CredentialType: kmip14.CredentialTypeUsernameAndPassword,
					CredentialValue: kmip.UsernameAndPasswordCredentialValue{
						Username: config.Username,
						Password: config.Password,
					},
				},
			},
		}
	}

	return kc, nil
}

func validate(config *Config) error {
	if _, found := versions[config.Version]; !found {
		return fmt.Errorf("%w: %+v", ErrKMIPVersionInvalid, config.Version)
	}

	if config.Hostname == "" && config.IP == "" {
		return ErrServerHostnameIPMissing
	}

	if config.Port == 0 {
		return ErrServerPortMissing
	}

	if config.RootCertificate == nil {
		return ErrRootCertMissing
	}

	if config.ClientCertificate == nil {
		return ErrClientCertMissing
	}

	if config.ClientPrivateKey == nil {
		return ErrClientKeyMissing
	}

	return nil
}

// sendRequest sends a request message to KMIP server.
func (kc *Client) sendRequest(payload interface{}, operation kmip14.Operation) (*kmip.ResponseBatchItem, *ttlv.Decoder, error) {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", kc.ip, kc.port), &kc.tlsConfig)
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()

	if _, certErr := conn.ConnectionState().PeerCertificates[0].Verify(x509.VerifyOptions{Roots: kc.tlsConfig.RootCAs}); certErr != nil {
		return nil, nil, certErr
	}

	requestMessage, err := ttlv.Marshal(kmip.RequestMessage{
		RequestHeader: kc.requestHeader,
		BatchItem: []kmip.RequestBatchItem{
			{
				Operation:      operation,
				RequestPayload: payload,
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	_, err = conn.Write(requestMessage)
	if err != nil {
		return nil, nil, err
	}

	ttlvDecoder := ttlv.NewDecoder(bufio.NewReader(conn))
	response, err := ttlvDecoder.NextTTLV()
	if err != nil {
		return nil, nil, err
	}

	var decodedResponse kmip.ResponseMessage
	err = ttlvDecoder.DecodeValue(&decodedResponse, response)
	if err != nil {
		return nil, nil, err
	}

	if decodedResponse.BatchItem[0].ResultStatus != kmip14.ResultStatusSuccess {
		return nil, nil, fmt.Errorf("%w: %s", ErrKMIPReqFailure, decodedResponse.BatchItem[0].ResultMessage)
	}

	return &decodedResponse.BatchItem[0], ttlvDecoder, nil
}

// GetSymmetricKey retrieves a symmetric key from KMIP server.
func (kc *Client) GetSymmetricKey(keyID string) ([]byte, error) {
	payload := GetRequest{
		UniqueIdentifier: kmip20.UniqueIdentifierValue{Text: keyID},
	}

	batchItem, decoder, err := kc.sendRequest(payload, kmip14.OperationGet)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrKMIPGetOpFailure, err)
	}

	var response GetResponse
	err = decoder.DecodeValue(&response, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrKMIPDecodeFailure, err)
	}

	var keyValue KeyValue
	err = decoder.DecodeValue(&keyValue, response.SymmetricKey.KeyBlock.KeyValue.(ttlv.TTLV))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrKMIPDecodeKeyBlockFailure, err)
	}

	return keyValue.KeyMaterial, nil
}

// CreateSymmetricKey creates a symmetric key on KMIP server.
func (kc *Client) CreateSymmetricKey(length int32) (*string, error) {
	var payload interface{}
	if kc.version.Major >= V20.Major {
		payload = CreateRequestV20{
			ObjectType: kmip20.ObjectTypeSymmetricKey,
			Attributes: Attributes{
				CryptographicAlgorithm: kmip14.CryptographicAlgorithmAES,
				CryptographicLength:    length,
				CryptographicUsageMask: kmip14.CryptographicUsageMaskEncrypt | kmip14.CryptographicUsageMaskDecrypt,
			},
		}
	} else {
		payload = kmip.CreateRequestPayload{
			ObjectType: kmip14.ObjectTypeSymmetricKey,
			TemplateAttribute: kmip.TemplateAttribute{
				Attribute: []kmip.Attribute{
					{
						AttributeName:  "Cryptographic Algorithm",
						AttributeValue: kmip14.CryptographicAlgorithmAES,
					},
					{
						AttributeName:  "Cryptographic Length",
						AttributeValue: length,
					},
					{
						AttributeName:  "Cryptographic Usage Mask",
						AttributeValue: kmip14.CryptographicUsageMaskEncrypt | kmip14.CryptographicUsageMaskDecrypt,
					},
				},
			},
		}
	}

	batchItem, decoder, err := kc.sendRequest(payload, kmip14.OperationCreate)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKMIPPerformCreateSymmetricKey, err)
	}

	var response CreateResponse
	err = decoder.DecodeValue(&response, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKMIPDecodeCreateSymmetricKey, err)
	}

	return &response.UniqueIdentifier, nil
}

// Encrypt encrypts data with an existing managed object stored by the KMIP server.
func (kc *Client) Encrypt(keyID string, data []byte) (*EncryptResponse, error) {
	payload := EncryptRequest{
		UniqueIdentifier: kmip20.UniqueIdentifierValue{Text: keyID},
		Data:             data,
	}

	batchItem, decoder, err := kc.sendRequest(payload, kmip14.OperationEncrypt)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKMIPPerformEncrypt, err)
	}

	var response EncryptResponse
	err = decoder.DecodeValue(&response, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKMIPDecodeEncrypt, err)
	}

	return &response, nil
}

// Decrypt decrypts data with an existing managed object stored by the KMIP server.
func (kc *Client) Decrypt(keyID string, data, iv []byte) (*DecryptResponse, error) {
	payLoad := DecryptRequest{
		UniqueIdentifier: kmip20.UniqueIdentifierValue{Text: keyID},
		Data:             data,
		IVCounterNonce:   iv,
	}

	batchItem, decoder, err := kc.sendRequest(payLoad, kmip14.OperationDecrypt)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKMIPPerformDecrypt, err)
	}

	var response DecryptResponse
	err = decoder.DecodeValue(&response, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrKMIPDecodeDecrypt, err)
	}

	return &response, nil
}
