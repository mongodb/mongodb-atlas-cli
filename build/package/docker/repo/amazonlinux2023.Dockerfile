FROM amazonlinux:2023

ARG package
ARG entrypoint
ARG server_version
ARG mongo_package
ARG mongo_repo

RUN printf "[${mongo_package}-${server_version}]\nname=MongoDB Repository\nbaseurl=${mongo_repo}/yum/amazon/2023/${mongo_package}/${server_version}/\$basearch/\ngpgcheck=1\nenabled=1\ngpgkey=https://www.mongodb.org/static/pgp/server-${server_version}.asc\n" > /etc/yum.repos.d/${mongo_package}-${server_version}.repo

RUN set -eux; \
    yum install -y ${package}

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
