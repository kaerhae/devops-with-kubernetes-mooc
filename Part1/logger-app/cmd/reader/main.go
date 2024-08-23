package main

import (
	"fmt"
	"log"
	"logger-app/api/handler"
	"logger-app/pkg/models"
	"net/http"
)

func main() {
	c := models.Config{}
	c.GetConfigFromEnv()
	if c.READER_PORT == "" {
		log.Fatal("Environment variables missing")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/current-status", handler.CurrentStatusHandler)
	mux.HandleFunc("/", handler.ReadFile)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.READER_PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
	}
}
