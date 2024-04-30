#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
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

STAGED_GO_FILES=$(git diff --cached --name-only | grep ".go$" | grep -v "mock")

for FILE in ${STAGED_GO_FILES}
do
    if [ -f "${FILE}" ]; then
        gofmt -w -s "${FILE}"
        goimports -w "${FILE}"
        git add "${FILE}"
    fi
done

if [[ -n "${STAGED_GO_FILES}" ]]; then
    set -o errexit
    go test --tags="unit integration" -race ./internal...
    make fix-lint
    set +o errexit
    for FILE in ${STAGED_GO_FILES}
    do
        git add "${FILE}"
    done

    make gen-docs > /dev/null
    git add docs
fi

STAGED_EVG_FILES=$(git diff --cached --name-only | grep "evergreen.yml$")

for FILE in ${STAGED_EVG_FILES}
do
    evergreen validate "${FILE}"
done

make check-library-owners
