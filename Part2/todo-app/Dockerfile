FROM openjdk:17-jdk-alpine

ARG VERSION
ENV VERSION=${VERSION}
WORKDIR /app
COPY ./build/libs/todo-app-${VERSION}.jar todo-app.jar

ENTRYPOINT [ "java", "-jar", "todo-app.jar" ]