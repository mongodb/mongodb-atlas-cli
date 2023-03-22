FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6

COPY <<EOF /etc/yum.repos.d/mongodb-org-6.0.repo
[mongodb-org-6.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/8/mongodb-org/6.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://pgp.mongodb.com/server-6.0.asc
EOF

RUN microdnf install shadow-utils yum jq &&\
    yum -y update &&\
    yum clean all &&\
    microdnf clean all

## Usefull tools
RUN yum install -y mongodb-database-tools mongodb-mongosh 

# Create the dedicated user
RUN useradd -ms /bin/bash atlas
USER atlas

COPY ./bin /usr/local/bin

CMD echo "Invoke the atlas cli with ... mongodb-atlas-cli"
