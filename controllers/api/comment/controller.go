package comment

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiCommentController struct {
	api.ApiController
}

func (c *ApiCommentController) GetCommentList() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(err)
		api.HandleApiError(c.Controller, err)
	})
	page, pageSize := c.GetPage()
	builder := service.CommentQueryBuilder{}
	builder.SetPage(service.PageOption{
		Page:     page,
		PageSize: pageSize,
	})
	if userId, err := c.GetInt64("user", 0); err == nil && userId != 0 {
		builder.SetUser(userId)
	}
	if gameId, err := c.GetInt64("game"); err == nil && gameId != 0 {
		builder.SetGame(gameId)
	}
	if goodId, err := c.GetInt64("good"); err == nil && goodId != 0 {
		builder.SetGood(goodId)
	}
	var commentList []*models.Comment
	count, err := builder.Query(&commentList)
	if err != nil {
		panic(err)
	}

	results := make([]interface{}, 0)
	for _, item := range commentList {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleData(&serializer.CommentSerializeModel{}, results, util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}

func (c *ApiCommentController) CreateComment() {
	var err error
	defer api.CheckError(func(e error) {
		beego.Error(err)
		api.HandleApiError(c.Controller, err)
	})

	claims, err := c.GetAuth()
	if err != nil {
		panic(err)
	}

	err = c.CheckPermission([]api.ApiPermissionInterface{
		CreateCommentPermission{},
	}, map[string]interface{}{
		"claims": *claims,
	})
	if err != nil {
		panic(err)
	}

	requestBodyModel := parser.CreateCommentModel{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody,&requestBodyModel)
	if err != nil {
		panic(api.ParseJsonDataError)
	}

	goodId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	comment, err := service.CreateComment(requestBodyModel.Content, requestBodyModel.Evaluation, &models.User{Id: claims.UserId}, &models.Good{Id: goodId})
	if err != nil {
		panic(err)
	}
	serializeModel := serializer.CommentSerializeModel{}
	c.Data["json"] = serializeModel.SerializeData(*comment,util.GetSiteAndPortUrl(c.Controller))
	c.ServeJSON()
}
