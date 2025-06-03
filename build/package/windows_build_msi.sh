#!/usr/bin/env bash
# Copyright 2020 MongoDB Inc
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

keyfile="${keyfile:-./build/ci/ssh_id}"
user="${user:-Administrator}"
hostsfile="${hostsfile:-./build/ci/hosts.json}"

echo "Packaging $PWD/build/package/msi.zip"
mkdir -p ./build/package/msi/bin
cp ./dist/windows_windows_amd64_v1/bin/atlas.exe ./build/package/msi/bin/atlas.exe
git tag --list 'atlascli/v*' --sort=-taggerdate | head -1 | cut -d 'v' -f 2 > ./build/package/msi/version.txt
cd ./build/package
zip -r msi.zip msi
rm -rf ./msi/version.txt ./msi/bin/atlas.exe
cd ../..

echo "Waiting for the Windows host to become available..."
build/ci/ssh-ready.sh -u "$user" -i "$keyfile" -h "$hostsfile"
host=$(jq -r '.[0].dns_name' "$hostsfile")

echo "Uploading $PWD/build/package/msi.zip to ${user}@${host}:/cygdrive/c/Users/Administrator/msi.zip"
scp -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no "${PWD}/build/package/msi.zip" "${user}@${host}:/cygdrive/c/Users/Administrator/msi.zip"
rm -rf ./build/package/msi.zip

echo "Building MSI on ${user}@${host}..."
ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "${user}@${host}" bash -c 'unzip -o "/cygdrive/c/Users/Administrator/msi.zip" -d "/cygdrive/c/Users/Administrator/msi" && rm -rf "/cygdrive/c/Users/Administrator/msi.zip" && cd "/cygdrive/c/Users/Administrator/msi" && ./generate-msi.sh'

VERSION_GIT="$(git tag --list "atlascli/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"
VERSION_NAME="$VERSION_GIT"
if [[ "${unstable-}" == "-unstable" ]]; then
	VERSION_NAME="$VERSION_GIT-next"
fi
MSI_FILE="${PWD}/bin/mongodb-atlas-cli_${VERSION_NAME}_windows_x86_64.msi"

echo "Downloading from ${user}@${host}:/cygdrive/c/Users/Administrator/msi/out.msi to ${MSI_FILE}"
scp -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no "${user}@${host}:/cygdrive/c/Users/Administrator/msi/out.msi" "${MSI_FILE}"

echo "Cleaning up ${user}@${host}:/cygdrive/c/Users/Administrator/msi"
ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "${user}@${host}" bash -c 'rm -rf "/cygdrive/c/Users/Administrator/msi"'

echo "Notarizing ${MSI_FILE}"
echo "${ARTIFACTORY_PASSWORD}" | podman login --password-stdin --username "${ARTIFACTORY_USERNAME}" artifactory.corp.mongodb.com

echo "GRS_CONFIG_USER1_USERNAME=${GRS_USERNAME}" > .env
echo "GRS_CONFIG_USER1_PASSWORD=${GRS_PASSWORD}" >> .env

echo "signing $MSI_FILE"
podman run \
	--env-file=.env \
	--rm \
	-v "$(pwd):$(pwd)" \
	-w "$(pwd)" \
	artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-jsign \
	/bin/bash -c "jsign --tsaurl http://timestamp.digicert.com -a ${AUTHENTICODE_KEY_NAME} \"$MSI_FILE\""

rm .env
