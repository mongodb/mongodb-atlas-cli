package client

import (
	"bufio"
	"fmt"

	"crypto/tls"
	"crypto/x509"

	kmip "github.com/gemalto/kmip-go"
	"github.com/gemalto/kmip-go/kmip14"
	"github.com/gemalto/kmip-go/kmip20"
	"github.com/gemalto/kmip-go/ttlv"
	"github.com/pkg/errors"
)

// Attributes key attributes required by Create request operation.
type Attributes struct {
	CryptographicAlgorithm kmip14.CryptographicAlgorithm
	CryptographicLength    int32
	CryptographicUsageMask kmip14.CryptographicUsageMask
}

// CreateRequestPayload used to Create symmetric key operation.
type CreateRequestPayload struct {
	ObjectType kmip20.ObjectType
	Attributes Attributes
}

// CreateResponsePayload response message for create operation.
type CreateResponsePayload struct {
	UniqueIdentifier string
}

// GetRequestPayload used for Get request operation.
type GetRequestPayload struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
}

// GetResponsePayload response of Get operation.
type GetResponsePayload struct {
	ObjectType       kmip14.ObjectType
	UniqueIdentifier string
	SymmetricKey     kmip.SymmetricKey
	PrivateKey       kmip.PrivateKey
}

// EncryptRequestPayload used for Encrypt request operation.
type EncryptRequestPayload struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
	Data             []byte
}

// EncryptResponsePayload response of Encrypt operation.
type EncryptResponsePayload struct {
	UniqueIdentifier string
	Data             []byte
	IVCounterNonce   []byte
}

// DecryptRequestPayload used for Decrypt request operation.
type DecryptRequestPayload struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
	Data             []byte
	IVCounterNonce   []byte
}

// DecryptResponsePayload response of Decrypt operation.
type DecryptResponsePayload struct {
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

var KmipV10 = Version{Major: 1, Minor: 0} // first KMIP version
var KmipV12 = Version{Major: 1, Minor: 2} //nolint:gomnd // KMIP version that implemented encrypt / decrypt
var KmipV20 = Version{Major: 2, Minor: 0} //nolint:gomnd // KMIP major version change (create operation signature changed)

var versions = map[Version]bool{KmipV10: true, KmipV12: true, KmipV20: true}

// cipherSuites is a list of enabled TLS 1.0–1.2 cipher suites.
var cipherSuites = []uint16{
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
}

// KMIPClient client used to communicate with a KMIP speaking server.
type KMIPClient struct {
	version       Version
	tlsConfig     tls.Config
	requestHeader kmip.RequestHeader
	ip            string
	port          int
}

// KMIPClientConfig structure used to configure a KMIP client.
type KMIPClientConfig struct {
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

// NewKmipClient creates a new KMIP client and initializes all the values required for establishing connection.
func NewKmipClient(config *KMIPClientConfig) (*KMIPClient, error) {
	if _, found := versions[config.Version]; !found {
		return nil, errors.Errorf("invalid KMIP version %+v", config.Version)
	}

	if config.IP == "" {
		return nil, errors.New("server IP is not provided")
	}

	if config.Port == 0 {
		return nil, errors.New("server port is not provided")
	}

	if config.ClientCertificate == nil {
		return nil, errors.New("client certificate is not provided")
	}

	if config.ClientPrivateKey == nil {
		return nil, errors.New("client private key is not provided")
	}

	if config.RootCertificate == nil {
		return nil, errors.New("root certificate is not provided")
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(config.RootCertificate)

	certificate, err := tls.X509KeyPair(config.ClientCertificate, config.ClientPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load client key and certificate")
	}

	hostname := config.Hostname
	if hostname == "" {
		hostname = config.IP
	}

	kc := &KMIPClient{
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
			ServerName:   hostname,
			CipherSuites: cipherSuites,
			RootCAs:      rootCAs,
			Certificates: []tls.Certificate{certificate},
			MinVersion:   tls.VersionTLS12,
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

// sendRequest sends a request message to KMIP server.
func (kc *KMIPClient) sendRequest(payload interface{}, operation kmip14.Operation) (*kmip.ResponseBatchItem, *ttlv.Decoder, error) {
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

	var responseMessage kmip.ResponseMessage
	err = ttlvDecoder.DecodeValue(&responseMessage, response)
	if err != nil {
		return nil, nil, err
	}

	if responseMessage.BatchItem[0].ResultStatus != kmip14.ResultStatusSuccess {
		return nil, nil, errors.Errorf("KMIP request failed with reason %s", responseMessage.BatchItem[0].ResultMessage)
	}

	return &responseMessage.BatchItem[0], ttlvDecoder, nil
}

// GetSymmetricKey retrieves a symmetric key from KMIP server.
func (kc *KMIPClient) GetSymmetricKey(keyID string) ([]byte, error) {
	payload := GetRequestPayload{
		UniqueIdentifier: kmip20.UniqueIdentifierValue{Text: keyID},
	}

	batchItem, decoder, err := kc.sendRequest(payload, kmip14.OperationGet)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform get operation")
	}

	var responsePayload GetResponsePayload
	err = decoder.DecodeValue(&responsePayload, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode get response payload")
	}

	var keyValue KeyValue
	err = decoder.DecodeValue(&keyValue, responsePayload.SymmetricKey.KeyBlock.KeyValue.(ttlv.TTLV))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode symmetric keyblock")
	}

	return keyValue.KeyMaterial, nil
}

// CreateSymmetricKey creates a symmetric key on KMIP server.
func (kc *KMIPClient) CreateSymmetricKey(length int32) (string, error) {
	var payload interface{}
	if kc.version.Major >= KmipV20.Major {
		payload = CreateRequestPayload{
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
		return "", errors.Wrap(err, "failed to perform create symmetric key operation")
	}

	var respPayload CreateResponsePayload
	err = decoder.DecodeValue(&respPayload, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return "", errors.Wrap(err, "failed to decode create symmetric key response payload")
	}

	return respPayload.UniqueIdentifier, nil
}

// Encrypt encrypts data with an existing managed object stored by the KMIP server.
func (kc *KMIPClient) Encrypt(keyID string, data []byte) (*EncryptResponsePayload, error) {
	payload := EncryptRequestPayload{
		UniqueIdentifier: kmip20.UniqueIdentifierValue{Text: keyID},
		Data:             data,
	}

	batchItem, decoder, err := kc.sendRequest(payload, kmip14.OperationEncrypt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform encrypt operation")
	}

	var respPayload EncryptResponsePayload
	err = decoder.DecodeValue(&respPayload, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode encrypt response payload")
	}

	return &respPayload, nil
}

// Decrypt decrypts data with an existing managed object stored by the KMIP server.
func (kc *KMIPClient) Decrypt(keyID string, data, iv []byte) (*DecryptResponsePayload, error) {
	payLoad := DecryptRequestPayload{
		UniqueIdentifier: kmip20.UniqueIdentifierValue{Text: keyID},
		Data:             data,
		IVCounterNonce:   iv,
	}

	batchItem, decoder, err := kc.sendRequest(payLoad, kmip14.OperationDecrypt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform decrypt operation")
	}

	var respPayload DecryptResponsePayload
	err = decoder.DecodeValue(&respPayload, batchItem.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode decrypt response payload")
	}

	return &respPayload, nil
}
