FROM registry.suse.com/suse/sle15

ARG url
ARG entrypoint

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.rpm \
    ${url}; \
    zypper in -y --allow-unsigned-rpm ./${entrypoint}.rpm; \
    rm ./${entrypoint}.rpm

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
