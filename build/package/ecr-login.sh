#!/usr/bin/env bash

# Copyright 2026 MongoDB Inc
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

# Logs the available container runtimes into the DevProd Platforms ECR registry, which hosts the
# Garasign signing images and the Docker Hub pull-through cache. Needs AWS credentials in the
# environment; in Evergreen these come from ec2.assume_role.
# See https://docs.devprod.prod.corp.mongodb.com/devprod-platforms-ecr

set -Eeou pipefail

ECR_REGISTRY=${ECR_REGISTRY:-901841024863.dkr.ecr.us-east-1.amazonaws.com}
ECR_REGION=${ECR_REGION:-us-east-1}

password=$(aws ecr get-login-password --region "${ECR_REGION}")

# Prepare the docker config so the token can be stored. Credential helpers need an unlocked keychain,
# which CI does not have. Drop them, keeping any credentials already in the file.
docker_config="${DOCKER_CONFIG:-${HOME}/.docker}/config.json"
mkdir -p "$(dirname "${docker_config}")"
if [[ ! -f "${docker_config}" ]]; then
  # Linux: no config yet so we seed it with an empty JSON object so the JSON edits below have something to parse.
  echo '{}' > "${docker_config}"
elif command -v python3 > /dev/null 2>&1; then
  # macOS: config ships with a helper configured. No-op elsewhere.
  python3 - "${docker_config}" <<'PY'
import json, sys

path = sys.argv[1]
with open(path) as f:
    config = json.load(f)

config.pop("credsStore", None)
config.pop("credHelpers", None)

with open(path, "w") as f:
    json.dump(config, f)
PY
fi

# Log in every container runtime installed on the host. Which one that is depends on the host:
# docker on ubuntu/debian/al2023/macOS, podman on rhel and the packaging hosts.
logged_in=false
for runtime in docker podman; do
  if ! command -v "${runtime}" > /dev/null 2>&1; then
    continue
  fi

  echo "Authenticating ${runtime} to ${ECR_REGISTRY}"
  if echo "${password}" | "${runtime}" login --username AWS --password-stdin "${ECR_REGISTRY}"; then
    logged_in=true
    continue
  fi

  # A podman failure is not fatal: macOS ships the podman client but no running machine, so its
  # login can never succeed, and docker is the runtime the tests use there. When podman is the only
  # runtime (rhel, packaging hosts) logged_in stays false and the check after the loop still exits.
  if [[ "${runtime}" != "docker" ]]; then
    echo "${runtime} login to ${ECR_REGISTRY} failed, skipping it"
    continue
  fi

  # macOS: docker resolves a default keychain helper even with none configured, so the login
  # above authenticates and then fails to store. Writing the token ourselves both stores it and
  # makes docker read from config.json from now on.
  echo "docker login failed to store the token, writing it to ${docker_config} instead"
  ECR_REGISTRY="${ECR_REGISTRY}" ECR_PASSWORD="${password}" python3 - "${docker_config}" <<'PY'
import base64, json, os, sys

path = sys.argv[1]
with open(path) as f:
    config = json.load(f)

auth = base64.b64encode(f"AWS:{os.environ['ECR_PASSWORD']}".encode()).decode()
config.setdefault("auths", {})[os.environ["ECR_REGISTRY"]] = {"auth": auth}

with open(path, "w") as f:
    json.dump(config, f)
PY
  chmod 600 "${docker_config}"
  logged_in=true
done

# Fail only if nothing authenticated
if [[ "${logged_in}" == "false" ]]; then
  echo "no usable container runtime could authenticate to ${ECR_REGISTRY}"
  exit 1
fi
