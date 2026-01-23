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

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLI_ROOT="${SCRIPT_DIR}/.."
CURRENT_SDK_RELEASE=$(cat "${CLI_ROOT}/.atlas-sdk-version")
SDK_MODULE_PATH="go.mongodb.org/atlas-sdk/${CURRENT_SDK_RELEASE}"
GO_MOD_BACKUP="${CLI_ROOT}/go.mod.backup"
GO_SUM_BACKUP="${CLI_ROOT}/go.sum.backup"
SDK_REPO_URL="https://github.com/mongodb/atlas-sdk-go.git"
SDK_TEMP_DIR=$(mktemp -d -t atlas-sdk-go-XXXXXX)

# Cleanup function to restore go.mod and go.sum, and remove temp SDK directory
cleanup() {
    local exit_code=$?
    cd "${CLI_ROOT}"
    if [ -f "${GO_MOD_BACKUP}" ]; then
        echo ""
        echo "==> Restoring original go.mod..."
        mv "${GO_MOD_BACKUP}" go.mod
        if [ -f "${GO_SUM_BACKUP}" ]; then
            mv "${GO_SUM_BACKUP}" go.sum
        fi
        echo "  Restored go.mod and go.sum"
    fi
    if [ -d "${SDK_TEMP_DIR}" ]; then
        echo "==> Cleaning up temporary SDK directory..."
        rm -rf "${SDK_TEMP_DIR}"
        echo "  Removed temporary SDK directory"
    fi
    # Exit with the original exit code
    exit $exit_code
}

# Set trap to ensure cleanup happens on exit
trap cleanup EXIT INT TERM

echo "==> Cloning SDK repository to temporary directory..."
echo "==> Temp directory: ${SDK_TEMP_DIR}"
git clone "${SDK_REPO_URL}" "${SDK_TEMP_DIR}"

echo "==> Switching SDK to dev-latest branch..."
cd "${SDK_TEMP_DIR}"

# Fetch and checkout dev-latest
git fetch origin dev-latest || git fetch origin
git checkout dev-latest

echo "==> Updating SDK module path to ${SDK_MODULE_PATH}..."

# Update go.mod module path if needed
if grep -q "^module github.com/mongodb/atlas-sdk-go" go.mod; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' "s|^module github.com/mongodb/atlas-sdk-go|module ${SDK_MODULE_PATH}|" go.mod
    else
        sed -i "s|^module github.com/mongodb/atlas-sdk-go|module ${SDK_MODULE_PATH}|" go.mod
    fi
    echo "  Updated go.mod module path"
elif grep -q "^module ${SDK_MODULE_PATH}" go.mod; then
    echo "  Module path already correct"
else
    echo "  Warning: Unexpected module path in go.mod"
fi

# Update all internal imports
echo "==> Updating internal imports..."
IMPORT_COUNT=$(grep -rl "github.com/mongodb/atlas-sdk-go" --include="*.go" . 2>/dev/null | wc -l | tr -d ' ' || echo "0")
if [ "${IMPORT_COUNT}" -gt 0 ]; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
        find . -name "*.go" -type f -exec sed -i '' "s|github.com/mongodb/atlas-sdk-go|${SDK_MODULE_PATH}|g" {} + 2>/dev/null || true
    else
        find . -name "*.go" -type f -exec sed -i "s|github.com/mongodb/atlas-sdk-go|${SDK_MODULE_PATH}|g" {} + 2>/dev/null || true
    fi
    echo "  Updated ${IMPORT_COUNT} files with new import paths"
else
    echo "  All imports already updated"
fi

cd "${CLI_ROOT}"

echo "==> Backing up go.mod and go.sum..."
cp go.mod "${GO_MOD_BACKUP}"
if [ -f go.sum ]; then
    cp go.sum "${GO_SUM_BACKUP}"
fi

echo "==> Updating go.mod with replace directive..."
# Remove existing replace directive if present
if grep -q "go.mongodb.org/atlas-sdk/${CURRENT_SDK_RELEASE} =>" go.mod; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' "/go.mongodb.org\/atlas-sdk\/${CURRENT_SDK_RELEASE} =>/d" go.mod
    else
        sed -i "/go.mongodb.org\/atlas-sdk\/${CURRENT_SDK_RELEASE} =>/d" go.mod
    fi
    echo "  Removed existing replace directive"
fi

# Append replace directive at the end of the file
echo "" >> go.mod
echo "replace (" >> go.mod
echo "	// Local SDK for testing (dev-latest branch)" >> go.mod
echo "	go.mongodb.org/atlas-sdk/${CURRENT_SDK_RELEASE} => ${SDK_TEMP_DIR}" >> go.mod
echo ")" >> go.mod
echo "  Added replace directive"

echo "==> Running go mod tidy..."
go mod tidy

echo "==> Building CLI with dev-latest SDK..."
echo "==> Current SDK commit: $(cd ${SDK_TEMP_DIR} && git rev-parse --short HEAD)"
BUILD_SUCCESS=false
if make build; then
    BUILD_SUCCESS=true
fi

if [ "$BUILD_SUCCESS" = true ]; then
    echo ""
    echo "==> Build successful! CLI built with SDK from dev-latest branch"
    echo "==> Binary available at: ./bin/atlas"
else
    echo ""
    echo "==> Build failed!"
    exit 1
fi

# Note: cleanup() will be called automatically via trap to restore go.mod
