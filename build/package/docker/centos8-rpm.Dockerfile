FROM artifactory.corp.mongodb.com/dockerhub/centos:8

ARG url
ARG entrypoint

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.rpm \
    ${url}; \
    rpm -U ${entrypoint}.rpm;

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
