package utils

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

func InitLogger(p syslog.Priority, useSyslog bool) *log.Logger {
	if useSyslog {
		l, err := syslog.NewLogger(p, log.Ldate|log.Ltime)
		if err != nil {
			log.Fatal("Error on initializing logger")
		}
		multiwriter := io.MultiWriter(os.Stdout, l.Writer())
		l.SetOutput(multiwriter)
		return l
	}

	return log.Default()
}
