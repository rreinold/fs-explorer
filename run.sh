#!/bin/bash

# Build, dockerize, and run fs-explorer

GOOS=linux go build fs-explorer.go
docker build -t fs-explorer:dev .
docker run -p 8080:8080 fs-explorer:dev
