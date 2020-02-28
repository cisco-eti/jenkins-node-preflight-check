#!/bin/bash
# Syntax build-docker.sh [-i|--image imagename]

DOCKER_IMAGE=nyota-gotemplate

while [[ $# -gt 0 ]]
do
    key="${1}"

    case ${key} in
    -i|--image)
        DOCKER_IMAGE="${2}"
        shift;shift
        ;;
    -h|--help)
        less README.md
        exit 0
        ;;
    -c|--code-coverage)
        CODE_COVERAGE=cc
        shift
        ;;
    *) # unknown
        echo Unknown Parameter $1
        exit 4
    esac
done

if [ -z "$DOCKER_BLD_ARGS" ]; then
    echo BUILDING DOCKER $DOCKER_IMAGE
    docker build -t $DOCKER_IMAGE -f build/Dockerfile .
else
    echo BUILDING DOCKER $DOCKER_IMAGE with --build-arg $DOCKER_BLD_ARGS
    docker build -f build/Dockerfile --build-arg $DOCKER_BLD_ARGS -t $DOCKER_IMAGE .
fi
