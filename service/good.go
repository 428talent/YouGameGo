package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)

func GetGoodById(goodId int) (*models.Good, error) {
	good := models.Good{Id: goodId}
	err := good.QueryById()
	return &good, err

}

func CreateGoodComment(user models.User, content string, goodId int, evaluation string) (*models.Good, *models.Comment, error) {
	err := security.CheckUserPermission(user, "CreateComment")
	if err != nil {
		return nil, nil, PermissionNotAccess
	}
	good := models.Good{Id: goodId}
	err = good.QueryById()
	if err != nil {
		return nil, nil, err
	}
	if !good.Enable {
		return nil, nil, ResourceNotEnable
	}
	err = good.ReadGame()
	if err != nil {
		return nil, nil, err
	}

	//检查用户是否购买该商品
	userBuyGoodOfGame, err := models.GetGoodList(func(o orm.QuerySeter) orm.QuerySeter {
		o.Filter("Users__User__Id", user.Id).Filter("game__id", good.Game).Filter("Id", good.Id)
		return o
	})
	if err != nil {
		return nil, nil, err
	}
	if len(userBuyGoodOfGame) == 0 {
		return nil, nil, UserNotBoughtGood
	}

	comment := models.Comment{
		User:       &user,
		Good:       &good,
		Enable:     true,
		Evaluation: evaluation,
		Content:    content,
	}
	//存储商品评论
	o := orm.NewOrm()
	err = comment.Save(o)
	if err != nil {
		return nil, nil, err
	}

	return &good, &comment, nil

}
