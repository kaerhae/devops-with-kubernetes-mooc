package models

import (
	"os"
)

type FilePaths struct {
	PING_FILE_PATH      string
	TIMESTAMP_FILE_PATH string
}

type Config struct {
	WRITER_PORT string
	READER_PORT string
	FilePaths   FilePaths
}

func (c *Config) GetConfigFromEnv() {
	c.READER_PORT = os.Getenv("READER_PORT")
	c.FilePaths.PING_FILE_PATH = os.Getenv("PING_FILE_PATH")
	c.FilePaths.TIMESTAMP_FILE_PATH = os.Getenv("TIMESTAMP_FILE_PATH")
	/* DEFAULT paths */
	if c.FilePaths.PING_FILE_PATH == "" {
		c.FilePaths.PING_FILE_PATH = "/usr/src/app/data/pings.txt"
	}
	if c.FilePaths.TIMESTAMP_FILE_PATH == "" {
		c.FilePaths.TIMESTAMP_FILE_PATH = "/usr/src/app/data/timestamps.txt"
	}
}
