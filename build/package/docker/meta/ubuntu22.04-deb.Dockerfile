FROM ubuntu:22.04

ARG url
ARG entrypoint
ARG server_version

RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		gnupg \
	; \
	if ! command -v ps > /dev/null; then \
		apt-get install -y --no-install-recommends procps; \
	fi; \
	curl -L https://www.mongodb.org/static/pgp/server-${server_version}.asc | apt-key add -; \
	echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/${server_version} multiverse" | tee /etc/apt/sources.list.d/mongodb-org-${server_version}.list; \
	rm -rf /var/lib/apt/lists/*

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.deb \
    ${url}; \
	apt-get update; \
	apt-get install -y ./${entrypoint}.deb; \
	rm -rf /var/lib/apt/lists/* ./${entrypoint}.deb

RUN mongosh --version
RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
