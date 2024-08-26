package main

import (
	"flag"
	"log/syslog"
	"logger-app/pkg/models"
	"logger-app/pkg/service"
	"logger-app/pkg/utils"
)

func main() {
	var useSyslog bool
	flag.BoolVar(&useSyslog, "syslog", false, "Enable syslog. By default app logs only to stdout")
	c := models.Config{}
	c.GetConfigFromEnv()
	l := utils.InitLogger(syslog.LOG_INFO, useSyslog)
	service.LogGeneratedStringToFile(l, c.FilePaths.TIMESTAMP_FILE_PATH)
}
