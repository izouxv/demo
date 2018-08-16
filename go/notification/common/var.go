package common

import (
	"github.com/jinzhu/gorm"
	"errors"
	)

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrDoesNotExist  = errors.New("object does not exist")
	ErrCanNotDelete  = errors.New("object should not delete ")
	InvalidArgument  = errors.New("invalid argument")
	PermissionDenied = errors.New("permission denied")
	ErrDataFailRule  = errors.New("data fail rule ")
	ErrUnknown       = errors.New("unknown")
)

var DB *gorm.DB
