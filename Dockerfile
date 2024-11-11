FROM golang:1.23.3-alpine3.20
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/command
WORKDIR /go/src/command
ADD . /go/src/command
