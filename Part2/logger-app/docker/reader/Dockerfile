FROM golang:1.22 as builder
WORKDIR /app
COPY . .
ENV BINARY_NAME=logger-reader

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/${BINARY_NAME} cmd/reader/main.go

FROM golang:1.22
WORKDIR /app
COPY --from=builder ./app/bin/${BINARY_NAME} .

ENTRYPOINT [ "./logger-reader" ]