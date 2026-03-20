#!/usr/bin/env bash
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
step " lint (make lint)" make lint
step "build (make build)" make build
step "unit tests (make unit-test)" make unit-test

if git diff main..HEAD --name-only | grep -q '^internal/cli/'; then
    step "regenerate docs (make gen-docs)" make gen-docs
else
    printf "%-40s %b\n" "regenerate docs:" "SKIP (no cli changes)"
fi

# check if any changes were done to files that have mockgen directives
if git diff main..HEAD --name-only | xargs grep -q 'mockgen' 2>/dev/null; then
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
