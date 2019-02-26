package service

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"yougame.com/yougame-server/mail"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/util"
)

var (
	UserExistError       = errors.New("user already exist")
	EmailAlreadyExist    = errors.New("email already exist")
	LoginUserFailed      = errors.New("user login failed")
	VerifyCodeInvalidate = errors.New("verify code invalidate")
)

func CreateUserAccount(username string, password string, email string) (*models.User, error) {
	o := orm.NewOrm()
	if o.QueryTable("auth_user").Filter("UserName", username).Exist() {
		return nil, UserExistError
	}
	if o.QueryTable("profile").Filter("email", email).Exist() {
		return nil, EmailAlreadyExist
	}
	encryptPassword, err := util.EncryptSha1WithSalt(password)
	if err != nil {
		return nil, err
	}

	err = o.Begin()
	if err != nil {
		return nil, err
	}
	dbTransaction := func() (*models.User, error) {
		profile := models.Profile{
			Nickname: username,
			Avatar:   "",
			Email:    email,
		}
		profileId, err := o.Insert(&profile)
		if err != nil {
			beego.Error(err)
			err = o.Rollback()
			return nil, err
		}

		user := &models.User{
			Username: username,
			Password: *encryptPassword,
			Enable:   true,
			Profile: &models.Profile{
				Id: int(profileId),
			},
		}
		_, err = o.Insert(user)
		if err != nil {
			return nil, err
		}
		wallet := &models.Wallet{User: user, Balance: 0}
		_, err = o.Insert(wallet)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	user, err := dbTransaction()
	if err != nil {
		err := o.Rollback()
		if err != nil {
			return nil, err
		}
	}
	err = o.Commit()
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	//send welcome mail
	err = mail.SendWelcomeEmail(user, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(userId int) *models.User {
	o := orm.NewOrm()
	user := models.User{Id: userId}
	err := o.Read(&user)
	if err != nil {
		return nil
	}
	return &user
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

func UpdateUserProfile(profile models.Profile, fields ...string) (*models.Profile, error) {
	o := orm.NewOrm()
	user := profile.User
	user.ReadProfile()
	profile.Id = user.Profile.Id
	err := profile.Update(o, fields...)
	if err != nil {
		return nil, err
	}
	userProfile, err := models.GetProfileByUser(int64(profile.User.Id))
	if err != nil {
		return nil, err
	}
	return userProfile, nil

}

func SendResetMail(username string) error {
	user := &models.User{
		Username: username,
	}
	o := orm.NewOrm()
	err := o.Read(user, "Username")
	if err != nil {
		return err
	}

	code, err := security.GenerateVerifyCode(user.Id, security.VerifyCodeTypeResetPassword)
	if err != nil {
		return err
	}

	err = user.ReadProfile()
	if err != nil {
		return err
	}

	err = mail.SendVerifyCodeEmail(user, user.Profile.Email, code)
	if err != nil {
		return err
	}

	return nil

}

func UpdatePassword(code int, rawPassword string) error {
	userId := security.GetVerifyCodeValue(security.VerifyCodeTypeResetPassword, code)
	if userId == 0 {
		return VerifyCodeInvalidate
	}

	o := orm.NewOrm()
	password, err := util.EncryptSha1WithSalt(rawPassword)
	if err != nil {
		return err
	}
	user := &models.User{
		Id:       userId,
		Password: *password,
	}

	_, err = o.Update(user, "password")
	if err != nil {
		return err
	}
	err = security.ClearVerifyCode(security.VerifyCodeTypeResetPassword, code)
	return err

}

type UserQueryBuilder struct {
	ResourceQueryBuilder
}

func (b *UserQueryBuilder) Query() (*int64, []*models.User, error) {
	condition := b.build()
	count, users, err := models.GetUserList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
	return &count, users, err
}

func (b *UserQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}
