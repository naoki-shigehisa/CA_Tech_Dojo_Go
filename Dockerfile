FROM golang:1.13.1-alpine

ENV GO111MODULE=on

RUN apk update && \
    apk add emacs curl git && \
    go get github.com/pilu/fresh && \
    go get github.com/gonum/floats && \
    go get github.com/gonum/matrix

WORKDIR /app
ADD . /app