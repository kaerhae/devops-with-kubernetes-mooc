package main

import (
	"fmt"
	"log"
	"logger-app/api/handler"
	"logger-app/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	c := models.Config{}
	c.GetConfigFromEnv()
	if c.READER_PORT == "" {
		log.Fatal("Environment variables missing")
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/", handler.ReadFile)
	mux.HandleFunc("/current-status", handler.CurrentStatusHandler)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.READER_PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
	}
}
