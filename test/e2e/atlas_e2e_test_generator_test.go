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
	"os/exec"
	"path"
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

type snapshotData struct {
	Body   []byte
	Status int
	Method string
	Path   string

	Headers map[string][]string
}

func (d snapshotData) Compare(v snapshotData) int {
	methodCmp := strings.Compare(d.Method, v.Method)
	if methodCmp != 0 {
		return methodCmp
	}

	pathCmp := strings.Compare(d.Path, v.Path)
	if pathCmp != 0 {
		return pathCmp
	}

	statusCmp := d.Status - v.Status
	if statusCmp != 0 {
		return statusCmp
	}

	return bytes.Compare(d.Body, v.Body)
}

func (d snapshotData) Write(w http.ResponseWriter) (int, error) {
	for k, v := range d.Headers {
		w.Header()[k] = v
	}

	w.WriteHeader(d.Status)

	return w.Write(d.Body)
}

// atlasE2ETestGenerator is about providing capabilities to provide projects and clusters for our e2e tests.
type atlasE2ETestGenerator struct {
	projectID           string
	projectName         string
	clusterName         string
	clusterRegion       string
	tier                string
	mDBVer              string
	enableBackup        bool
	firstProcess        *atlasv2.ApiHostViewAtlas
	t                   *testing.T
	fileIDs             map[string]int
	memoryMap           map[string]any
	lastData            *snapshotData
	currentSnapshotMode snapshotMode
}

// Log formats its arguments using default formatting, analogous to Println,
// and records the text in the error log. For tests, the text will be printed only if
// the test fails or the -test.v flag is set. For benchmarks, the text is always
// printed to avoid having performance depend on the value of the -test.v flag.
func (g *atlasE2ETestGenerator) Log(args ...any) {
	g.t.Log(args...)
}

// Logf formats its arguments according to the format, analogous to Printf, and
// records the text in the error log. A final newline is added if not provided. For
// tests, the text will be printed only if the test fails or the -test.v flag is
// set. For benchmarks, the text is always printed to avoid having performance
// depend on the value of the -test.v flag.
func (g *atlasE2ETestGenerator) Logf(format string, args ...any) {
	g.t.Logf(format, args...)
}

// newAtlasE2ETestGenerator creates a new instance of atlasE2ETestGenerator struct.
func newAtlasE2ETestGenerator(t *testing.T, opts ...func(g *atlasE2ETestGenerator)) *atlasE2ETestGenerator {
	t.Helper()
	g := &atlasE2ETestGenerator{t: t, currentSnapshotMode: snapshotModeSkip, fileIDs: map[string]int{}, memoryMap: map[string]any{}}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func withBackup() func(g *atlasE2ETestGenerator) {
	return func(g *atlasE2ETestGenerator) {
		g.enableBackup = true
	}
}

func withSnapshot() func(g *atlasE2ETestGenerator) {
	return func(g *atlasE2ETestGenerator) {
		g.snapshotServer()
	}
}

// generateProject generates a new project and also registers its deletion on test cleanup.
func (g *atlasE2ETestGenerator) generateProject(prefix string) {
	g.t.Helper()

	if g.projectID != "" {
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

	g.projectID, err = createProject(g.projectName)
	if err != nil {
		g.t.Fatalf("unexpected error creating project: %v", err)
	}
	g.Logf("projectID=%s", g.projectID)
	g.Logf("projectName=%s", g.projectName)
	if g.projectID == "" {
		g.t.Fatal("projectID not created")
	}

	if skipCleanup() {
		return
	}

	g.t.Cleanup(func() {
		deleteProjectWithRetry(g.t, g.projectID)
	})
}

func (g *atlasE2ETestGenerator) generateFlexCluster() {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	g.clusterName = g.memory("generateFlexClusterName", must(RandClusterName())).(string)

	err := deployFlexClusterForProject(g.projectID, g.clusterName)
	if err != nil {
		g.t.Fatalf("unexpected error deploying flex cluster: %v", err)
	}
	g.t.Logf("flexClusterName=%s", g.clusterName)

	if skipCleanup() {
		return
	}

	g.t.Cleanup(func() {
		_ = deleteClusterForProject(g.projectID, g.clusterName)
	})
}

func (g *atlasE2ETestGenerator) generateClusterWithPrefix(prefix string) {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	var err error
	if g.tier == "" {
		g.tier = e2eTier()
	}

	if g.mDBVer == "" {
		mdbVersion, e := MongoDBMajorVersion()
		require.NoError(g.t, e)

		g.mDBVer = mdbVersion
	}

	g.clusterName = g.memory(prefix+"GenerateClusterName", must(RandClusterNameWithPrefix(prefix))).(string)

	g.clusterRegion, err = deployClusterForProject(g.projectID, g.clusterName, g.tier, g.mDBVer, g.enableBackup)
	if err != nil {
		g.Logf("projectID=%q, clusterName=%q", g.projectID, g.clusterName)
		g.t.Errorf("unexpected error deploying cluster: %v", err)
	}
	g.t.Logf("clusterName=%s", g.clusterName)

	if skipCleanup() {
		return
	}

	g.t.Cleanup(func() {
		g.Logf("Cluster cleanup %q\n", g.projectID)
		if e := deleteClusterForProject(g.projectID, g.clusterName); e != nil {
			g.t.Errorf("unexpected error deleting cluster: %v", e)
		}
	})
}

// generateCluster generates a new cluster and also registers its deletion on test cleanup.
func (g *atlasE2ETestGenerator) generateCluster() {
	g.generateClusterWithPrefix("cluster")
}

// generateProjectAndCluster calls both generateProject and generateCluster.
func (g *atlasE2ETestGenerator) generateProjectAndCluster(prefix string) {
	g.t.Helper()

	g.generateProject(prefix)
	g.generateClusterWithPrefix(prefix)
}

// newAvailableRegion returns the first region for the provider/tier.
func (g *atlasE2ETestGenerator) newAvailableRegion(tier, provider string) (string, error) {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	return newAvailableRegion(g.projectID, tier, provider)
}

// getHostnameAndPort returns hostname:port from the first process.
func (g *atlasE2ETestGenerator) getHostnameAndPort() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return *p.Id, nil
}

// getHostname returns the hostname of first process.
func (g *atlasE2ETestGenerator) getHostname() (string, error) {
	g.t.Helper()

	p, err := g.getFirstProcess()
	if err != nil {
		return "", err
	}

	return *p.Hostname, nil
}

// getFirstProcess returns the first process of the project.
func (g *atlasE2ETestGenerator) getFirstProcess() (*atlasv2.ApiHostViewAtlas, error) {
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
func (g *atlasE2ETestGenerator) getProcesses() ([]atlasv2.ApiHostViewAtlas, error) {
	g.t.Helper()

	if g.projectID == "" {
		g.t.Fatal("unexpected error: project must be generated")
	}

	resp, err := g.runCommand(
		processesEntity,
		"list",
		"--projectId",
		g.projectID,
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

// runCommand runs a command on atlascli.
func (g *atlasE2ETestGenerator) runCommand(args ...string) ([]byte, error) {
	g.t.Helper()

	cliPath, err := AtlasCLIBin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath, args...)

	cmd.Env = os.Environ()
	return RunAndGetStdOut(cmd)
}

func skipCleanup() bool {
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

func (g *atlasE2ETestGenerator) snapshotBaseDir() string {
	g.t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		g.t.Fatal(err)
	}

	if strings.HasSuffix(dir, "test/e2e") {
		dir = path.Join(dir, ".snapshots")
	} else {
		dir = path.Join(dir, "test/e2e/.snapshots")
	}

	return dir
}

func (g *atlasE2ETestGenerator) snapshotDir() string {
	g.t.Helper()

	dir := g.snapshotBaseDir()

	dir = path.Join(dir, g.t.Name())

	return dir
}

func (g *atlasE2ETestGenerator) enforceDir(filename string) {
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

func (g *atlasE2ETestGenerator) snapshotBaseName(r *http.Request) string {
	g.t.Helper()

	dir := g.snapshotDir()

	return fmt.Sprintf("%s/%s_%s", dir, r.Method, strings.ReplaceAll(strings.ReplaceAll(r.URL.Path, "/", "_"), ":", "_"))
}

func (g *atlasE2ETestGenerator) snapshotName(r *http.Request) string {
	g.t.Helper()

	baseName := g.snapshotBaseName(r)

	g.fileIDs[baseName]++

	id := g.fileIDs[baseName]

	fileName := fmt.Sprintf("%s_%d.json", baseName, id)

	return fileName
}

func (g *atlasE2ETestGenerator) snapshotNameStepBack(r *http.Request) {
	g.t.Helper()

	baseName := g.snapshotBaseName(r)

	g.fileIDs[baseName] -= 2
	if g.fileIDs[baseName] < 0 {
		g.t.Fatal("no previous snapshot")
	}
}

func updateSnapshots() bool {
	return isTrue(os.Getenv(updateSnapshotsEnvVarKey))
}

func skipSnapshots() bool {
	return os.Getenv(updateSnapshotsEnvVarKey) == "skip"
}

func (g *atlasE2ETestGenerator) loadMemory() {
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

func (g *atlasE2ETestGenerator) storeMemory() {
	g.t.Helper()

	dir := g.snapshotDir()
	filename := path.Join(dir, "memory.json")

	buf, err := json.Marshal(g.memoryMap)
	if err != nil {
		g.t.Fatal(err)
	}

	g.enforceDir(filename)

	if err := os.WriteFile(filename, buf, 0600); err != nil {
		g.t.Fatal(err)
	}
}

func (g *atlasE2ETestGenerator) readSnapshot(r *http.Request) snapshotData {
	g.t.Helper()

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

	var data snapshotData
	if err := json.Unmarshal(buf, &data); err != nil {
		g.t.Fatal(err)
	}

	return data
}

func (g *atlasE2ETestGenerator) snapshotServer() {
	g.t.Helper()

	if skipSnapshots() {
		return
	}

	targetURI := os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL")

	targetURL, err := url.Parse(targetURI)
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
		if resp.StatusCode == http.StatusUnauthorized && resp.Header.Get("Www-Authenticate") != "" {
			return nil // skip 401
		}

		data := snapshotData{
			Path:    resp.Request.URL.Path,
			Method:  resp.Request.Method,
			Status:  resp.StatusCode,
			Headers: resp.Header,
		}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, resp.Body); err != nil {
			return err
		}
		resp.Body.Close()
		resp.Body = io.NopCloser(&buf)

		data.Body = buf.Bytes()

		if g.lastData != nil && data.Compare(*g.lastData) == 0 {
			return nil // skip same content
		}

		g.lastData = &data

		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		filename := g.snapshotName(resp.Request)
		g.t.Logf("writing snapshot at %q", filename)
		g.enforceDir(filename)
		return os.WriteFile(filename, out, 0600)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if g.currentSnapshotMode == snapshotModeUpdate {
			r.Host = targetURL.Host
			proxy.ServeHTTP(w, r)
			return
		}

		data := g.readSnapshot(r)
		if _, err := data.Write(w); err != nil {
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

func (g *atlasE2ETestGenerator) memory(key string, value any) any {
	g.t.Helper()

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

func (g *atlasE2ETestGenerator) memoryFunc(key string, value any, marshal func(value any) ([]byte, error), unmarshal func([]byte) (any, error)) any {
	g.t.Helper()

	switch g.currentSnapshotMode {
	case snapshotModeSkip:
		return value
	case snapshotModeUpdate:
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

func (g *atlasE2ETestGenerator) memoryRand(key string, n int64) *big.Int {
	g.t.Helper()

	r, ok := g.memoryFunc(key, must(RandInt(n)), func(value any) ([]byte, error) {
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
