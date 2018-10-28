package util

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/config"
	"io"
)

func EncryptSha1(data string) (*string, error) {
	t := sha1.New()
	_, err := io.WriteString(t, data)
	if err != nil {
		return nil, err
	}
	enString := fmt.Sprintf("%x", t.Sum(nil))
	return &enString, nil
}

func EncryptSha1WithSalt(data string) (*string, error) {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		return nil, err
	}
	salt := appConfig.String("salt")

	enData, err := EncryptSha1(data + salt)
	if err != nil {
		return nil, err
	}
	return enData, nil
}
