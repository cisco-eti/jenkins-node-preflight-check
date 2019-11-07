FROM golang:latest

ARG BINARY_NAME=go-template
ARG SRC_DIR=$GOPATH/src/sqbu-github.cisco.com/Nyota/$BINARY_NAME

RUN apt-get update && \
    apt-get install -y \
        git \
        openssh-server \
        gcc \
        g++ \
        supervisor

COPY . $SRC_DIR

WORKDIR $SRC_DIR

RUN sh $SRC_DIR/githooks/gofmt_check

RUN go build && \
    go test

#RUN rm -rf /go/src/sqbu-github.cisco.com

EXPOSE 5000

ENTRYPOINT ["./go-template"]
