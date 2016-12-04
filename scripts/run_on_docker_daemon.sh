#!/bin/bash

# 注:ポート番号は固定
docker run -d \
  --name $1 \
  -p 8080:8080 \
  -e SAKURA_IOT_ECHO_PATH \
  -e SAKURA_IOT_ECHO_HOSTNAME \
  -e SAKURA_IOT_ECHO_SECRET \
  -e SAKURA_IOT_ECHO_DEBUG \
  $1:latest ${@:2}