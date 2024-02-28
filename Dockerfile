# syntax=docker/dockerfile:1.3-labs
FROM registry.access.redhat.com/ubi8/ubi-minimal:8.9
ENV MONGODB_ATLAS_IS_CONTAINERIZED=true

COPY <<EOF /etc/yum.repos.d/mongodb-org-x86_64-6.0.repo
[mongodb-org-x86_64-6.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/8/mongodb-org/6.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-6.0.asc
EOF

COPY <<EOF /etc/yum.repos.d/mongodb-org-aarch64-6.0.repo
[mongodb-org-aarch64-6.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/8/mongodb-org/6.0/aarch64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-6.0.asc
EOF

RUN microdnf -y install jq yum &&\
    yum -y update &&\
    yum install -y mongodb-atlas &&\
    yum clean all &&\
    microdnf clean all &&\
    rm -rf /var/cache

CMD ["tail", "-f", "/dev/null"]
