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

package quickstart

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/usage"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func (opts *Opts) createCluster() error {
	if err := opts.askClusterOptions(); err != nil {
		return err
	}
	if _, err := opts.store.CreateCluster(opts.newCluster()); err != nil {
		return err
	}

	return nil
}

func (opts *Opts) askClusterOptions() error {
	var qs []*survey.Question

	if opts.ClusterName == "" {
		opts.ClusterName = opts.defaultName
		qs = append(qs, newClusterNameQuestion(opts.ClusterName))
	}

	if opts.Provider == "" {
		qs = append(qs, newClusterProviderQuestion())
	}

	if opts.Provider == "" || opts.ClusterName == "" || opts.Region == "" {
		fmt.Print(`
[Set up your Atlas cluster]
`)
	}

	if err := survey.Ask(qs, opts); err != nil {
		return err
	}

	// We need the provider to ask for the region
	if opts.Region == "" {
		return opts.askClusterRegion()
	}
	return nil
}

func (opts *Opts) askClusterRegion() error {
	regions, err := opts.defaultRegions()
	if err != nil {
		return err
	}
	regionQ := newRegionQuestions(regions)
	return survey.AskOne(regionQ, &opts.Region, survey.WithValidator(survey.Required))
}

func newRegionQuestions(defaultRegions []string) survey.Prompt {
	return &survey.Select{
		Message: "Cloud Provider Region",
		Help:    usage.Region,
		Options: defaultRegions,
	}
}

func (opts *Opts) newCluster() *atlas.Cluster {
	diskSizeGB := atlas.DefaultDiskSizeGB[strings.ToUpper(tenant)][opts.tier]
	mdbVersion, _ := cli.DefaultMongoDBMajorVersion()
	cluster := &atlas.Cluster{
		GroupID:             opts.ConfigProjectID(),
		ClusterType:         replicaSet,
		ReplicationSpecs:    []atlas.ReplicationSpec{opts.newReplicationSpec()},
		ProviderSettings:    opts.newProviderSettings(),
		MongoDBMajorVersion: mdbVersion,
		DiskSizeGB:          &diskSizeGB,
		Name:                opts.ClusterName,
		Labels: []atlas.Label{
			{
				Key:   "Infrastructure Tool",
				Value: "MongoDB CLI Quickstart",
			},
		},
	}

	return cluster
}

const (
	shards   = 1
	members  = 3
	zoneName = "Zone 1"
)

func (opts *Opts) newReplicationSpec() atlas.ReplicationSpec {
	var (
		readOnlyNodes int64
		priority      int64 = 7
		shards        int64 = shards
		members       int64 = members
	)
	replicationSpec := atlas.ReplicationSpec{
		NumShards: &shards,
		ZoneName:  zoneName,
		RegionsConfig: map[string]atlas.RegionsConfig{
			opts.Region: {
				ReadOnlyNodes:  &readOnlyNodes,
				ElectableNodes: &members,
				Priority:       &priority,
			},
		},
	}
	return replicationSpec
}

const (
	tenant  = "TENANT"
	atlasM5 = "M5"
)

func (opts *Opts) newProviderSettings() *atlas.ProviderSettings {
	providerName := opts.providerName()

	var backingProviderName string
	if providerName == tenant {
		backingProviderName = opts.Provider
	}

	return &atlas.ProviderSettings{
		InstanceSizeName:    opts.tier,
		ProviderName:        providerName,
		RegionName:          opts.Region,
		BackingProviderName: backingProviderName,
	}
}

func (opts *Opts) providerName() string {
	if opts.tier == atlasM2 || opts.tier == atlasM5 {
		return tenant
	}
	return opts.Provider
}

func (opts *Opts) defaultRegions() ([]string, error) {
	cloudProviders, err := opts.store.CloudProviderRegions(
		opts.ConfigProjectID(),
		opts.tier,
		[]*string{&opts.Provider},
	)

	if err != nil {
		return nil, err
	}

	if len(cloudProviders.Results) == 0 || len(cloudProviders.Results[0].InstanceSizes) == 0 {
		return nil, errors.New("no regions available")
	}

	availableRegions := cloudProviders.Results[0].InstanceSizes[0].AvailableRegions

	defaultRegions := make([]string, 0, len(availableRegions))
	popularRegionIndex := search.DefaultRegion(availableRegions)

	if popularRegionIndex != -1 {
		// the most popular region must be the first in the list
		popularRegion := availableRegions[popularRegionIndex]
		defaultRegions = append(defaultRegions, popularRegion.Name)

		// remove popular region from availableRegions
		availableRegions = append(availableRegions[:popularRegionIndex], availableRegions[popularRegionIndex+1:]...)
	}

	for _, v := range availableRegions {
		defaultRegions = append(defaultRegions, v.Name)
	}

	return defaultRegions, nil
}
