package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	PORT string
}

func (c *Config) GetConfigFromEnv() {
	c.PORT = os.Getenv("PORT")

	if c.PORT == "" {
		log.Fatal("Environment variables missing")
	}
}

func main() {
	var useSyslog bool
	flag.BoolVar(&useSyslog, "syslog", false, "Enable syslog. By default app logs only to stdout")
	/* Before Exercise 1.07, gen and log random string */
	// l := InitLogger(syslog.LOG_INFO, useSyslog)
	// LogGeneratedString(l)

	/* Exercise 1.07 --> */
	c := Config{}
	c.GetConfigFromEnv()

	mux := http.NewServeMux()
	mux.HandleFunc("/", CurrentStatusHandler)

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", c.PORT),
		mux,
	)
	if err != nil {
		log.Fatal("Error on starting server in port")
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

func CurrentStatusHandler(w http.ResponseWriter, req *http.Request) {
	r := GenerateRandomString()
	ts := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(w, fmt.Sprintf("%s %s", ts, r))
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
