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

ssh -i "$keyfile" -o ConnectTimeout=10 -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt "${user}@${host}" echo "SSH connection to $host successful."
