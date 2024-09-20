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

if [[ "${DOCS_ATLAS_CLI_TOKEN:?}" == "" ]]; then
    echo "missing \$DOCS_ATLAS_CLI_TOKEN"
    exit 1
fi

if [[ "${CLOUD_DOCS_TOKEN:?}" == "" ]]; then
    echo "missing \$CLOUD_DOCS_TOKEN"
    exit 1
fi

if [[ "${EVERGREEN_USER:?}" == "" ]]; then
    echo "missing \$EVERGREEN_USER"
    exit 1
fi

if [[ "${EVERGREEN_API_KEY:?}" == "" ]]; then
    echo "missing \$EVERGREEN_API_KEY"
    exit 1
fi

if [[ "${DOCS_SLACK_CHANNEL:?}" == "" ]]; then
    echo "missing \$DOCS_SLACK_CHANNEL"
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

echo "https://x-access-token:${DOCS_ATLAS_CLI_TOKEN:?}@github.com" > .git-credentials
echo "https://x-access-token:${DOCS_ATLAS_CLI_TOKEN:?}@api.github.com" >> .git-credentials

docker run \
    --name docs-atlas-cli \
    -v "${PWD}:/usr/src/app" \
    -v "${PWD}/.gitconfig:/root/.gitconfig" \
    -v "${PWD}/.git-credentials:/root/.git-credentials" \
    -e COPYBARA_WORKFLOW=docs-atlas-cli \
    -e "COPYBARA_OPTIONS=--github-api-bearer-auth true" \
    google/copybara

DOCS_ATLAS_CLI_PR_URL=$(docker logs docs-atlas-cli 2>&1 | grep "/pull/" | sed -E 's/^.*(https\:.*)$/\1/')

echo "https://x-access-token:${CLOUD_DOCS_TOKEN:?}@github.com" > .git-credentials
echo "https://x-access-token:${CLOUD_DOCS_TOKEN:?}@api.github.com" >> .git-credentials

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

TARGET="filipe.menezes" #$DOCS_SLACK_CHANNEL
MSG="[TESTING PLEASE IGNORE] Hey team :wave: ${DOCS_ATLAS_CLI_PR_URL} is ready for review :thankyou:"
curl --header "Api-User:${EVERGREEN_USER:?}" \
    --header "Api-Key:${EVERGREEN_API_KEY:?}" \
    --request POST "https://evergreen.mongodb.com/rest/v2/notifications/slack" \
    --data "{\"target\":\"$TARGET\",\"msg\":\"$MSG\"}"

MSG="[TESTING PLEASE IGNORE] Hey team :wave: ${CLOUD_DOCS_PR_URL} is ready for review :thankyou:"
curl --header "Api-User:${EVERGREEN_USER:?}" \
    --header "Api-Key:${EVERGREEN_API_KEY:?}" \
    --request POST "https://evergreen.mongodb.com/rest/v2/notifications/slack" \
    --data "{\"target\":\"$TARGET\",\"msg\":\"$MSG\"}"
