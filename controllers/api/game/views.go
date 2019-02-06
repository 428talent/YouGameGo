package game

import (
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type AdminGetGameView struct {
	Controller *Controller
}

func (v *AdminGetGameView) Render() interface{} {
	template := serializer.NewGameTemplate(serializer.AdminGameTemplateType)
	gameId, err := strconv.Atoi(v.Controller.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	queryBuilder := service.GameQueryBuilder{}
	queryBuilder.InId(gameId)
	count, result, err := queryBuilder.Query()
	if err != nil {
		panic(err)
	}
	if *count == 0 {
		panic(api.ResourceNotFoundError)
	}
	gameModel := result[0]
	if gameModel.Band != nil{
		err = gameModel.ReadGameBand()
		if err != nil {
			panic(err)
		}
	}else{
		gameModel.Band = &models.Image{
			Path:"",
		}
	}

	template.Serialize(gameModel, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	})
	return template
}

type DefaultGetGameView struct {
	Controller *Controller
}

func (v *DefaultGetGameView) Render() interface{} {
	template := serializer.NewGameTemplate(serializer.DefaultGameTemplateType)
	gameId, err := strconv.Atoi(v.Controller.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	queryBuilder := service.GameQueryBuilder{}
	queryBuilder.InId(gameId)
	count, result, err := queryBuilder.Query()
	if err != nil {
		panic(err)
	}
	if *count == 0 {
		panic(api.ResourceNotFoundError)
	}
	gameModel := result[0]

	template.Serialize(gameModel, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(v.Controller.Controller),
	})
	return template
}
