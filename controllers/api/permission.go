package api

import "github.com/pkg/errors"

var PermissionDeniedError = errors.New("permission denied")

type ApiPermissionInterface interface {
	CheckPermission(context map[string]interface{}) bool
}



