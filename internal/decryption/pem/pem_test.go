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

//go:build unit
// +build unit

package pem

import (
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const DummyCA = `-----BEGIN CERTIFICATE-----
MIIDVDCCAjwCCQDYPsYgBDwJyDANBgkqhkiG9w0BAQsFADBsMQswCQYDVQQGEwJJ
RTEKMAgGA1UECAwBRDEPMA0GA1UEBwwGRHVibGluMRUwEwYDVQQKDAxETyBOT1Qg
VFJVU1QxFTATBgNVBAsMDERPIE5PVCBUUlVTVDESMBAGA1UEAwwJbG9jYWxob3N0
MB4XDTIyMDQxODIyMzgyOFoXDTIzMDQxODIyMzgyOFowbDELMAkGA1UEBhMCSUUx
CjAIBgNVBAgMAUQxDzANBgNVBAcMBkR1YmxpbjEVMBMGA1UECgwMRE8gTk9UIFRS
VVNUMRUwEwYDVQQLDAxETyBOT1QgVFJVU1QxEjAQBgNVBAMMCWxvY2FsaG9zdDCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALpZmIdbDcPzeZsY1u22rRPm
5Mlub0R2rIg3edyXZqFolRRSYCW8BVqCLHztLtSLCkXWkhW36vIVNvX4qSo9BNRS
Nf+JqCRG0+nFjhNm6G6IQuv6rtlzOkIusUnFvKRIigwFytCKWezbMMNBVZP3wq7R
xO0uT7LjHzalDa1MXM1BJQPUWWwlP9YuLY7vvzFS7urKUcoMV4xrHx655VZLEz27
TH6lh7OmIsrtK3nCSqtkrXdqFxRebZvkeWqoW/BH6ixFhUpo2YYO+UsKUIJw3eFi
m+s7QvKoahG0R4gfSUuZJXYkhRBHlRpmQNP4XBaGK+Bb2EYOt2R9O8k8If7e1tsC
AwEAATANBgkqhkiG9w0BAQsFAAOCAQEAmb/xEYiwBvQYqsMl8vtST9kJfgnkzmhm
MwcAhBN1NDmWI43XH0hgD5O/VRGZa77dbUiSuw9QdLl2gByV0FjIJC8k+UgpGFe3
3tQtocNU8+Rn/NnjRTpCrNediAQ3IsCfmBfNBA951QbZozhMqg/keZSGKMu4qPNW
mrpzi/Q+bJWVGXThcPbqrNIN1Pve9QEzwGl9zfdyV306VmMUy/zTt/HJ9YbQhNcj
DRNzqy7EhCPgJmOR3GTiOPKTA+WkbW2XjG9SnZlAlfpR9e30hzVeuEIlHqBldCLU
MKL8BfxVw2hLoaxdJFMZG9oxX228aVRKjx26Id93xmkzvoMOm3RYSQ==
-----END CERTIFICATE-----
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAulmYh1sNw/N5mxjW7batE+bkyW5vRHasiDd53JdmoWiVFFJg
JbwFWoIsfO0u1IsKRdaSFbfq8hU29fipKj0E1FI1/4moJEbT6cWOE2bobohC6/qu
2XM6Qi6xScW8pEiKDAXK0IpZ7Nsww0FVk/fCrtHE7S5PsuMfNqUNrUxczUElA9RZ
bCU/1i4tju+/MVLu6spRygxXjGsfHrnlVksTPbtMfqWHs6Yiyu0recJKq2Std2oX
FF5tm+R5aqhb8EfqLEWFSmjZhg75SwpQgnDd4WKb6ztC8qhqEbRHiB9JS5kldiSF
EEeVGmZA0/hcFoYr4FvYRg63ZH07yTwh/t7W2wIDAQABAoIBAHvuVBtIufnkda5p
MZ88CxUeTG+OC1+r0QhyZJAI/I9B30t5kUnnJmRQCEg48RkXMwjJL7RT8WN4Kmoh
KlfV5t1Ro2nD4dfmZs6kvN04ZkIESwVnXVtuX4NeTDe00sUrHOvr+jsAl9eG2oIw
dDqI3qenCGF4mDZvB/YuhM8I5vr1QpVeCLJLQOwgXcqq1rngqG+iRuTtsj1bbE4D
y/4qpryExNMdn42kY650B9zrEoPcdoV6XcXF0m5xxkgGas/7lnSzA2P0LZU8BVS6
hcPpJZykSmOB4NtF0hDd4VCguCsDzud6SZ7t2Aot43tU79uCC9zIUSVL494PYDef
81QbE1ECgYEA3vuYUQjiDqUOlGpVcmZYvMGcJaUvGg4x08H5mQJVTUyS5OOWoDBR
4tbGULLBjFGwRtl/2exo9hDrzrr9XReB6G2gUu3XAZCUUBNay4w2APnz1Geo8LmX
VhjVgCcgnTOZSivf/gUNhwAZlj+I8oXJ3Y3Et0XIO3f6JG3WiPuMbCkCgYEA1fFk
9yKWuZYAH0cZtpbn4Vm1pjslJkNCy0BGg+pueFDXY24lbPQMoB9uFOcRMWKG5V57
If1uPoPDZR1k4i5s7/Cw+oWL3HMkDr/slwhEqbAKISbfhvqA06cELkNu0VV0cRgW
rAFJNAc++hzYvf7Xny3ophTAG3rf3mIfAmD5S2MCgYAkiPCyBlSTtbOn2axabC6J
7ucYu/H1wPGlEplE2r8DRVKkMi4R3Rjto+cmfcN8rD3HvgdWu4ePGcKpQrYUtK9S
V/P24oVh+kByxlkQFM8cZdfvq3RgzOfg8Xy53K9ZUoUBRCMVSdqnjfqjRZG4uvcS
WBItPT/LjqLrqRuHoj+l0QKBgQCioCkodsFt9yjGncxc8B75PLEI2CKoEC7Aw24W
rmgkywa/DSYjyOuj9+A8wVxfVs7FoeklcDiSCqTHwu1BxRqH1UUiWctz2o5JK/jS
4bUX67n3c04sk1TEDkvuQtIFC9lEcpQhUaTsiKmFg9H5srMCy+nx/Qn+mYt8xsdd
jotRkwKBgQCJfuUc0NS0DVMoS0Ln/7NhICwktjynVDRYDFNtx8Me1aTI92zAcneN
Jtil9C2r7AWekZKiQUhlO4xfFjGex0hU7yF55YzPiQmQFptJm0gGT6t0jcP91ajk
oZRRTayjJmMafVvzNTGF5F3GX6zjnI9Ryyzxo7pJ3+PNwmwsH/rqtg==
-----END RSA PRIVATE KEY-----
`

const DummyCert = `-----BEGIN CERTIFICATE-----
MIIENTCCAx2gAwIBAgIJAPWNjXbYMr7lMA0GCSqGSIb3DQEBCwUAMGwxCzAJBgNV
BAYTAklFMQowCAYDVQQIDAFEMQ8wDQYDVQQHDAZEdWJsaW4xFTATBgNVBAoMDERP
IE5PVCBUUlVTVDEVMBMGA1UECwwMRE8gTk9UIFRSVVNUMRIwEAYDVQQDDAlsb2Nh
bGhvc3QwHhcNMjIwNDE5MTYxNDI5WhcNMjMwOTAxMTYxNDI5WjBsMQswCQYDVQQG
EwJJRTEKMAgGA1UECAwBRDEPMA0GA1UEBwwGRHVibGluMRUwEwYDVQQKDAxETyBO
T1QgVFJVU1QxFTATBgNVBAsMDERPIE5PVCBUUlVTVDESMBAGA1UEAwwJbG9jYWxo
b3N0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv6tWJkTr99TuxWN2
ih7uXVIbjRCd1pLTvmoZxHee4TYbs7zwHCzanbTeqQ2LOZlrqHLwmJ9E+xrkDSsB
mlDfI3J9f5dIBeEZAZDP9GcZ64KCLq4PgdQV0YLPiuwYyEuIPZrDkNY7weVqBpk9
oEf4HLktxHx+zbsp6/SxAMKCYBTcy8wioccdLI8lBLJeVOl/KsuxfkGILoH+ryl5
qBdYGeZzGnOjU4cJVFOCvJ7zJDn2ASGghO7JbmKPotr/NeY0MXEKJR4zHIHyYvRh
Kit5V5bq3DJw5kp0TFkVpjhRaMaLkaP8w97bEvaOthV5fJB94WG44eEuYhuO/xyY
h2SLEwIDAQABo4HZMIHWMIGGBgNVHSMEfzB9oXCkbjBsMQswCQYDVQQGEwJJRTEK
MAgGA1UECAwBRDEPMA0GA1UEBwwGRHVibGluMRUwEwYDVQQKDAxETyBOT1QgVFJV
U1QxFTATBgNVBAsMDERPIE5PVCBUUlVTVDESMBAGA1UEAwwJbG9jYWxob3N0ggkA
2D7GIAQ8CcgwCQYDVR0TBAIwADALBgNVHQ8EBAMCBaAwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMBQGA1UdEQQNMAuCCWxvY2FsaG9zdDANBgkqhkiG9w0B
AQsFAAOCAQEAgKINT8ASLnG/k/+H68iqoPfb49melXKtRiVG5jYlCN8P7v3Yj/AT
m3Wbq/cGayd2sewh4UgvkmUWEuw6OCBsORT/E9+teq7G/XbWK6YGpc7WCzJT0kJD
8sOK2LuRegPM7gEoIZ5KBycVBxB3mLkIyiOeFpCK+ZoW8gd9Ug2ZNK4YAyMDFfW9
yJ7hJThLZmckaMZBY83yrSD3BTevLN22cWphj9Sna7BW+7c5Pqw3W9i4YO4wSmwU
J1FPS2VF0Pz5ORDNp5fgz2JVS4b3k2IQ0dEIXQW3OeBO1i7p+frUOroQFu8ZXLac
romOggcaq3uWOek9yP+3XusUjXWJ3ZPPsA==
-----END CERTIFICATE-----
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAv6tWJkTr99TuxWN2ih7uXVIbjRCd1pLTvmoZxHee4TYbs7zw
HCzanbTeqQ2LOZlrqHLwmJ9E+xrkDSsBmlDfI3J9f5dIBeEZAZDP9GcZ64KCLq4P
gdQV0YLPiuwYyEuIPZrDkNY7weVqBpk9oEf4HLktxHx+zbsp6/SxAMKCYBTcy8wi
occdLI8lBLJeVOl/KsuxfkGILoH+ryl5qBdYGeZzGnOjU4cJVFOCvJ7zJDn2ASGg
hO7JbmKPotr/NeY0MXEKJR4zHIHyYvRhKit5V5bq3DJw5kp0TFkVpjhRaMaLkaP8
w97bEvaOthV5fJB94WG44eEuYhuO/xyYh2SLEwIDAQABAoIBAG+aRD9kQkG8Kouk
rpEeEY0lEgXBdyZJuCFXhklvnYNlDhxKF0VQmLrbZgrpZ/fR7W3X/1/e3TuZHNDO
CdDg5gytzYVNgRJrTzQqLewRXHZVN5gWodDmvQ2RLWemsYdu85VrWBQtqf3spx/Q
eqGpRj7zVELkerEwGejaQXY/y8pFV2tPc/oz2l6v/dV4PNL2USvArvtJIXoVuvk/
W3Ee84YKNOM4H/L9gYQGhTwRJraeqa2vmR7RT49sT/jHoeFR7g99QNY13dZiRvXd
N9wsZXHf8RXaB8flX6MLMc/o0Ojnj19ZgjA5/t09T3LPjA9WDt+nGbDkVzR1Xy0X
jqGXTQECgYEA/By6jM198AQFA+bNUGhAxxOHBP3r+ivLB3cUZQNPU+j9yIAxMx4S
N5O644udkVnPRP99tRIWHCVkZBuH80dUNQrKjYW2Ao/fFDS1eGzrJREUsBzKmVNu
0MKHDi7OYFuPzaKExgxptx9meQ21zsS+mgmgoQ3E6qCxPfKuouyXrQkCgYEAwp/+
zlTS3Bn6v/COvViemgXCgjzDtGp2NQ1maTjx2MsPjn9WfTq7ocMeTWIEC/nS6gbx
6I+mdC04Wd1AMOc+6+2UBcbcuyWy4QYuSX6lG7hbkXACNZaiRSHYVFpchNgGYm+n
F7z1F8YTz1nPMrjhCV9jihbjNtDlFhqNi5nm2jsCgYEAhQpEB3mJM9drLhvlzMC3
LlbHsYKtvF7PzSixwnx0qDsTcXL0g50iz+FNhjZu9/0Eu8x3cc4RjNjOmWVN4LuL
XFJNgVFGMyPo/Kiz+tC/ZdgVqroGz9KPb+q3imx4y7CFumZA2qJCRzhywv7RKkP4
sSDTeyng+E/EOISQU7m2cMECgYEAh96aTAD7k5yvaP/PJnCPiIcs2y8AkRshmrfY
Hu0aKXbZTWmoP5SZGLzWkr8yhAnMLITcrLZcRg6roFDNV1aYnqwlAkNqJVyUHHPs
LHK1YTy68DV51V9ruUd/dqP+ot8M1fuMcw3/LLGjcsYH2CkpMRneq7B+vu3mgB/Z
YPP4LbECgYEAzg2ZEbuI907Y100oXb/UUZGlGBw8DeMXFpvWKL/Jb/iJrpLcIa/e
vHgO9rgkXKNvGIr+cRlsPPc9W/hAtjwQ39YBe/GjTXAfjbLfpOjOOwDKEhq1OwZ9
U7dWGIYfntsGNMmigGYyUY8+RtrhyaUURJJ9OlJ68w1wEh/BRdCFFSQ=
-----END RSA PRIVATE KEY-----
`

/* #nosec */
const DummyPwd = "njs5Ndl1HllX1I2"
const DummyEncryptedCert = `-----BEGIN CERTIFICATE-----
MIIFNTCCBB2gAwIBAgIJAPWNjXbYMr7kMA0GCSqGSIb3DQEBCwUAMGwxCzAJBgNV
BAYTAklFMQowCAYDVQQIDAFEMQ8wDQYDVQQHDAZEdWJsaW4xFTATBgNVBAoMDERP
IE5PVCBUUlVTVDEVMBMGA1UECwwMRE8gTk9UIFRSVVNUMRIwEAYDVQQDDAlsb2Nh
bGhvc3QwHhcNMjIwNDE4MjIzODMwWhcNMjMwODMxMjIzODMwWjBsMQswCQYDVQQG
EwJJRTEKMAgGA1UECAwBRDEPMA0GA1UEBwwGRHVibGluMRUwEwYDVQQKDAxETyBO
T1QgVFJVU1QxFTATBgNVBAsMDERPIE5PVCBUUlVTVDESMBAGA1UEAwwJbG9jYWxo
b3N0MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA3Gj8zsC1WmpoO98K
YzL4ivYVu3uO4HV8TsMy1wGj6GjPyFdh6f8/FIyiZCencS0qyG6ftz7ib30V4jqx
ex02CtoAdpsbu4HmBczPtURlb6+o/C+SOcrfJuOJ2zB8hcBjGEJvtZ6GzoczAK27
5Fu5B4DbKDD/IZzoXj/GB+NTN4nIWjLS9Sg9gFgFCgbVpHVYh9VZ1zlFrY4dVKEC
tFWKC7+ntjtq0Pek63CsZAetd6QvLa05AQKDnolM/hky216yZWWu+GnN9abhdy+9
BDfFiqNSJc/3IsI4SytNNvtjnZFuRIbYX/B3EMWhrEurnyIoaMSMFKgJuuSbALcF
dj01pGpJEqChAadHh6YqZM2Utex/nP/04vllun8ebM0qIjg9XdJLzckjgX6GP4Ro
FxbJwkd5cx/5iVqqkDgOYTYztIeJE3IYXGPCHlclAPVEw6x6A3bpJMZ1jeNXBocd
NWmxYmpQEEcfoelc+WppfRr7CqJ7SVmP+MzGIrS+/egI5dPTqUUBV025cb059L3j
zUL+7UtYtHGEAZFy+Xr/NeQimkOnkIlUdqn6RCPgkiHyHraRjO8Ni2JGkFEVoA2p
X6LceabBglR8scIjhJehI7djUaEhRZ0wYHlzNApv4q4HB02VUsMKRtDNgl+TyvRK
M6FaAiZ+DAhDGsPMvZdL5u8aXB8CAwEAAaOB2TCB1jCBhgYDVR0jBH8wfaFwpG4w
bDELMAkGA1UEBhMCSUUxCjAIBgNVBAgMAUQxDzANBgNVBAcMBkR1YmxpbjEVMBMG
A1UECgwMRE8gTk9UIFRSVVNUMRUwEwYDVQQLDAxETyBOT1QgVFJVU1QxEjAQBgNV
BAMMCWxvY2FsaG9zdIIJANg+xiAEPAnIMAkGA1UdEwQCMAAwCwYDVR0PBAQDAgWg
MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAUBgNVHREEDTALgglsb2Nh
bGhvc3QwDQYJKoZIhvcNAQELBQADggEBAFrrP7RWF+IU/nzFmZJLU6x7zLCGBDuA
+w66OO14NSqpuk4BF5fR+BDudfV9oeAoEUHHD4vIJ77nbqWVdGPUlJkSqXJBFMVA
Pd0+RCMjsMYreH92O9uaJH3/BJzn9teAR0ueoejbDu5UA5Q5qaqP98qV8O8ZBrFC
efSA36jUDVJRK+F7Kt8QQ4BkIAusLzzVa28qY2TZ8KY3bDgqam08gLZ1spJflxrp
MC8HsX53iOH+Fk3khF2VJDkTl9dUnlL4jxzuH+lTK3vKwtzpDPWiQCxu0QJuJhJc
t7xyRO0lwwnpe5w8+qqTlirrR4dT3VYZFe4goMynZeqEgSSM9i5Fjb8=
-----END CERTIFICATE-----
-----BEGIN ENCRYPTED PRIVATE KEY-----
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
IXbLbUL9NcvLyZnehDa7vPWIoQ==
-----END ENCRYPTED PRIVATE KEY-----
`

func TestValidateBlocks(t *testing.T) {
	t.Run("CA with private key and cert blocks is valid", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		_ = afero.WriteFile(fs, "pemfile", []byte(DummyCA), 0600)
		pem := &pemDecoderValidator{fs: fs}

		isEncrypted, err := pem.ValidateBlocks("pemfile")

		assert.NoError(t, err)
		assert.False(t, isEncrypted)
	})

	t.Run("client cert with private key and cert blocks is valid", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		_ = afero.WriteFile(fs, "pemfile", []byte(DummyCert), 0600)
		pem := &pemDecoderValidator{fs: fs}

		isEncrypted, err := pem.ValidateBlocks("pemfile")

		assert.NoError(t, err)
		assert.False(t, isEncrypted)
	})

	t.Run("client cert with encrypted private key and cert blocks is valid", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		_ = afero.WriteFile(fs, "pemfile", []byte(DummyEncryptedCert), 0600)
		pem := &pemDecoderValidator{fs: fs}

		isEncrypted, err := pem.ValidateBlocks("pemfile")

		assert.NoError(t, err)
		assert.True(t, isEncrypted)
	})

	t.Run("client cert without cert block is not valid", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		certContent := strings.ReplaceAll(DummyCert, string(CertificateBlock), "DUMMY PEM BLOCK TYPE")
		_ = afero.WriteFile(fs, "pemfile", []byte(certContent), 0600)
		pem := &pemDecoderValidator{fs: fs}

		isEncrypted, err := pem.ValidateBlocks("pemfile")

		assert.Error(t, err)
		assert.False(t, isEncrypted)
	})

	t.Run("client cert without private key block is not valid", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		certContent := strings.ReplaceAll(DummyCert, string(RSAPrivateKeyBlock), "DUMMY PEM BLOCK TYPE")
		_ = afero.WriteFile(fs, "pemfile", []byte(certContent), 0600)
		pem := &pemDecoderValidator{fs: fs}

		isEncrypted, err := pem.ValidateBlocks("pemfile")

		assert.Error(t, err)
		assert.False(t, isEncrypted)
	})
}

func TestDecode(t *testing.T) {
	t.Run("decode CA", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		_ = afero.WriteFile(fs, "pemfile", []byte(DummyCA), 0600)
		pem := &pemDecoderValidator{fs: fs}

		cert, privateKey, err := pem.Decode("pemfile", "")

		assert.NoError(t, err)
		assert.Contains(t, string(cert), CertificateBlock)
		assert.Contains(t, string(privateKey), RSAPrivateKeyBlock)
	})

	t.Run("decode client cert with encrypted private key using wrong password returns error", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		_ = afero.WriteFile(fs, "pemfile", []byte(DummyEncryptedCert), 0600)
		pem := &pemDecoderValidator{fs: fs}

		cert, privateKey, err := pem.Decode("pemfile", "wrong pwd")

		assert.Error(t, err)
		assert.Nil(t, cert)
		assert.Nil(t, privateKey)
	})

	t.Run("decode client cert with encrypted private key using correct password is successful", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		_ = afero.WriteFile(fs, "pemfile", []byte(DummyEncryptedCert), 0600)
		pem := &pemDecoderValidator{fs: fs}

		cert, privateKey, err := pem.Decode("pemfile", DummyPwd)

		assert.NoError(t, err)
		assert.Contains(t, string(cert), CertificateBlock)
		assert.Contains(t, string(privateKey), RSAPrivateKeyBlock)
	})
}
