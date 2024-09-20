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

if [[ "${TAG:?}" == "" ]]; then
    echo "missing \$TAG"
    exit 1
fi

if [[ "${docs_atlas_cli_token:?}" == "" ]]; then
    echo "missing \$docs_atlas_cli_token"
    exit 1
fi

if [[ "${cloud_docs_token:?}" == "" ]]; then
    echo "missing \$cloud_docs_token"
    exit 1
fi

if [[ "${evergreen_user:?}" == "" ]]; then
    echo "missing \$evergreen_user"
    exit 1
fi

if [[ "${evergreen_api_key:?}" == "" ]]; then
    echo "missing \$evergreen_api_key"
    exit 1
fi

if [[ "${docs_slack_channel:?}" == "" ]]; then
    echo "missing \$docs_slack_channel"
    exit 1
fi

cat <<EOF > .gitconfig
[user]
    name = apix-bot[bot]
    email = 168195273+apix-bot[bot]@users.noreply.github.com
[credential]
    helper = store
EOF

envsubst < copy.bara.sky.template > copy.bara.sky

echo "https://x-access-token:${docs_atlas_cli_token:?}@github.com" > .git-credentials
echo "https://x-access-token:${docs_atlas_cli_token:?}@api.github.com" >> .git-credentials

docker run \
    --name docs-atlas-cli \
    -v "${PWD}:/usr/src/app" \
    -v "${PWD}/.gitconfig:/root/.gitconfig" \
    -v "${PWD}/.git-credentials:/root/.git-credentials" \
    -e COPYBARA_WORKFLOW=docs-atlas-cli \
    -e "COPYBARA_OPTIONS=--github-api-bearer-auth true" \
    google/copybara

DOCS_ATLAS_CLI_PR_URL=$(docker logs docs-atlas-cli 2>&1 | grep "/pull/" | sed -E 's/^.*(https\:.*)$/\1/')

echo "https://x-access-token:${cloud_docs_token:?}@github.com" > .git-credentials
echo "https://x-access-token:${cloud_docs_token:?}@api.github.com" >> .git-credentials

docker run \
    --name cloud-docs \
    -v "${PWD}:/usr/src/app" \
    -v "${PWD}/.gitconfig:/root/.gitconfig" \
    -v "${PWD}/.git-credentials:/root/.git-credentials" \
    -e COPYBARA_WORKFLOW=cloud-docs \
    -e "COPYBARA_OPTIONS=--github-api-bearer-auth true" \
    google/copybara

CLOUD_DOCS_PR_URL=$(docker logs cloud-docs 2>&1 | grep "/pull/" | sed -E 's/^.*(https\:.*)$/\1/')

rm -rf .git-credentials .gitconfig copy.bara.sky
docker rm -f cloud-docs docs-atlas-cli

TARGET="filipe.menezes" #$docs_slack_channel
MSG="[TESTING PLEASE IGNORE] Hey team :wave: ${DOCS_ATLAS_CLI_PR_URL} is ready for review :thankyou:"
curl --header "Api-User:${evergreen_user:?}" \
    --header "Api-Key:${evergreen_api_key:?}" \
    --request POST "https://evergreen.mongodb.com/rest/v2/notifications/slack" \
    --data "{\"target\":\"$TARGET\",\"msg\":\"$MSG\"}"

MSG="[TESTING PLEASE IGNORE] Hey team :wave: ${CLOUD_DOCS_PR_URL} is ready for review :thankyou:"
curl --header "Api-User:${evergreen_user:?}" \
    --header "Api-Key:${evergreen_api_key:?}" \
    --request POST "https://evergreen.mongodb.com/rest/v2/notifications/slack" \
    --data "{\"target\":\"$TARGET\",\"msg\":\"$MSG\"}"
