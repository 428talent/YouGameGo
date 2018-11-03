package security

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"letauth/security"
	"yougame.com/yougame-server/models"
)

type UserClaims struct {
	jwt.StandardClaims
	UserId int
}

func ParseAuthHeader(c beego.Controller) (*UserClaims, error) {
	jwtToken := c.Ctx.Request.Header.Get("Authorization")
	beego.Debug(jwtToken)
	var claims UserClaims
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	userClaims := token.Claims.(*UserClaims)
	beego.Debug(userClaims.UserId)
	return &claims, nil
}
func ParseAuthString(jwtToken string) (*UserClaims, error) {
	beego.Debug(jwtToken)
	var claims UserClaims
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	userClaims := token.Claims.(*UserClaims)
	beego.Debug(userClaims.UserId)
	return &claims, nil
}
func GenerateJWTSign(user *models.User) (*string, error) {
	now := jwt.TimeFunc()
	expire := now.AddDate(0, 0, 15)
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Username,
			ExpiresAt: expire.Unix(),
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "LetAuth",
			Subject:   "All",
		},
		UserId: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(security.AppSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
