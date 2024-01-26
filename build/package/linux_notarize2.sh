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

set -Eeou pipefail

# Notarize generated binaries with GPG and replace the original binary with the notarized one
# This depends on binaries being generated in a goreleaser manner and gon being set up.
# goreleaser should already take care of calling this script as a hook.


echo "here2 ${LINUX_FILE}"
if [[ -f "${LINUX_FILE}" ]]; then
  echo "YESSSSSSSSSSS here2"
  echo "${ARTIFACTORY_PASSWORD}" | docker login --password-stdin --username "${ARTIFACTORY_USERNAME}" artifactory.corp.mongodb.com

  echo ANDREA > dist/test.asc

#	echo "notarizing Linux binary ${AMD64_TAR_FILE}"
#  docker run \
#    -e GRS_CONFIG_USER1_USERNAME="${GRS_USERNAME}" \
#    -e GRS_CONFIG_USER1_PASSWORD="${GRS_PASSWORD}" \
#    --rm \
#    -v "$(pwd)":"$(pwd)" \
#    -w "$(pwd)" \
#    artifactory.corp.mongodb.com/release-tools-container-registry-local/garasign-gpg \
#    /bin/bash -c "gpgloader && gpg --yes -v --armor --detach-sign ${LINUX_FILE}"
fi

echo "Signing completed."
