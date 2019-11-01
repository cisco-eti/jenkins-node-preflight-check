FROM golang:latest

ARG BINARY_NAME=gotemplate

RUN apt-get update && \
    apt-get install -y \
        git \
        openssh-server \
        gcc \
        g++ \
        supervisor

COPY . /go/src/sqbu-github.cisco.com/Nyota/gotemplate

WORKDIR /go/src/sqbu-github.cisco.com/Nyota/gotemplate

RUN sh githooks/gofmt_check

RUN go build && \
    go test

RUN rm -rf /go/src/sqbu-github.cisco.com

EXPOSE 5000

ENTRYPOINT ["/go/bin/gotemplate"]
