package handler

import (
	"encoding/json"
	"logger-app/pkg/models"
	"logger-app/pkg/service"
	"net/http"
	"os"
)

func ReadFile(w http.ResponseWriter, req *http.Request) {
	c := models.Config{}
	c.GetConfigFromEnv()

	w.Header().Set("Content-Type", "application/json")

	content, err := os.ReadFile(c.FILE_PATH)
	if err != nil {
		errorRes := models.ErrorResponse{
			Status:  500,
			Message: "Error while reading data",
		}
		json.NewEncoder(w).Encode(errorRes)
	}
	res := models.ReadResponse{
		Hash:    service.GenerateRandomString(),
		Content: string(content),
	}
	json.NewEncoder(w).Encode(res)
}
