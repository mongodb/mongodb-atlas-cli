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

if [[ "${tool_name:?}" != atlascli ]]; then
	echo "invalid tool '${tool_name:?}' expected 'atlascli'"
	exit 1
fi

VERSION="$(git tag --list "${tool_name:?}/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"

FILE_EXT=deb
if [[ "${image-}" =~ "rpm" ]]; then
	FILE_EXT=rpm
fi

URL=https://mongodb-mongocli-build.s3.amazonaws.com/${project-}/dist/${revision-}_${created_at-}/mongodb-atlas_${VERSION}-next_linux_x86_64.${FILE_EXT}
ENTRYPOINT=atlas

docker build \
	--build-arg url="${URL-}" \
	--build-arg entrypoint="${ENTRYPOINT-}" \
	--build-arg server_version="5.0" \
	-t "${tool_name-}-${image-}" \
	-f "${image-}.Dockerfile" .
