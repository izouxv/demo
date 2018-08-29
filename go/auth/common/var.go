package common

import (
	"errors"
	"auth/go-drbac/drbac"
)

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrDoesNotExist  = errors.New("object does not exist")
	ErrCanNotDelete  = errors.New("object should not delete ")
)

// 最大深度
var MaxDepth int32 = 6

var Drbac *drbac.DrbacServer