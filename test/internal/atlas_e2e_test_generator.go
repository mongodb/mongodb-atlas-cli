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

package internal

import (
	"bufio"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/sha1" //nolint:gosec // no need to be secure just replacing long filenames for windows
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
	"os/exec"
	"path"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

const updateSnapshotsEnvVarKey = "UPDATE_SNAPSHOTS"

type snapshotMode int

const (
	snapshotModeReplay snapshotMode = iota
	snapshotModeUpdate
	snapshotModeSkip
)

func decompress(r *http.Response) error {
	var err error
	shouldRemoveHeaders := false

	for _, encoding := range r.Header["Content-Encoding"] {
		reader := r.Body
		decompression := true
		switch encoding {
		case "gzip", "x-gzip":
			reader, err = gzip.NewReader(r.Body)
			if err != nil {
				return err
			}
		case "deflate":
			reader = flate.NewReader(r.Body)
		default:
			decompression = false
		}

		if decompression {
			shouldRemoveHeaders = true
			buf := new(bytes.Buffer)

			for {
				const bufferSize = 1024
				if _, err := io.CopyN(buf, reader, bufferSize); err != nil {
					if err == io.EOF {
						break
					}
					return err
				}
			}

			if err := reader.Close(); err != nil {
				return err
			}

			r.Body = io.NopCloser(buf)
			r.ContentLength = int64(buf.Len())
			r.Header["Content-Length"] = []string{strconv.FormatInt(r.ContentLength, 10)}
		}
	}

	if shouldRemoveHeaders {
		removeContentEncoding(r)
	}

	return nil
}

func removeContentEncoding(r *http.Response) {
	delete(r.Header, "Content-Encoding")

	if len(r.Header["Vary"]) == 0 {
		return
	}

	r.Header["Vary"] = slices.DeleteFunc(r.Header["Vary"], func(s string) bool {
		return strings.EqualFold(s, "accept-encoding")
	})
	if len(r.Header["Vary"]) == 0 {
		delete(r.Header, "Vary")
		return
	}
}

func compareSnapshots(a *http.Response, b *http.Response) int {
	methodCmp := strings.Compare(a.Request.Method, b.Request.Method)
	if methodCmp != 0 {
		return methodCmp
	}

	pathCmp := strings.Compare(a.Request.URL.Path, b.Request.URL.Path)
	if pathCmp != 0 {
		return pathCmp
	}

	statusCmp := a.StatusCode - b.StatusCode
	if statusCmp != 0 {
		return statusCmp
	}

	aBody, err := io.ReadAll(a.Body)
	if err != nil {
		return 0
	}
	a.Body.Close()
	a.Body = io.NopCloser(bytes.NewReader(aBody))

	bBody, err := io.ReadAll(b.Body)
	if err != nil {
		return 0
	}
	b.Body.Close()
	b.Body = io.NopCloser(bytes.NewReader(bBody))

	return bytes.Compare(aBody, bBody)
}

// AtlasE2ETestGenerator is about providing capabilities to provide projects and clusters for our e2e tests.
type AtlasE2ETestGenerator struct {
	ProjectID           string
	projectName         string
	ClusterName         string
	clusterRegion       string
	Tier                string
	MDBVer              string
	enableBackup        bool
	firstProcess        *atlasv2.ApiHostViewAtlas
	t                   *testing.T
	fileIDs             map[string]int
	memoryMap           map[string]any
	lastSnapshot        *http.Response
	currentSnapshotMode snapshotMode
	testName            string
	skipSnapshots       func(snapshot *http.Response, prevSnapshot *http.Response) bool
	snapshotNameFunc    func(r *http.Request) string
	snapshotTargetURI   string
}

// Log formats its arguments using default formatting, analogous to Println,
// and records the text in the error log. For tests, the text will be printed only if
// the test fails or the -test.v flag is set. For benchmarks, the text is always
// printed to avoid having performance depend on the value of the -test.v flag.
func (g *AtlasE2ETestGenerator) Log(args ...any) {
	g.t.Log(args...)
}

// Logf formats its arguments according to the format, analogous to Printf, and
// records the text in the error log. A final newline is added if not provided. For
// tests, the text will be printed only if the test fails or the -test.v flag is
// set. For benchmarks, the text is always printed to avoid having performance
// depend on the value of the -test.v flag.
func (g *AtlasE2ETestGenerator) Logf(format string, args ...any) {
	g.t.Logf(format, args...)
}

// newAtlasE2ETestGenerator creates a new instance of AtlasE2ETestGenerator struct.
func NewAtlasE2ETestGenerator(t *testing.T, opts ...func(g *AtlasE2ETestGenerator)) *AtlasE2ETestGenerator {
	t.Helper()
	g := &AtlasE2ETestGenerator{
		t:                   t,
		testName:            t.Name(),
		currentSnapshotMode: snapshotModeSkip,
		skipSnapshots:       compositeSnapshotSkipFunc(Skip401Snapshots, SkipSimilarSnapshots),
		fileIDs:             map[string]int{},
		memoryMap:           map[string]any{},
		snapshotNameFunc:    defaultSnapshotBaseName,
		snapshotTargetURI:   os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL"),
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func WithBackup() func(g *AtlasE2ETestGenerator) {
	return func(g *AtlasE2ETestGenerator) {
		g.enableBackup = true
	}
}

func WithSnapshot() func(g *AtlasE2ETestGenerator) {
	return func(g *AtlasE2ETestGenerator) {
		g.snapshotServer()
	}
}

func WithSnapshotSkip(f ...func(*http.Response, *http.Response) bool) func(g *AtlasE2ETestGenerator) {
	return func(g *AtlasE2ETestGenerator) {
		g.skipSnapshots = compositeSnapshotSkipFunc(f...)
	}
}

func compositeSnapshotSkipFunc(f ...func(*http.Response, *http.Response) bool) func(*http.Response, *http.Response) bool {
	return func(snapshot *http.Response, prevSnapshot *http.Response) bool {
		for _, fn := range f {
			if fn == nil {
				continue
			}
			if fn(snapshot, prevSnapshot) {
				return true
			}
		}
		return false
	}
}

func WithSnapshotNameFunc(f func(*http.Request) string) func(g *AtlasE2ETestGenerator) {
	return func(g *AtlasE2ETestGenerator) {
		g.snapshotNameFunc = f
	}
}

func (g *AtlasE2ETestGenerator) Run(name string, f func(t *testing.T)) {
	g.t.Helper()

	g.t.Run(name, func(t *testing.T) {
		t.Helper()

		g.testName = t.Name()
		g.lastSnapshot = nil

		t.Cleanup(func() {
			g.testName = g.t.Name()
		})

		f(t)
	})
}

// GenerateProject generates a new project and also registers its deletion on test cleanup.
func (g *AtlasE2ETestGenerator) GenerateProject(prefix string) {
	g.t.Helper()

	if g.ProjectID != "" {
		g.t.Fatal("unexpected error: project was already generated")
	}

	var err error
	if prefix == "" {
		g.projectName, err = RandProjectName()
	} else {
		g.projectName, err = RandProjectNameWithPrefix(prefix)
	}
	if err != nil {
		g.t.Fatalf("unexpected error: %v", err)
	}

	g.ProjectID, err = createProject(g.projectName)
	if err != nil {
		g.t.Fatalf("unexpected error creating project: %v", err)
	}
	g.Logf("projectID=%s", g.ProjectID)
	g.Logf("projectName=%s", g.projectName)
	if g.ProjectID == "" {
		g.t.Fatal("projectID not created")
	}

	if SkipCleanup() {
		return
	}

	g.t.Cleanup(func() {
		deleteProjectWithRetry(g.t, g.ProjectID)
	})
}

func (g *AtlasE2ETestGenerator) generateClusterWithPrefix(prefix string) {
	g.t.Helper()

	if g.ProjectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	var err error
	if g.Tier == "" {
		g.Tier = E2eTier()
	}

	if g.MDBVer == "" {
		mdbVersion, e := MongoDBMajorVersion()
		require.NoError(g.t, e)

		g.MDBVer = mdbVersion
	}

	g.ClusterName = g.Memory(prefix+"GenerateClusterName", Must(RandClusterNameWithPrefix(prefix))).(string)

	g.clusterRegion, err = deployClusterForProject(g.ProjectID, g.ClusterName, g.Tier, g.MDBVer, g.enableBackup)
	if err != nil {
		g.Logf("projectID=%q, clusterName=%q", g.ProjectID, g.ClusterName)
		g.t.Errorf("unexpected error deploying cluster: %v", err)
	}
	g.t.Logf("clusterName=%s", g.ClusterName)

	if SkipCleanup() {
		return
	}

	g.t.Cleanup(func() {
		g.Logf("Cluster cleanup %q\n", g.ProjectID)
		if e := DeleteClusterForProject(g.ProjectID, g.ClusterName); e != nil {
			g.t.Errorf("unexpected error deleting cluster: %v", e)
		}
	})
}

// GenerateCluster generates a new cluster and also registers its deletion on test cleanup.
func (g *AtlasE2ETestGenerator) GenerateCluster() {
	g.generateClusterWithPrefix("cluster")
}

// generateProjectAndCluster calls both generateProject and generateCluster.
func (g *AtlasE2ETestGenerator) GenerateProjectAndCluster(prefix string) {
	g.t.Helper()

	g.GenerateProject(prefix)
	g.generateClusterWithPrefix(prefix)
}

// NewAvailableRegion returns the first region for the provider/tier.
func (g *AtlasE2ETestGenerator) NewAvailableRegion(tier, provider string) (string, error) {
	g.t.Helper()

	if g.ProjectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	return NewAvailableRegion(g.ProjectID, tier, provider)
}

// GetHostnameAndPort returns hostname:port from the first process.
func (g *AtlasE2ETestGenerator) GetHostnameAndPort() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return *p.Id, nil
}

// GetHostname returns the hostname of first process.
func (g *AtlasE2ETestGenerator) GetHostname() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	return *p.Hostname, nil
}

// getFirstProcess returns the first process of the project.
func (g *AtlasE2ETestGenerator) getFirstProcess() (*atlasv2.ApiHostViewAtlas, error) {
	g.t.Helper()

	if g.firstProcess != nil {
		return g.firstProcess, nil
	}

	processes, err := g.getProcesses()
	if err != nil {
		return nil, err
	}
	g.firstProcess = &processes[0]

	return g.firstProcess, nil
}

// getHostname returns the list of processes.
func (g *AtlasE2ETestGenerator) getProcesses() ([]atlasv2.ApiHostViewAtlas, error) {
	g.t.Helper()

	if g.ProjectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	resp, err := g.RunCommand(
		processesEntity,
		"list",
		"--projectId",
		g.ProjectID,
		"-o=json",
	)
	if err != nil {
		return nil, err
	}

	var processes *atlasv2.PaginatedHostViewAtlas
	if err := json.Unmarshal(resp, &processes); err != nil {
		g.t.Fatalf("unexpected error: project must be generated %s - %s", err, resp)
	}

	if len(processes.GetResults()) == 0 {
		g.t.Fatalf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	return processes.GetResults(), nil
}

// RunCommand runs a command on atlascli.
func (g *AtlasE2ETestGenerator) RunCommand(args ...string) ([]byte, error) {
	g.t.Helper()

	cliPath, err := AtlasCLIBin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath, args...)

	cmd.Env = os.Environ()
	return RunAndGetStdOut(cmd)
}

func SkipCleanup() bool {
	return isTrue(os.Getenv("E2E_SKIP_CLEANUP"))
}

func isTrue(s string) bool {
	switch s {
	case "t", "T", "true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES", "1":
		return true
	default:
		return false
	}
}

func (g *AtlasE2ETestGenerator) snapshotBaseDir() string {
	g.t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		g.t.Fatal(err)
	}

	if strings.HasSuffix(dir, "test/e2e") {
		dir = path.Join(dir, "testdata/.snapshots")
	} else {
		dir = path.Join(dir, "test/e2e/testdata/.snapshots")
	}

	return dir
}

func (g *AtlasE2ETestGenerator) snapshotDir() string {
	g.t.Helper()

	dir := g.snapshotBaseDir()

	dir = path.Join(dir, g.testName)

	return dir
}

func (g *AtlasE2ETestGenerator) enforceDir(filename string) {
	g.t.Helper()

	dir := path.Dir(filename)

	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			g.t.Fatal(err)
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			g.t.Fatal(err)
		}
	}
}

func defaultSnapshotBaseName(r *http.Request) string {
	return fmt.Sprintf("%s_%s", r.Method, strings.ReplaceAll(strings.ReplaceAll(r.URL.Path, "/", "_"), ":", "_"))
}

func SnapshotHashedName(r *http.Request) string {
	defaultSnapshotBaseName := defaultSnapshotBaseName(r)
	hash := fmt.Sprintf("%x", sha1.Sum([]byte(defaultSnapshotBaseName))) //nolint:gosec // no need to be secure just replacing long filenames for windows
	return hash
}

func (g *AtlasE2ETestGenerator) maskString(s string) string {
	o := s
	o = strings.ReplaceAll(o, os.Getenv("MONGODB_ATLAS_ORG_ID"), "a0123456789abcdef012345a")
	o = strings.ReplaceAll(o, os.Getenv("MONGODB_ATLAS_PROJECT_ID"), "b0123456789abcdef012345b")
	o = strings.ReplaceAll(o, os.Getenv("IDENTITY_PROVIDER_ID"), "d0123456789abcdef012345d")
	o = strings.ReplaceAll(o, os.Getenv("E2E_CLOUD_ROLE_ID"), "c0123456789abcdef012345c")
	o = strings.ReplaceAll(o, os.Getenv("E2E_FLEX_INSTANCE_NAME"), "test-flex")
	o = strings.ReplaceAll(o, os.Getenv("E2E_TEST_BUCKET"), "test-bucket")
	o = strings.ReplaceAll(o, g.snapshotTargetURI, "http://localhost:8080/")
	return o
}

func (g *AtlasE2ETestGenerator) prepareRequest(r *http.Request) {
	g.t.Helper()
	var err error
	r.URL, err = url.Parse(g.maskString(r.URL.String()))
	if err != nil {
		g.t.Fatal(err)
	}
}

func (g *AtlasE2ETestGenerator) fileKey(r *http.Request) string {
	g.t.Helper()

	return fmt.Sprintf("%s/%s", g.testName, g.snapshotNameFunc(r))
}

func (g *AtlasE2ETestGenerator) snapshotName(r *http.Request) string {
	g.t.Helper()

	dir := g.snapshotDir()
	baseName := g.snapshotNameFunc(r)

	key := g.fileKey(r)

	g.fileIDs[key]++

	id := g.fileIDs[key]

	fileName := path.Join(dir, fmt.Sprintf("%s_%d.snaphost", baseName, id))

	return fileName
}

func (g *AtlasE2ETestGenerator) snapshotNameStepBack(r *http.Request) {
	g.t.Helper()

	key := g.fileKey(r)

	g.fileIDs[key] -= 2
	if g.fileIDs[key] < 0 {
		g.t.Fatal("no previous snapshot")
	}
}

func updateSnapshots() bool {
	return isTrue(os.Getenv(updateSnapshotsEnvVarKey))
}

func skipSnapshots() bool {
	return os.Getenv(updateSnapshotsEnvVarKey) == "skip"
}

func (g *AtlasE2ETestGenerator) loadMemory() {
	g.t.Helper()

	dir := g.snapshotDir()
	filename := path.Join(dir, "memory.json")

	buf, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			g.memoryMap = map[string]any{}
			return
		}
		g.t.Fatal(err)
	}

	if err := json.Unmarshal(buf, &g.memoryMap); err != nil {
		g.t.Fatal(err)
	}
}

func (g *AtlasE2ETestGenerator) storeMemory() {
	g.t.Helper()

	dir := g.snapshotDir()
	filename := path.Join(dir, "memory.json")

	buf, err := json.Marshal(g.memoryMap)
	if err != nil {
		g.t.Fatal(err)
	}

	g.enforceDir(filename)

	const permission = 0600
	if err := os.WriteFile(filename, buf, permission); err != nil {
		g.t.Fatal(err)
	}
}

func (g *AtlasE2ETestGenerator) prepareSnapshot(r *http.Response) *http.Response {
	g.t.Helper()

	buf, err := httputil.DumpResponse(r, true)
	if err != nil {
		g.t.Fatal(err)
	}

	req := r.Request
	g.prepareRequest(req)

	reader := bufio.NewReader(bytes.NewReader(buf))
	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		g.t.Fatal(err)
	}
	resp.Body = io.NopCloser(reader)

	if err := decompress(resp); err != nil {
		g.t.Fatal(err)
	}

	if resp.ContentLength > 0 && strings.Contains(resp.Header.Get("Content-Type"), "json") {
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			g.t.Fatal(err)
		}

		buf = []byte(g.maskString(string(buf)))
		resp.Body = io.NopCloser(bytes.NewReader(buf))
		resp.ContentLength = int64(len(buf))
		resp.Header["Content-Length"] = []string{strconv.FormatInt(resp.ContentLength, 10)}

		for k, mv := range resp.Header {
			for i, v := range mv {
				resp.Header[k][i] = g.maskString(v)
			}
		}
	}

	return resp
}

func (g *AtlasE2ETestGenerator) storeSnapshot(r *http.Response) {
	g.t.Helper()

	out, err := httputil.DumpResponse(r, true)
	if err != nil {
		g.t.Fatal(err)
	}

	filename := g.snapshotName(r.Request)
	g.t.Logf("writing snapshot at %q", filename)
	g.enforceDir(filename)
	const permission = 0600
	if err := os.WriteFile(filename, out, permission); err != nil {
		g.t.Fatal(err)
	}
}

func (g *AtlasE2ETestGenerator) readSnapshot(r *http.Request) *http.Response {
	g.t.Helper()

	g.prepareRequest(r)

	filename := g.snapshotName(r)

	g.t.Logf("reading snapshot from %q", filename)
	buf, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			g.snapshotNameStepBack(r)
			return g.readSnapshot(r)
		}
		g.t.Fatal(err)
	}

	reader := bufio.NewReader(bytes.NewReader(buf))
	resp, err := http.ReadResponse(reader, r)
	if err != nil {
		g.t.Fatal(err)
	}

	return resp
}

func Skip401Snapshots(snapshot *http.Response, _ *http.Response) bool {
	return snapshot.StatusCode == http.StatusUnauthorized && snapshot.Header.Get("Www-Authenticate") != ""
}

func SkipSimilarSnapshots(snapshot *http.Response, prevSnapshot *http.Response) bool {
	return prevSnapshot != nil && compareSnapshots(snapshot, prevSnapshot) == 0
}

func (g *AtlasE2ETestGenerator) snapshotServer() {
	g.t.Helper()

	if skipSnapshots() {
		return
	}

	targetURL, err := url.Parse(g.snapshotTargetURI)
	if err != nil {
		g.t.Fatal(err)
	}

	if updateSnapshots() {
		g.currentSnapshotMode = snapshotModeUpdate

		dir := g.snapshotDir()
		_ = os.RemoveAll(dir)
	} else {
		g.currentSnapshotMode = snapshotModeReplay

		g.loadMemory()
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ModifyResponse = func(resp *http.Response) error {
		snapshot := g.prepareSnapshot(resp)

		if g.skipSnapshots(snapshot, g.lastSnapshot) {
			return nil // skip
		}

		g.lastSnapshot = snapshot

		g.storeSnapshot(snapshot)

		return nil
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if g.currentSnapshotMode == snapshotModeUpdate {
			r.Host = targetURL.Host
			proxy.ServeHTTP(w, r)
			return
		}

		data := g.readSnapshot(r)

		for k, v := range data.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(data.StatusCode)

		if _, err := io.Copy(w, data.Body); err != nil {
			g.t.Fatal(err)
		}
		if err := data.Body.Close(); err != nil {
			g.t.Fatal(err)
		}
	}))

	g.t.Cleanup(func() {
		if g.currentSnapshotMode == snapshotModeUpdate {
			g.storeMemory()
		}
		server.Close()
	})

	g.t.Setenv("MONGODB_ATLAS_OPS_MANAGER_URL", server.URL)
}

func (g *AtlasE2ETestGenerator) Memory(key string, value any) any {
	g.t.Helper()

	if key == "" {
		g.t.Fatal("key cannot be empty")
	}

	key = fmt.Sprintf("%s/%s", g.testName, key)

	switch g.currentSnapshotMode {
	case snapshotModeSkip:
		return value
	case snapshotModeUpdate:
		_, ok := g.memoryMap[key]
		if ok {
			g.t.Fatalf("memory key %q already exists", key)
		}
		g.memoryMap[key] = value
		return value
	case snapshotModeReplay:
		data, ok := g.memoryMap[key]
		if !ok {
			g.t.Fatalf("memory key %q not found", key)
		}
		return data
	default:
		g.t.Fatalf("unexpected snapshot mode: %v", g.currentSnapshotMode)
		return nil
	}
}

func (g *AtlasE2ETestGenerator) MemoryFunc(key string, value any, marshal func(value any) ([]byte, error), unmarshal func([]byte) (any, error)) any {
	g.t.Helper()

	if key == "" {
		g.t.Fatal("key cannot be empty")
	}

	key = fmt.Sprintf("%s/%s", g.testName, key)

	switch g.currentSnapshotMode {
	case snapshotModeSkip:
		return value
	case snapshotModeUpdate:
		_, ok := g.memoryMap[key]
		if ok {
			g.t.Fatalf("memory key %q already exists", key)
		}
		data, err := marshal(value)
		if err != nil {
			g.t.Fatalf("marshal: %v", err)
		}
		g.memoryMap[key] = base64.StdEncoding.EncodeToString(data)
		return value
	case snapshotModeReplay:
		data, ok := g.memoryMap[key]
		if !ok {
			g.t.Fatalf("memory key %q not found", key)
		}
		buf, err := base64.StdEncoding.DecodeString(data.(string))
		if err != nil {
			g.t.Fatalf("decode: %v", err)
		}
		r, err := unmarshal(buf)
		if err != nil {
			g.t.Fatalf("unmarshal: %v", err)
		}
		return r
	default:
		g.t.Fatalf("unexpected snapshot mode: %v", g.currentSnapshotMode)
		return nil
	}
}

func (g *AtlasE2ETestGenerator) MemoryRand(key string, n int64) *big.Int {
	g.t.Helper()

	r, ok := g.MemoryFunc(key, Must(RandInt(n)), func(value any) ([]byte, error) {
		i := value.(*big.Int)
		return i.Bytes(), nil
	}, func(buf []byte) (any, error) {
		return big.NewInt(0).SetBytes(buf), nil
	}).(*big.Int)

	if !ok {
		g.t.Fatalf("unexpected error: %v", r)
	}
	return r
}
