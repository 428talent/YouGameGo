package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type WishListQueryBuilder struct {
	ids        []interface{}
	userIds    []interface{}
	gameIds    []interface{}
	pageOption PageOption
	enable     string
}

func (builder *WishListQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	count, result, err := builder.GetWishList()
	return &count, result, err
}

func (builder *WishListQueryBuilder) InId(id ...interface{}) {
	builder.ids = append(builder.ids, id...)
}

func (builder *WishListQueryBuilder) SetPage(page int64, pageSize int64) {
	builder.pageOption = PageOption{
		Page:     page,
		PageSize: pageSize,
	}
}

func (builder *WishListQueryBuilder) WithGame(gameId ...interface{}) {
	builder.gameIds = append(builder.gameIds, gameId...)
}

func (builder *WishListQueryBuilder) BelongToUser(userId ...interface{}) *WishListQueryBuilder {
	builder.userIds = append(builder.userIds, userId...)
	return builder
}

func (builder *WishListQueryBuilder) WithEnable(visibility string) {
	builder.enable = visibility
}
func (builder *WishListQueryBuilder) GetWishList() (int64, []*models.WishList, error) {
	cond := orm.NewCondition()
	if len(builder.userIds) > 0 {
		cond = cond.And("user_id__in", builder.userIds...)
	}
	if len(builder.ids) > 0 {
		cond = cond.And("id__in", builder.ids...)
	}
	if len(builder.enable) > 0 {
		switch builder.enable {
		case "visit":
			cond = cond.And("enable", true)
		case "remove":
			cond = cond.And("enable", false)
		}

	}
	if len(builder.gameIds) > 0 {
		cond = cond.And("game_id__in", builder.gameIds...)
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
