# Releasing

## Daily snapshots

We generate a snapshot build via the `goreleaser_snaphot` and `go_msi_snapshot` evergreen tasks,
these tasks run on master and can be patched at any time.

- **goreleaser_snaphot:** used with `goreleaser` to generate linux, mac and windows builds, the mac build will also be signed and notarized
- **go_msi_snapshot:** used with `go-msi` to generate a Windows msi installer

## Stable release

Stable releases are managed by the [Release](.github/workflows/release.yaml) GitHub Actions
workflow.

To cut a release:

1. Go to **Actions → Release → Run workflow**, pick the `master` branch, and enter the
   version to release as `X.Y.Z` (without the leading `v`).
2. The workflow:
   - verifies the required checks are green on the release commit — the required PR checks
     on `master` plus a few release-specific Evergreen variants,
   - performs the Jira bookkeeping PCT used to do: creates the `atlascli-<version>` fix
     version (renaming the rolling `next-atlascli-release` and recreating it), creates the
     CLOUDP release ticket, moves any still-open tickets back to `next-atlascli-release`,
     and opens and links the DOCSP release-notes ticket,
   - pushes the `vX.Y.Z` tag, which triggers the [evergreen](build/ci/release.yml) release
     pipeline to build, sign and publish.
3. Once the GitHub release is published, [close-release.yaml](.github/workflows/close-release.yaml)
   marks the fix version released and resolves the release ticket.

> [!NOTE]
> The workflow is idempotent: re-running it reuses an existing release ticket and release
> notes ticket rather than creating duplicates.

## Package Managers

Package Managers are published after a stable release happens, in which binaries are stored in 
GitHub releases and also uploaded to our download center (https://www.mongodb.com/try/download/atlascli.

* [Chocolatey](http://chocolatey.org) release is triggered in https://github.com/mongodb-forks/chocolatey-packages/, 
  the GitHub Action will trigger every weekday at 4pm (UTC) to check if there are any new releases 
  in https://github.com/mongodb/mongodb-atlas-cli/releases/.

* [Homebrew](http://brew.sh/) release is triggered in https://github.com/Homebrew/homebrew-core/, 
  which is maintained by the homebrew community, not MongoDB.

* Yum and Apt are handled internally via evergreen tasks `push_stable_atlascli_generate` and `push_stable_mongocli_generate`.

## Docker Image

Our Docker image release for AtlasCLI is managed through the [docker-release.yml](.github/workflows/docker-release.yml)  workflow. This process is automated to run daily, ensuring the latest versions of the image dependencies are updated.
![github_action](https://github.com/mongodb/mongodb-atlas-cli/assets/5663078/fd54ccda-7794-4139-af92-dbde0c278e78)

### Release Steps

#### Step 1: SBOM Generation

An SBOM Lite is generated using the maintained purls.txt file and uploaded to Kundukto. The SBOM Lite is also included as an artifact in the public repository release.

#### Step 2: Build and Stage

The AtlasCLI Docker image is built from the ([Dockerfile](Dockerfile)) and tagged in three ways:
`latest`, `vX.Y.Z` (reflecting the latest release version, e.g., `v1.22.0`), and `vX.Y.Z-date` 
(adding the current date, e.g., `v1.22.0-2024-01-01`). 
This image is initially published to a staging registry to prepare for signature in the next step.

#### Step 3: Sign and Publish

We retrieve the image from the staging registry and use its [OCI index](https://github.com/opencontainers/image-spec/blob/main/image-index.md) to identify the three
relevant digests. Each digest is signed using [cosign](https://github.com/sigstore/cosign), and the corresponding signature is stored 
in the MongoDB cosign repository. The signed image is then pushed to the public repository.

#### Step 4: Verify Signature

The Docker image's signature is verified to confirm its authenticity.

#### Step 5: Trace Artifacts

Papertrail runs to record detailed metadata about the release for improved traceability and accountability.

#### Step 6: Compliance Reporting

A GitHub workflow generates a compliance report after the release and opens a PR with the report.
