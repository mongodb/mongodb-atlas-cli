// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/clusters/connect"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	errGeneratingConnectionString = errors.New("error generating connection string for deployment")
	errConnectingDeployment       = errors.New("error connecting to deployment")
	errDecodingUUID               = errors.New("error decoding uuid")
	errCastingUUID                = errors.New("error casting uuid")
	errUUIDNotFound               = errors.New("error uuid is not found")
	errEmptyUUID                  = errors.New("error uuid is empty")
)

//go:generate go tool go.uber.org/mock/mockgen -destination=../../../mocks/mock_deployment_opts_telemetry.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options DeploymentTelemetry
type DeploymentTelemetry interface {
	AppendDeploymentType()
	AppendDeploymentUUID()
	AppendClusterWideScalingMode()
	AppendIndependentShardScalingMode()
}

func NewDeploymentTypeTelemetry(opts *DeploymentOpts) DeploymentTelemetry {
	return opts
}

func (opts *DeploymentOpts) collectUUID(ctx context.Context) error {
	if opts.DeploymentUUID == "" {
		cnnStr, err := opts.ConnectionString(ctx)
		if err != nil {
			return fmt.Errorf("%w: %w", errGeneratingConnectionString, err)
		}

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(cnnStr))
		if err != nil {
			return fmt.Errorf("%w: %w", errConnectingDeployment, err)
		}

		defer func() {
			_ = client.Disconnect(ctx)
		}()

		atlascli := map[string]any{}

		if err := client.Database("admin").Collection("atlascli").FindOne(ctx, bson.M{}).Decode(&atlascli); err != nil {
			return fmt.Errorf("%w: %w", errDecodingUUID, err)
		}

		if atlascli["uuid"] == nil {
			return errUUIDNotFound
		}

		uuid, ok := atlascli["uuid"].(string)
		if !ok {
			return errCastingUUID
		}

		if uuid == "" {
			return errEmptyUUID
		}

		opts.DeploymentUUID = uuid
	}

	telemetry.AppendOption(telemetry.WithDeploymentUUID(opts.DeploymentUUID))

	return nil
}

func (opts *DeploymentOpts) AppendDeploymentType() {
	var deploymentType string
	if opts.IsLocalDeploymentType() {
		deploymentType = LocalCluster
	} else if opts.IsAtlasDeploymentType() {
		deploymentType = connect.AtlasCluster
	}
	if deploymentType != "" {
		telemetry.AppendOption(telemetry.WithDeploymentType(deploymentType))
	}
}

func (opts *DeploymentOpts) AppendDeploymentUUID() {
	if opts.IsAtlasDeploymentType() || opts.DeploymentType == "" || opts.DeploymentUUID != "" {
		return
	}

	if err := opts.collectUUID(context.TODO()); err != nil {
		_, _ = log.Debugf("error collecting deployment uuid: %v", err)
	}
}

func (opts *DeploymentOpts) AppendClusterWideScalingMode() {
	if opts.IsLocalDeploymentType() {
		return
	}
	telemetry.AppendOption(telemetry.WithDetectedAutoScalingMode("clusterWideScaling"))
}

func (opts *DeploymentOpts) AppendIndependentShardScalingMode() {
	if opts.IsLocalDeploymentType() {
		return
	}
	telemetry.AppendOption(telemetry.WithDetectedAutoScalingMode("independentShardScaling"))
}
