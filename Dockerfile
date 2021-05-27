FROM docker.io/bitnami/minideb:buster
LABEL maintainer "MongoDB <containers@mongodb.com>"

ENV HOME="/" \
    OS_ARCH="amd64" \
    OS_FLAVOUR="debian-10" \
    OS_NAME="linux"

# Install required system packages and dependencies
RUN install_packages ca-certificates curl gzip jq procps tar wget
RUN MCLI_TAG=$(curl -sL --header "Accept: application/json" https://github.com/mongodb/mongocli/releases/latest | jq -r '.["tag_name"]') && \
    MCLI_VERSION=$(echo $MCLI_TAG | cut -dv -f2) && \
    MCLI_DEB="mongocli_${MCLI_VERSION}_linux_x86_64.deb" && \
    curl -OL https://github.com/mongodb/mongocli/releases/download/${MCLI_TAG}/${MCLI_DEB} && \
    echo "About to install mongocli from: ${MCLI_DEB}" && \
    dpkg -i ${MCLI_DEB}

USER 1001
ENTRYPOINT [ "mongocli" ]
CMD [ "--help" ]
