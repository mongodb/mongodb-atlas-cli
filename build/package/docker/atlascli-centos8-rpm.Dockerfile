FROM centos:8

ARG revision
ARG created_at
ARG version

ENV URL=https://mongodb-mongocli-build.s3.amazonaws.com/mongocli-master/dist/${revision}_${created_at}/mongodb-atlas-cli_${version}-next_linux_x86_64.rpm

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output atlas.rpm \
    ${URL}; \
    rpm -U atlas.rpm;

RUN atlas --version

ENTRYPOINT [ "atlas" ]
