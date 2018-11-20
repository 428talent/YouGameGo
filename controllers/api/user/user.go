package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
	"yougame.com/yougame-server/validate"

	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
)

type ApiUserController struct {
	api.ApiController
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
	var err error
	defer api.CheckError(func(e error) {
		if validateError, ok := e.(*validate.ValidateError); ok {
			validateError.BuildResponse().ServerError(c.Controller)
			return
		}
		switch err {
		case service.UserExistError:
			UserExistError.ServerError(c.Controller)
			return
		default:
			api.HandleApiError(c.Controller, err)
		}
	})
	var requestBody parser.CreateUserRequestStruct
	err = requestBody.Parse(c.Ctx.Input.RequestBody)
	if err != nil {
		panic(err)
	}
	err = validate.ValidateData(requestBody)
	if err != nil {
		panic(err)
	}
	userId, err := models.CreateUserAccount(requestBody.Username, requestBody.Password)
	if err != nil {
		panic(err)
	}
	c.Data["json"] = serializer.CommonApiResponseBody{
		Success: true,
		Payload: CreateUserResponsePayload{
			Username: requestBody.Username,
			Id:       *userId,
		},
	}

}

func (c *ApiUserController) UserLogin() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(err)
		if e == service.LoginUserFailed {
			api.AuthFailedError.ServerError(c.Controller)
			return
		}
		api.HandleApiError(c.Controller, err)
	})
	var requestStruct = parser.GetTokenRequestStruct{}
	requestData, err := requestStruct.ParseGetTokenRequestBody(c.Ctx.Input.RequestBody)
	if err != nil {
		panic(err)
	}

	signString, user, err := service.UserLogin(requestData.LoginName, requestData.Password)
	if err != nil {
		panic(err)
	}
	responseBody := serializer.UserLoginResponseBody{}
	c.Data["json"] = responseBody.Serialize(signString, *user)
	c.ServeJSON()

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
func (c *ApiUserController) UploadAvatar() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(err)
		if e == service.LoginUserFailed {
			api.AuthFailedError.ServerError(c.Controller)
			return
		}
		api.HandleApiError(c.Controller, err)
	})
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	claims, err := c.GetAuth()
	if err != nil {
		panic(err)
	}
	if claims == nil {
		panic(api.ClaimsNoFoundError)
	}
	user, err := models.GetUserById(userId)
	if err != nil {
		panic(err)
	}

	//check user is itself
	if user.Id != claims.UserId {
		panic(api.PermissionDeniedError)
	}

	f, h, err := c.GetFile("avatar")
	if err != nil {
		beego.Error(err)
	}
	defer f.Close()
	models.ReadProfile(user)
	err = os.Remove(user.Profile.Avatar)
	path := "static/upload/user/avatar/" + util.EncodeFileName(h.Filename)
	err = c.SaveToFile("avatar", path)
	if err != nil {
		panic(err)
	}
	err = user.Profile.SaveAvatar(path)
	if err != nil {
		panic(err)
	}
	serializeData, err := serializer.SerializeUserObject(*user, serializer.SerializeUser{})
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = serializeData
	c.ServeJSON()
}

func (c *ApiUserController) ChangeUserProfile() {
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	claims, err := c.GetAuth()
	if err != nil {
		panic(err)
	}
	if claims == nil {
		panic(api.ClaimsNoFoundError)
	}
	var requestData parser.ChangeProfileRequestStruct
	err = parser.ParseReqeustBody(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		panic(api.ParseJsonDataError)
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		panic(err)
	}
	if user.Id != claims.UserId {
		panic(api.PermissionDeniedError)
	}

	if err = user.Profile.ChangeUserProfile("", requestData.Nickname); err != nil {
		panic(err)
	}
	serializeData, err := serializer.SerializeUserObject(*user, serializer.SerializeUser{})
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = serializeData
	c.ServeJSON()
}
func (c *ApiUserController) UploadJsonAvatar() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		switch e {
		default:
			api.HandleApiError(c.Controller, e)
			return
		}
	})
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(err)

	}
	if claims == nil {
		panic(service.NoAuthError)
	}

	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		panic(err)
	}

	requestBodyStruct := parser.UploadUserAvatarRequestStruct{}
	err = requestBodyStruct.Parse(c.Ctx.Input.RequestBody)
	if err != nil {
		panic(err)
	}
	filename := util.EncodeFileName(fmt.Sprintf("user_avatar_%d", claims.UserId))
	var filePath string
	switch requestBodyStruct.ImageType {
	case "jpg":
		filePath, err = util.Base64toJpg(requestBodyStruct.Avatar, filename)
	case "png":
		filePath, err = util.Base64toPng(requestBodyStruct.Avatar, filename)
	}
	err = service.UpdateUserAvatar(claims.UserId, filePath)
	if err != nil {
		panic(err)
	}
	serializeData, err := serializer.SerializeUserObject(*user, serializer.SerializeUser{})
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = serializeData
	c.ServeJSON()
}

func (c *ApiUserController) GetUserWishList() {
	var err error
	defer api.CheckError(func(e error) {
		api.HandleApiError(c.Controller, e)
	})

	page, pageSize := c.GetPage()
	queryContext := make(map[string]interface{})
	beego.Debug(c.Ctx.Input.Param(":id"))
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	if userId > 0 {
		queryContext["user"] = int64(userId)
	}else{
		panic(api.ParseJsonDataError)
	}
	count, wishlist, err := service.GetWishList(queryContext, page, pageSize)
	if err != nil {
		panic(err)
	}

	results := make([]interface{}, 0)
	for _, item := range wishlist {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleData(&serializer.WishListModel{}, results, util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}
