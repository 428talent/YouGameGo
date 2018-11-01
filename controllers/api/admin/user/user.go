package user

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"

	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
)

type ApiUserController struct {
	beego.Controller
}

func RegisterUserApiRouter() {
	beego.Router("/api/users", &ApiUserController{}, "post:CreateUser")
	beego.Router("/api/user/auth", &ApiUserController{}, "post:UserLogin")
}

type CreateUserResponsePayload struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
}

func (c *ApiUserController) CreateUser() {
	var requestBody parser.CreateUserRequestStruct
	err := requestBody.Parse(c.Ctx.Input.RequestBody)
	if err != nil {
		panic(err)
	}
	userId, err := models.CreateUserAccount(requestBody.Username, requestBody.Password)
	if err != nil {
		panic(err)
	}
	c.ServeJSON()
	defer func() {
		troubleMaker := recover()
		if troubleMaker != nil {
			err = troubleMaker.(error)
			switch err {
			case service.UserExistError:
				UserExistError.ServerError(c.Controller)
			default:
				api.HandleApiError(c.Controller, err)
			}
		} else {
			c.Data["json"] = serializer.CommonApiResponseBody{
				Success: true,
				Payload: CreateUserResponsePayload{
					Username: requestBody.Username,
					Id:       *userId,
				},
			}
		}
	}()
}

func (c *ApiUserController) UserLogin() {
	var requestStruct = parser.GetTokenRequestStruct{}
	requestData, err := requestStruct.ParseGetTokenRequestBody(c.Ctx.Input.RequestBody)
	if err != nil {
		panic(err)
	}

	signString,user,err :=  service.UserLogin(requestData.LoginName,requestData.Password)
	if err != nil {
		panic(err)
	}
	defer func() {
		troubleMaker := recover()
		if troubleMaker != nil {
			err = troubleMaker.(error)
			switch err {
			case service.LoginUserFailed:
				api.AuthFailedError.ServerError(c.Controller)
			default:
				api.HandleApiError(c.Controller, err)
			}
		} else {
			responseBody := serializer.UserLoginResponseBody{}
			c.Data["json"] = responseBody.Serialize(signString, *user)
			c.ServeJSON()
		}
	}()

}

//type GetUserResponseBody struct {
//	Id        int    `json:"id"`
//	Username  string `json:"username"`
//	LastLogin *int64 `json:"last_login"`
//	CreateAt  *int64 `json:"create_at"`
//}
//
//func SerializerUser(user *models.User) *GetUserResponseBody {
//	createAt := user.Created.Unix()
//	lastLogin := user.LastLogin.Unix()
//
//	serializerData := &GetUserResponseBody{
//		Id:        user.Id,
//		Username:  user.Username,
//		LastLogin: &lastLogin,
//		CreateAt:  &createAt,
//	}
//
//	if (user.LastLogin == time.Time{}) {
//		serializerData.LastLogin = nil
//	}
//	return serializerData
//}

func (c *ApiUserController) GetUser() {
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		return
	}
	user, err := models.GetUserById(userId)
	if err != nil {
		beego.Error(err)
		return
	}
	serializeData, err := serializer.SerializeUserObject(*user, serializer.SerializeUser{})
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = serializeData
	c.ServeJSON()
}

//
//
//func (c *ApiUserController) GetUserList() {
//	page, pageSize := util.ReadPageParam(c.Controller)
//	count, userList, err := models.GetAllUser(page, pageSize)
//	if err != nil {
//		beego.Error(err)
//		return
//	}
//	var serializeData []interface{}
//	for _, data := range userList {
//		serializeData = append(serializeData, c.Serialize(*data))
//	}
//	c.Data["json"] = common.PageResponse{
//		Count:    *count,
//		PageSize: pageSize,
//		Page:     page,
//		Result:   serializeData,
//	}
//	c.ServeJSON()
//}
//
//
//func (c *ApiUserController) UploadAvatar() {
//	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	claims, err := security.ParseAuthHeader(c.Controller)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	c.ControllerContext = controllers.ControllerContext{
//		AuthClaims: *claims,
//	}
//
//	user, err := models.GetUserById(userId)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	f, h, err := c.GetFile("avatar")
//	if err != nil {
//		beego.Error(err)
//	}
//	defer f.Close()
//	models.ReadProfile(user)
//	err = os.Remove(user.Profile.Avatar)
//	path := "static/upload/avatar/" + util.EncodeFileName(h.Filename)
//	err = c.SaveToFile("avatar", path)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	err = user.Profile.SaveAvatar(path)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	data := c.Serialize(*user)
//	c.Data["json"] = &data
//	c.ServeJSON()
//}
//
//
//func (c *ApiUserController) ChangeUserProfile() {
//	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	claims, err := security.ParseAuthHeader(c.Controller)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	c.ControllerContext = controllers.ControllerContext{
//		AuthClaims: *claims,
//	}
//	var requestData request.ChangeUserProfileRequestBody
//	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//
//	err = requestData.ValidRequest()
//	if err != nil {
//		if validateError, ok := err.(*APIError.ValidateError); ok {
//			validateError.BuildResponse().ServerError(c.Controller, 400)
//			return
//		}
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//
//	user, err := models.GetUserById(userId)
//	if err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	if err = user.Profile.ChangeUserProfile(requestData.Email, requestData.Nickname); err != nil {
//		beego.Error(err)
//		controllers.AbortServerError(c.Controller)
//		return
//	}
//	data := c.Serialize(*user)
//	c.Data["json"] = &data
//	c.ServeJSON()
//}
