package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type WishList struct {
	Id      int
	UserId  int
	Enable  bool
	Game    *Game     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (w *WishList) Query(id int64) error {
	o := orm.NewOrm()
	w.Id = int(id)
	err := o.Read(w)
	return err
}

func (w *WishList) Save(o orm.Ormer) error {
	_, err := o.Insert(w)
	return err
}

func (w *WishList) Delete(o orm.Ormer) error {
	_, err := o.Delete(w)
	return err
}

func (w *WishList) Update(id int64, o orm.Ormer, fields ...string) error {
	w.Id = int(id)
	_, err := o.Update(w, fields...)
	return err
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

func DeleteWishList(filter func(o orm.QuerySeter) orm.QuerySeter) error {
	o := orm.NewOrm()
	seter := o.QueryTable("wish_list")
	_, err := filter(seter).Update(orm.Params{"enable": false})
	if err != nil {
		return err
	}
	return err
}

func (w *WishList) ReadGame() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(w.Game, "Game")
	return err
}
