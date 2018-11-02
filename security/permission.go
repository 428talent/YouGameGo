package security

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type QueryPermissionAccessResult struct {
	count int
}

func CheckClaimsPermission(claims UserClaims, permissionName string) error {
	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.LoadRelated(user, "UserGroups", 5)
	if err != nil {
		return err
	}
	result := new(models.Permission)
	err = o.Raw(`select distinct permission.*
		from permission
        	inner join user_group_permissions on permission.id = user_group_permissions.permission_id
        	inner join user_group on user_group.id = user_group_permissions.user_group_id
        	inner join auth_user_user_groups on auth_user_user_groups.user_group_id = user_group.id
		where auth_user_id = ? and permission.name = ?`, claims.UserId, permissionName).QueryRow(&result)
	if err != nil {
		return err
	}
	return nil
}
