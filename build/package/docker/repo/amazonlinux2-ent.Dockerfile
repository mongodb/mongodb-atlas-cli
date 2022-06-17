FROM amazonlinux:2

ARG package
ARG entrypoint
ARG server_version

RUN printf "[mongodb-enterprise-${server_version}]\nname=MongoDB Enterprise Repository\nbaseurl=https://repo.mongodb.com/yum/amazon/2/mongodb-enterprise/${server_version}/\$basearch/\ngpgcheck=1\nenabled=1\ngpgkey=https://www.mongodb.org/static/pgp/server-${server_version}.asc\n" > /etc/yum.repos.d/mongodb-enterprise-${server_version}.repo

RUN set -eux; \
    yum install -y ${package}

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
