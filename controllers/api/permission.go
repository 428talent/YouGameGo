package api

import "errors"

var PermissionDeniedError = errors.New("permission denied")

type PermissionInterface interface {
	CheckPermission(context map[string]interface{}) bool
}
