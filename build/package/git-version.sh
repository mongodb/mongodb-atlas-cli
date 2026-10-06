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

set -Eeou pipefail

# Single source of truth for deriving the Atlas CLI version from git tags, so
# the tag format lives in one place. Both the legacy "atlascli/vX.Y.Z" tags and
# the new "vX.Y.Z" tags are supported during the GoReleaser Pro -> OSS
# migration; the newest matching tag wins.
#
# Usage:
#   git-version.sh [version|tag|describe]
#
#   version   (default)  latest released version, e.g. 1.59.0
#   tag                  latest release tag, e.g. atlascli/v1.59.0 or v1.60.0
#   describe             version from `git describe`, including the commit
#                        distance for unreleased builds, e.g.
#                        1.59.0-24-g37a2b9bb8

tag_patterns=('atlascli/v*' 'v*')

# Strips the optional "atlascli/" prefix and the leading "v" from a tag.
version_from_tag() {
	cut -d 'v' -f 2
}

case "${1:-version}" in
	tag)
		git tag --list "${tag_patterns[@]}" --sort=-taggerdate | head -1
		;;
	describe)
		git describe --match 'atlascli/v*' --match 'v*' | version_from_tag
		;;
	version)
		git tag --list "${tag_patterns[@]}" --sort=-taggerdate | head -1 | version_from_tag
		;;
	*)
		echo "usage: $(basename "$0") [version|tag|describe]" >&2
		exit 2
		;;
esac
