export GO111MODULE=on
export GOPRIVATE="wwwin-github.cisco.com"
export GONOPROXY="github.com,gopkg.in,go.uber.org"
export GOPROXY=https://${artifactory_user}:${artifactory_password}@engci-maven-master.cisco.com/artifactory/api/go/nyota-go

go get ./...
golangci-lint run