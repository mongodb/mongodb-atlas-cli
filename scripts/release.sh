#!/usr/bin/env bash

set -euo pipefail

VERSION="${1-}"

if [[ -z "${VERSION}" ]]; then
    echo "Please provide a version"
    exit 1
fi

if [[ "${VERSION}" == v* ]]; then
    echo "Please omit the 'v' when using this script"
    exit 1
fi

read -p "Are you sure to release v${VERSION}? " -n 1 -r
echo    # (optional) move to a new line
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    exit 1
fi

git tag -a -s "v${VERSION}" -m "v${VERSION}"
git push origin "v${VERSION}"
evergreen patch -p mongocli-master -y -d "Release ${VERSION}" -v release_publish -v release_msi -t all -f
