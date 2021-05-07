FROM golang:1.13.1-alpine

ENV GO111MODULE=on
ENV GOPATH=

RUN apk update && \
    apk add emacs curl git

WORKDIR /go
ADD . /go