package common

import "github.com/pkg/errors"

var UploadDir   = "./upload/"
var ServerAddress string

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrDoesNotExist  = errors.New("object does not exist")
	ErrCanNotDelete  = errors.New("object should not delete ")
	InvalidArgument  = errors.New("invalid argument")
	PermissionDenied = errors.New("permission denied")
)

const (
	FilesUrl =  "/v1.0/file"
	FileUrl  =  "/v1.0/file/:fid"
)

