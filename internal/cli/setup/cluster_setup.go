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
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var ErrNoRegions = errors.New("no regions found for the cloud provider")

func (opts *Opts) createCluster() error {
	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
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

	if opts.shouldAskForValue(flag.ClusterName) || opts.shouldAskForValue(flag.Provider) || opts.shouldAskForValue(flag.Region) {
		fmt.Print(`
[Set up your Atlas cluster]
`)
	}

	if err := telemetry.TrackAsk(qs, opts); err != nil {
		return err
	}

	// We need the provider to ask for the region
	if opts.shouldAskForValue(flag.Region) {
		return opts.askClusterRegion()
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

func newRegionQuestions(defaultRegions []string) survey.Prompt {
	return &survey.Select{
		Message: "Cloud Provider Region",
		Help:    usage.Region,
		Options: defaultRegions,
	}
}

func defaultDiskSizeGB(provider, tier string) float64 {
	return atlas.DefaultDiskSizeGB[strings.ToUpper(provider)][tier]
}

func (opts *Opts) newCluster() *atlasv2.AdvancedClusterDescription {
	cluster := &atlasv2.AdvancedClusterDescription{
		GroupId:                      pointer.Get(opts.ConfigProjectID()),
		ClusterType:                  pointer.Get(replicaSet),
		ReplicationSpecs:             &[]atlasv2.ReplicationSpec{opts.newAdvanceReplicationSpec()},
		Name:                         &opts.ClusterName,
		TerminationProtectionEnabled: &opts.EnableTerminationProtection,
	}

	if len(opts.Tag) > 0 {
		var tags []atlasv2.ResourceTag
		for k, v := range opts.Tag {
			if k != "" && v != "" {
				tags = append(tags, atlasv2.ResourceTag{Key: k, Value: v})
			}
		}
		cluster.Tags = &tags
	}

	if opts.providerName() != tenant {
		diskSizeGB := defaultDiskSizeGB(opts.providerName(), opts.Tier)
		mdbVersion, _ := cli.DefaultMongoDBMajorVersion()
		cluster.DiskSizeGB = &diskSizeGB
		cluster.MongoDBMajorVersion = &mdbVersion
	}

	return cluster
}

var (
	shards   = 1
	zoneName = "Zone 1"
)

func (opts *Opts) newAdvanceReplicationSpec() atlasv2.ReplicationSpec {
	return atlasv2.ReplicationSpec{
		NumShards:     &shards,
		ZoneName:      &zoneName,
		RegionConfigs: &[]atlasv2.CloudRegionConfig{opts.newAdvancedRegionConfig()},
	}
}

const (
	tenant  = "TENANT"
	atlasM2 = "M2"
	atlasM5 = "M5"
)

func (opts *Opts) newAdvancedRegionConfig() atlasv2.CloudRegionConfig {
	providerName := opts.providerName()

	priority := 7
	regionConfig := atlasv2.CloudRegionConfig{
		ProviderName: &providerName,
		Priority:     &priority,
		RegionName:   &opts.Region,
	}

	regionConfig.ElectableSpecs = &atlasv2.HardwareSpec{
		InstanceSize: &opts.Tier,
	}
	members := 3
	if providerName == tenant {
		regionConfig.BackingProviderName = &opts.Provider
	} else {
		regionConfig.ElectableSpecs.NodeCount = &members
	}

	return regionConfig
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
