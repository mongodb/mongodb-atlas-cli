package e2e_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

	filename := fmt.Sprintf("%s/%s_%s", dir, r.Method, strings.ReplaceAll(r.URL.Path, "/", "_"))

	if fileIDs[t.Name()] == nil {
		fileIDs[t.Name()] = map[string]int{}
	}
	fileIDs[t.Name()][filename]++

	id := fileIDs[t.Name()][filename]

	return fmt.Sprintf("%s_%d.json", filename, id)
}

func snapshotServer(t *testing.T) {
	t.Helper()

	targetURI := os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL")

	targetURL, err := url.Parse(targetURI)
	if err != nil {
		t.Fatal(err)
	}

	updateSnapshots := os.Getenv("UPDATE_SNAPSHOTS") == "true"

	if updateSnapshots {
		dir := snapshotDir(t)
		_ = os.RemoveAll(dir)
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
		if updateSnapshots {
			r.Host = targetURL.Host
			proxy.ServeHTTP(w, r)
			return
		}

		filename := snapshotName(t, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var data map[string][]string
		if err := json.Unmarshal(buf, &data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)

		if data["__body__"] != nil {
			body, err := base64.StdEncoding.DecodeString(data["__body__"][0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = w.Write(body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}))

	t.Cleanup(func() {
		server.Close()
	})

	t.Setenv("MONGODB_ATLAS_OPS_MANAGER_URL", server.URL)
}
