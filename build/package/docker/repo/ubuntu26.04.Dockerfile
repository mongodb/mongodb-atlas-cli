FROM 901841024863.dkr.ecr.us-east-1.amazonaws.com/dockerhub/library/ubuntu:26.04

ARG package
ARG entrypoint
ARG server_version
ARG pgp_server_version
ARG mongo_package
ARG mongo_repo

RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		gnupg \
		apt-transport-https \
	; \
	if ! command -v ps > /dev/null; then \
		apt-get install -y --no-install-recommends procps; \
	fi; \
	install -d -m 0755 /usr/share/keyrings; \
	curl -L https://www.mongodb.org/static/pgp/server-${pgp_server_version}.asc \
		| gpg --dearmor -o /usr/share/keyrings/mongodb-server-${pgp_server_version}.gpg; \
	echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-${pgp_server_version}.gpg ] ${mongo_repo}/apt/ubuntu resolute/${mongo_package}/${server_version} multiverse" | tee /etc/apt/sources.list.d/${mongo_package}-${server_version}.list; \
	apt-get update; \
	apt-get install -y --no-install-recommends ${package}; \
	rm -rf /var/lib/apt/lists/*

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
