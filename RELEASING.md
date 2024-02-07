# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed and notarized
- **go_msi_snapshot:** used with `go-msi` to generate a Windows msi installer

## Stable release

Stable releases are now managed by internal tooling (PCT)

Use the instructions bellow as a fallback.

## Docker Image
We use [docker-release.yml](.github/workflows/docker-release.yml) to release the AtlasCLI Docker image.
![githubaction](https://github.com/mongodb/mongodb-atlas-cli/assets/5663078/08da2575-b10d-4469-8604-0302d557c349)
The release process runs daily and use the follow logic: 
- Step 1: Build and publish docker image to staging registry:
  In this step, we generate the AtlasCLI docker image ([Dockerfile](Dockerfile)) with 3 tags: `latest`, `vX.Y.Z` (this is the latest AtlasCLI release version, ie `v1.22.0`) and `vX.Y.Z-date` (this is the latest AtlasCLI release version with the today date, ie `v1.22.0-2024-01-01`).
  We need to release to the staging registry first to allow us to sign all the different digests generated in the [OCI index](https://github.com/opencontainers/image-spec/blob/main/image-index.md) in the Step 2.
- Step 2: Sign and Publish docker image: TO BE CONTINUED
  
  


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


# Generate the SBOM
The Software Bill of Materials (SBOM) is a description of the components that make up a software artifact.

## Atlas CLI Binary
We use `go version` to generate the SBOM for Atlas CLI binaries. You can generate the SBOM via the following command:
```bash
go version -m <path_to_atlasCLI_binary>
```

## MongoCLI Binary
We use `go version` to generate the SBOM for MongoCLI binaries. You can generate the SBOM via the following command:
```bash
go version -m <path_to_mongoCLI_binary>
```

## Atlas CLI Docker image
We use `docker sbom` to generate the SBOM for the AtlasCLI docker image. You can generate the SBOM via the following command:
```bash
docker sbom mongodb/atlas:latest
```
