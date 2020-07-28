# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks are run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed an notarized
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
3. Run the evergreen releasing task, ie: `evergreen patch -p mongocli-master -y -d "Release v1.0.0" -v release_publish -v release_msi -t all -f`
4. Open a PR to update `server_version` in [build/ci/evergreen.yml](build/ci/evergreen.yml). The number does not matter as long as it's higher than the previous one.
