FROM 901841024863.dkr.ecr.us-east-1.amazonaws.com/dockerhub/library/debian:12-slim

ARG url
ARG entrypoint

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

RUN set -eux; \
    curl --silent --show-error --fail --location --retry 3 \
    --output ${entrypoint}.deb \
    ${url}; \
    dpkg -i ${entrypoint}.deb;

RUN ${entrypoint} --version

ENV ENTRY=${entrypoint}

ENTRYPOINT $ENTRY
