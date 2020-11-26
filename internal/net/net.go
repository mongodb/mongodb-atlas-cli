// Copyright 2020 MongoDB Inc
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
package net

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
)

// APIURIs is the URIs of the services used by newIPAddress to get the client's public IP.
var APIURIs = []string{
	"https://api.ipify.org",
	"http://myexternalip.com/raw",
	"http://ipinfo.io/ip",
	"http://ipecho.net/plain",
	"http://icanhazip.com",
	"http://ifconfig.me/ip",
	"http://ident.me",
	"http://checkip.amazonaws.com",
	"http://bot.whatismyipaddress.com",
	"http://whatismyip.akamai.com",
	"http://wgetip.com",
	"http://ip.tyk.nu",
}

func IPAddress() (string, error) {
	return ipAddress(APIURIs)
}

// IPAddress returns client's public ip
func ipAddress(services []string) (string, error) {
	publicIP := ""
	for _, uri := range services {
		publicIP = ipAddressFromService(uri)
		if publicIP != "" {
			break
		}
	}

	if publicIP == "" {
		return publicIP, errors.New("error in finding your public IP")
	}

	return publicIP, nil
}

func ipAddressFromService(uri string) string {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		uri,
		nil,
	)

	req.Header.Add("Accept", "application/json")

	if err == nil {
		res, err := http.DefaultClient.Do(req)

		if err == nil {
			responseBytes, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err == nil {
				return string(responseBytes)
			}
		}
	}
	return ""
}
