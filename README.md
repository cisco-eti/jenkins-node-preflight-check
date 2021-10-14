# ETI sre-go-helloworld

[![Build Status](https://engci-private-sjc.cisco.com/jenkins/eti-sre/buildStatus/icon?job=SRE%2FProjects%2Fsre-go-helloworld%2Fbuild%2Fsre-go-helloworld%2Fmaster)](https://engci-private-sjc.cisco.com/jenkins/eti-sre/job/SRE/job/Projects/job/sre-go-helloworld/job/build/job/sre-go-helloworld/job/master/)

[API Try-out Swagger UI Docs](https://wwwin-github.cisco.com/pages/eti/sre-go-helloworld)

Project layout based on this [link](https://github.com/golang-standards/project-layout)

This repository provides a project structure for restful go based microservice. Example used
for go microservice is a simple http server app.

## Jenkins Pipeline

[Jenkins](https://engci-private-sjc.cisco.com/jenkins/eti-sre/job/SRE/job/pipeline/job/sre-go-helloworld/)

## Additional Setup Instructions for Lab VM before build

Skip this step if is not building on Lab VM (i.e. rcdn6-vmXX-YYY).

export DOCKER_BLD_ARGS="https_proxy=http://proxy.esl.cisco.com:80"

Create a credentials file for artifactory in ~/.nyota/credentials with content:

```bash
artifactory_user=yourCECID
artifactory_password=yourPASSWORD
```

and run after:

```bash
chmod 400 /Users/jegarnie/.nyota/credentials
```

## Build and Run Docker

```bash
./build-docker.sh

docker run --name postgres -e POSTGRES_DB=helloworld -e POSTGRES_PASSWORD=strongpassword -d postgres
docker run -it -p 5000:5000 -e DB_CONNECTION_INFO=/tmp/dbconfig.json -e DB_NAME=helloworld -v $PWD/build/:/tmp/  --link postgres:postgre sre-go-helloworld
```

## Instructions for Developers

### Install git hook to autoformat and run tests

From the main directory, run:

```bash
ln -s $(pwd)/githooks/pre-commit .git/hooks/pre-commit`
```

### Different directories

### /src/pkg

  This will contain the libraries internal to the app

### /docs

  If you need the swag tool, install it using command:

```bash
   go get -u github.com/swaggo/swag/cmd/swag
```

  This contains the rest api specifications in JSON/ yaml. This specifications
  would be used for api documentation. Generated from handler comments using command:

```bash
swag init --parseDependency --parseInternal
```

  Then, if you want to play with the API, run the swagger docker container using:

```bash
docker run --name swagger_ui -p 8080:8080 -d swaggerapi/swagger-ui
```

Open a browser to the swagger UI [http://localhost:8080/](http://localhost:8080/), put http://localhost:5000/docs in the input
field and click on explore button.

### /test

  This contains all the tests and dependent mock functions.

### /deployments

  This contains system and container orchestration, deployment configurations and templates.

### /githooks

  This will have git hooks e.g. go_fmt, staticcheck etc.

### main.go

  This is the main package for the microservice.
