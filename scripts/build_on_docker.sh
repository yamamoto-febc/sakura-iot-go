#!/bin/bash

set -e

DOCKER_IMAGE_NAME="$1-build"
DOCKER_CONTAINER_NAME="$1-build-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -t $DOCKER_IMAGE_NAME -f scripts/Dockerfile.build .

docker run --name $DOCKER_CONTAINER_NAME \
       -e SAKURA_IOT_ECHO_PATH \
       -e SAKURA_IOT_ECHO_HOSTNAME \
       -e SAKURA_IOT_ECHO_PORT \
       -e SAKURA_IOT_ECHO_SECRET \
       -e SAKURA_IOT_ECHO_DEBUG \
       -e TESTARGS \
       $DOCKER_IMAGE_NAME make "$@"
if [[ "$@" == *"build"* ]]; then
  docker cp $DOCKER_CONTAINER_NAME:`docker inspect -f "{{ .Config.WorkingDir  }}" $DOCKER_CONTAINER_NAME`/bin ./
fi
docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
