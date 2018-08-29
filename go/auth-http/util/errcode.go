package util

import (
	"errors"
)

var (
	ErrOK = errors.New("OK")									//成功
	ErrUnknown  = errors.New("Unknown ")						//未知异常（系统异常）
	ErrInvalidArgument  = errors.New("InvalidArgument")		//参数错误
	ErrNotFound  = errors.New("NotFound")						//NotFound
	ErrAlreadyExists  = errors.New("AlreadyExists")			//AlreadyExists
	ErrPermissionDenied  = errors.New("PermissionDenied")		//拒绝访问
	ErrUnauthenticated  = errors.New("Unauthenticated ")		//账号校验失败

	//账号
	ErrUsername  = errors.New("Username is incorrect or empty ")		//用户名错误
	ErrPassword  = errors.New("Password is incorrect or empty ")		//密码错误
	ErrToken  = errors.New("Token is nil or invalid ")					//用户未登录或token失效
	ErrAccountDisableToUse  = errors.New("Account_Disable_to_Use")		//账号禁用
	ErrAccountNotActive  = errors.New("Account_Not_Active")			//账号未激活
	ErrAccountException  = errors.New("Account_Exception")				//账号异常


)
