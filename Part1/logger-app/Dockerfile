FROM golang:1.22
WORKDIR /app
COPY . .
ENV BINARY_NAME=logger
RUN go build -o ${BINARY_NAME} main.go

CMD [ "./logger" ]