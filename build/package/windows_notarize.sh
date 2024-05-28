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

VERSION_GIT="$(git tag --list "atlascli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"
VERSION_NAME="$VERSION_GIT"
if [[ "${unstable-}" == "-unstable" ]]; then
	VERSION_NAME="$VERSION_GIT-next"
fi

EXE_FILE="bin/atlas.exe"
MSI_FILE="bin/mongodb-atlas-cli_${VERSION_NAME}_windows_x86_64.msi"

if [[ -f "$EXE_FILE" && -f "$MSI_FILE" ]]; then
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
		/bin/bash -c "jsign --tsaurl http://timestamp.digicert.com -a mongo-authenticode-2021 \"$EXE_FILE\""
	
	echo "signing $MSI_FILE"
	podman run \
		--env-file=.env \
		--rm \
		-v "$(pwd):$(pwd)" \
		-w "$(pwd)" \
		artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-jsign \
		/bin/bash -c "jsign --tsaurl http://timestamp.digicert.com -a mongo-authenticode-2021 \"$MSI_FILE\""
	
	rm .env
fi
