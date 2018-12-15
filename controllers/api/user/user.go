package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/controllers/api/order"
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
	user, err := models.CreateUserAccount(requestBody.Username, requestBody.Password)
	if err != nil {
		panic(err)
	}
	serializerModel := serializer.UserSerializerModel{}

	c.Data["json"] = serializerModel.Serialize(*user, util.GetSiteAndPortUrl(c.Controller))
	c.ServeJSON()
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
	serializerModel := serializer.UserSerializerModel{}
	c.Data["json"] = serializerModel.Serialize(*user, util.GetSiteAndPortUrl(c.Controller))
	c.ServeJSON()
}

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
	serializerModel := serializer.UserProfileModel{}
	c.Data["json"] = serializerModel.Serialize(*user.Profile,util.GetSiteAndPortUrl(c.Controller))
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
func (c *ApiUserController) GetOrderList() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		api.HandleApiError(c.Controller, e)
	})
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(err)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}
	orderUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	page, pageSize := util.ParsePageRequest(c.Controller)
	permissionContext := map[string]interface{}{
		"claims":      *claims,
		"orderUserId": orderUserId,
	}
	permissions := []api.ApiPermissionInterface{
		order.GetOwnOrderPermission{},
	}
	err = c.CheckPermission(permissions, permissionContext)
	if err != nil {
		panic(api.PermissionDeniedError)
	}
	//query filter
	builder := service.GetOrderListBuilder{}
	builder.SetUser(int64(claims.UserId))
	if states := c.GetStrings("state"); len(states) > 0 {
		builder.SetState(states)
	}
	count, orders, err := service.GetOrderList(builder)
	if err != nil {
		panic(err)
	}
	results := make([]interface{}, 0)
	for _, item := range orders {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleData(&serializer.OrderModel{}, results, util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}
func (c *ApiUserController) GetUserWishList() {
	var err error
	defer api.CheckError(func(e error) {
		api.HandleApiError(c.Controller, e)
	})

	queryBuilder := service.WishListQueryBuilder{}
	page, pageSize := c.GetPage()
	queryBuilder.WithPage(service.PageOption{
		Page:     page,
		PageSize: pageSize,
	})
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	if userId > 0 {
		queryBuilder.BelongToUser(userId)
	} else {
		panic(api.ParseJsonDataError)
	}
	count, wishlist, err := queryBuilder.GetWishList()
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
