#!/bin/bash

VERSION=$(cat VERSION.txt)

# READER
docker build -f docker/reader/Dockerfile -t kaerhae/devops-with-kube-logger-reader:v$VERSION .

docker push kaerhae/devops-with-kube-logger-reader:v$VERSION

# WRITER
docker build -f docker/writer/Dockerfile -t kaerhae/devops-with-kube-logger-writer:v$VERSION .

docker push kaerhae/devops-with-kube-logger-writer:v$VERSION
