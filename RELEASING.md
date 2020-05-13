# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks are run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed an notarized
- **go_msi_snapshot:** used with `go-msi` to generate a Windows msi installer

## Stable release

To generate a new stable release you must follow the following steps:

1. Tag a new version, ie: `git tag -a -s v1.0.0 -m "v1.0.0"`
2. Publish the new tag, ie `git push origin v1.0.0`
3. Run the evergreen releasing task, ie: `evergreen patch -p mongocli-master -y -d "Release v1.0.0" -v release -t release`
