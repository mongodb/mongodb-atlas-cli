FROM amazonlinux:2

ARG package
ARG entrypoint
ARG server_version

RUN printf "[mongodb-org-${server_version}]\nname=MongoDB Repository\nbaseurl=https://repo.mongodb.org/yum/amazon/2/mongodb-org/${server_version}/x86_64/\ngpgcheck=1\nenabled=1\ngpgkey=https://www.mongodb.org/static/pgp/server-${server_version}.asc\n" > /etc/yum.repos.d/mongodb-org-${server_version}.repo

RUN set -eux; \
    yum install -y ${package}

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
