FROM debian:12-slim

ARG package
ARG entrypoint
ARG server_version
ARG mongo_package
ARG mongo_repo

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
	echo "deb [ arch=amd64,arm64 ] ${mongo_repo}/apt/debian bookworm/${mongo_package}/${server_version} main" | tee /etc/apt/sources.list.d/${mongo_package}-${server_version}.list; \
	apt-get update; \
	apt-get install -y --no-install-recommends ${package}; \
	rm -rf /var/lib/apt/lists/*

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
