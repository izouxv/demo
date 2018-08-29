package util

import (
	"google.golang.org/grpc/status"
)

var (
	Successful = status.Error(10000,"成功")
	SystemError = status.Error(10001,"系统异常")

	TokenIsInvalid = status.Error(10002,"Token失效或未登录，请重新登录")
	PermissionDenied = status.Error(10003,"权限不足，拒绝访问")
	UserKickedOut = status.Error(10004,"用户被踢出")
	GetContextUserInfoError = status.Error(10005,"登陆信息失效，请重新登录")
	BodyIsIncorrectOrEmpty = status.Error(10006,"Body为空或输入有误")
	InvalidArgument = status.Error(10007,"求参数有误")
	JsonError = status.Error(10008,"Json有误")
	URLDoesNotExist = status.Error(10009,"不存在使用该域名的租户")

	UsernameIsIncorrectOrEmpty = status.Error(20001,"Username为空或格式有误")
	PasswordIsIncorrectOrEmpty = status.Error(20002,"Password为空或输入有误")
	NicknameIsIncorrectOrEmpty = status.Error(20003,"Nickname为空或输入有误")
	AccountDisableToUse = status.Error(20004,"账号不可用，请联系管理员")
	AccountNotActive = status.Error(20005,"账号未激活，请通过邀请邮件激活")
	AccountException = status.Error(20006,"账号异常，请联系管理员")
	UserDoesNotExist = status.Error(20007,"用户不存在")
	UsernameAndPasswordError = status.Error(20008,"用户名或密码错误")
	PasswordTokenIsIncorrectOrEmpty = status.Error(20010,"重置密码Token输入有误或失效")
	UserAlreadyExist = status.Error(200011,"用户已存在")


	TidIsIncorrectOrEmpty = status.Error(30001,"Tid为空或格式有误")
	TenantNameIsIncorrectOrEmpty = status.Error(30002, "TenantName为空或格式有误")
	CanNotDeleteTenant = status.Error(30003, "存在子租户，无法删除")
	TenantSystemError = status.Error(30004,"租户被禁用或系统异常")
	TenantAlreadyActivated  = status.Error(30005,"租户已激活，无需再次邀请")

	RoleNameIsIncorrectOrEmpty = status.Error(40001,"RoleName为空或格式有误")
	MidsIsIncorrectOrEmpty = status.Error(40002,"Mids为空或输入有误")
	RidIsIncorrectOrEmpty = status.Error(40003,"Rid为空或输入有误")


	DidIsIncorrectOrEmpty = status.Error(50001,"Did为空或格式有误")
	DomainNameIsIncorrectOrEmpty = status.Error(50002,"DomainName为空或格式有误")

	NotFind = status.Error(60001,"未找到请求的资源")

	InsufficientAccountBalance = status.Error(70001,"余额不足")
)
