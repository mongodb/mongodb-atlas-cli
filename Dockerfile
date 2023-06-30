FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6

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

RUN microdnf install shadow-utils jq yum yum-utils &&\
    yum -y update &&\
    yum clean all &&\
    microdnf clean all

RUN yum install -y mongodb-atlas

ENTRYPOINT ["/bin/bash"]