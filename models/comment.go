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
	Id      int
	Good    *Good `orm:"rel(fk)"`
	User    *User `orm:"rel(fk)"`
	Content string
	Rating  int
	Enable  bool `mapstructure:"enable"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func DeleteCommentMultiple(filter func(o orm.QuerySeter) orm.QuerySeter) error {
	o := orm.NewOrm()
	setter := filter(o.QueryTable("comment"))
	_, err := setter.Update(orm.Params{
		"enable": false,
	})
	return err
}


func (comment *Comment) Query(id int64) error {
	comment.Id = int(id)
	err := orm.NewOrm().Read(comment)
	return err
}

func (comment *Comment) Delete(o orm.Ormer) error {
	comment.Enable = false
	_, err := o.Update(comment, "enable")
	return err
}

func (comment *Comment) Update(id int64, o orm.Ormer, fields ...string) error {
	comment.Id = int(id)
	_, err := o.Update(comment, fields...)
	return err
}

func (comment *Comment) Save(o orm.Ormer) error {
	_, err := o.Insert(comment)
	return err
}

func GetCommentList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Comment, error) {
	o := orm.NewOrm()
	seter := o.QueryTable("comment")
	var result []*Comment
	_, err := filter(seter).All(&result)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, result, err
}
func GetGameCommentCount(gameId int) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("comment").Filter("Good__Game__id", gameId).Filter("enable", true).Count()
}

func GetGameCommentWithEvaluationCount(gameId int, evaluation string) (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("comment").Filter("Good__Game__id", gameId).Filter("evaluation", evaluation).Filter("enable", true).Count()
}

type CommentRatingCountResult struct {
	Rating int64 `json:"rating"`
	Count  int64 `json:"count"`
}

func GetGameRatingCount(gameId int) ([]*CommentRatingCountResult, error) {
	o := orm.NewOrm()

	var resultList []*CommentRatingCountResult
	sql := `select comment.rating,count(*) as count
from comment
       inner join good
       inner join game
where
      good.game_id = game.id AND good.id = comment.good_id AND game_id = ?
group by rating`
	_, err := o.Raw(sql, gameId).QueryRows(&resultList)
	if err != nil {
		return nil, err
	}
	return resultList, nil
}
