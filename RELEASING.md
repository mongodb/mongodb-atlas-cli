# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed and notarized
- **go_msi_snapshot:** used with `go-msi` to generate a Windows msi installer

## Stable release

Stable releases are now managed by internal tooling (PCT)

Use the instructions bellow as a fallback.

## Package Managers

Package Managers are published after a stable release happens, in which binaries are stored in github releases and also uploaded to our download center (https://www.mongodb.com/try/download/atlascli.

* [Chocolatey](http://chocolatey.org) release is triggered in https://github.com/mongodb-forks/chocolatey-packages/, the Github Action will trigger every weekday at 4pm (UTC) to check if there are any new releases in https://github.com/mongodb/mongodb-atlas-cli/releases/.

* [Homebrew](http://brew.sh/) release is triggered in https://github.com/Homebrew/homebrew-core/, which is not maintained by MongoDB rather by homebrew community.

* Yum and Apt are handled internally via evergreen tasks `push_stable_atlascli_generate` and `push_stable_mongocli_generate`.

## Docker Image
Our Docker image release for AtlasCLI is managed through the [docker-release.yml](.github/workflows/docker-release.yml)  workflow. This process is automated to run daily, ensuring the latest versions of the image dependencies are updated.
![github_action](https://github.com/mongodb/mongodb-atlas-cli/assets/5663078/fd54ccda-7794-4139-af92-dbde0c278e78)
### Release Steps
#### Step 1: Build and Stage
The AtlasCLI Docker image is built from the ([Dockerfile](Dockerfile)) and tagged in three ways: `latest`, `vX.Y.Z` (reflecting the latest release version, e.g., `v1.22.0`), and `vX.Y.Z-date` (adding the current date, e.g., `v1.22.0-2024-01-01`). This image is initially published to a staging registry to prepare for signature in the next step.

#### Step 2: Sign and Publish
We retrieve the image from the staging registry and use its [OCI index](https://github.com/opencontainers/image-spec/blob/main/image-index.md) to identify the three relevant digests. Each digest is signed using [cosign](https://github.com/sigstore/cosign), and the corresponding signature is stored in the MongoDB cosign repository. The signed image is then pushed to the public repository.

#### Step 3: Verify Signature
The final step involves verifying the Docker image's signature to confirm its authenticity.



## Deprecated
_**Note:** This action will only publish a release for [maintainers of the cli](https://github.com/orgs/mongodb/teams/mongocli)_

To manually generate a new stable release you can run:


```bash
./scripts/release.sh atlascli/v1.0.0
```

**Note:** Please use the `atlascli/vX.Y.Z` format for the version to release 

This will do the following things:
1. Tag a new version, ie: `git tag -a -s atlascli/v1.0.0 -m "atlascli/v1.0.0"`
2. Publish the new tag, ie `git push origin atlascli/v1.0.0`
3. The [evergreen](build/ci/release.yml) release task will run after a tag event from master.
4. If everything goes smoothly the release will be published in the [releases page](https://github.com/mongodb/mongodb-atlas-cli/releases), and [download center](https://www.mongodb.com/try/download/atlascli).


# Generate the SBOM
The Software Bill of Materials (SBOM) is a description of the components that make up a software artifact.

## Atlas CLI Binary
We use `go version` to generate the SBOM for Atlas CLI binaries. You can generate the SBOM via the following command:
```bash
go version -m <path_to_atlasCLI_binary>
```

## Atlas CLI Docker image
We use `docker sbom` to generate the SBOM for the Atlas CLI docker image. You can generate the SBOM via the following command:
```bash
docker sbom mongodb/atlas:latest
```
