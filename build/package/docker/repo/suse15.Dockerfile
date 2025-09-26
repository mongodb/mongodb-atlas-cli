FROM registry.suse.com/suse/sle15

ARG package
ARG entrypoint
ARG server_version
ARG pgp_server_version
ARG mongo_package
ARG mongo_repo

RUN rpm --import https://pgp.mongodb.com/server-${server_version}.asc

RUN zypper addrepo --gpgcheck "${mongo_repo}/zypper/suse/15/${mongo_package}/${server_version}/x86_64/" mongodb

RUN set -eux; \
    zypper in -y ${package}

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
