# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed and notarized
- **go_msi_snapshot:** used with `go-msi` to generate a Windows msi installer

## Stable release

To generate a new stable release you can run:

```bash
./scripts/release.sh 1.0.0
```

**Note:** Please omit the `v` from the version to release 

This will do the following things:
1. Tag a new version, ie: `git tag -a -s v1.0.0 -m "v1.0.0"`
2. Publish the new tag, ie `git push origin v1.0.0`
3. The evergreen releasing tasks will run as part of the tag process.
4. The release will be set as draft in the [releases page](https://github.com/mongodb/mongocli/releases), review the release and if it looks ok, publish it.
