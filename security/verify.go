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

func getNamespace(userId int, verifyType string) string {
	return fmt.Sprintf("verify:%s:%d", verifyType, userId)
}
func GenerateVerifyCode(userId int, verifyType string) (int, error) {
	code := generateRandomCode()
	_, err := database.RedisClient.Set(getNamespace(userId, verifyType), code, time.Minute*15).Result()
	return code, err
}

func CheckGenerateVerifyCodeValidate(userId int, verifyType string, code int) bool {
	cacheCode, err := database.RedisClient.Get(getNamespace(userId, verifyType)).Result()
	if err != nil {
		return false
	}

	verifyCode, err := strconv.Atoi(cacheCode)
	if err != nil {
		return false
	}

	if verifyCode == code {
		return true
	}
	return false

}

func generateRandomCode() int {
	rand.Seed(time.Now().UnixNano())
	return 1000 + rand.Intn(9999-1000)
}
