FROM registry.suse.com/bci/bci-base:16.0

ARG url
ARG entrypoint
ARG server_version
ARG pgp_server_version

# server-9.asc has no trailing newline, which rpm on SLES 16 rejects
RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output mongodb.asc \
    https://pgp.mongodb.com/server-${pgp_server_version}.asc; \
    echo >> mongodb.asc; \
    rpm --import mongodb.asc; \
    rm mongodb.asc

RUN zypper addrepo --gpgcheck "https://repo.mongodb.org/zypper/suse/16/mongodb-org/${server_version}/x86_64/" mongodb

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.rpm \
    ${url}; \
    zypper in -y --allow-unsigned-rpm ./${entrypoint}.rpm; \
    rm ./${entrypoint}.rpm

RUN mongosh --version
RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
