package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type WishList struct {
	Id      int
	UserId  int
	Enable bool
	Game    *Game     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func SaveWishList(wishList *WishList) error {
	o := orm.NewOrm()
	id, err := o.Insert(wishList)
	if err != nil {
		return err
	}
	wishList.Id = int(id)
	return nil
}

func GetWishList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*WishList, error) {
	o := orm.NewOrm()
	var wishlist []*WishList
	seter := o.QueryTable("wish_list")
	_, err := filter(seter).All(&wishlist)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, wishlist, err
}

func (w *WishList) ReadGame() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(w.Game, "Game")
	return err
}
