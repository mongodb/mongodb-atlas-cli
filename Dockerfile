# Build stage: compile atlas CLI from source with current Go toolchain
FROM golang:1.26 AS builder

ARG TARGETARCH
ARG GIT_SHA=""
ARG ATLAS_VERSION=""

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build \
    -trimpath -mod=readonly \
    -ldflags "-s -w \
      -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version.GitCommit=${GIT_SHA} \
      -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version.Version=${ATLAS_VERSION}" \
    -o /atlas ./cmd/atlas

# Runtime stage
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.7

ENV MONGODB_ATLAS_IS_CONTAINERIZED=true

RUN microdnf -y install jq && \
    microdnf clean all && \
    rm -rf /var/cache

COPY --from=builder /atlas /usr/bin/atlas

CMD ["tail", "-f", "/dev/null"]
