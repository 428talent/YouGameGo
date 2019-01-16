package service

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/serializer"
)

var (
	CommentExistError = errors.New("comment already exist")
)

func CreateComment(content string, evaluate string, user *models.User, good *models.Good) (*models.Comment, error) {
	o := orm.NewOrm()

	if o.QueryTable("comment").Filter("user_id", user.Id).Filter("good_id", good.Id).Exist() {
		return nil, CommentExistError
	}
	comment := models.Comment{
		Content: content,
		Good:    good,
		User:    user,
	}
	commentId, err := o.Insert(&comment)
	if err != nil {
		return nil, err
	}
	comment.Id = int(commentId)
	return &comment, nil
}

type CommentQueryBuilder struct {
	ResourceQueryBuilder
	gameIds []interface{}
	userId  []interface{}
	goodId  []interface{}
}

func (builder *CommentQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return builder.Query()
}

func (builder *CommentQueryBuilder) SetGame(gameId ...interface{}) *CommentQueryBuilder {
	builder.gameIds = append(builder.gameIds, gameId...)
	return builder

}

func (builder *CommentQueryBuilder) SetUser(userId ...interface{}) *CommentQueryBuilder {
	builder.userId = append(builder.userId, userId...)
	return builder

}

func (builder *CommentQueryBuilder) SetGood(goodId ...interface{}) *CommentQueryBuilder {
	builder.goodId = append(builder.goodId, goodId...)
	return builder

}

func (builder *CommentQueryBuilder) Query() (*int64, []*models.Comment, error) {
	condition := builder.build()
	if len(builder.gameIds) != 0 {
		condition = condition.And("good__game__id__in", builder.gameIds...)
	}
	if len(builder.goodId) != 0 {
		condition = condition.And("good_id__in", builder.goodId...)
	}
	if len(builder.userId) != 0 {
		condition = condition.And("user_id__in", builder.userId...)
	}
	return models.GetCommentList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			querySetter = querySetter.OrderBy(builder.orders...)
		}
		return querySetter
	})

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

func GetCommentSummary(gameId int) (*serializer.CommentSummarySerializeTemplate, error) {

	ratingCount, err := models.GetGameRatingCount(gameId)
	if err != nil {
		return nil, err
	}
	result := &serializer.CommentSummarySerializeTemplate{
		Rating: ratingCount,
	}
	return result, nil
}
