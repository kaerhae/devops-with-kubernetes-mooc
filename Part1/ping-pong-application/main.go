package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var COUNTER int

type Config struct {
	PORT string
}

func (c *Config) GetConfigFromEnv() {
	c.PORT = os.Getenv("PORT")

	if c.PORT == "" {
		log.Fatal("Environment variables missing")
	}
}

func main() {
	c := Config{}
	c.GetConfigFromEnv()

	mux := http.NewServeMux()
	mux.HandleFunc("/pingpong", CurrentStatusHandler)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
	}
}

func CurrentStatusHandler(w http.ResponseWriter, req *http.Request) {
	COUNTER = COUNTER + 1
	io.WriteString(w, fmt.Sprintf("pong %d", COUNTER))
}
