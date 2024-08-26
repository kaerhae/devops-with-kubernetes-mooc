FROM golang:1.22
WORKDIR /app
COPY . .
ENV BINARY_NAME=ping-pong-application

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ${BINARY_NAME} main.go

ENTRYPOINT [ "./ping-pong-application" ]