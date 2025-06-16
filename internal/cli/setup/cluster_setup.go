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

package setup

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"golang.org/x/mod/semver"
)

var ErrNoRegions = errors.New("no regions found for the cloud provider")
var ErrNoVersions = errors.New("no mongodb versions found for the cloud provider")
var deprecatedClusterWideScalingMessage = "If you require more flexibility in shard scaling, consider using --autoScalingMode independentShardScaling."

const (
	tenant  = "TENANT"
	atlasM2 = "M2"
	atlasM5 = "M5"
)

func (opts *Opts) createCluster() error {
	if opts.AutoScalingMode == clusterWideScaling {
		_, _ = fmt.Fprintln(os.Stderr, deprecatedClusterWideScalingMessage)

		if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
			return err
		}
		return nil
	}

	if _, err := opts.store.CreateClusterLatest(opts.newClusterLatest()); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) askClusterOptions() error {
	var qs []*survey.Question

	if opts.shouldAskForValue(flag.ClusterName) {
		if opts.ClusterName == "" {
			opts.ClusterName = opts.defaultName
		}
		qs = append(qs, newClusterNameQuestion(opts.ClusterName))
	}

	if opts.shouldAskForValue(flag.Provider) {
		qs = append(qs, newClusterProviderQuestion())
	}

	if opts.shouldAskForValue(flag.ClusterName) || opts.shouldAskForValue(flag.Provider) || opts.shouldAskForValue(flag.Region) || opts.shouldAskForValue(flag.MDBVersion) {
		fmt.Print(`
[Set up your Atlas cluster]
`)
	}

	if err := telemetry.TrackAsk(qs, opts); err != nil {
		return err
	}

	// We need the provider to ask for the region
	if opts.shouldAskForValue(flag.Region) {
		if err := opts.askClusterRegion(); err != nil {
			return err
		}
	}

	if opts.shouldAskForValue(flag.MDBVersion) {
		return opts.askClusterMDBVersion()
	}

	return nil
}

func (opts *Opts) askClusterRegion() error {
	regions, err := opts.defaultRegions()
	if err != nil {
		return err
	}

	if len(regions) == 0 {
		return fmt.Errorf("%w: %v", ErrNoRegions, opts.Provider)
	}

	regionQ := newRegionQuestions(regions)
	return telemetry.TrackAskOne(regionQ, &opts.Region, survey.WithValidator(survey.Required))
}

func (opts *Opts) askClusterMDBVersion() error {
	if opts.providerName() == tenant {
		return nil
	}

	versions, defaultVersion, err := opts.mdbVersions(opts.Tier)
	if err != nil {
		return err
	}

	if len(versions) == 0 {
		return nil
	}

	return telemetry.TrackAskOne(newMDBVersionQuestion(versions, defaultVersion), &opts.MDBVersion, survey.WithValidator(survey.Required))
}

func newRegionQuestions(defaultRegions []string) survey.Prompt {
	return &survey.Select{
		Message: "Cloud Provider Region",
		Help:    usage.Region,
		Options: defaultRegions,
	}
}

func newMDBVersionQuestion(versions []string, defaultVersion string) survey.Prompt {
	return &survey.Select{
		Message: "Mongodb Version",
		Help:    usage.MDBVersion,
		Options: versions,
		Default: defaultVersion,
	}
}

func providerName(tier, provider string) string {
	if tier == DefaultAtlasTier || tier == atlasM2 || tier == atlasM5 {
		return tenant
	}
	return strings.ToUpper(provider)
}

func (opts *Opts) providerName() string {
	return providerName(opts.Tier, opts.Provider)
}

func (opts *clusterSettings) providerName() string {
	return providerName(opts.Tier, opts.Provider)
}

// Regions overlap.
func (opts *Opts) defaultRegions() ([]string, error) {
	cloudProviders, err := opts.store.CloudProviderRegions(
		opts.ConfigProjectID(),
		opts.Tier,
		[]string{opts.Provider},
	)

	if err != nil {
		return nil, err
	}

	if len(cloudProviders.GetResults()) == 0 || len(cloudProviders.GetResults()[0].GetInstanceSizes()) == 0 {
		return nil, errors.New("no regions available")
	}

	availableRegions := cloudProviders.GetResults()[0].GetInstanceSizes()[0].GetAvailableRegions()

	defaultRegions := make([]string, 0, len(availableRegions))
	popularRegionIndex := search.DefaultRegion(availableRegions)

	if popularRegionIndex != -1 {
		// the most popular region must be the first in the list
		popularRegion := availableRegions[popularRegionIndex]
		defaultRegions = append(defaultRegions, popularRegion.GetName())

		// remove popular region from availableRegions
		availableRegions = append(availableRegions[:popularRegionIndex], availableRegions[popularRegionIndex+1:]...)
	}

	for _, v := range availableRegions {
		defaultRegions = append(defaultRegions, v.GetName())
	}

	return defaultRegions, nil
}

func (opts *Opts) mdbVersions(instanceSize string) ([]string, string, error) {
	const maxItems = 500
	versions, err := opts.store.MDBVersions(
		opts.ConfigProjectID(),
		&store.MDBVersionListOptions{
			ListOptions: store.ListOptions{
				ItemsPerPage: maxItems,
			},
			InstanceSize: &instanceSize,
		},
	)

	if err != nil {
		return nil, "", err
	}

	if len(versions.GetResults()) == 0 {
		return nil, "", errors.New("no versions available")
	}

	results := map[string]bool{}
	defaultResult := ""
	for _, v := range versions.GetResults() {
		if v.GetVersion() == "" {
			continue
		}
		results[v.GetVersion()] = true
		if v.GetDefaultStatus() == "DEFAULT" {
			defaultResult = v.GetVersion()
		}
	}

	keys := make([]string, 0, len(results))
	for k := range results {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a string, b string) int {
		return semver.Compare("v"+a, "v"+b)
	})

	return keys, defaultResult, nil
}
