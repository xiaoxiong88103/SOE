#!/bin/bash

# 检查是否有参数传递给脚本
if [ "$1" = "dev" ]; then
  cd ..
  docker build -f build/dev-dockerfile -t soe:dev .
else
  cd ..
  docker build -f build/build-dockerfile -t soe:build .
fi
