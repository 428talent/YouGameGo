package service

import "errors"

var (
	NoAuthError = errors.New("auth failed")
)