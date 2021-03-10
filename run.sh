#!/bin/bash

# Build, dockerize, and run fs-explorer
# $1: HOST_DIR: The directory on local machine to host

docker build -t fs-explorer:dev .
docker run -p 8080:8080 -v "$1":/foo fs-explorer:dev
