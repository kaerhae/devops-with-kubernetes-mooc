package models

import (
	"os"
)

type FilePaths struct {
	TIMESTAMP_FILE_PATH string
}

type Config struct {
	WRITER_PORT string
	READER_PORT string
	PING_URL    string
	FilePaths   FilePaths
}

func (c *Config) GetConfigFromEnv() {
	c.READER_PORT = os.Getenv("READER_PORT")
	c.FilePaths.TIMESTAMP_FILE_PATH = os.Getenv("TIMESTAMP_FILE_PATH")
	c.PING_URL = os.Getenv("PING_URL")
	/* DEFAULTs */
	if c.FilePaths.TIMESTAMP_FILE_PATH == "" {
		c.FilePaths.TIMESTAMP_FILE_PATH = "/usr/src/app/data/timestamps.txt"
	}

	if c.PING_URL == "" {
		c.PING_URL = "http://localhost:9001/pongs"
	}
}
