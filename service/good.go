package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
)
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
	_, userBuyGoodOfGame, err := models.GetGoodList(func(o orm.QuerySeter) orm.QuerySeter {
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
		User:   &user,
		Good:   &good,
		Enable: true,

		Content: content,
	}
	//存储商品评论
	o := orm.NewOrm()
	err = comment.Save(o)
	if err != nil {
		return nil, nil, err
	}

	return &good, &comment, nil

}

type GoodQueryBuilder struct {
	ResourceQueryBuilder
	gameIds        []interface{}
	gameCommentIds []interface{}
	goodIds        []interface{}
}

func (q *GoodQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return q.Query()
}

func (q *GoodQueryBuilder) InGameId(gameId ...interface{}) {
	q.gameIds = append(q.gameIds, gameId...)
}

func (q *GoodQueryBuilder) WithGameCommentGood(gameId ...interface{}) {
	q.gameCommentIds = append(q.gameCommentIds, gameId...)
}

func (q *GoodQueryBuilder) buildCondition() *orm.Condition {
	condition := q.build()
	if len(q.gameIds) > 0 {
		condition = condition.And("game_id__in", q.gameIds...)
	}
	return condition
}

func (q *GoodQueryBuilder) Query() (*int64, []*models.Good, error) {
	condition := q.buildCondition()
	if len(q.gameCommentIds) > 0 {
		commentQuery := CommentQueryBuilder{}
		commentQuery.SetGame(q.gameCommentIds...)
		_, result, err := commentQuery.Query()
		if err != nil {
			return nil, nil, err
		}
		if len(result)  != 0{
			goodIds := make([]interface{}, 0)
			for _, comment := range result {
				goodIds = append(goodIds, comment.Good.Id)
			}
			condition = condition.And("id__in", goodIds...)
		}

	}
	return models.GetGoodList(func(o orm.QuerySeter) orm.QuerySeter {
		setter := o.SetCond(condition).Limit(q.pageOption.PageSize).Offset(q.pageOption.Offset())
		if len(q.orders) > 0 {
			setter = setter.OrderBy(q.orders...)
		}
		return setter
	})
}

func (q *GoodQueryBuilder) Delete() error {
	condition := q.buildCondition()
	return models.DeleteGoodMultiple(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(condition)
	})
}

