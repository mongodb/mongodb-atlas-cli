FROM registry.suse.com/suse/sle15

ARG url
ARG entrypoint
ARG server_version
ARG pgp_server_version

RUN rpm --import https://pgp.mongodb.com/server-${pgp_server_version}.asc

RUN zypper addrepo --gpgcheck "https://repo.mongodb.org/zypper/suse/15/mongodb-org/${server_version}/x86_64/" mongodb

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
