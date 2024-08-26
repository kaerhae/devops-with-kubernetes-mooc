package service

import (
	"fmt"
	"log"
	"logger-app/pkg/utils"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomString() string {
	randomString := uuid.New()
	return randomString.String()
}

func GenerateTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GenerateRandomStringWithTimestamp() string {
	return fmt.Sprintf(
		"%s %s",
		GenerateTimestamp(),
		GenerateRandomString(),
	)
}

func LogGeneratedStringToFile(l *log.Logger, filename string) {
	for {
		r := GenerateRandomString()
		l.Println(r)
		utils.CreateFileIfNotExist(filename)
		utils.OverwriteFile(
			filename,
			GenerateTimestamp(),
		)
		time.Sleep(5 * time.Second)
	}
}
