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
var memoryMap = map[string]map[string]any{}
var lastData = map[string]map[string][]string{}

func snapshotBaseDir(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if strings.HasSuffix(dir, "test/e2e") {
		dir = path.Join(dir, ".snapshots")
	} else {
		dir = path.Join(dir, "test/e2e/.snapshots")
	}

	return dir
}

func snapshotDir(t *testing.T) string {
	t.Helper()

	testName, _ := splitTestName(t)

	dir := snapshotBaseDir(t)

	dir = path.Join(dir, testName)

	return dir
}

func enforceDir(t *testing.T, filename string) {
	t.Helper()

	dir := path.Dir(filename)

	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			t.Fatal(err)
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Fatal(err)
		}
	}
}

func snapshotBaseName(t *testing.T, r *http.Request) string {
	t.Helper()

	dir := snapshotDir(t)

	return fmt.Sprintf("%s/%s_%s", dir, r.Method, strings.ReplaceAll(strings.ReplaceAll(r.URL.Path, "/", "_"), ":", "_"))
}

func snapshotName(t *testing.T, r *http.Request) string {
	t.Helper()

	baseName := snapshotBaseName(t, r)

	testName, _ := splitTestName(t)

	if fileIDs[testName] == nil {
		fileIDs[testName] = map[string]int{}
	}
	fileIDs[testName][baseName]++

	id := fileIDs[testName][baseName]

	fileName := fmt.Sprintf("%s_%d.json", baseName, id)

	return fileName
}

func snapshotNameStepBack(t *testing.T, r *http.Request) {
	t.Helper()

	baseName := snapshotBaseName(t, r)

	testName, _ := splitTestName(t)

	if fileIDs[testName] == nil {
		fileIDs[testName] = map[string]int{}
	}
	fileIDs[testName][baseName] -= 2
	if fileIDs[testName][baseName] < 0 {
		t.Fatal("no previous snapshot")
	}
}

func updateSnapshots() bool {
	return isTrue(os.Getenv("UPDATE_SNAPSHOTS"))
}

func skipSnapshots() bool {
	return os.Getenv("UPDATE_SNAPSHOTS") == "skip"
}

func loadMemory(t *testing.T) {
	t.Helper()

	dir := snapshotDir(t)
	filename := path.Join(dir, "memory.json")

	buf, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			memoryMap[t.Name()] = map[string]any{}
			return
		}
		t.Fatal(err)
	}

	var data map[string]any
	if err := json.Unmarshal(buf, &data); err != nil {
		t.Fatal(err)
	}
	memoryMap[t.Name()] = data
}

func storeMemory(t *testing.T) {
	t.Helper()

	dir := snapshotDir(t)
	filename := path.Join(dir, "memory.json")
	testName, _ := splitTestName(t)

	buf, err := json.Marshal(memoryMap[testName])
	if err != nil {
		t.Fatal(err)
	}

	enforceDir(t, filename)

	if err := os.WriteFile(filename, buf, 0600); err != nil {
		t.Fatal(err)
	}
}

func readSnapshot(t *testing.T, r *http.Request) []byte {
	t.Helper()

	filename := snapshotName(t, r)

	t.Logf("reading snapshot from %q", filename)
	buf, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			snapshotNameStepBack(t, r)
			return readSnapshot(t, r)
		}
		t.Fatal(err)
	}

	return buf
}

func snapshotServer(t *testing.T) {
	t.Helper()

	if skipSnapshots() {
		return
	}

	targetURI := os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL")

	targetURL, err := url.Parse(targetURI)
	if err != nil {
		t.Fatal(err)
	}

	if updateSnapshots() {
		dir := snapshotDir(t)
		_ = os.RemoveAll(dir)
	} else {
		loadMemory(t)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == http.StatusUnauthorized && resp.Header.Get("Www-Authenticate") != "" {
			return nil // skip 401
		}

		testName, _ := splitTestName(t)

		data := map[string][]string{}
		for k, v := range resp.Header {
			data[k] = v
		}

		data["__path__"] = []string{resp.Request.URL.Path}

		data["__method__"] = []string{resp.Request.Method}

		data["__status__"] = []string{strconv.Itoa(resp.StatusCode)}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, resp.Body); err != nil {
			return err
		}
		resp.Body.Close()
		resp.Body = io.NopCloser(&buf)

		data["__body__"] = []string{base64.StdEncoding.EncodeToString(buf.Bytes())}

		if lastData[testName] != nil && data["__method__"][0] == lastData[testName]["__method__"][0] && data["__path__"][0] == lastData[testName]["__path__"][0] && data["__status__"][0] == lastData[testName]["__status__"][0] && data["__body__"][0] == lastData[testName]["__body__"][0] {
			return nil // skip same content
		}

		lastData[testName] = data

		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		filename := snapshotName(t, resp.Request)
		t.Logf("writing snapshot at %q", filename)
		enforceDir(t, filename)
		return os.WriteFile(filename, out, 0600)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if updateSnapshots() {
			r.Host = targetURL.Host
			proxy.ServeHTTP(w, r)
			return
		}

		buf := readSnapshot(t, r)
		var data map[string][]string
		if err := json.Unmarshal(buf, &data); err != nil {
			t.Fatal(err)
		}

		for k, v := range data {
			if k == "__body__" || k == "__status__" || k == "__method__" || k == "__path__" {
				continue
			}
			w.Header()[k] = v
		}

		status, err := strconv.Atoi(data["__status__"][0])
		if err != nil {
			t.Fatal(err)
		}
		w.WriteHeader(status)

		if data["__body__"] != nil {
			body, err := base64.StdEncoding.DecodeString(data["__body__"][0])
			if err != nil {
				t.Fatal(err)
			}
			_, err = w.Write(body)
			if err != nil {
				t.Fatal(err)
			}
		}
	}))

	t.Cleanup(func() {
		if updateSnapshots() {
			storeMemory(t)
		}
		server.Close()
	})

	t.Setenv("MONGODB_ATLAS_OPS_MANAGER_URL", server.URL)
}

func splitTestName(t *testing.T) (string, string) {
	t.Helper()

	parts := strings.SplitN(t.Name(), "/", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func memory[T any](t *testing.T, key string, value T) T {
	t.Helper()

	if skipSnapshots() {
		return value
	}

	testName, subTestName := splitTestName(t)
	if subTestName != "" {
		key = subTestName + "/" + key
	}

	if memoryMap[testName] == nil {
		memoryMap[testName] = map[string]any{}
	}
	if updateSnapshots() {
		_, ok := memoryMap[testName][key]
		if ok {
			t.Fatalf("memory key %q already exists", key)
		}
		memoryMap[testName][key] = value
		return value
	}
	data, ok := memoryMap[testName][key]
	if !ok {
		t.Fatalf("memory key %q not found", key)
	}
	r, ok := data.(T)
	if !ok {
		t.Fatalf("memory key %q has unexpected type %T", key, data)
	}
	return r
}

func memoryFunc[T any](t *testing.T, key string, value T, marshal func(value T) ([]byte, error), unmarshal func([]byte) (T, error)) T {
	t.Helper()

	if skipSnapshots() {
		return value
	}

	testName, subTestName := splitTestName(t)
	if subTestName != "" {
		key = subTestName + "/" + key
	}
	if memoryMap[testName] == nil {
		memoryMap[testName] = map[string]any{}
	}
	if updateSnapshots() {
		data, err := marshal(value)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		memoryMap[testName][key] = base64.StdEncoding.EncodeToString(data)
		return value
	}
	data, ok := memoryMap[testName][key]
	if !ok {
		t.Fatalf("memory key %q not found", key)
	}
	buf, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	r, err := unmarshal(buf)
	if err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	return r
}

func memoryRand(t *testing.T, key string, n int64) *big.Int { //nolint:unparam // in case there are more than one random values in the same test
	t.Helper()

	return memoryFunc(t, key, must(RandInt(n)), func(value *big.Int) ([]byte, error) {
		return value.Bytes(), nil
	}, func(data []byte) (*big.Int, error) {
		return new(big.Int).SetBytes(data), nil
	})
}

func setup(t *testing.T) {
	t.Helper()

	snapshotServer(t)
}
