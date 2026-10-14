FROM registry.suse.com/bci/bci-base:16.0

ARG package
ARG entrypoint
ARG server_version
ARG pgp_server_version
ARG mongo_package
ARG mongo_repo

# server-9.asc has no trailing newline, which rpm on SLES 16 rejects
RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output mongodb.asc \
    https://pgp.mongodb.com/server-${pgp_server_version}.asc; \
    echo >> mongodb.asc; \
    rpm --import mongodb.asc; \
    rm mongodb.asc

RUN zypper addrepo --gpgcheck "${mongo_repo}/zypper/suse/16/${mongo_package}/${server_version}/x86_64/" mongodb

RUN set -eux; \
    zypper in -y ${package}

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
