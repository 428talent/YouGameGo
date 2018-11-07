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
