FROM centos:8

ARG url
ARG entrypoint
ARG server_version

RUN cat <<EOF > /etc/yum.repos.d/mongodb-org-${server_version}.repo
[mongodb-org-${server_version}]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/\$releasever/mongodb-org/${server_version}/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://pgp.mongodb.com/server-${server_version}.asc
EOF

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.rpm \
    ${url}; \
    yum install -y ./${entrypoint}.rpm; \
    rm ./${entrypoint}.rpm

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
