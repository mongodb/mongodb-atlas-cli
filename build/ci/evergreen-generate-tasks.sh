#!/usr/bin/env bash

# Copyright 2022 MongoDB Inc
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

args=(-tasks "${task_type:?}")
if [[ -n "${e2e_variant:-}" ]]; then
  args+=(-variant "${e2e_variant}")
fi
if [[ -n "${e2e_task_prefix:-}" ]]; then
  args+=(-task-prefix "${e2e_task_prefix}")
fi
if [[ -n "${e2e_include_groups:-}" ]]; then
  args+=(-include-groups "${e2e_include_groups}")
fi
if [[ -n "${e2e_exclude_groups:-}" ]]; then
  args+=(-exclude-groups "${e2e_exclude_groups}")
fi
if [[ -n "${e2e_exclude:-}" ]]; then
  args+=(-exclude "${e2e_exclude}")
fi
if [[ -n "${e2e_profile:-}" ]]; then
  args+=(-profile "${e2e_profile}")
fi

go run tools/cmd/genevergreen/main.go "${args[@]}" >tasks.json
