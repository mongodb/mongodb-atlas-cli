#!/usr/bin/env bash

# Copyright 2025 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -Eeou pipefail

export WORKDIR=${workdir:?}

# Authenticate Docker to AWS ECR
aws ecr get-login-password --region us-east-1 | podman login --username AWS --password-stdin 901841024863.dkr.ecr.us-east-1.amazonaws.com

echo "Generating SBOMs..."
podman run --rm \
  -v "$WORKDIR/src/github.com/mongodb/mongodb-atlas-cli:/pwd" \
  901841024863.dkr.ecr.us-east-1.amazonaws.com/release-infrastructure/silkbomb:2.0 \
  update \
  --purls /pwd/build/package/purls.txt \
  --sbom-out /pwd/sbom.json
  