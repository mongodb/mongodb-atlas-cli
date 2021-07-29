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
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type httpClient interface {
	PostForm(string, url.Values) (*http.Response, error)
}

// PostForm makes an POST request by serializing input parameters as a form and parsing the response
// of the same type.
func PostForm(c httpClient, u string, params url.Values, v interface{}) (*atlas.Response, error) {
	resp, err := c.PostForm(u, params)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Ensure the response body is fully read and closed
		// before we reconnect, so that we reuse the same TCP connection.
		// Close the previous response's body. But read at least some of
		// the body so if it's small the underlying TCP connection will be
		// re-used. No need to check for errors: if it fails, the Transport
		// won't reuse it anyway.
		const maxBodySlurpSize = 2 << 10
		if resp.ContentLength == -1 || resp.ContentLength <= maxBodySlurpSize {
			_, _ = io.CopyN(ioutil.Discard, resp.Body, maxBodySlurpSize)
		}

		resp.Body.Close()
	}()

	r := &atlas.Response{Response: resp}
	if err2 := atlas.CheckResponse(resp); err2 != nil {
		return r, err2
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if errors.Is(decErr, io.EOF) {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return r, err
}
