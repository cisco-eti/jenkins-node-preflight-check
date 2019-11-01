# go-template for Nyota microservices
Project layout based on this [link](https://github.com/golang-standards/project-layout)

# Instructions for Developers

## Install git hook to autoformat and run tests

From the main directory, run `ln -s $(pwd)/githooks/pre-commit .git/hooks/pre-commit`

## Build and Run
```
go get -v ./...

go build

./gotemplate-microservice
```

