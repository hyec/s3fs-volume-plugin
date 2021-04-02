ARG UBUNTU_VERSION
FROM --platform=${TARGETPLATFORM:-linux/amd64} ubuntu:${UBUNTU_VERSION:-20.04} as base

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt update && \
    apt install -y cifs-utils curl tini && \
    apt clean && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION
FROM --platform=${TARGETPLATFORM:-linux/amd64} golang:${GO_VERSION:-1.15.10}-alpine as dev

RUN apk update && apk add git

COPY cifs-volume-plugin/ /go/src/github.com/marcelo-ochoa/docker-volume-plugins/cifs-volume-plugin
COPY mounted-volume/ /go/src/github.com/marcelo-ochoa/docker-volume-plugins/mounted-volume

RUN cd /go/src/github.com/marcelo-ochoa/docker-volume-plugins/cifs-volume-plugin && \
    CGO_ENABLED=0 GOOS=linux go get

FROM base

COPY --from=dev /go/bin/cifs-volume-plugin /
