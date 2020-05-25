#!/usr/bin/env bash

. ./notary_env.sh

set -Eeou pipefail

# --version needs to match the mongodb server version to publish to the right repo
# 4.X goes to the 4.x repo
# any *-rc version goes to testing repo
# everything else goes to development repo
./curator \
    --level debug \
    repo submit \
    --service "${barque_url}" \
    --config build/ci/repo_config.yml \
    --distro "${distro}" \
    --edition "${edition}" \
    --version "${server_version}" \
    --arch x86_64 \
    --packages "https://s3.amazonaws.com/mongodb-mongocli-build/${project}/dist/${revision}_${created_at}/${ext}.tgz"
