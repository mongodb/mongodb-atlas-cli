SSDLC Compliance Report: Atlas CLI ${VERSION}
=================================================================

- Release Creators: ${AUTHORS}
- Created On:       ${DATE}

Overview:

- **Product and Release Name**

    - Atlas CLI v${VERSION}, ${DATE}.

- **Process Document**
  - http://go/how-we-develop-software-doc

- **Tool used to track third party vulnerabilities**
  - [Kondukto](https://arcticglow.kondukto.io/)

- **Dependency Information**
  - See SBOMS Lite manifests (CycloneDX in JSON format) for `Intel` and `ARM` are to be found [here](.)
  - See [instructions on how the SBOMs are generated or how to generate them manually](../../dev/image-sboms.md)

- **Static Analysis Report**
  - No SAST findings. Our CI system blocks merges on any SAST findings.${IGNORED_VULNERABILITIES}

- **Release Signature Report**
  - Image signatures enforced by CI pipeline.
  - See [Signature verification instructions here](../../dev/signed-images.md)
  - Self-verification shortcut:
    ```shell
    make verify IMG=mongodb/mongodb-atlas-cli:${VERSION} SIGNATURE_REPO=mongodb/signatures
    ```

- **Security Testing Report**
  - Available as needed from Cloud Security.

- **Security Assessment Report**
  - Available as needed from Cloud Security.

Assumptions and attestations:

- Internal processes are used to ensure CVEs are identified and mitigated within SLAs.

- All Operator images are signed by MongoDB, with signatures stored at `docker.io/mongodb/signatures`.