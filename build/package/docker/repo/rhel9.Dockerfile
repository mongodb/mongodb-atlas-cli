FROM registry.access.redhat.com/ubi9/ubi

ARG package
ARG entrypoint
ARG server_version
ARG mongo_package
ARG mongo_repo

RUN rm -rf /etc/yum.repos.d/*

RUN printf "[${mongo_package}-${server_version}]\nname=MongoDB Repository\nbaseurl=${mongo_repo}/yum/redhat/\$releasever/${mongo_package}/${server_version}/\$basearch/\ngpgcheck=1\nenabled=1\ngpgkey=https://www.mongodb.org/static/pgp/server-${server_version}.asc\n" > /etc/yum.repos.d/${mongo_package}-${server_version}.repo

RUN set -eux; \
    yum install -y ${package}

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
