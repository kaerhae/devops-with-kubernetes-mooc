FROM golang:1.22 as builder
WORKDIR /app
COPY . .
ENV BINARY_NAME=logger-writer

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/${BINARY_NAME} cmd/writer/main.go

FROM golang:1.22
WORKDIR /app
COPY --from=builder ./app/bin/${BINARY_NAME} .

ENTRYPOINT [ "./logger-writer" ]