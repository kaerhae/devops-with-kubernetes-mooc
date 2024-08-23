package models

import (
	"os"
)

type Config struct {
	WRITER_PORT string
	READER_PORT string
	FILE_PATH   string
}

func (c *Config) GetConfigFromEnv() {
	c.READER_PORT = os.Getenv("READER_PORT")
	c.FILE_PATH = os.Getenv("FILE_PATH")
	/* DEFAULT /usr/src/app/data */
	if c.FILE_PATH == "" {
		c.FILE_PATH = "/usr/src/app/data"
	}
}
