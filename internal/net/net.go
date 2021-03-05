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
	"io/ioutil"
	"net/http"
	"strings"
)

const APIURI = "http://checkip.amazonaws.com"

func IPAddress() string {
	return ipAddressFromAPI(APIURI)
}

// ipAddressFromAPI gets the client's public ip by calling the input endpoint
func ipAddressFromAPI(uri string) string {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		uri,
		nil,
	)

	if err != nil {
		return ""
	}

	req.Header.Add("Accept", "text/plain")
	res, err := http.DefaultClient.Do(req)

	defer func() {
		_ = res.Body.Close()
	}()

	if err == nil {
		responseBytes, err1 := ioutil.ReadAll(res.Body)
		if err1 == nil {
			return strings.TrimSpace(string(responseBytes))
		}
	}

	return ""
}
