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

RUN go get github.com/stretchr/testify
RUN go mod download

RUN sh $SRC_DIR/githooks/gofmt_check

RUN latest=$(curl -s https://api.github.com/repos/dominikh/go-tools/releases/latest) && \
tagname=$(echo "$latest" | grep "tag_name" | cut -d':' -f2 | cut -d'"' -f2) && \
[ "$(uname)" = Darwin ] && system=darwin || system=linux; \
echo tagname == $tagname system == $system && \
curl -L https://github.com/dominikh/go-tools/releases/download/${tagname}/staticcheck_${system}_amd64.tar.gz > /tmp/staticcheck_${system}_amd64.tar.gz && \
ls -l /tmp/staticcheck_${system}_amd64.tar.gz && \
tar xzf /tmp/staticcheck_${system}_amd64.tar.gz && \
cd staticcheck && \
cp staticcheck /usr/local/bin/ && \
chmod +x /usr/local/bin/staticcheck && \
rm -rf /tmp/staticcheck_*

RUN sh $SRC_DIR/githooks/staticcheck_gotemplate

RUN go install && \
    go test ./...

WORKDIR $GOPATH

RUN rm -rf /go/src/sqbu-github.cisco.com

EXPOSE 5000

ENTRYPOINT ["./bin/go-template"]
