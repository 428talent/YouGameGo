package service

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/util"
)

var (
	UserExistError  = errors.New("user already exist")
	LoginUserFailed = errors.New("user login failed")
)

func CreateUserAccount(username string, password string) (*int64, error) {
	o := orm.NewOrm()
	if o.QueryTable("auth_user").Filter("UserName", username).Exist() {
		panic(UserExistError)
	}
	encryptPassword, err := util.EncryptSha1WithSalt(password)
	if err != nil {
		panic(err)
	}
	err = o.Begin()
	if err != nil {
		panic(err)
	}

	profile := models.Profile{
		Nickname: username,
		Avatar:   "",
		Email:    "",
	}
	profileId, err := o.Insert(&profile)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Username: username,
		Password: *encryptPassword,
		Enable:   true,
		Profile: &models.Profile{
			Id: int(profileId),
		},
	}
	userId, err := o.Insert(&user)
	defer func() {
		troubleMaker := recover()
		if troubleMaker != nil {
			err = troubleMaker.(error)
			err = o.Rollback()
		}
	}()
	return &userId, err
}

// 用户登录
func UserLogin(username string, password string) (string, *models.User, error) {
	o := orm.NewOrm()
	user := models.User{
		Username: username,
		Password: password,
	}
	if !models.CheckUserValidate(&user) {
		return "", nil, LoginUserFailed
	}
	user.LastLogin = time.Now()
	_, err := o.Update(&user, "LastLogin")
	if err != nil {
		return "", nil, err
	}

	signString, err := security.GenerateJWTSign(&user)
	if err != nil {
		return "", nil, err
	}
	return *signString, &user, err
}

func UpdateUserAvatar(uid int, path string) error {
	o := orm.NewOrm()
	var profile models.Profile
	err := o.QueryTable("profile").Filter("User__Id", uid).One(&profile)
	if err != nil {
		return err
	}
	err = util.DeleteFile(profile.Avatar)
	if err != nil {
		return err
	}
	profile.Avatar = path
	_, err = o.Update(&profile)
	return nil
}
