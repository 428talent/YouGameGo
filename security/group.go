package security

import "yougame.com/yougame-server/models"

const (
	UserGroupAdmin = "admin"
	Anonymous      = "anonymous"
)

func CheckUserGroup(user *models.User, expect string) bool {
	if user == nil {
		return false
	}
	if user.UserGroups == nil {
		err := user.ReadUserGroup()
		if err != nil {
			return false
		}
	}
	for _, group := range user.UserGroups {
		if group.Name == expect {
			return true
		}
	}
	return false
}
