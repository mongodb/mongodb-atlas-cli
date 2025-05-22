#!/bin/bash
# Copyright 2025 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -eu

release_date=${DATE:-$(date -u '+%Y-%m-%d')}

export DATE="${release_date}"

if [ -z "${AUTHOR:-}" ]; then
  AUTHOR=$(git config user.name)
fi

if [ -z "${VERSION:-}" ]; then
  VERSION=$(git tag --list 'atlascli/v*' --sort=-taggerdate | head -1 | cut -d 'v' -f 2)
fi

export AUTHOR
export VERSION

echo "Generating SSDLC checklist for AtlasCLI version ${VERSION}, author ${AUTHOR} and release date ${DATE}..."

# Ensure AtlasCLI version directory exists
mkdir -p "compliance/v${VERSION}"

envsubst < docs/releases/ssdlc-compliance.template.md \
  > "compliance/v${VERSION}/ssdlc-compliance-${VERSION}.md"

echo "SDLC checklist ready. Files in compliance/v${VERSION}/:"
ls -l "compliance/v${VERSION}/"

echo "Printing the generated report:"
cat "compliance/v${VERSION}/ssdlc-compliance-${VERSION}.md"
