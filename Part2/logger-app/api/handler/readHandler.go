package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"logger-app/pkg/client"
	"logger-app/pkg/models"
	"logger-app/pkg/service"
	"net/http"
	"os"
)

func ReadFile(w http.ResponseWriter, req *http.Request) {
	c := models.Config{}
	c.GetConfigFromEnv()
	w.Header().Set("Content-Type", "application/json")

	ts, err := os.ReadFile(c.FilePaths.TIMESTAMP_FILE_PATH)
	if err != nil {
		errorRes := models.ErrorResponse{
			Status:  500,
			Message: "Error while reading timestamp data",
		}
		json.NewEncoder(w).Encode(errorRes)
		return

	}

	apiClient := client.InitClient()
	res, err := apiClient.GetPongs()
	if err != nil {
		fmt.Printf("Error: %v", err)
		errorRes := models.ErrorResponse{
			Status:  500,
			Message: "Error while reading pong data",
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	content := fmt.Sprintf(
		"%s: %s\n Ping / Pongs: %d",
		ts,
		service.GenerateRandomString(),
		res.Pongs,
	)
	io.WriteString(w, content)
	return
}
