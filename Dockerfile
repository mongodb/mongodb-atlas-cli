# syntax=docker/dockerfile:1.3-labs
#
# On M1 macOS, `docker buildx build --platform=linux/amd64 .`
#

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6

COPY <<EOF /etc/yum.repos.d/mongodb-org-6.0.repo
[mongodb-org-6.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/8/mongodb-org/6.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://pgp.mongodb.com/server-6.0.asc
EOF

RUN microdnf install yum &&\
    yum -y update &&\
    yum clean all &&\
    microdnf clean all

RUN yum install -y mongodb-atlas

CMD echo "Invoke the atlas cli with ... mongodb-atlas-cli"
