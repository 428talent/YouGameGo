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
