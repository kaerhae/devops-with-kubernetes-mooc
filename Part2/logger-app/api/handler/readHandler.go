package handler

import (
	"encoding/json"
	"fmt"
	"io"
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

	pings, err := os.ReadFile(c.FilePaths.PING_FILE_PATH)
	if err != nil {
		errorRes := models.ErrorResponse{
			Status:  500,
			Message: "Error while reading ping data",
		}
		fmt.Println(err)
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	content := fmt.Sprintf("%s: %s\n Ping / Pongs: %s", ts, service.GenerateRandomString(), pings)
	io.WriteString(w, content)
}
