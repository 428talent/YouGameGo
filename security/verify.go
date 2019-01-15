package security

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"yougame.com/yougame-server/database"
)

const (
	VerifyCodeTypeResetPassword = "ResetPassword"
)

func getNamespace(code int, verifyType string) string {
	return fmt.Sprintf("verify:%s:%d", verifyType, code)
}
func GenerateVerifyCode(userId int, verifyType string) (int, error) {
	code := generateRandomCode()
	_, err := database.RedisClient.Set(getNamespace(code, verifyType), userId, time.Minute*15).Result()
	return code, err
}

func GetVerifyCodeValue(verifyType string, code int) int {
	cacheCode, err := database.RedisClient.Get(getNamespace(code, verifyType)).Result()
	if err != nil {
		return 0
	}

	cacheUserId, err := strconv.Atoi(cacheCode)
	if err != nil {
		return 0
	}

	return cacheUserId
}

func ClearVerifyCode(verifyType string, code int) error {
	_, err := database.RedisClient.Del(getNamespace(code, verifyType)).Result()
	return err
}

func generateRandomCode() int {
	rand.Seed(time.Now().UnixNano())
	return 100000 + rand.Intn(999999-100000)
}
