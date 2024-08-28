#!/bin/bash

./gradlew build

VERSION=$(./gradlew -q pV)

docker build --build-arg VERSION=$VERSION -t kaerhae/devops-with-kube-todo-backend:v$VERSION .

docker push kaerhae/devops-with-kube-todo-backend:v$VERSION