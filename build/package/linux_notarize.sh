#!/usr/bin/env bash

# Copyright 2024 MongoDB Inc
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

#set -Eeou pipefail
#
## Notarize generated binaries with GPG and replace the original binary with the notarized one
## This depends on binaries being generated in a goreleaser manner and gon being set up.
## goreleaser should already take care of calling this script as a hook.
#
#AMD64_DEB_FILE="./dist/linux_linux_amd64_v1/bin/mongocli/mongodbcli_${VERSION}-next_linux_x86_64.deb"
#ARM64_DEB_FILE="./dist/linux_linux_arm64/bin/mongocli/mongodbcli_${VERSION}-next_linux_arm64.deb"
#AMD64_RPM_FILE="./dist/linux_linux_amd64_v1/bin/mongocli/mongodbcli_${VERSION}-next_linux_x86_64.rpm"
#ARM64_RPM_FILE="./dist/linux_linux_arm64/bin/mongocli/mongodbcli_${VERSION}-next_linux_arm64.rpm"
#AMD64_TAR_FILE="./dist/linux_linux_amd64_v1/bin/mongocli/mongodbcli_${VERSION}-next_linux_x86_64.tar.gz"
#ARM64_TAR_FILE="./dist/linux_linux_arm64/bin/mongocli/mongodbcli_${VERSION}-next_linux_arm64.tar.gz"
#
#echo "here1"
#if [[ "${TOOL_NAME:?}" == atlascli ]]; then
#  AMD64_DEB_FILE="./dist/mongodb-atlas-cli_${VERSION}_linux_x86_64.deb"
#  ARM64_DEB_FILE="./dist/mongodb-atlas-cli_${VERSION}_linux_arm64.deb"
#  AMD64_RPM_FILE="./dist/mongodb-atlas-cli_${VERSION}_linux_x86_64.rpm"
#  ARM64_RPM_FILE="./dist/mongodb-atlas-cli_${VERSION}_linux_arm64.rpm"
#  AMD64_TAR_FILE="dist/mongodb-atlas-cli_${VERSION}_linux_x86_64.tar.gz"
#  ARM64_TAR_FILE="dist/mongodb-atlas-cli_${VERSION}_linux_arm64.tar.gz"
#fi
#
#echo "here2"
#if [[ -f "${AMD64_DEB_FILE}" && -f "${ARM64_DEB_FILE}" &&  -f "${AMD64_RPM_FILE}" && -f "${ARM64_RPM_FILE}" ]]; then
#  echo "Yessssssss!!!!!!!"
#  echo "${ARTIFACTORY_PASSWORD}" | docker login --password-stdin --username "${ARTIFACTORY_USERNAME}" artifactory.corp.mongodb.com
#
#  echo "GRS_CONFIG_USER1_USERNAME=${GRS_USERNAME}" >> "signing-envfile"
#  echo "GRS_CONFIG_USER1_PASSWORD=${GRS_PASSWORD}" >> "signing-envfile"
#
#	echo "notarizing Linux binary ${AMD64_DEB_FILE}"
#  docker run \
#    --env-file=signing-envfile \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${AMD64_DEB_FILE}"
#
#	echo "notarizing Linux binary ${ARM64_DEB_FILE}"
#  docker run \
#    --env-file=signing-envfile \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${ARM64_DEB_FILE}"
#
#	echo "notarizing Linux binary ${AMD64_RPM_FILE}"
#  docker run \
#    --env-file=signing-envfile \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${AMD64_RPM_FILE}"
#
#	echo "notarizing Linux binary ${ARM64_RPM_FILE}"
#  docker run \
#    --env-file=signing-envfile \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${ARM64_RPM_FILE}"
#fi
#
#echo "HERE tar ${AMD64_TAR_FILE} ${ARM64_TAR_FILE}"
#if [[ -f "${AMD64_TAR_FILE}" &&  -f "${ARM64_TAR_FILE}" ]]; then
#  echo "${ARTIFACTORY_PASSWORD}" | docker login --password-stdin --username "${ARTIFACTORY_USERNAME}" artifactory.corp.mongodb.com
#
#	echo "notarizing Linux binary ${AMD64_TAR_FILE}"
#  docker run \
#    -e GRS_CONFIG_USER1_USERNAME="${GRS_USERNAME}" \
#    -e GRS_CONFIG_USER1_PASSWORD="${GRS_PASSWORD}" \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${AMD64_TAR_FILE}"
#
#	echo "notarizing Linux binary ${ARM64_TAR_FILE}"
#  docker run \
#    -e GRS_CONFIG_USER1_USERNAME="${GRS_USERNAME}" \
#    -e GRS_CONFIG_USER1_PASSWORD="${GRS_PASSWORD}" \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${ARM64_TAR_FILE}"
#fi

echo "Signing completed."
