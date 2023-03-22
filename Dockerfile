FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6

RUN microdnf install shadow-utils jq yum yum-utils &&\
    yum -y update &&\
    yum clean all &&\
    microdnf clean all

## Gives users ability to install other tools based of that image
RUN yum-config-manager --add-repo https://repo.mongodb.org/yum/redhat/8/mongodb-org/6.0/x86_64/

# Create the dedicated user
RUN useradd -ms /bin/bash atlas
USER atlas

COPY ./bin /usr/local/bin

CMD echo "Usage: invoke image with 'atlas'" 
