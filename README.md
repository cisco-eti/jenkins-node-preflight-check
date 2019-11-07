# go-template for Nyota microservices
Project layout based on this [link](https://github.com/golang-standards/project-layout)

This repository provides a project structure for restful go based microservice. Example used
for go microservice is a simple http server app.

# Different directories
## /pkg
  This will contain the libraries that will be used by external applications.
  If the packages are meant to be private(not to be imported by other projects),
  place them under “/internal” and not under /pkg.
## /api
  This contains the rest api specifications in JSON/ yaml. This specifications
  would be used for api documentation.
## /test
  This contains all the tests and dependent mock functions.
## /deployments
  This contains system and container orchestration, deployment configurations and templates.
## /githooks
  This will have git hooks e.g. go_fmt, staticcheck etc.
## app.go
  This is the main package for the microservice.

# Instructions for Developers

## Install git hook to autoformat and run tests

From the main directory, run `ln -s $(pwd)/githooks/pre-commit .git/hooks/pre-commit`

## Build and Run
```
go get -v ./...

go build

./gotemplate-microservice
```

## Run tests

go test -v ./...

## Build and Run Docker

./build-docker.sh

docker run -it -p 5000:5000 nyota-gotemplate

##
