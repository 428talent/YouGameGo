package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

func CreateComment(content string, evaluate string, user *models.User, good *models.Good) (*models.Comment, error) {
	o := orm.NewOrm()
	comment := models.Comment{
		Content:    content,
		Evaluation: evaluate,
		Good:       good,
		User:       user,
	}
	commentId, err := o.Insert(comment)
	if err != nil {
		return nil, err
	}
	comment.Id = int(commentId)
	return &comment, nil
}

type CommentQueryBuilder struct {
	gameId int64
	userId int64
	goodId int64
	page PageOption
}

func (builder *CommentQueryBuilder) build() *orm.Condition {
	cond := orm.NewCondition()
	if builder.page.Page == 0 {
		builder.page.Page = 1
	}
	if builder.page.PageSize == 0 {
		builder.page.PageSize = 10
	}
	if builder.gameId != 0 {
		cond = cond.And("good__game__id",builder.gameId)
	}
	if builder.goodId != 0 {
		cond = cond.And("good_id",builder.gameId)
	}
	if builder.userId != 0 {
		cond = cond.And("user_id",builder.userId)
	}
	return cond
}

func (builder *CommentQueryBuilder) SetPage(pageOption PageOption) *CommentQueryBuilder {
	builder.page = pageOption
	return builder
}
func (builder *CommentQueryBuilder) SetGame(gameId int64) *CommentQueryBuilder {
	builder.gameId = gameId
	return builder

}

func (builder *CommentQueryBuilder) SetUser(userId int64) *CommentQueryBuilder {
	builder.userId = userId
	return builder

}

func (builder *CommentQueryBuilder) SetGood(goodId int64) *CommentQueryBuilder {
	builder.goodId = goodId
	return builder

}

func (builder *CommentQueryBuilder) Query(md interface{}) (int64, error) {
	modelStruct :=  models.Comment{}
	count, err :=modelStruct.GetList(func(o orm.QuerySeter) orm.QuerySeter {
		cond := builder.build()
		return o.SetCond(cond).Limit(builder.page.PageSize).Offset((builder.page.Page - 1) * builder.page.PageSize)
	}, md)
	return count, err
}


func GetGameCommentStatistics(gameId int) (count int64, positive int64, negative int64, err error) {
	if count, err = models.GetGameCommentCount(gameId); err != nil {
		return
	}
	if positive, err = models.GetGameCommentWithEvaluationCount(gameId, models.EvaluationPositive); err != nil {
		return
	}
	if negative, err = models.GetGameCommentWithEvaluationCount(gameId, models.EvaluationNegative); err != nil {
		return
	}
	return
}
