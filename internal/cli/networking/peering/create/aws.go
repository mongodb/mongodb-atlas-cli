// Copyright 2020 MongoDB Inc
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

package create

import (
	"context"
	"fmt"
	"strings"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=aws_mock_test.go -package=create . AWSPeeringConnectionCreator

type AWSPeeringConnectionCreator interface {
	AWSContainers(string) ([]atlasv2.CloudProviderContainer, error)
	CreateContainer(string, *atlasv2.CloudProviderContainer) (*atlasv2.CloudProviderContainer, error)
	CreatePeeringConnection(string, *atlasv2.BaseNetworkPeeringConnectionSettings) (*atlasv2.BaseNetworkPeeringConnectionSettings, error)
}

type AWSOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	region              string
	routeTableCidrBlock string
	accountID           string
	vpcID               string
	atlasCIDRBlock      string
	store               AWSPeeringConnectionCreator
}

func (opts *AWSOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *AWSOpts) Run() error {
	opts.region = normalizeAtlasRegion(opts.region)
	container, err := opts.containerExists()
	if err != nil {
		return err
	}

	if container == nil {
		var err2 error
		r, err2 := opts.store.CreateContainer(opts.ConfigProjectID(), opts.newContainer())
		container = r
		if err2 != nil {
			return err2
		}
	}
	r, err := opts.store.CreatePeeringConnection(opts.ConfigProjectID(), opts.newPeer(*container.Id))
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *AWSOpts) containerExists() (*atlasv2.CloudProviderContainer, error) {
	r, err := opts.store.AWSContainers(opts.ConfigProjectID())
	if err != nil {
		return nil, err
	}
	for i := range r {
		if r[i].GetRegionName() == opts.region {
			return &r[i], nil
		}
	}
	return nil, nil
}

func (opts *AWSOpts) newContainer() *atlasv2.CloudProviderContainer {
	return &atlasv2.CloudProviderContainer{
		AtlasCidrBlock: &opts.atlasCIDRBlock,
		RegionName:     &opts.region,
		ProviderName:   pointer.Get("AWS"),
	}
}

func normalizeAtlasRegion(region string) string {
	region = strings.ToUpper(region)
	return strings.ReplaceAll(region, "-", "_")
}

func (opts *AWSOpts) newPeer(containerID string) *atlasv2.BaseNetworkPeeringConnectionSettings {
	provider := "AWS"
	region := strings.ToLower(opts.region)
	region = strings.ReplaceAll(region, "_", "-")
	return &atlasv2.BaseNetworkPeeringConnectionSettings{
		ProviderName:        &provider,
		AccepterRegionName:  &region,
		AwsAccountId:        &opts.accountID,
		ContainerId:         containerID,
		RouteTableCidrBlock: &opts.routeTableCidrBlock,
		VpcId:               &opts.vpcID,
	}
}

// AwsBuilder atlas networking peering create aws
// --accepterRegionName accepterRegionName: Specifies the region where the peer VPC resides.
// --awsAccountId awsAccountId: Account ID of the owner of the peer VPC.
// --containerId containerId: Unique identifier of the Atlas VPC container for the region.
// --routeTableCidrBlock routeTableCidrBlock: 	Peer VPC CIDR block or subnet.
// --vpcID vpcID: Unique identifier of the peer VPC.
// --projectId projectId: ID of the project
// Create a network peering with AWS, this command will internally check if a container already exists for the provider and region and if it does then we’ll use that,
// if it does not exists we’ll try to create one and use it,
// there can only be one container per provider and region.
func AwsBuilder() *cobra.Command {
	opts := &AWSOpts{}
	cmd := &cobra.Command{
		Use:   "aws",
		Short: "Create a network peering connection between the Atlas VPC and your AWS VPC.",
		Long:  longDesc + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Args: require.NoArgs,
		Example: `  # Create a network peering connection between the Atlas VPC in CIDR block 192.168.0.0/24 and your AWS VPC in CIDR block 10.0.0.0/24 for AWS account number 854333054055:
  atlas networking peering create aws --accountId 854333054055 --atlasCidrBlock 192.168.0.0/24 --region us-east-1 --routeTableCidrBlock 10.0.0.0/24 --vpcId vpc-078ac381aa90e1e63`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.accountID, flag.AccountID, "", usage.AccountID)
	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.ContainerRegion)
	cmd.Flags().StringVar(&opts.routeTableCidrBlock, flag.RouteTableCidrBlock, "", usage.RouteTableCidrBlock)
	cmd.Flags().StringVar(&opts.vpcID, flag.VpcID, "", usage.VpcID)
	cmd.Flags().StringVar(&opts.atlasCIDRBlock, flag.AtlasCIDRBlock, "", usage.AtlasCIDRBlock)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.AccountID)
	_ = cmd.MarkFlagRequired(flag.RouteTableCidrBlock)
	_ = cmd.MarkFlagRequired(flag.VpcID)
	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
