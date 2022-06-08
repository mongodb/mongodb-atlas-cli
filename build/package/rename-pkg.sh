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


VERSION="$(git tag --list "${tool_name:?}/v*" --sort=taggerdate | tail -1 | cut -d "v" -f 2)"

if [[ "${unstable-}" == "-unstable" ]]; then
  VERSION="${VERSION}-next"
fi
FILENAME="${package_name-}_${VERSION}_linux_x86_64"
META_FILENAME="${meta_package_name-}_${VERSION}_linux_x86_64"

cd dist

mkdir yum apt

function rename {
  FROM=$1
  TO=$2
  ECHO "Renaming ${FROM} to ${TO}"
  mv "$FROM" "$TO"
}

# we could generate a similar name with goreleaser but we want to keep the vars evg compatible to use later
if [[ "${package_name:?}" == mongocli ]]; then
  rename "${FILENAME}.deb" "apt/mongodb-cli${unstable-}_${VERSION}${latest_deb-}_amd64.deb"
  rename "${FILENAME}.rpm" "yum/mongodb-cli${unstable-}-${VERSION}${latest_rpm-}.x86_64.rpm"
else
  rename "${FILENAME}.deb" "apt/mongodb-atlas-cli${unstable-}_${VERSION}${latest_deb-}_amd64.deb"
  rename "${FILENAME}.rpm" "yum/mongodb-atlas-cli${unstable-}-${VERSION}${latest_rpm-}.x86_64.rpm"

  rename "${META_FILENAME}.deb" "apt/mongodb-atlas${unstable-}_${VERSION}${latest_deb-}.deb"
  rename "${META_FILENAME}.rpm" "yum/mongodb-atlas${unstable-}-${VERSION}${latest_rpm-}.rpm"
fi

