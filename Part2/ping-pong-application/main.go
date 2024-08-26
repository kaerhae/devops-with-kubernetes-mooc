package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type PongResponse struct {
	Pongs int `json:"pongs"`
}

type Config struct {
	PORT      string
	FILE_PATH string
}

type ServerRouter interface {
	Ping(w http.ResponseWriter, req *http.Request)
	Pongs(w http.ResponseWriter, req *http.Request)
}

type serverRouter struct {
	Config Config
}

// Ping implements ServerRouter.
// When receives request, adds to counter, which count amount of pings made.
// Returns amount of Pongs.
func (s *serverRouter) Ping(w http.ResponseWriter, req *http.Request) {
	createFileIfNotExist(s.Config.FILE_PATH)
	counter, err := readFile(s.Config.FILE_PATH)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error: %v", err))
	}
	counter = counter + 1
	overwriteFile(s.Config.FILE_PATH, strconv.Itoa(counter))
	p := PongResponse{
		Pongs: counter,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// Pongs implements ServerRouter.
// Returns amount of Pings made to application.
func (s *serverRouter) Pongs(w http.ResponseWriter, req *http.Request) {
	count, err := readFile(s.Config.FILE_PATH)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error: %v", err))
	}
	p := PongResponse{
		Pongs: count,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func createFileIfNotExist(filename string) error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		return nil
	}

	return nil
}

func readFile(filename string) (int, error) {
	c, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	/* If file is empty, return zero requests */
	if len(c) == 0 {
		return 0, nil
	}

	r, err := strconv.Atoi(string(c))
	if err != nil {
		return 0, err
	}
	return r, nil
}

func overwriteFile(filename string, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(content))
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func InitServerRouter() ServerRouter {
	c := Config{}
	c.GetConfigFromEnv()
	return &serverRouter{
		Config: c,
	}
}

func (c *Config) GetConfigFromEnv() {
	c.PORT = os.Getenv("PORT")
	c.FILE_PATH = os.Getenv("PING_FILE_PATH")
	if c.PORT == "" || c.FILE_PATH == "" {
		log.Fatal("Environment variables missing")
	}
}

func main() {
	c := Config{}
	c.GetConfigFromEnv()
	s := InitServerRouter()
	mux := http.NewServeMux()
	mux.HandleFunc("/pingpong", s.Ping)
	mux.HandleFunc("/pongs", s.Pongs)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
	}
}
