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

export NOTARY_KEY_NAME
export BARQUE_USERNAME
export BARQUE_API_KEY
case "${NOTARY_KEY_NAME}" in
server-4.6)
	export NOTARY_TOKEN=${signing_auth_token_46:?}
	;;
server-5.0)
	export NOTARY_TOKEN=${signing_auth_token_50:?}
	;;
server-6.0)
	export NOTARY_TOKEN=${signing_auth_token_60:?}
	;;
server-7.0)
	export NOTARY_TOKEN=${signing_auth_token_70:?}
	;;
esac

set -Eeou pipefail

# Confirm package is there
echo "Confirm package is there https://s3.amazonaws.com/mongodb-mongocli-build/${project:?}/dist/${revision:?}_${created_at:?}/atlascli-${ext:?}-${arch:?}.tgz"
curl -fLO --show-error "https://s3.amazonaws.com/mongodb-mongocli-build/${project:?}/dist/${revision:?}_${created_at:?}/atlascli-${ext:?}-${arch:?}.tgz"

# --version needs to match the mongodb server version to publish to the right repo
# 4.X goes to the 4.x repo
# any *-rc version goes to testing repo
# everything else goes to development repo
curator \
	--level debug \
	repo submit \
	--service "${barque_url:?}" \
	--config build/ci/repo_config.yaml \
	--distro "${distro:?}" \
	--edition "${edition:?}" \
	--version "${server_version:?}" \
	--arch "${arch:?}" \
	--packages "https://s3.amazonaws.com/mongodb-mongocli-build/${project:?}/dist/${revision:?}_${created_at:?}/atlascli-${ext:?}-${arch:?}.tgz"
