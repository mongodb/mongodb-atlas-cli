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

package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1" //nolint:gosec // required for notary service
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

const (
	pbkdf2Iterations = 1000
	pbkdf2KeyLength  = 16
)

func generateAuthTokenString(authTokenPassword string) string {
	salt := []byte(authTokenPassword)
	startIndex := len(salt)/2 - 1 //nolint:gomnd // half the length of the salt
	for i := startIndex; i >= 0; i-- {
		opp := len(salt) - 1 - i
		salt[i], salt[opp] = salt[opp], salt[i]
	}

	authKey := pbkdf2.Key([]byte(authTokenPassword), salt, pbkdf2Iterations, pbkdf2KeyLength, sha1.New)
	dateStr := time.Now().String()
	signedData := hmac.New(sha1.New, authKey)
	signedData.Write([]byte(dateStr))
	rawSignature := signedData.Sum(nil)
	return fmt.Sprintf("%x%s", rawSignature, dateStr)
}

func makeRequest(ctx context.Context, url, method, contentType string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func generateMultiPartForm(filePath string, fields map[string]string) (buffer *bytes.Buffer, contentType string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, "", err
	}
	for k, v := range fields {
		err = writer.WriteField(k, v)
		if err != nil {
			return nil, "", err
		}
	}

	return body, writer.FormDataContentType(), nil
}

func sign(notaryURL, filePath, notarySigningKey, notarySigningComment, notaryAuthToken string) error {
	buffer, contentType, err := generateMultiPartForm(filePath, map[string]string{
		"key":        notarySigningKey,
		"comment":    notarySigningComment,
		"auth_token": notaryAuthToken,
	})
	if err != nil {
		return err
	}

	signingURL := fmt.Sprintf("%s/api/sign", notaryURL)
	body, err := makeRequest(context.Background(), signingURL, "POST", contentType, buffer)
	if err != nil {
		return err
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return err
	}

	downloadURL := jsonResponse["permalink"].(string)

	tmpDir := os.TempDir()
	if err != nil {
		return err
	}

	tmpFile := filepath.Join(tmpDir, path.Base(filePath))

	err = downloadFile(context.Background(), notaryURL+downloadURL, tmpFile)
	if err != nil {
		return err
	}

	return os.Rename(tmpFile, filePath)
}

func downloadFile(ctx context.Context, url, localPath string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("download failed: %s", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return os.WriteFile(localPath, b, 0600) //nolint:gomnd // read only permission
}

func main() {
	key := os.Getenv("NOTARY_SIGNING_KEY")
	comment := os.Getenv("NOTARY_SIGNING_COMMENT")
	authToken := os.Getenv("NOTARY_AUTH_TOKEN")
	url := os.Getenv("NOTARY_URL")

	var filePath string

	flag.StringVar(&filePath, "file", "", "file to sign")
	flag.Parse()

	if key == "" {
		log.Fatalln("You must specify $NOTARY_SIGNING_KEY")
	}

	if comment == "" {
		log.Fatalln("You must specify $NOTARY_SIGNING_COMMENT")
	}

	if authToken == "" {
		log.Fatalln("You must specify $NOTARY_AUTH_TOKEN")
	}

	if url == "" {
		log.Fatalln("You must specify $NOTARY_URL")
	}

	if filePath == "" {
		log.Fatalln("You must specify a file to sign")
	}

	err := sign(url, filePath, key, comment, generateAuthTokenString(authToken))
	if err != nil {
		log.Fatalln(err)
	}
}
