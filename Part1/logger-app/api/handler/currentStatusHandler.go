package handler

import (
	"io"
	"logger-app/pkg/service"
	"net/http"
)

func CurrentStatusHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(
		w,
		service.GenerateRandomStringWithTimestamp(),
	)
}
