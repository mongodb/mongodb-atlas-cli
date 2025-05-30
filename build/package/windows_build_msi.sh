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

echo "$PWD"
ls -la

HOSTNAME=$(jq -r '.[0].dns_name' < hosts.json)
USERNAME=Administrator

identity_file=~/.ssh/mcipacker.pem
attempts=0
connection_attempts=25

while ! ssh \
    -i "$identity_file" \
    -o ConnectTimeout=10 \
    -o ForwardAgent=yes \
    -o IdentitiesOnly=yes \
    -o StrictHostKeyChecking=no \
    "$(printf "%s@%s" "$USERNAME" "$HOSTNAME")" \
    exit
do
    ((attempts++))
    [ "$attempts" -ge "$connection_attempts" ] && printf "SSH connection attempt %d/%d failed." "$attempts" "$connection_attempts" && exit 1
    printf "SSH connection attempt %d/%d failed. Retrying...\n" "$attempts" "$connection_attempts"
    sleep 10
done

ssh \
    -i "$identity_file" \
    -o ConnectTimeout=10 \
    -o ForwardAgent=yes \
    -o IdentitiesOnly=yes \
    -o StrictHostKeyChecking=no \
    "$(printf "%s@%s" "$USERNAME" "$HOSTNAME")" \
    bash -c 'echo "echo from remote host"'
