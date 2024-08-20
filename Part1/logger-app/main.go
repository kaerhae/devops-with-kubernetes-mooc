package main

import (
	"flag"
	"io"
	"log"
	"log/syslog"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	var useSyslog bool
	flag.BoolVar(&useSyslog, "syslog", false, "Enable syslog. By default app logs only to stdout")
	l := InitLogger(syslog.LOG_INFO, useSyslog)
	LogGeneratedString(l)
}

func GenerateRandomString() string {
	randomString := uuid.New()
	return randomString.String()
}

func LogGeneratedString(l *log.Logger) {
	for {
		l.Println(GenerateRandomString())
		time.Sleep(5 * time.Second)
	}
}

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
