FROM openjdk:17-jdk-alpine

ARG VERSION
ENV VERSION=${VERSION}
WORKDIR /app
COPY ./build/libs/todo-backend-${VERSION}.jar todo-backend.jar

ENTRYPOINT [ "java", "-jar", "todo-backend.jar" ]