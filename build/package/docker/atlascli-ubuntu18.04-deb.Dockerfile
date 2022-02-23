FROM ubuntu:18.04

ARG revision
ARG created_at
ARG version

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

ENV URL=https://mongodb-mongocli-build.s3.amazonaws.com/mongocli-master/dist/${revision}_${created_at}/mongodb-atlas-cli_${version}-next_linux_x86_64.deb

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output atlas.deb \
    ${URL}; \
    dpkg -i atlas.deb;

RUN atlas --version

ENTRYPOINT [ "atlas" ]
