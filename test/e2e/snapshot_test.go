// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e_test

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

var fileIDs = map[string]map[string]int{}
var randInts = map[string]map[int64]int64{}
var currentTestName string

func snapshotDir(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if strings.HasSuffix(dir, "test/e2e") {
		dir = path.Join(dir, "snapshots", t.Name())
	} else {
		dir = path.Join(dir, "test/e2e/snapshots", t.Name())
	}

	if err := enforceSnapshotDir(dir); err != nil {
		t.Fatal(err)
	}

	return dir
}

func enforceSnapshotDir(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func snapshotName(t *testing.T, r *http.Request) string {
	t.Helper()

	dir := snapshotDir(t)

	filename := fmt.Sprintf("%s/%s_%s", dir, r.Method, strings.ReplaceAll(strings.ReplaceAll(r.URL.Path, "/", "_"), ":", "_"))

	if fileIDs[t.Name()] == nil {
		fileIDs[t.Name()] = map[string]int{}
	}
	fileIDs[t.Name()][filename]++

	id := fileIDs[t.Name()][filename]

	return fmt.Sprintf("%s_%d.json", filename, id)
}

func updateSnapshots() bool {
	return os.Getenv("UPDATE_SNAPSHOTS") == "true"
}

func RandInt(maximum int64) (int64, error) {
	if updateSnapshots() {
		i, err := rand.Int(rand.Reader, big.NewInt(maximum))
		if err != nil {
			return 0, err
		}

		r := i.Int64()

		if randInts[currentTestName] == nil {
			randInts[currentTestName] = map[int64]int64{}
		}

		randInts[currentTestName][0]++

		id := randInts[currentTestName][0]
		randInts[currentTestName][id] = r

		return r, nil
	}

	randInts[currentTestName][0]++
	id := randInts[currentTestName][0]

	return randInts[currentTestName][id], nil
}

func snapshotServer(t *testing.T) {
	t.Helper()

	targetURI := os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL")

	targetURL, err := url.Parse(targetURI)
	if err != nil {
		t.Fatal(err)
	}

	dir := snapshotDir(t)
	randIntsFilename := path.Join(dir, "randInts.json")

	if updateSnapshots() {
		randInts[t.Name()] = map[int64]int64{}
		_ = os.RemoveAll(dir)
	} else {
		buf, err := os.ReadFile(randIntsFilename)
		if err != nil {
			t.Fatal(err)
		}

		var data map[int64]int64
		if err := json.Unmarshal(buf, &data); err != nil {
			t.Fatal(err)
		}
		data[0] = 0

		randInts[t.Name()] = data
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ModifyResponse = func(resp *http.Response) error {
		filename := snapshotName(t, resp.Request)

		data := map[string][]string{}
		for k, v := range resp.Header {
			data[k] = v
		}

		data["__status__"] = []string{strconv.Itoa(resp.StatusCode)}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, resp.Body); err != nil {
			return err
		}
		resp.Body.Close()
		resp.Body = io.NopCloser(&buf)

		data["__body__"] = []string{base64.StdEncoding.EncodeToString(buf.Bytes())}

		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		t.Logf("writing snapshot at %q", filename)
		return os.WriteFile(filename, out, 0600)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if updateSnapshots() {
			r.Host = targetURL.Host
			proxy.ServeHTTP(w, r)
			return
		}

		filename := snapshotName(t, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Logf("reading snapshot from %q", filename)
		buf, err := os.ReadFile(filename)
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, `{"errorCode":"NOT_FOUND","error":"snapshot not found"}`, http.StatusNotFound)
				return
			}
			http.Error(w, fmt.Sprintf(`{"errorCode":"UNKNOWN_ERROR","error":%q}`, err.Error()), http.StatusInternalServerError)
			return
		}
		var data map[string][]string
		if err := json.Unmarshal(buf, &data); err != nil {
			http.Error(w, fmt.Sprintf(`{"errorCode":"UNKNOWN_ERROR","error":%q}`, err.Error()), http.StatusInternalServerError)
			return
		}

		for k, v := range data {
			if k == "__body__" || k == "__status__" || k == "Content-Length" {
				continue
			}
			w.Header()[k] = v
		}

		status, err := strconv.Atoi(data["__status__"][0])
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"errorCode":"UNKNOWN_ERROR","error":%q}`, err.Error()), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)

		if data["__body__"] != nil {
			body, err := base64.StdEncoding.DecodeString(data["__body__"][0])
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"errorCode":"UNKNOWN_ERROR","error":%q}`, err.Error()), http.StatusInternalServerError)
				return
			}
			_, err = w.Write(body)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"errorCode":"UNKNOWN_ERROR","error":%q}`, err.Error()), http.StatusInternalServerError)
				return
			}
		}
	}))

	t.Cleanup(func() {
		if updateSnapshots() {
			out, err := json.MarshalIndent(randInts[t.Name()], "", "  ")

			if err != nil {
				t.Fatal(err)
			}

			err = os.WriteFile(randIntsFilename, out, 0600)

			if err != nil {
				t.Fatal(err)
			}
		}
		server.Close()
	})

	t.Setenv("MONGODB_ATLAS_OPS_MANAGER_URL", server.URL)
}

func setup(t *testing.T) {
	t.Helper()

	currentTestName = t.Name()

	snapshotServer(t)
}
