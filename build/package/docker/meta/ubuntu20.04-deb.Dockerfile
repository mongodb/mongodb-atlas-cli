FROM ubuntu:20.04

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
	curl -O https://www.mongodb.org/static/pgp/server-${server_version}.asc; \
	apt-key add server-${server_version}.asc; \
	rm server-${server_version}.asc; \
	echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/${server_version} multiverse" | tee /etc/apt/sources.list.d/mongodb-org-${server_version}.list

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.deb \
    ${url}; \
	apt-get update; \
	apt-get install ./${entrypoint}.deb; \
	rm -rf /var/lib/apt/lists/*

RUN mongosh --version
RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
