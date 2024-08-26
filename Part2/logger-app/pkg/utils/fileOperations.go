package utils

import (
	"errors"
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

func OverwriteFile(filename string, content string) error {
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
