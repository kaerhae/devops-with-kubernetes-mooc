#!/bin/bash

VERSION=$(cat VERSION.txt)

docker build -t kaerhae/devops-with-kube-logger:v$VERSION .

docker push kaerhae/devops-with-kube-logger:v$VERSION