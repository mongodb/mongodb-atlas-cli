# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed and notarized
- **go_msi_snapshot:** used with `go-msi` to generate a Windows msi installer

## Stable release

Stable releases are now managed by internal tooling (PCT)

Use the instructions bellow as a fallback.

### Deprecated
_**Note:** This action will only publish a release for [maintainers of the cli](https://github.com/orgs/mongodb/teams/mongocli)_

To manually generate a new stable release you can run:


```bash
./scripts/release.sh atlascli/v1.0.0
```

**Note:** Please use the `atlascli/vX.Y.Z` or `mongocli/vX.Y.Z` format for the version to release 

This will do the following things:
1. Tag a new version, ie: `git tag -a -s atlascli/v1.0.0 -m "atlascli/v1.0.0"`
2. Publish the new tag, ie `git push origin atlascli/v1.0.0`
3. The [evergreen](build/ci/release.yml) release task will run after a tag event from master.
4. If everything goes smoothly the release will be published in the [releases page](https://github.com/mongodb/mongodb-atlas-cli/releases), and [download center](https://www.mongodb.com/try/download/mongocli).
