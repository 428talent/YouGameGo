package service

import "errors"

var (
	NoAuthError         = errors.New("auth failed")
	PermissionNotAccess = errors.New("permission not access")
	ResourceNotEnable   = errors.New("resource not enable")
	UserNotBoughtGood   = errors.New("user not bought good")
	NotFound            = errors.New("not found resources")
)
