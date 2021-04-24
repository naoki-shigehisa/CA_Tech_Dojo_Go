FROM golang:1.13.1-alpine

ENV GO111MODULE=on
ENV GOPATH=

RUN apk update && \
    apk add emacs curl git && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/jinzhu/gorm

WORKDIR /go
ADD . /go