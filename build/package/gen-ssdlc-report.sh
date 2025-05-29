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

if [ "${AUGMENTED_REPORT}" = "true" ]; then
  target_dir="."
  file_name="ssdlc-compliance-${VERSION}-${DATE}.md"
  SBOM_TEXT="  - See Augmented SBOM manifests (CycloneDX in JSON format):
      - This file has been provided along with this report under the name 'linux_amd64_augmented_sbom_v${VERSION}.json'
      - Please note that this file was generated on ${DATE} and may not reflect the latest security information of all third party dependencies."

else # If not augmented, generate the standard report
  target_dir="compliance/v${VERSION}"
  file_name="ssdlc-compliance-${VERSION}.md"
  SBOM_TEXT="  - See SBOM Lite manifests (CycloneDX in JSON format):
      - https://github.com/mongodb/mongodb-atlas-cli/releases/download/atlascli%2Fv${VERSION}/sbom.json"
  # Ensure AtlasCLI version directory exists
  mkdir -p "${target_dir}"
fi

export AUTHOR
export VERSION
export SBOM_TEXT

echo "Generating SSDLC checklist for AtlasCLI version ${VERSION}, author ${AUTHOR} and release date ${DATE}..."

envsubst < docs/releases/ssdlc-compliance.template.md \
  > "${target_dir}/${file_name}"

echo "SDLC checklist ready. Files in ${target_dir}/:"
ls -l "${target_dir}/"

echo "Printing the generated report:"
cat "${target_dir}/${file_name}"
