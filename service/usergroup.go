package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type UserGroupQueryBuilder struct {
	ResourceQueryBuilder
}

func (b *UserGroupQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *UserGroupQueryBuilder) Query() (*int64, []*models.UserGroup, error) {
	condition := b.build()
	count, groups, err := models.GetUserGroupList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
	return &count, groups, err
}

func AddUserGroupPermission(groupId int, ids []int) error {
	query := UserGroupQueryBuilder{}
	query.InId(groupId)
	count, result, err := query.Query()
	if err != nil {
		return err
	}
	if *count != 1 {
		return NotFound
	}
	userGroup := result[0]
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}
	transaction := func() error {
		m2m := o.QueryM2M(userGroup, "Permissions")
		permissions := make([]*models.Permission, 0)
		for _, permissionId := range ids {
			permissions = append(permissions, &models.Permission{Id: permissionId})
		}
		_, err := m2m.Add(permissions)
		return err
	}
	err = transaction()
	if err != nil {
		roolErr := o.Rollback()
		if roolErr != nil {
			return roolErr
		}
		return err
	} else {
		err = o.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveUserGroupPermission(groupId int, ids []int) error {
	query := UserGroupQueryBuilder{}
	query.InId(groupId)
	count, result, err := query.Query()
	if err != nil {
		return err
	}
	if *count != 1 {
		return NotFound
	}
	userGroup := result[0]
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}
	transaction := func() error {
		m2m := o.QueryM2M(userGroup, "Permissions")
		permissions := make([]*models.Permission, 0)
		for _, permissionId := range ids {
			permissions = append(permissions, &models.Permission{Id: permissionId})
		}
		_, err := m2m.Remove(permissions)
		return err
	}
	err = transaction()
	if err != nil {
		roolErr := o.Rollback()
		if roolErr != nil {
			return roolErr
		}
		return err
	} else {
		err = o.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func AddUserGroupUsers(userGroupId int, ids []int) error {
	query := UserGroupQueryBuilder{}
	query.InId(userGroupId)
	count, result, err := query.Query()
	if err != nil {
		return err
	}
	if *count != 1 {
		return NotFound
	}
	userGroup := result[0]
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}
	transaction := func() error {
		m2m := o.QueryM2M(userGroup, "Users")
		users := make([]*models.User, 0)
		for _, userId := range ids {
			users = append(users, &models.User{Id: userId})
		}
		_, err := m2m.Add(users)
		return err
	}
	err = transaction()
	if err != nil {
		roolErr := o.Rollback()
		if roolErr != nil {
			return roolErr
		}
		return err
	} else {
		err = o.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveUserGroupUsers(userGroupId int, ids []int) error {
	query := UserGroupQueryBuilder{}
	query.InId(userGroupId)
	count, result, err := query.Query()
	if err != nil {
		return err
	}
	if *count != 1 {
		return NotFound
	}
	userGroup := result[0]
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}
	transaction := func() error {
		m2m := o.QueryM2M(userGroup, "Users")
		users := make([]*models.User, 0)
		for _, userId := range ids {
			users = append(users, &models.User{Id: userId})
		}
		_, err := m2m.Remove(users)
		return err
	}
	err = transaction()
	if err != nil {
		roolErr := o.Rollback()
		if roolErr != nil {
			return roolErr
		}
		return err
	} else {
		err = o.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}
