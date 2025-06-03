#!/usr/bin/env bash

# Copyright 2022 MongoDB Inc
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

EXE_FILE="dist/windows_windows_amd64_v1/bin/atlas.exe"

if [[ -f "$EXE_FILE" ]]; then
	echo "${ARTIFACTORY_PASSWORD}" | podman login --password-stdin --username "${ARTIFACTORY_USERNAME}" artifactory.corp.mongodb.com

	echo "GRS_CONFIG_USER1_USERNAME=${GRS_USERNAME}" > .env
	echo "GRS_CONFIG_USER1_PASSWORD=${GRS_PASSWORD}" >> .env

	echo "signing $EXE_FILE"
	podman run \
		--env-file=.env \
		--rm \
		-v "$(pwd):$(pwd)" \
		-w "$(pwd)" \
		artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-jsign \
		/bin/bash -c "jsign --tsaurl http://timestamp.digicert.com -a ${AUTHENTICODE_KEY_NAME} \"$EXE_FILE\""
	
	rm .env
fi
