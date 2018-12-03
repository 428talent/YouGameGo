package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type WishListQueryBuilder struct {
	ids        []interface{}
	userIds    []interface{}
	pageOption PageOption
	isOnlyEnable bool
}

func (builder *WishListQueryBuilder) InId(ids ...interface{}) *WishListQueryBuilder {
	builder.ids = append(builder.ids, ids...)
	return builder
}

func (builder *WishListQueryBuilder) BelongToUser(userId ...interface{}) *WishListQueryBuilder {
	builder.userIds = append(builder.userIds, userId...)
	return builder
}

func (builder *WishListQueryBuilder) WithPage(option PageOption) *WishListQueryBuilder {
	builder.pageOption = option
	return builder
}
func (builder *WishListQueryBuilder) OnlyEnable(isOnlyEnable bool) *WishListQueryBuilder {
	builder.isOnlyEnable = isOnlyEnable
	return builder
}
func (builder *WishListQueryBuilder) GetWishList() (int64, []*models.WishList, error) {
	cond := orm.NewCondition()
	if len(builder.userIds) > 0 {
		cond = cond.And("user_id__in", builder.userIds...)
	}
	if len(builder.ids) > 0 {
		cond = cond.And("id__in", builder.ids...)
	}
	if builder.isOnlyEnable {
		cond = cond.And("enable", true)
	}
	count, wishlist, err := models.GetWishList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(cond).Limit(builder.pageOption.PageSize).Offset((builder.pageOption.Page - 1) * builder.pageOption.PageSize)
	})
	return *count, wishlist, err
}

func (builder *WishListQueryBuilder) DeleteWishList() error {
	cond := orm.NewCondition()
	if len(builder.userIds) > 0 {
		cond = cond.And("user_id__in", builder.userIds...)
	}
	if len(builder.ids) > 0 {
		cond = cond.And("id__in", builder.ids...)
	}
	err := models.DeleteWishList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(cond)
	})
	return err
}
