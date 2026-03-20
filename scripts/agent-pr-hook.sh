#!/usr/bin/env bash
# Copyright 2026 MongoDB Inc
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
set -euo pipefail

PASS="\033[32mOK\033[0m"
FAIL="\033[31mFAIL\033[0m"
failed=0

step() {
    local name="$1"
    shift
    if output=$("$@" 2>&1); then
        printf "%-40s %b\n" "$name:" "$PASS"
    else
        printf "%-40s %b\n" "$name:" "$FAIL"
        echo "$output"
        echo ""
        failed=1
    fi
}

branch=$(git rev-parse --abbrev-ref HEAD)
if [[ "$branch" =~ ^CLOUDP-[0-9]+$ ]]; then
    printf "%-40s %b\n" "verify branch naming ($branch):" "$PASS"
else
    printf "%-40s %b\n" "verify branch naming ($branch):" "$FAIL"
    echo "  Branch name must be a JIRA ticket ID (e.g. CLOUDP-12345)."
    echo "  Rename with: git branch -m CLOUDP-XXXXX"
    echo ""
    failed=1
fi

step "get dependencies (make setup)" make setup
step "format code (make fmt)" make fmt
step "lint (make lint)" make lint
step "build (make build)" make build
step "unit tests (make unit-test)" make unit-test

if git diff master..HEAD --name-only | grep -q '^internal/cli/'; then
    step "regenerate docs (make gen-docs)" make gen-docs
else
    printf "%-40s %b\n" "regenerate docs:" "SKIP (no cli changes)"
fi

# check if any changes were done to files that have mockgen directives
if git diff master..HEAD --name-only -- '*.go' | xargs grep -q 'mockgen' 2>/dev/null; then
    step "regenerate mocks (make gen-mocks)" make gen-mocks
else
    printf "%-40s %b\n" "regenerate mocks:" "SKIP (no mockgen changes)"
fi

if [ "$failed" -ne 0 ]; then
    echo ""
    echo "Some checks failed. Fix the issues above before opening a PR."
    exit 1
fi

echo ""
echo "All checks passed."
