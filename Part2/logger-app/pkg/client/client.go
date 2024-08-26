package client

import (
	"encoding/json"
	"io"
	"logger-app/pkg/models"
	"net/http"
)

type client struct {
	Client *http.Client
	Config models.Config
}
type Client interface {
	GetPongs() (models.PongResponse, error)
}

func InitClient() Client {
	c := models.Config{}
	c.GetConfigFromEnv()

	return &client{
		Client: http.DefaultClient,
		Config: c,
	}
}

func (c *client) GetPongs() (models.PongResponse, error) {
	resp, err := c.Client.Get(c.Config.PING_URL)
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
