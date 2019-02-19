package util

import (
	"github.com/astaxie/beego/config"
	"os"
)

func GetConfigItem(envKey string, configKey string, config config.Configer, defaultValue string) string{
	value := os.Getenv(envKey)
	if len(value) == 0{
		value = config.DefaultString(configKey, defaultValue)
	}
	return value
}
