package utils

import (
	"errors"
	"fmt"
	"os"
)

func CreateFileIfNotExist(filename string) error {
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

func WriteToFile(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(fmt.Sprintf("%s\n", content))); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
