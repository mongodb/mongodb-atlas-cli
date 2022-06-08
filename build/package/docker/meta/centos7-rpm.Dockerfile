FROM centos:7

ARG url
ARG entrypoint
ARG server_version

RUN echo $'[mongodb-org-${server_version}] \n\
name=MongoDB Repository \n\
baseurl=https://repo.mongodb.org/yum/redhat/\$releasever/mongodb-org/${server_version}/x86_64/ \n\
gpgcheck=1 \n\
enabled=1 \n\
gpgkey=https://pgp.mongodb.com/server-${server_version}.asc \n\
' > /etc/yum.repos.d/mongodb-org-${server_version}.repo

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.rpm \
    ${url}; \
    yum install -y ./${entrypoint}.rpm; \
    rm ./${entrypoint}.rpm

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
