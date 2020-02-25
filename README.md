# go-template for Nyota microservices
Project layout based on this [link](https://github.com/golang-standards/project-layout)

This repository provides a project structure for restful go based microservice. Example used
for go microservice is a simple http server app.

# Jenkins Pipeline
[Jenkins](https://engci-private-rcdn.cisco.com/jenkins/bsft-jenkins1/job/FlowerBed/job/WFL/job/go-template/)

# Additional Setup Instructions for Lab VM before build

Skip this step if is not building on Lab VM (i.e. rcdn6-vmXX-YYY).

export DOCKER_BLD_ARGS="https_proxy=http://proxy.esl.cisco.com:80"

# Build and Run Docker

./build-docker.sh

docker run -it -p 5000:5000 nyota-gotemplate

# Instructions for Developers

## Install git hook to autoformat and run tests

From the main directory, run `ln -s $(pwd)/githooks/pre-commit .git/hooks/pre-commit`

## Different directories

## /src/pkg
  This will contain the libraries internal to the app

## /docs
  This contains the rest api specifications in JSON/ yaml. This specifications
  would be used for api documentation. Generated from handler comments using command:
```bash 
swag init
```

## /test
  This contains all the tests and dependent mock functions.

## /deployments
  This contains system and container orchestration, deployment configurations and templates.

## /githooks
  This will have git hooks e.g. go_fmt, staticcheck etc.

## main.go
  This is the main package for the microservice.


