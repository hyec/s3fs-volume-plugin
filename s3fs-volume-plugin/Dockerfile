FROM ubuntu:20.04 as base
ARG TINI_VERSION="v0.19.0"
# armhf
ARG TPLATFORM="amd64"
ENV DEBIAN_FRONTEND="noninteractive"

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-${TPLATFORM} /tini
RUN chmod +x /tini

RUN apt update && \
    apt dist-upgrade -y && \
    apt install -y s3fs curl && \
    apt clean && rm -rf /var/lib/apt/lists/*

FROM base as dev
ARG GO_VERSION="1.15.2"
# armv6l
ARG GPLATFORM="amd64"

RUN curl --silent -L https://dl.google.com/go/go${GO_VERSION}.linux-${GPLATFORM}.tar.gz | tar -C /usr/local -zxf - && \
    apt update && apt install -y git gcc && apt clean && rm -rf /var/lib/apt/lists/*

COPY s3fs-volume-plugin/ /root/go/src/github.com/marcelo-ochoa/docker-volume-plugins/s3fs-volume-plugin
COPY mounted-volume/ /root/go/src/github.com/marcelo-ochoa/docker-volume-plugins/mounted-volume

RUN cd /root/go/src/github.com/marcelo-ochoa/docker-volume-plugins/s3fs-volume-plugin && \
    /usr/local/go/bin/go get

FROM base

COPY --from=dev /root/go/bin/s3fs-volume-plugin /
