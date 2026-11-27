// Copyright 2026 MongoDB Inc
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
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"slices"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/cmd/genevergreen/generate"
	"gopkg.in/yaml.v3"
)

const fetchTimeout = 5 * time.Minute

var (
	gaVersionRe = regexp.MustCompile(`^(\d+\.\d+)\.\d+$`)
	jiraRe      = regexp.MustCompile(`^[A-Z]+-\d+$`)

	// feedArchs are the architectures Atlas CLI publishes; distros lists the per-distro set.
	feedArchs = []string{"x86_64", "aarch64"}

	// targetDistro maps feed targets to Atlas CLI distro IDs where they differ.
	targetDistro = map[string]string{
		"rhel8":  "rhel80",
		"rhel82": "rhel80",
		"rhel9":  "rhel90",
		"rhel93": "rhel90",
	}

	// skippedTargets are feed targets Atlas CLI never publishes to.
	skippedTargets = []string{
		"macos",
		"windows",
		"rhel70",
		"amazon2",
		"debian10",
		"debian11",
		"suse12",
		"ubuntu1804",
		"ubuntu2004",
	}
)

type feed struct {
	Versions []release `json:"versions"`
}

type release struct {
	Version           string     `json:"version"`
	ProductionRelease bool       `json:"production_release"`
	Downloads         []download `json:"downloads"`
}

type download struct {
	Arch   string `json:"arch"`
	Target string `json:"target"`
}

type exception struct {
	Version string `yaml:"version"`
	Distro  string `yaml:"distro"`
	Action  string `yaml:"action"`
	Jira    string `yaml:"jira"`
}

const (
	actionAdd    = "add"
	actionRemove = "remove"
)

// pair is one version and distro, e.g. 8.0/rhel10.
type pair struct {
	Version string
	Distro  string
}

func (p pair) String() string {
	return p.Version + "/" + p.Distro
}

type result struct {
	Matrix        generate.ServerMatrix
	AddedVersions []string
	Added         []pair
	Removed       []pair
	// Unknown lists feed targets Atlas CLI doesn't publish to and doesn't skip. Never added.
	Unknown []pair
}

func fetchFeed(ctx context.Context, url string) (*feed, error) {
	ctx, cancel := context.WithTimeout(ctx, fetchTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req) //nolint:gosec // G704: URL is the -feed flag set by the operator
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	var f feed
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", url, err)
	}
	return &f, nil
}

// feedTargets returns the union of published targets across each minor version's GA patches.
func feedTargets(f *feed) (map[string]map[string]bool, error) {
	targets := map[string]map[string]bool{}
	for _, r := range f.Versions {
		if !r.ProductionRelease {
			continue
		}
		m := gaVersionRe.FindStringSubmatch(r.Version)
		if m == nil {
			return nil, fmt.Errorf("GA release has invalid version %q", r.Version)
		}
		minor := m[1]
		if targets[minor] == nil {
			targets[minor] = map[string]bool{}
		}
		for _, d := range r.Downloads {
			if !slices.Contains(feedArchs, d.Arch) {
				continue
			}
			t := d.Target
			if mapped, ok := targetDistro[t]; ok {
				t = mapped
			}
			targets[minor][t] = true
		}
	}
	if len(targets) == 0 {
		return nil, errors.New("feed has no GA releases")
	}
	return targets, nil
}

// update merges feed distros into current and applies exceptions. It never removes a pair the feed dropped.
func update(current generate.ServerMatrix, f *feed, exceptions []exception) (*result, error) {
	targets, err := feedTargets(f)
	if err != nil {
		return nil, err
	}
	merged, unknown, err := mergeTargets(current, targets)
	if err != nil {
		return nil, err
	}
	if err := applyExceptions(merged, exceptions); err != nil {
		return nil, err
	}
	for sv := range merged {
		slices.Sort(merged[sv])
	}
	if err := merged.Validate(); err != nil {
		return nil, err
	}
	res := diff(current, merged)
	res.Unknown = unknown
	return res, nil
}

// mergeTargets adds feed distros Atlas CLI publishes to, for versions in current or newer than its newest.
func mergeTargets(current generate.ServerMatrix, targets map[string]map[string]bool) (generate.ServerMatrix, []pair, error) {
	versions := current.Versions()
	newest := versions[len(versions)-1]
	publish := generate.PublishDistros()

	merged := generate.ServerMatrix{}
	for sv, ds := range current {
		merged[sv] = slices.Clone(ds)
	}
	var unknown []pair
	for minor, ts := range targets {
		_, existing := current[minor]
		if !existing && generate.CompareMinorVersions(minor, newest) <= 0 {
			continue
		}
		var kept []string
		for t := range ts {
			switch {
			case slices.Contains(skippedTargets, t):
			case slices.Contains(publish, t):
				kept = append(kept, t)
			default:
				unknown = append(unknown, pair{minor, t})
			}
		}
		if !existing && len(kept) == 0 {
			return nil, nil, fmt.Errorf("new version %s has no distro Atlas CLI publishes to", minor)
		}
		for _, d := range kept {
			if !merged.Supports(minor, d) {
				merged[minor] = append(merged[minor], d)
			}
		}
	}
	slices.SortFunc(unknown, comparePairs)
	return merged, unknown, nil
}

func diff(current, merged generate.ServerMatrix) *result {
	res := &result{Matrix: merged}
	for _, sv := range merged.Versions() {
		if _, ok := current[sv]; !ok {
			res.AddedVersions = append(res.AddedVersions, sv)
		}
		for _, d := range merged[sv] {
			if !current.Supports(sv, d) {
				res.Added = append(res.Added, pair{sv, d})
			}
		}
	}
	for _, sv := range current.Versions() {
		for _, d := range current[sv] {
			if !merged.Supports(sv, d) {
				res.Removed = append(res.Removed, pair{sv, d})
			}
		}
	}
	return res
}

func applyExceptions(m generate.ServerMatrix, exceptions []exception) error {
	for _, e := range exceptions {
		ds, ok := m[e.Version]
		if !ok {
			return fmt.Errorf("exception %s: version not in the matrix", pair{e.Version, e.Distro})
		}
		switch e.Action {
		case actionAdd:
			if !slices.Contains(ds, e.Distro) {
				m[e.Version] = append(ds, e.Distro)
			}
		case actionRemove:
			m[e.Version] = slices.DeleteFunc(ds, func(d string) bool { return d == e.Distro })
		}
	}
	return nil
}

func parseExceptions(b []byte) ([]exception, error) {
	dec := yaml.NewDecoder(bytes.NewReader(b))
	dec.KnownFields(true)
	var exceptions []exception
	if err := dec.Decode(&exceptions); err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return nil, errors.New("exceptions must be a single YAML document")
	}
	publish := generate.PublishDistros()
	seen := map[pair]bool{}
	for _, e := range exceptions {
		p := pair{e.Version, e.Distro}
		if !generate.IsMinorVersion(e.Version) {
			return nil, fmt.Errorf("exception %s: invalid version, want major.minor", p)
		}
		if !slices.Contains(publish, e.Distro) {
			return nil, fmt.Errorf("exception %s: unknown distro", p)
		}
		if e.Action != actionAdd && e.Action != actionRemove {
			return nil, fmt.Errorf("exception %s: invalid action %q", p, e.Action)
		}
		if !jiraRe.MatchString(e.Jira) {
			return nil, fmt.Errorf("exception %s: invalid jira %q", p, e.Jira)
		}
		if seen[p] {
			return nil, fmt.Errorf("exception %s: duplicate", p)
		}
		seen[p] = true
	}
	return exceptions, nil
}

func comparePairs(a, b pair) int {
	if c := generate.CompareMinorVersions(a.Version, b.Version); c != 0 {
		return c
	}
	return cmp.Compare(a.Distro, b.Distro)
}

// marshalMatrix writes versions oldest first, so 10.0 sorts after 9.0.
func marshalMatrix(m generate.ServerMatrix) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{\n")
	for i, sv := range m.Versions() {
		key, err := json.Marshal(sv)
		if err != nil {
			return nil, err
		}
		ds, err := json.MarshalIndent(m[sv], "  ", "  ")
		if err != nil {
			return nil, err
		}
		b.WriteString("  ")
		b.Write(key)
		b.WriteString(": ")
		b.Write(ds)
		if i < len(m)-1 {
			b.WriteByte(',')
		}
		b.WriteByte('\n')
	}
	b.WriteString("}\n")
	return b.Bytes(), nil
}
