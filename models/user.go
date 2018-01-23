package models

import (
	"time"
)
//用户模块
type User struct {
	Id        int
	Username  string
	Password  string
	LastLogin time.Time
	Enable    bool
	Created   time.Time
	Updated   time.Time
}
func (u *User) TableName() string {
	return "auth_user"
}