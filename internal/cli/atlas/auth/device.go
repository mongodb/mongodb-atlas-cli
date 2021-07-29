// Copyright 2021 MongoDB Inc
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

package auth

import (
	"errors"
	"net/url"
	"strings"
	"time"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

// CodeResponse holds information about the authorization-in-progress.
type CodeResponse struct {
	// The user verification code is displayed on the device so the user can enter the code in a browser.
	UserCode string `json:"user_code"`
	// The verification URL where users need to enter the UserCode.
	VerificationURI string `json:"verification_uri"`

	// The device verification code is 40 characters and used to verify the device.
	DeviceCode string `json:"device_code"`
	// The number of seconds before the DeviceCode and UserCode expire.
	ExpiresIn int `json:"expires_in"`
	// The minimum number of seconds that must pass before you can make a new access token request to
	// complete the device authorization.
	Interval int `json:"interval"`

	timeNow   func() time.Time
	timeSleep func(time.Duration)
}

// RequestCode initiates the authorization flow by requesting a code from uri.
func RequestCode(c httpClient, uri, clientID string, scopes []string) (*CodeResponse, error) {
	var r *CodeResponse
	_, err := PostForm(c,
		uri,
		url.Values{
			"client_id": {clientID},
			"scope":     {strings.Join(scopes, " ")},
		},
		&r,
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

const grantType = "urn:ietf:params:oauth:grant-type:device_code"

// ErrTimeout is thrown when polling the server for the granted token has timed out.
var ErrTimeout = errors.New("authentication timed out")

// PollToken polls the server at pollURL until an access token is granted or denied.
func PollToken(c httpClient, pollURL, clientID string, code *CodeResponse) (*AccessToken, error) {
	timeNow := code.timeNow
	if timeNow == nil {
		timeNow = time.Now
	}
	timeSleep := code.timeSleep
	if timeSleep == nil {
		timeSleep = time.Sleep
	}

	checkInterval := time.Duration(code.Interval) * time.Second
	expiresAt := timeNow().Add(time.Duration(code.ExpiresIn) * time.Second)

	for {
		timeSleep(checkInterval)
		var r *AccessToken
		_, err := PostForm(
			c,
			pollURL, url.Values{
				"client_id":   {clientID},
				"device_code": {code.DeviceCode},
				"grant_type":  {grantType},
			}, &r)
		var target *atlas.ErrorResponse
		if errors.As(err, &target) && target.ErrorCode == "THROTTLED" {
			continue
		}
		if err != nil {
			return nil, err
		}

		if timeNow().After(expiresAt) {
			return nil, ErrTimeout
		}
		return r, nil
	}
}
