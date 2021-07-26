FROM debian:9-slim

ARG revision
ARG created_at
ARG mcli_version

RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends \
	  ca-certificates \
		curl \
	; \
	if ! command -v ps > /dev/null; then \
		apt-get install -y --no-install-recommends procps; \
	fi; \
	rm -rf /var/lib/apt/lists/*

ENV MCLI_URL=https://mongodb-mongocli-build.s3.amazonaws.com/mongocli-master/dist/${revision}_${created_at}/mongocli_${mcli_version}-next_linux_x86_64.deb

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output mongocli.deb \
    ${MCLI_URL}; \
    dpkg -i mongocli.deb;

RUN mongocli --version

ENTRYPOINT [ "mongocli" ]
