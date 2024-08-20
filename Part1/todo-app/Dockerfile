FROM openjdk:17-jdk-alpine

WORKDIR /app
COPY ./build/libs/todo-app-0.0.1.jar todo-app-0.0.1.jar

ENTRYPOINT [ "java", "-jar", "todo-app-0.0.1.jar" ]