FROM centos:8

ARG revision
ARG created_at
ARG mcli_version

ENV MCLI_URL=https://mongodb-mongocli-build.s3.amazonaws.com/mongocli-master/dist/${revision}_${created_at}/mongocli_${mcli_version}-next_linux_x86_64.rpm

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output mongocli.rpm \
    ${MCLI_URL}; \
    rpm -U mongocli.rpm;

RUN mongocli --version

ENTRYPOINT [ "mongocli" ]
