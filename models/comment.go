package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	EvaluationNegative = "Negative"
	EvaluationPositive = "Positive"
)

type Comment struct {
	Id         int
	Good       *Good `orm:"rel(fk)"`
	User       *User `orm:"rel(fk)"`
	Content    string
	Evaluation string
	Enable     bool
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func (comment *Comment) Save(o orm.Ormer) error {
	_, err := o.Insert(comment)
	return err
}

func GetCommentList(filter func(o orm.QuerySeter) orm.QuerySeter) ([]*Comment, error) {
	o := orm.NewOrm()
	var commentList []*Comment
	seter := o.QueryTable("comment")
	_, err := filter(seter).All(&commentList)
	return commentList, err
}
func (comment *Comment) GetList(filter func(o orm.QuerySeter) orm.QuerySeter, md interface{}) (count int64, err error) {
	o := orm.NewOrm()
	seter := o.QueryTable("comment")
	_, err = filter(seter).All(md)
	if err != nil {
		return
	}
	count, err = filter(seter).Count()
	return
}
func GetGameCommentCount(gameId int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("comment").Filter("Good__Game__id", gameId).Filter("enable", true).Count()
}

func GetGameCommentWithEvaluationCount(gameId int, evaluation string) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("comment").Filter("Good__Game__id", gameId).Filter("evaluation", evaluation).Filter("enable", true).Count()
}
