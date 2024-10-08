package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var COUNTER int

type Config struct {
	PORT      string
	FILE_PATH string
}

type ServerRouter interface {
	Ping(w http.ResponseWriter, req *http.Request)
}

type serverRouter struct {
	Config Config
}

// Ping implements ServerRouter.
func (s *serverRouter) Ping(w http.ResponseWriter, req *http.Request) {
	createFileIfNotExist(s.Config.FILE_PATH)
	counter, err := readFile(s.Config.FILE_PATH)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Error: %v", err))
	}
	counter = counter + 1
	overwriteFile(s.Config.FILE_PATH, strconv.Itoa(counter))
	io.WriteString(w, fmt.Sprintf("pong %d", counter))
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

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
	}
}
