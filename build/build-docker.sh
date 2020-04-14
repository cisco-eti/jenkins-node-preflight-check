#!/bin/bash
# Syntax build-docker.sh [-i|--image imagename]

PROJECT=go-template
DOCKER_IMAGE=nyota-gotemplate
H_OUT=index.html
S_OUT=staticanalysis.txt


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
    -s|--static-analysis)
        STATIC_ANALYSIS=sa
        shift
        ;;
    *) # unknown
        echo Unknown Parameter $1
        exit 4
    esac
done

echo BUILDING DOCKER $DOCKER_IMAGE
docker build -t $DOCKER_IMAGE -f build/Dockerfile --build-arg HTML_OUT=$H_OUT --build-arg CODE_COVERAGE=$CODE_COVERAGE --build-arg STATIC_ANALYSIS=$STATIC_ANALYSIS --build-arg SA_OUT=$S_OUT .

if [ "${CODE_COVERAGE}" = "cc" ] ; then
    # extract the H_OUT file from the docker image just created
    id=$(docker create $DOCKER_IMAGE)
    docker cp $id:/go/src/sqbu-github.cisco.com/Nyota/${PROJECT}/$H_OUT .
    docker rm -v $id
    if [ ! -d "pipeline/lib" ] ; then
        echo "Your coverage HTML report is in $H_OUT"
    fi
fi

if [ "${STATIC_ANALYSIS}" = "sa" ] ; then
    # extract the S_OUT file from the docker image just created
    id=$(docker create $DOCKER_IMAGE)
    docker cp $id:/go/src/sqbu-github.cisco.com/Nyota/${PROJECT}/$S_OUT .
    docker rm -v $id

    if [ ! -d "pipeline/lib" ] ; then
        echo "Your static analysis report is in $S_OUT"
    fi
fi
