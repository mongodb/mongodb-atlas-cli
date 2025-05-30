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

build/ci/ssh-ready.sh -u "$user" -i "$keyfile" -h "$hostsfile"

host=$(jq -r '.[0].dns_name' "$hostsfile")

mkdir -p ./build/package/msi/bin

cp bin/atlas.exe ./build/package/msi/bin/atlas.exe

cd ./build/package
git tag --list 'atlascli/v*' --sort=-taggerdate | head -1 | cut -d 'v' -f 2 > msi/version.txt
zip -r msi.zip msi >/dev/null 2>&1
cd ../..

scp -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "${user}@${host}" "build/package/msi.zip:/cygdrive/c/Users/Administrator/msi.zip"

ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no "${user}@${host}" bash -c 'unzip -o "/cygdrive/c/Users/Administrator/msi.zip" -d "/cygdrive/c/Users/Administrator/msi" && rm -rf "/cygdrive/c/Users/Administrator/msi.zip" && cd "/cygdrive/c/Users/Administrator/msi" && ./generate-msi.sh'

scp -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "${user}@${host}" "/cygdrive/c/Users/Administrator/msi/out.msi:${PWD}/build/package/msi/out.msi"

ls -la ./build/package/msi/
