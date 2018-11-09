package util

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func EncodeFileName(fileName string) string {
	nowString := time.Now().String()
	ext := filepath.Ext(fileName)
	return fmt.Sprintf("%x%s", md5.Sum([]byte(fileName+nowString)), ext)
}

func DeleteFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	err := os.Remove(path)
	return err
}
