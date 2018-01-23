package api_admin_user

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	"time"
	"you_game_go/controllers/api/admin"
	"you_game_go/models"
)

type UserLoginRequest struct {
	Username string
	Password string
}
type UserLoginResponse struct {
	UserId int
	Token  string
}
type UserLoginController struct {
	beego.Controller
}

func (c *UserLoginController) Post() {
	requestData := UserLoginRequest{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)

	if err != nil {
		beego.Error(err)
	}

	user := models.User{
		Username: requestData.Username,
		Password: requestData.Password,
	}

	isValidate := models.CheckUserValidate(&user)

	if isValidate {
		beego.Info(user)
		loginToken := models.Token{
			UserId:    user.Id,
			LoginTime: time.Now(),
		}
		userToken := uuid.New()
		models.InsertTokenToRedis(&loginToken, userToken)
		c.Ctx.SetCookie("token", userToken)
		responseBody := UserLoginResponse{
			UserId: user.Id,
			Token:  userToken,
		}
		c.Data["json"] = responseBody
	} else {
		apiError := api_admin.ApiError{
			Message:   "用户名或密码错误",
			ErrorCode: 2001,
		}
		c.Data["json"] = apiError
		c.Ctx.Output.SetStatus(401)
	}
	c.ServeJSON()
}
