package client

import (
	"encoding/json"
	"io"
	"logger-app/pkg/models"
	"net/http"
)

var (
	url string = ""
)

type client struct {
	Client *http.Client
}
type Client interface {
	GetPongs() (models.PongResponse, error)
}

func InitClient() Client {
	return &client{
		Client: http.DefaultClient,
	}
}

func (c *client) GetPongs() (models.PongResponse, error) {
	resp, err := c.Client.Get(url)
	if err != nil {
		return models.PongResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.PongResponse{}, err
	}
	var pongResponse models.PongResponse
	err = json.Unmarshal(body, &pongResponse)
	if err != nil {
		return models.PongResponse{}, err
	}

	return pongResponse, nil
}
