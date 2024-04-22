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

import "context"

func (opts *DeploymentOpts) RemoveLocal(ctx context.Context) error {
	volumes := []string{opts.LocalMongodDataVolume(), opts.LocalMongotDataVolume(), opts.LocalMongoMetricsVolume()}

	if c, _ := opts.PodmanClient.ContainerInspect(ctx, opts.LocalMongodHostname()); c != nil {
		for _, m := range c[0].Mounts {
			if m.Name != opts.LocalMongodDataVolume() {
				volumes = append(volumes, m.Name)
				break
			}
		}
	}

	if _, errRemove := opts.PodmanClient.RemoveContainers(ctx, opts.LocalMongodHostname(), opts.LocalMongotHostname()); errRemove != nil {
		return errRemove
	}

	if _, errRemove := opts.PodmanClient.RemoveNetworks(ctx, opts.LocalNetworkName()); errRemove != nil {
		return errRemove
	}

	if _, errRemove := opts.PodmanClient.RemoveVolumes(ctx, volumes...); errRemove != nil {
		return errRemove
	}

	return nil
}
