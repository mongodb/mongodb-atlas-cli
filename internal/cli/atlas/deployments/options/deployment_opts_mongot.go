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
	"fmt"
	"time"
)

const (
	waitMongotTimeout  = 5 * time.Minute
	internalMongotPort = 27027
)

func (*DeploymentOpts) InternalMongotAddress(ip string) string {
	return fmt.Sprintf("%s:%d", ip, internalMongotPort)
}

func (opts *DeploymentOpts) MongotIP(ctx context.Context) (string, error) {
	containers, err := opts.PodmanClient.ContainerInspect(ctx, opts.LocalMongotHostname())
	if err != nil {
		return "", err
	}
	c := containers[0]
	n, ok := c.NetworkSettings.Networks[opts.LocalNetworkName()]
	if !ok {
		return "", ErrDeploymentNotFound
	}
	return n.IPAddress, nil
}

func (opts *DeploymentOpts) WaitForMongot(parentCtx context.Context, ip string) error {
	ctx, cancel := context.WithTimeout(parentCtx, waitMongotTimeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := opts.PodmanClient.Exec(ctx, opts.LocalMongodHostname(), "/bin/sh", "-c", fmt.Sprintf("mongosh %s --eval \"db.adminCommand('ping')\"", opts.InternalMongotAddress(ip))); err == nil { // ping was successful
				return nil
			}
		}
	}
}
