# syntax=docker/dockerfile:1.3-labs
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.8
ENV MONGODB_ATLAS_IS_CONTAINERIZED=true

COPY <<EOF /etc/yum.repos.d/mongodb-org-x86_64-8.0.repo
[mongodb-org-x86_64-8.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/9/mongodb-org/8.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-8.0.asc
EOF

COPY <<EOF /etc/yum.repos.d/mongodb-org-aarch64-8.0.repo
[mongodb-org-aarch64-8.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/9/mongodb-org/8.0/aarch64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-8.0.asc
EOF

RUN microdnf -y install jq &&\
    microdnf -y update &&\
    microdnf -y install mongodb-atlas &&\
    microdnf clean all &&\
    rm -rf /var/cache &&\
    rpm -e --nodeps \
      curl-minimal libcurl-minimal libxml2 \
      glib2 libarchive libgcrypt systemd-libs libsolv libnghttp2 \
      libblkid libmount libsmartcols libuuid \
      gnupg2 krb5-libs openldap sed libyaml &&\
    ! command -v curl &&\
    ! command -v yum

CMD ["tail", "-f", "/dev/null"]
