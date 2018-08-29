package util

type reply struct {
	Code   int32       `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

var CodeMap = map[string]reply{
	"Successful":   {10000, "成功", nil},
	"SystemError": {10001, "系统异常", nil},
	//系统
	"TokenIsInvalid":         		{10002, "Token失效或未登录，请重新登录", nil},
	"PermissionDenied":          	{10003, "权限不足，拒绝访问", nil},
	"UserKickedOut":              	{10004, "用户被踢出", nil},
	"GetContextUserInfoError": 		{10005, "登陆信息失效，请重新登录", nil},
	"BodyIsIncorrectOrEmpty":		{10006, "Body为空或输入有误", nil},
	"InvalidArgument":				{10007, "请求参数有误", nil},
	"JsonError":					{10008, "Json有误", nil},
	"URLDoesNotExist":				{10009, "不存在使用该域名的租户", nil},

	//账号
	"UsernameIsIncorrectOrEmpty":     {20001, "Username为空或格式有误", nil},
	"PasswordIsIncorrectOrEmpty":     {20002, "Password为空或输入有误", nil},
	"NicknameIsIncorrectOrEmpty":     {20003, "Nickname为空或输入有误", nil},
	"AccountDisableToUse":    		  {20004, "账号不可用，请联系管理员", nil},
	"AccountNotActive":    			  {20005, "账号未激活，请通过邀请邮件激活", nil},
	"AccountException":    			  {20006, "账号异常，请联系管理员", nil},
	"UserDoesNotExist":               {20007, "用户不存在", nil},
	"UsernameAndPasswordError":       {20008, "用户名或密码错误", nil},
	"NewPasswordIsIncorrectOrEmpty":  {20009, "新密码为空或输入有误", nil},
	"PasswordTokenIsIncorrectOrEmpty":{20010, "重置密码Token输入有误或失效", nil},
	"UserAlreadyExist":				  {20011, "用户已存在", nil},
	"UidIsIncorrectOrEmpty":     	  {20012, "Uid为空或格式有误", nil},

	//租户
	"TidIsIncorrectOrEmpty":     		{30001, "Tid为空或格式有误", nil},
	"TenantNameIsIncorrectOrEmpty":     {30002, "TenantName为空或格式有误", nil},
	"CanNotDeleteTenant":     			{30003, "存在子租户，无法删除", nil},
	"TenantSystemError":     			{30004, "租户已被禁用，请联系管理员", nil},
	"TenantAlreadyActivated":     		{30005, "租户已激活，无需再次邀请", nil},

	//角色
	"RoleNameIsIncorrectOrEmpty":     	{40001, "RoleName为空或格式有误", nil},
	"MidsIsIncorrectOrEmpty":			{40002, "Mids为空或输入有误", nil},
	"RidIsIncorrectOrEmpty":         	{40003, "Rid为空或输入有误", nil},

	//域
	"DidIsIncorrectOrEmpty":     		{50001, "Did为空或格式有误", nil},
	"DomainNameIsIncorrectOrEmpty":     {50002, "DomainName为空或格式有误", nil},

	//其他
	"NotFind":     						{60001, "未找到请求的资源", nil},
	"PerpageIsIncorrectOrEmpty":     	{60002, "Perpage为空或格式有误", nil},
	"PageIsIncorrectOrEmpty":     		{60003, "Page为空或格式有误", nil},


}

func GetCodeAndMsg(msg string) reply {
	return CodeMap[msg]
}
