FROM centos:8

ARG url
ARG entrypoint
ARG server_version

RUN rm -rf /etc/yum.repos.d/*

RUN printf "[mongodb-org-${server_version}]\nname=MongoDB Repository\nbaseurl=https://repo.mongodb.org/yum/redhat/\$releasever/mongodb-org/${server_version}/x86_64/\ngpgcheck=1\nenabled=1\ngpgkey=https://pgp.mongodb.com/server-${server_version}.asc\n" > /etc/yum.repos.d/mongodb-org-${server_version}.repo

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.rpm \
    ${url}; \
    yum install -y ./${entrypoint}.rpm; \
    rm ./${entrypoint}.rpm

RUN mongosh --version
RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
