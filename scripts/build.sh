#!/usr/bin/env bash

ROOT_PATH=$(realpath $(dirname $0)/..)

cd ${ROOT_PATH}

if test ! -f .buildrc; then
    echo ".buildrc file not found, you can copy .buildrc.example to .buildrc and modify it"
    exit 1
fi

source .buildrc

set -x

docker build -f docker/Dockerfile -t ${IMAGE_NAME}:${IMAGE_VERSION} -t ${IMAGE_NAME}:latest .
