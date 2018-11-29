package comment

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
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
	page,pageSize := c.GetPage()
	builder := service.CommentQueryBuilder{}
	builder.SetPage(service.PageOption{
		Page:page,
		PageSize:pageSize,
	})
	if userId,err := c.GetInt64("user",0);err == nil && userId != 0{
		builder.SetUser(userId)
	}
	if gameId,err := c.GetInt64("game");err == nil && gameId != 0{
		builder.SetGame(gameId)
	}
	if goodId,err := c.GetInt64("good");err == nil && goodId != 0{
		builder.SetGood(goodId)
	}
	var commentList []*models.Comment
	count,err := builder.Query(&commentList)
	if err != nil {
		panic(err)
	}

	results := make([]interface{}, 0)
	for _, item := range commentList {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleData(&serializer.CommentSerializeModel{}, results, util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList,count,page,pageSize)
}
