package api

import (
	. "file-server/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strings"
)

var errToCode = map[error]codes.Code{
	ErrAlreadyExists: codes.AlreadyExists,
	ErrDoesNotExist:  codes.NotFound,
	InvalidArgument:  codes.InvalidArgument,
	PermissionDenied: codes.PermissionDenied,
}

func errToRPCError(err error) error {
	if strings.ContainsAny(err.Error(), "PRIMARY") {
		err = ErrAlreadyExists
	}
	cause := errors.Cause(err)
	code, ok := errToCode[cause]
	if !ok {
		code = codes.Unknown
	}

	return grpc.Errorf(code, cause.Error())
}
